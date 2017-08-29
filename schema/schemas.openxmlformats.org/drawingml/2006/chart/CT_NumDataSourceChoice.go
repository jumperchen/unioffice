// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_NumDataSourceChoice struct {
	NumRef *CT_NumRef
	NumLit *CT_NumData
}

func NewCT_NumDataSourceChoice() *CT_NumDataSourceChoice {
	ret := &CT_NumDataSourceChoice{}
	return ret
}
func (m *CT_NumDataSourceChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NumRef != nil {
		senumRef := xml.StartElement{Name: xml.Name{Local: "numRef"}}
		e.EncodeElement(m.NumRef, senumRef)
	}
	if m.NumLit != nil {
		senumLit := xml.StartElement{Name: xml.Name{Local: "numLit"}}
		e.EncodeElement(m.NumLit, senumLit)
	}
	return nil
}
func (m *CT_NumDataSourceChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NumDataSourceChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "numRef":
				m.NumRef = NewCT_NumRef()
				if err := d.DecodeElement(m.NumRef, &el); err != nil {
					return err
				}
			case "numLit":
				m.NumLit = NewCT_NumData()
				if err := d.DecodeElement(m.NumLit, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumDataSourceChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumDataSourceChoice) Validate() error {
	return m.ValidateWithPath("CT_NumDataSourceChoice")
}
func (m *CT_NumDataSourceChoice) ValidateWithPath(path string) error {
	if m.NumRef != nil {
		if err := m.NumRef.ValidateWithPath(path + "/NumRef"); err != nil {
			return err
		}
	}
	if m.NumLit != nil {
		if err := m.NumLit.ValidateWithPath(path + "/NumLit"); err != nil {
			return err
		}
	}
	return nil
}
