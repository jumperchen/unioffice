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

type CT_PivotCaches struct {
	// PivotCache
	PivotCache []*CT_PivotCache
}

func NewCT_PivotCaches() *CT_PivotCaches {
	ret := &CT_PivotCaches{}
	return ret
}
func (m *CT_PivotCaches) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sepivotCache := xml.StartElement{Name: xml.Name{Local: "x:pivotCache"}}
	e.EncodeElement(m.PivotCache, sepivotCache)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotCaches) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PivotCaches:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pivotCache":
				tmp := NewCT_PivotCache()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotCache = append(m.PivotCache, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotCaches
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotCaches) Validate() error {
	return m.ValidateWithPath("CT_PivotCaches")
}
func (m *CT_PivotCaches) ValidateWithPath(path string) error {
	for i, v := range m.PivotCache {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotCache[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
