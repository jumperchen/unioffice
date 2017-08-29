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

type CT_FFCheckBoxChoice struct {
	Size     *CT_HpsMeasure
	SizeAuto *CT_OnOff
}

func NewCT_FFCheckBoxChoice() *CT_FFCheckBoxChoice {
	ret := &CT_FFCheckBoxChoice{}
	return ret
}
func (m *CT_FFCheckBoxChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Size != nil {
		sesize := xml.StartElement{Name: xml.Name{Local: "w:size"}}
		e.EncodeElement(m.Size, sesize)
	}
	if m.SizeAuto != nil {
		sesizeAuto := xml.StartElement{Name: xml.Name{Local: "w:sizeAuto"}}
		e.EncodeElement(m.SizeAuto, sesizeAuto)
	}
	return nil
}
func (m *CT_FFCheckBoxChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFCheckBoxChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "size":
				m.Size = NewCT_HpsMeasure()
				if err := d.DecodeElement(m.Size, &el); err != nil {
					return err
				}
			case "sizeAuto":
				m.SizeAuto = NewCT_OnOff()
				if err := d.DecodeElement(m.SizeAuto, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFCheckBoxChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FFCheckBoxChoice) Validate() error {
	return m.ValidateWithPath("CT_FFCheckBoxChoice")
}
func (m *CT_FFCheckBoxChoice) ValidateWithPath(path string) error {
	if m.Size != nil {
		if err := m.Size.ValidateWithPath(path + "/Size"); err != nil {
			return err
		}
	}
	if m.SizeAuto != nil {
		if err := m.SizeAuto.ValidateWithPath(path + "/SizeAuto"); err != nil {
			return err
		}
	}
	return nil
}
