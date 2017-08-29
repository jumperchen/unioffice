// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package relationships

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Relationships struct {
	Relationship []*Relationship
}

func NewCT_Relationships() *CT_Relationships {
	ret := &CT_Relationships{}
	return ret
}
func (m *CT_Relationships) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Relationship != nil {
		seRelationship := xml.StartElement{Name: xml.Name{Local: "Relationship"}}
		e.EncodeElement(m.Relationship, seRelationship)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Relationships) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Relationships:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Relationship":
				tmp := NewRelationship()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Relationship = append(m.Relationship, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Relationships
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Relationships) Validate() error {
	return m.ValidateWithPath("CT_Relationships")
}
func (m *CT_Relationships) ValidateWithPath(path string) error {
	for i, v := range m.Relationship {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Relationship[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
