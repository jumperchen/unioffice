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

type CT_Boolean struct {
	// Value
	VAttr bool
	// Unused Item
	UAttr *bool
	// Calculated Item
	FAttr *bool
	// Caption
	CAttr *string
	// Member Property Count
	CpAttr *uint32
	// Member Property Indexes
	X []*CT_X
}

func NewCT_Boolean() *CT_Boolean {
	ret := &CT_Boolean{}
	return ret
}
func (m *CT_Boolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
		Value: fmt.Sprintf("%v", m.VAttr)})
	if m.UAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "u"},
			Value: fmt.Sprintf("%v", *m.UAttr)})
	}
	if m.FAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "f"},
			Value: fmt.Sprintf("%v", *m.FAttr)})
	}
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	if m.CpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cp"},
			Value: fmt.Sprintf("%v", *m.CpAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "x:x"}}
		e.EncodeElement(m.X, sex)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Boolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VAttr = parsed
		}
		if attr.Name.Local == "u" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UAttr = &parsed
		}
		if attr.Name.Local == "f" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FAttr = &parsed
		}
		if attr.Name.Local == "c" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CAttr = &parsed
		}
		if attr.Name.Local == "cp" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CpAttr = &pt
		}
	}
lCT_Boolean:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "x":
				tmp := NewCT_X()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Boolean
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Boolean) Validate() error {
	return m.ValidateWithPath("CT_Boolean")
}
func (m *CT_Boolean) ValidateWithPath(path string) error {
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
