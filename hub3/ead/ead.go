package ead

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"sync/atomic"

	c "github.com/delving/rapid-saas/config"
)

func init() {
	path := c.Config.EAD.CacheDir
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Unable to create cache dir; %s", err)
		}
	}

}

// NodeConfig holds all the configuration options fo generating Archive Nodes
type NodeConfig struct {
	Sparse  bool
	Counter *NodeCounter
}

// NewNodeConfig creates a new NodeConfig
func NewNodeConfig(ctx context.Context, sparse bool) *NodeConfig {
	return &NodeConfig{
		Counter: &NodeCounter{},
		Sparse:  sparse,
	}
}

// NodeCounter is a concurrency safe counter for number of Nodes processed
type NodeCounter struct {
	counter uint64
}

// Increment increments the count by one
func (nc *NodeCounter) Increment() {
	atomic.AddUint64(&nc.counter, 1)
}

// GetCount returns the snapshot of the current count
func (nc *NodeCounter) GetCount() uint64 {
	return atomic.LoadUint64(&nc.counter)
}

// NewNodeList converts the Archival Description Level to a Nodelist
// Nodelist is an optimized lossless Protocol Buffer container.
func (dsc *Cdsc) NewNodeList(cfg *NodeConfig) (*NodeList, uint64, error) {
	nl := &NodeList{}
	nl.Type = dsc.Attrtype
	for _, label := range dsc.Chead {
		nl.Label = append(nl.Label, label.Head)
	}

	for _, nn := range dsc.Nested {
		node, err := nn.NewNode(cfg)
		if err != nil {
			return nil, 0, err
		}
		nl.Nodes = append(nl.Nodes, node)
	}
	return nl, cfg.Counter.GetCount(), nil
}

// Sparse creates a sparse version of the list of Archive Nodes
func (nl *NodeList) Sparse() {
	Sparsify(nl.Nodes)
}

// Sparse creates a sparse version of Header
func (h *Header) Sparse() {
	if h.GetDateAsLabel() {
		h.DateAsLabel = false
		for _, date := range h.GetDate() {
			h.Label = append(h.Label, date.GetLabel())
		}
	}
	h.Date = nil
	h.ID = nil
	h.Physdesc = ""
}

// Sparsify is a recursive function that creates a Sparse representation
// of a list of Nodes. This is mostly used to efficiently create Tree Views
// of the Archive C-Levels
func Sparsify(nodes []*Node) {
	for _, n := range nodes {
		n.HTML = ""
		n.CTag = ""
		n.Header.Sparse()
		if len(n.Nodes) != 0 {
			Sparsify(n.Nodes)
		}

	}
}

// NewNodeID converts a unitid field from the EAD did to a NodeID
func (ui *Cunitid) NewNodeID() (*NodeID, error) {
	id := &NodeID{
		ID:       ui.ID,
		TypeID:   ui.Attridentifier,
		Type:     ui.Attrtype,
		Audience: ui.Attraudience,
	}
	return id, nil
}

// NewNodeIDs extract Unit Identifiers from the EAD did
func (cdid *Cdid) NewNodeIDs(sparse bool) ([]*NodeID, string, error) {
	ids := []*NodeID{}
	var invertoryNumber string
	for _, unitid := range cdid.Cunitid {
		id, err := unitid.NewNodeID()
		if err != nil {
			return nil, "", err
		}
		switch id.GetType() {
		case "ABS", "series_code", "":
			invertoryNumber = id.GetID()
		}
		if !sparse {
			ids = append(ids, id)
		}
	}
	return ids, invertoryNumber, nil
}

// NewNodeDate extract date infomation frme the EAD unitdate
func (date *Cunitdate) NewNodeDate() (*NodeDate, error) {
	nDate := &NodeDate{
		Calendar: date.Attrcalendar,
		Era:      date.Attrera,
		Normal:   date.Attrnormal,
		Label:    date.Date,
	}
	return nDate, nil
}

// NewHeader creates an Archival Header
func (cdid *Cdid) NewHeader(sparse bool) (*Header, error) {
	header := &Header{}
	if cdid.Cphysdesc != nil && !sparse {
		header.Physdesc = cdid.Cphysdesc.PhyscDesc
	}

	for _, label := range cdid.Cunittitle {
		if len(label.Cunitdate) != 0 {
			header.DateAsLabel = true
			for _, date := range label.Cunitdate {
				nodeDate, err := date.NewNodeDate()
				if err != nil {
					return nil, err
				}
				switch sparse {
				case true:
					header.Label = append(header.Label, nodeDate.GetLabel())
					header.DateAsLabel = false
				case false:
					header.Date = append(header.Date, nodeDate)
				}
			}
			continue
		}
		header.Label = append(header.Label, label.Title)
	}

	if !sparse {
		for _, date := range cdid.Cunitdate {
			nodeDate, err := date.NewNodeDate()
			if err != nil {
				return nil, err
			}
			header.Date = append(header.Date, nodeDate)
		}
	}

	nodeIDs, inventoryID, err := cdid.NewNodeIDs(sparse)
	if err != nil {
		return nil, err
	}
	if inventoryID != "" {
		header.InventoryNumber = inventoryID
	}
	header.ID = append(header.ID, nodeIDs...)

	return header, nil
}

// NewNode converts EAD c01 to a Archival Node
func (c *Cc01) NewNode(cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:    c.XMLName.Local,
		Depth:   int32(1),
		Type:    c.Attrlevel,
		SubType: c.Attrotherlevel,
		Order:   cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDS := []string{header.GetInventoryNumber()}

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDS, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc02) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc03) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc04) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc05) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc06) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc07) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc08) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc09) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc10) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc11) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	parentIDs = append(parentIDs, header.GetInventoryNumber())

	if len(c.Nested) != 0 {
		for _, nn := range c.Nested {
			n, err := nn.NewNode(parentIDs, cfg)
			if err != nil {
				return nil, err
			}
			node.Nodes = append(node.Nodes, n)
		}
	}
	return node, nil
}

// NewNode converts EAD nested cLevel to an Archival Node
func (c *Cc12) NewNode(parentIDs []string, cfg *NodeConfig) (*Node, error) {
	cfg.Counter.Increment()
	node := &Node{
		CTag:      c.XMLName.Local,
		Depth:     int32(len(parentIDs) + 1),
		Type:      c.Attrlevel,
		SubType:   c.Attrotherlevel,
		ParentIDs: parentIDs,
		Order:     cfg.Counter.GetCount(),
	}

	// add header
	if cfg.Sparse {
		node.CTag = ""
	}
	header, err := c.Cdid.NewHeader(cfg.Sparse)
	if err != nil {
		return nil, err
	}
	node.Header = header

	// add scope content
	if c.Cscopecontent != nil && !cfg.Sparse {
		html, err := xml.Marshal(c.Cscopecontent.Cp)
		if err != nil {
			return nil, err
		}
		node.HTML = string(html)
	}

	// add nested
	//parentIDs = append(parentIDs, header.GetInventoryNumber())

	//if len(c.Nested) != 0 {
	//for _, nn := range c.Nested {
	//n, err := nn.NewNode(parentIDs, cfg)
	//if err != nil {
	//return nil, err
	//}
	//node.Nodes = append(node.Nodes, n)
	//}
	//}
	return node, nil
}
