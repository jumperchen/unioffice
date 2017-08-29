// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Styles struct {
	// Document Default Paragraph and Run Properties
	DocDefaults *CT_DocDefaults
	// Latent Style Information
	LatentStyles *CT_LatentStyles
	// Style Definition
	Style []*CT_Style
}

func NewCT_Styles() *CT_Styles {
	ret := &CT_Styles{}
	return ret
}
func (m *CT_Styles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.DocDefaults != nil {
		sedocDefaults := xml.StartElement{Name: xml.Name{Local: "w:docDefaults"}}
		e.EncodeElement(m.DocDefaults, sedocDefaults)
	}
	if m.LatentStyles != nil {
		selatentStyles := xml.StartElement{Name: xml.Name{Local: "w:latentStyles"}}
		e.EncodeElement(m.LatentStyles, selatentStyles)
	}
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "w:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Styles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Styles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "docDefaults":
				m.DocDefaults = NewCT_DocDefaults()
				if err := d.DecodeElement(m.DocDefaults, &el); err != nil {
					return err
				}
			case "latentStyles":
				m.LatentStyles = NewCT_LatentStyles()
				if err := d.DecodeElement(m.LatentStyles, &el); err != nil {
					return err
				}
			case "style":
				tmp := NewCT_Style()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Style = append(m.Style, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Styles
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Styles) Validate() error {
	return m.ValidateWithPath("CT_Styles")
}
func (m *CT_Styles) ValidateWithPath(path string) error {
	if m.DocDefaults != nil {
		if err := m.DocDefaults.ValidateWithPath(path + "/DocDefaults"); err != nil {
			return err
		}
	}
	if m.LatentStyles != nil {
		if err := m.LatentStyles.ValidateWithPath(path + "/LatentStyles"); err != nil {
			return err
		}
	}
	for i, v := range m.Style {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Style[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
