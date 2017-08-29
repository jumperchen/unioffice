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

type CT_Acc struct {
	AccPr *CT_AccPr
	E     *CT_OMathArg
}

func NewCT_Acc() *CT_Acc {
	ret := &CT_Acc{}
	ret.E = NewCT_OMathArg()
	return ret
}
func (m *CT_Acc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.AccPr != nil {
		seaccPr := xml.StartElement{Name: xml.Name{Local: "m:accPr"}}
		e.EncodeElement(m.AccPr, seaccPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Acc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_Acc:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "accPr":
				m.AccPr = NewCT_AccPr()
				if err := d.DecodeElement(m.AccPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Acc
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Acc) Validate() error {
	return m.ValidateWithPath("CT_Acc")
}
func (m *CT_Acc) ValidateWithPath(path string) error {
	if m.AccPr != nil {
		if err := m.AccPr.ValidateWithPath(path + "/AccPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
