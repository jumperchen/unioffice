// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_StretchInfoProperties struct {
	FillRect *CT_RelativeRect
}

func NewCT_StretchInfoProperties() *CT_StretchInfoProperties {
	ret := &CT_StretchInfoProperties{}
	return ret
}
func (m *CT_StretchInfoProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.FillRect != nil {
		sefillRect := xml.StartElement{Name: xml.Name{Local: "a:fillRect"}}
		e.EncodeElement(m.FillRect, sefillRect)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_StretchInfoProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_StretchInfoProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fillRect":
				m.FillRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.FillRect, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StretchInfoProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_StretchInfoProperties) Validate() error {
	return m.ValidateWithPath("CT_StretchInfoProperties")
}
func (m *CT_StretchInfoProperties) ValidateWithPath(path string) error {
	if m.FillRect != nil {
		if err := m.FillRect.ValidateWithPath(path + "/FillRect"); err != nil {
			return err
		}
	}
	return nil
}
