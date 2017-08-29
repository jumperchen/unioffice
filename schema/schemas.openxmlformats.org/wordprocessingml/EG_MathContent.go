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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type EG_MathContent struct {
	OMathPara *math.OMathPara
	OMath     *math.OMath
}

func NewEG_MathContent() *EG_MathContent {
	ret := &EG_MathContent{}
	return ret
}
func (m *EG_MathContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.OMathPara != nil {
		seoMathPara := xml.StartElement{Name: xml.Name{Local: "m:oMathPara"}}
		e.EncodeElement(m.OMathPara, seoMathPara)
	}
	if m.OMath != nil {
		seoMath := xml.StartElement{Name: xml.Name{Local: "m:oMath"}}
		e.EncodeElement(m.OMath, seoMath)
	}
	return nil
}
func (m *EG_MathContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_MathContent:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oMathPara":
				m.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(m.OMathPara, &el); err != nil {
					return err
				}
			case "oMath":
				m.OMath = math.NewOMath()
				if err := d.DecodeElement(m.OMath, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_MathContent
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_MathContent) Validate() error {
	return m.ValidateWithPath("EG_MathContent")
}
func (m *EG_MathContent) ValidateWithPath(path string) error {
	if m.OMathPara != nil {
		if err := m.OMathPara.ValidateWithPath(path + "/OMathPara"); err != nil {
			return err
		}
	}
	if m.OMath != nil {
		if err := m.OMath.ValidateWithPath(path + "/OMath"); err != nil {
			return err
		}
	}
	return nil
}
