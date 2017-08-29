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

type CT_GvmlConnectorNonVisual struct {
	CNvPr      *CT_NonVisualDrawingProps
	CNvCxnSpPr *CT_NonVisualConnectorProperties
}

func NewCT_GvmlConnectorNonVisual() *CT_GvmlConnectorNonVisual {
	ret := &CT_GvmlConnectorNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvCxnSpPr = NewCT_NonVisualConnectorProperties()
	return ret
}
func (m *CT_GvmlConnectorNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "a:cNvCxnSpPr"}}
	e.EncodeElement(m.CNvCxnSpPr, secNvCxnSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GvmlConnectorNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvCxnSpPr = NewCT_NonVisualConnectorProperties()
lCT_GvmlConnectorNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvCxnSpPr":
				if err := d.DecodeElement(m.CNvCxnSpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlConnectorNonVisual
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GvmlConnectorNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlConnectorNonVisual")
}
func (m *CT_GvmlConnectorNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvCxnSpPr.ValidateWithPath(path + "/CNvCxnSpPr"); err != nil {
		return err
	}
	return nil
}
