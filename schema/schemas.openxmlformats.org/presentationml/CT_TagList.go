// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TagList struct {
	// Programmable Extensibility Tag
	Tag []*CT_StringTag
}

func NewCT_TagList() *CT_TagList {
	ret := &CT_TagList{}
	return ret
}
func (m *CT_TagList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Tag != nil {
		setag := xml.StartElement{Name: xml.Name{Local: "p:tag"}}
		e.EncodeElement(m.Tag, setag)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TagList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TagList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tag":
				tmp := NewCT_StringTag()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tag = append(m.Tag, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TagList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TagList) Validate() error {
	return m.ValidateWithPath("CT_TagList")
}
func (m *CT_TagList) ValidateWithPath(path string) error {
	for i, v := range m.Tag {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tag[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
