// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_MetadataRecord struct {
	// Metadata Record Type Index
	TAttr uint32
	// Metadata Record Value Index
	VAttr uint32
}

func NewCT_MetadataRecord() *CT_MetadataRecord {
	ret := &CT_MetadataRecord{}
	return ret
}
func (m *CT_MetadataRecord) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
		Value: fmt.Sprintf("%v", m.VAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MetadataRecord) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.TAttr = uint32(parsed)
		}
		if attr.Name.Local == "v" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.VAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MetadataRecord: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_MetadataRecord) Validate() error {
	return m.ValidateWithPath("CT_MetadataRecord")
}
func (m *CT_MetadataRecord) ValidateWithPath(path string) error {
	return nil
}
