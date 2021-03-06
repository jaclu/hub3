// Copyright 2017 Delving B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/OneOfOne/xxhash"
	"github.com/allegro/bigcache"
	c "github.com/delving/hub3/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var httpCache *bigcache.BigCache

// CachedResponse stores the request in the Cache.
// This object is always return from the CacheRequest
type CachedResponse struct {
	Body        []byte
	StatusCode  int
	ContentType string
}

func RegisterCache(r chi.Router) {
	// init the big cache
	eviction := time.Duration(c.Config.Cache.LifeWindowMinutes) * time.Minute
	config := bigcache.DefaultConfig(eviction)
	config.HardMaxCacheSize = c.Config.Cache.HardMaxCacheSize
	config.MaxEntrySize = c.Config.Cache.MaxEntrySize

	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to start bigCache implementation: %#v", err))
	}

	httpCache = cache

	r.Get("/api/cache/stats", cacheStats)
	r.Handle(fmt.Sprintf("%s/*", c.Config.Cache.APIPrefix), cacheHandler())
}

// PrepareCacheRequest modifies the request for the remote call
// It returns the unique hash from the request that is used as the cacheKey
func PrepareCacheRequest(r *http.Request) (cacheKey string, err error) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		domain = c.Config.Cache.CacheDomain
	} else {
		u, err := url.Parse(domain)
		if err != nil {
			log.Printf("Unable to parse domain %s due to: %s", domain, err)
			return "", err
		}

		domain = u.Host

		params := r.URL.Query()
		params.Del("domain")
		r.URL.RawQuery = params.Encode()
	}

	r.URL.Host = domain
	r.Host = domain
	r.RequestURI = ""
	r.URL.Scheme = "https"
	if c.Config.Cache.StripPrefix {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, c.Config.Cache.APIPrefix)
	}

	method := r.Method
	path := r.URL.EscapedPath()
	contentType := r.Header.Get("Content-Type")
	var b bytes.Buffer
	b.WriteString(method + path + contentType + r.URL.RawQuery)

	//dump, _ := httputil.DumpRequest(r, true)
	//fmt.Printf("%s\n", dump)
	if r.Body != nil {
		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			log.Printf("Unable to read body for creating cache key: %#v", readErr)
			return "", readErr
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		b.Write(body)
	}

	hash := xxhash.Checksum64(b.Bytes())
	return fmt.Sprintf("%016x", hash), nil
}

func getCachedRequest(r *http.Request) (cr *CachedResponse, err error) {
	cacheKey, err := PrepareCacheRequest(r)
	fmt.Printf("%s\n", cacheKey)

	if err != nil {
		return cr, err
	}

	entry, err := httpCache.Get(cacheKey)
	if err == nil {
		dec := gob.NewDecoder(bytes.NewBuffer(entry))
		cacheCr := CachedResponse{}
		err = dec.Decode(&cacheCr)
		if err != nil {
			log.Printf("Unable to decode cached entry: %#v\n", err)
			return cr, err
		}
		return &cacheCr, err
	}

	// Create New http Transport
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: transCfg,
	}
	log.Printf(r.URL.String(), r.Method)

	resp, err := netClient.Do(r)

	// TODO handle context cancelled properly
	if err != nil || resp.Body == nil {
		log.Printf("Error in proxy query: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	cr = &CachedResponse{}
	cr.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Unable to read the response body with error: %s", err)
	}
	cr.StatusCode = resp.StatusCode
	cr.ContentType = resp.Header.Get("Content-Type")

	// set cache key

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err = enc.Encode(cr)
	if err != nil {
		return cr, err
	}
	err = httpCache.Set(cacheKey, b.Bytes())
	if err != nil {
		log.Printf("Unable to set cache for cacheKey: %#v", cacheKey)
		return cr, err
	}
	return
}

func cacheRequest(w http.ResponseWriter, r *http.Request) {
	cr, err := getCachedRequest(r)

	if err != nil {
		log.Printf("Unable to cache request: %#v", err)
		return
	}

	w.Header().Set("Content-Type", cr.ContentType)
	w.WriteHeader(cr.StatusCode)

	_, err = w.Write(cr.Body)
	if err != nil {
		log.Printf("unable to write the response body to the response: %#v", err)
		return
	}
	return
}

func cacheHandler() http.HandlerFunc {
	return http.HandlerFunc(cacheRequest)
}

func cacheStats(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, httpCache.Stats())
	return
}
