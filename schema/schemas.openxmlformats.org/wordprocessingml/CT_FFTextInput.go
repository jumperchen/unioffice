// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_FFTextInput struct {
	// Text Box Form Field Type
	Type *CT_FFTextType
	// Default Text Box Form Field String
	Default *CT_String
	// Text Box Form Field Maximum Length
	MaxLength *CT_DecimalNumber
	// Text Box Form Field Formatting
	Format *CT_String
}

func NewCT_FFTextInput() *CT_FFTextInput {
	ret := &CT_FFTextInput{}
	return ret
}
func (m *CT_FFTextInput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "w:CT_FFTextInput"
	e.EncodeToken(start)
	start.Attr = nil
	if m.Type != nil {
		setype := xml.StartElement{Name: xml.Name{Local: "w:type"}}
		e.EncodeElement(m.Type, setype)
	}
	if m.Default != nil {
		sedefault := xml.StartElement{Name: xml.Name{Local: "w:default"}}
		e.EncodeElement(m.Default, sedefault)
	}
	if m.MaxLength != nil {
		semaxLength := xml.StartElement{Name: xml.Name{Local: "w:maxLength"}}
		e.EncodeElement(m.MaxLength, semaxLength)
	}
	if m.Format != nil {
		seformat := xml.StartElement{Name: xml.Name{Local: "w:format"}}
		e.EncodeElement(m.Format, seformat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FFTextInput) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFTextInput:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "type":
				m.Type = NewCT_FFTextType()
				if err := d.DecodeElement(m.Type, &el); err != nil {
					return err
				}
			case "default":
				m.Default = NewCT_String()
				if err := d.DecodeElement(m.Default, &el); err != nil {
					return err
				}
			case "maxLength":
				m.MaxLength = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.MaxLength, &el); err != nil {
					return err
				}
			case "format":
				m.Format = NewCT_String()
				if err := d.DecodeElement(m.Format, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFTextInput
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FFTextInput) Validate() error {
	return m.ValidateWithPath("CT_FFTextInput")
}
func (m *CT_FFTextInput) ValidateWithPath(path string) error {
	if m.Type != nil {
		if err := m.Type.ValidateWithPath(path + "/Type"); err != nil {
			return err
		}
	}
	if m.Default != nil {
		if err := m.Default.ValidateWithPath(path + "/Default"); err != nil {
			return err
		}
	}
	if m.MaxLength != nil {
		if err := m.MaxLength.ValidateWithPath(path + "/MaxLength"); err != nil {
			return err
		}
	}
	if m.Format != nil {
		if err := m.Format.ValidateWithPath(path + "/Format"); err != nil {
			return err
		}
	}
	return nil
}
