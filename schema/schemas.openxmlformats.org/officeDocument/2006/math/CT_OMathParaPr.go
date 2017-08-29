// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_OMathParaPr struct {
	Jc *CT_OMathJc
}

func NewCT_OMathParaPr() *CT_OMathParaPr {
	ret := &CT_OMathParaPr{}
	return ret
}
func (m *CT_OMathParaPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Jc != nil {
		sejc := xml.StartElement{Name: xml.Name{Local: "m:jc"}}
		e.EncodeElement(m.Jc, sejc)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OMathParaPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OMathParaPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "jc":
				m.Jc = NewCT_OMathJc()
				if err := d.DecodeElement(m.Jc, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OMathParaPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OMathParaPr) Validate() error {
	return m.ValidateWithPath("CT_OMathParaPr")
}
func (m *CT_OMathParaPr) ValidateWithPath(path string) error {
	if m.Jc != nil {
		if err := m.Jc.ValidateWithPath(path + "/Jc"); err != nil {
			return err
		}
	}
	return nil
}
