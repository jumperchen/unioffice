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

type CT_CustomWorkbookViews struct {
	// Custom Workbook View
	CustomWorkbookView []*CT_CustomWorkbookView
}

func NewCT_CustomWorkbookViews() *CT_CustomWorkbookViews {
	ret := &CT_CustomWorkbookViews{}
	return ret
}
func (m *CT_CustomWorkbookViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secustomWorkbookView := xml.StartElement{Name: xml.Name{Local: "x:customWorkbookView"}}
	e.EncodeElement(m.CustomWorkbookView, secustomWorkbookView)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomWorkbookViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomWorkbookViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customWorkbookView":
				tmp := NewCT_CustomWorkbookView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustomWorkbookView = append(m.CustomWorkbookView, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomWorkbookViews
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomWorkbookViews) Validate() error {
	return m.ValidateWithPath("CT_CustomWorkbookViews")
}
func (m *CT_CustomWorkbookViews) ValidateWithPath(path string) error {
	for i, v := range m.CustomWorkbookView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustomWorkbookView[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
