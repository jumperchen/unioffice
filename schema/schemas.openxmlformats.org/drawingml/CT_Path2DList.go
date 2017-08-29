// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Path2DList struct {
	Path []*CT_Path2D
}

func NewCT_Path2DList() *CT_Path2DList {
	ret := &CT_Path2DList{}
	return ret
}
func (m *CT_Path2DList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Path != nil {
		sepath := xml.StartElement{Name: xml.Name{Local: "a:path"}}
		e.EncodeElement(m.Path, sepath)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Path2DList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Path2DList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "path":
				tmp := NewCT_Path2D()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Path = append(m.Path, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2DList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Path2DList) Validate() error {
	return m.ValidateWithPath("CT_Path2DList")
}
func (m *CT_Path2DList) ValidateWithPath(path string) error {
	for i, v := range m.Path {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Path[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
