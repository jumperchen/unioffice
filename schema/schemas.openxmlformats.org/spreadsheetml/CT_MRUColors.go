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
	"log"
)

type CT_MRUColors struct {
	// Color
	Color []*CT_Color
}

func NewCT_MRUColors() *CT_MRUColors {
	ret := &CT_MRUColors{}
	return ret
}
func (m *CT_MRUColors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secolor := xml.StartElement{Name: xml.Name{Local: "x:color"}}
	e.EncodeElement(m.Color, secolor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MRUColors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MRUColors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "color":
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MRUColors
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MRUColors) Validate() error {
	return m.ValidateWithPath("CT_MRUColors")
}
func (m *CT_MRUColors) ValidateWithPath(path string) error {
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
