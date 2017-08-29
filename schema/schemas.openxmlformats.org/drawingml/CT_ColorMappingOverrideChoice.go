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

type CT_ColorMappingOverrideChoice struct {
	MasterClrMapping   *CT_EmptyElement
	OverrideClrMapping *CT_ColorMapping
}

func NewCT_ColorMappingOverrideChoice() *CT_ColorMappingOverrideChoice {
	ret := &CT_ColorMappingOverrideChoice{}
	return ret
}
func (m *CT_ColorMappingOverrideChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MasterClrMapping != nil {
		semasterClrMapping := xml.StartElement{Name: xml.Name{Local: "a:masterClrMapping"}}
		e.EncodeElement(m.MasterClrMapping, semasterClrMapping)
	}
	if m.OverrideClrMapping != nil {
		seoverrideClrMapping := xml.StartElement{Name: xml.Name{Local: "a:overrideClrMapping"}}
		e.EncodeElement(m.OverrideClrMapping, seoverrideClrMapping)
	}
	return nil
}
func (m *CT_ColorMappingOverrideChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorMappingOverrideChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "masterClrMapping":
				m.MasterClrMapping = NewCT_EmptyElement()
				if err := d.DecodeElement(m.MasterClrMapping, &el); err != nil {
					return err
				}
			case "overrideClrMapping":
				m.OverrideClrMapping = NewCT_ColorMapping()
				if err := d.DecodeElement(m.OverrideClrMapping, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorMappingOverrideChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ColorMappingOverrideChoice) Validate() error {
	return m.ValidateWithPath("CT_ColorMappingOverrideChoice")
}
func (m *CT_ColorMappingOverrideChoice) ValidateWithPath(path string) error {
	if m.MasterClrMapping != nil {
		if err := m.MasterClrMapping.ValidateWithPath(path + "/MasterClrMapping"); err != nil {
			return err
		}
	}
	if m.OverrideClrMapping != nil {
		if err := m.OverrideClrMapping.ValidateWithPath(path + "/OverrideClrMapping"); err != nil {
			return err
		}
	}
	return nil
}
