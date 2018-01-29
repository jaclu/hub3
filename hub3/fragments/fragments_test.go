// Copyright © 2017 Delving B.V. <info@delving.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fragments_test

import (
	fmt "fmt"
	"net/url"
	"reflect"

	. "bitbucket.org/delving/rapid/hub3/fragments"
	r "github.com/deiu/rdf2go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// testGraph creates a dummy graph for testing
func testGraph() *r.Graph {
	baseUri := "https://rapid.org/resource"

	g := r.NewGraph(baseUri)
	g.Add(r.NewTriple(r.NewResource("a"), r.NewResource("b"), r.NewResource("c")))
	//r.NewTriple(r.NewResource("a"), r.NewResource("title"), r.NewLiteral("title")),
	//r.NewTriple(r.NewResource("a"), r.NewResource("subject"), r.NewLiteralWithLanguage("subject", "nl")),
	return g
}

// URIRef is a function to create an RDFLiteal resource
func URIRef(uri string) r.Term {
	return r.NewResource(uri)
}

// Literal is a utility function to create a RDF literal
func Literal(value string, language string, dataType ObjectXSDType) r.Term {
	if language != "" {
		return r.NewLiteralWithLanguage(value, language)
	}
	return r.NewLiteral(value)
}

var _ = Describe("Fragments", func() {

	Describe("When receiving a Triple", func() {

		spec := "test-spec"
		rev := int32(1)
		ng := "urn:1/graph"
		fg := FragmentGraph{
			Spec:          spec,
			Revision:      rev,
			NamedGraphURI: ng,
		}
		Context("with an object resource", func() {

			t := r.NewTriple(URIRef("urn:1"), URIRef("urn:subject"), URIRef("urn:target"))
			f, err := fg.CreateFragment(t)

			It("should have a spec", func() {
				Expect(t).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
				Expect(f).ToNot(BeNil())
				Expect(f.GetSpec()).To(Equal(spec))

				Expect(f.GetPredicate()).To(Equal("urn:subject"))
			})

			It("should have a revision number", func() {
				Expect(f.GetRevision()).To(Equal(rev))
			})

			It("should have a NamedGraphURI", func() {
				Expect(f.GetNamedGraphURI()).To(Equal(ng))
			})

			It("should have a subject without <>", func() {
				r := f.GetSubject()
				Expect(r).To(Equal("urn:1"))
				Expect(r).ToNot(HaveSuffix("%s", ">"))
				Expect(r).ToNot(HavePrefix("%s", "<"))
			})

			It("should have predicate without <>", func() {
				r := f.GetPredicate()
				Expect(r).To(Equal("urn:subject"))
				Expect(r).ToNot(HaveSuffix("%s", ">"))
				Expect(r).ToNot(HavePrefix("%s", "<"))
			})

			It("should have an object", func() {
				r := f.GetObject()
				Expect(r).To(Equal("urn:target"))
				Expect(r).ToNot(HaveSuffix("%s", ">"))
				Expect(r).ToNot(HavePrefix("%s", "<"))
				fmt.Println(reflect.TypeOf(r))
				fmt.Println(t.String())
			})

			It("should have resource as objecttype", func() {
				t := f.GetObjectType()
				Expect(t).ToNot(BeNil())
				Expect(t).To(Equal(ObjectType_RESOURCE))
			})

		})

		Context("when receiving a triple with a literal object", func() {

			t := r.NewTriple(URIRef("urn:1"), URIRef("urn:subject"), Literal("river", "", ObjectXSDType_STRING))
			f, err := fg.CreateFragment(t)

			It("should have literal as objecttype", func() {
				Expect(err).ToNot(HaveOccurred())
				t := f.GetObjectType()
				Expect(t).ToNot(BeNil())
				Expect(t).To(Equal(ObjectType_LITERAL))
			})

			It("should have no language", func() {
				Expect(f.Language).To(Equal(""))
			})

			It("should have string as datatype", func() {
				Expect(f.DataType).To(Equal(ObjectXSDType_STRING))
			})
		})

		Context("when receiving a triple with a literal and language", func() {
			t := r.NewTriple(URIRef("urn:1"), URIRef("urn:subject"), Literal("river", "en", ObjectXSDType_STRING))
			f, err := fg.CreateFragment(t)

			It("should have a language", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Language).To(Equal("en"))
			})

			It("should have string as datatype", func() {
				Expect(f.DataType).To(Equal(ObjectXSDType_STRING))
			})
		})
	})

	Describe("creating a new FragmentRequest", func() {

		Context("directly", func() {

			It("should have no triple pattern set", func() {
				fr := NewFragmentRequest()
				Expect(fr).ToNot(BeNil())
				Expect(fr.GetSubject()).To(BeEmpty())
				Expect(fr.GetPredicate()).To(BeEmpty())
				Expect(fr.GetObject()).To(BeEmpty())
				Expect(fr.GetLanguage()).To(BeEmpty())
			})

			It("should have a non-zero page start", func() {
				fr := NewFragmentRequest()
				Expect(fr.GetPage()).To(Equal(int32(1)))
			})
		})

		Context("parsing from url.Values", func() {

			It("should ignore empty values", func() {
				fr := NewFragmentRequest()
				v := url.Values{}
				v.Add("subject", "urn:1")
				v.Add("predicate", "")
				v.Add("object", "")
				v.Add("language", "")
				err := fr.ParseQueryString(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(fr.GetSubject()).To(Equal("urn:1"))
				Expect(fr.GetPredicate()).To(BeEmpty())
				Expect(fr.GetObject()).To(BeEmpty())
				Expect(fr.GetLanguage()).To(BeEmpty())
			})
		})
	})
})
