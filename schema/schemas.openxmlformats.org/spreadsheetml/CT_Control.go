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
	"strconv"
)

type CT_Control struct {
	// Shape Id
	ShapeIdAttr uint32
	IdAttr      string
	// Control Name
	NameAttr *string
	// Embedded Control Properties
	ControlPr *CT_ControlPr
}

func NewCT_Control() *CT_Control {
	ret := &CT_Control{}
	return ret
}
func (m *CT_Control) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shapeId"},
		Value: fmt.Sprintf("%v", m.ShapeIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ControlPr != nil {
		secontrolPr := xml.StartElement{Name: xml.Name{Local: "x:controlPr"}}
		e.EncodeElement(m.ControlPr, secontrolPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Control) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "shapeId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ShapeIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
	}
lCT_Control:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "controlPr":
				m.ControlPr = NewCT_ControlPr()
				if err := d.DecodeElement(m.ControlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Control
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Control) Validate() error {
	return m.ValidateWithPath("CT_Control")
}
func (m *CT_Control) ValidateWithPath(path string) error {
	if m.ControlPr != nil {
		if err := m.ControlPr.ValidateWithPath(path + "/ControlPr"); err != nil {
			return err
		}
	}
	return nil
}
