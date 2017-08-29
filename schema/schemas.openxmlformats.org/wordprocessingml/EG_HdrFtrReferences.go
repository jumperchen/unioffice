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

type EG_HdrFtrReferences struct {
	// Header Reference
	HeaderReference *CT_HdrFtrRef
	// Footer Reference
	FooterReference *CT_HdrFtrRef
}

func NewEG_HdrFtrReferences() *EG_HdrFtrReferences {
	ret := &EG_HdrFtrReferences{}
	return ret
}
func (m *EG_HdrFtrReferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.HeaderReference != nil {
		seheaderReference := xml.StartElement{Name: xml.Name{Local: "w:headerReference"}}
		e.EncodeElement(m.HeaderReference, seheaderReference)
	}
	if m.FooterReference != nil {
		sefooterReference := xml.StartElement{Name: xml.Name{Local: "w:footerReference"}}
		e.EncodeElement(m.FooterReference, sefooterReference)
	}
	return nil
}
func (m *EG_HdrFtrReferences) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_HdrFtrReferences:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "headerReference":
				m.HeaderReference = NewCT_HdrFtrRef()
				if err := d.DecodeElement(m.HeaderReference, &el); err != nil {
					return err
				}
			case "footerReference":
				m.FooterReference = NewCT_HdrFtrRef()
				if err := d.DecodeElement(m.FooterReference, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_HdrFtrReferences
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_HdrFtrReferences) Validate() error {
	return m.ValidateWithPath("EG_HdrFtrReferences")
}
func (m *EG_HdrFtrReferences) ValidateWithPath(path string) error {
	if m.HeaderReference != nil {
		if err := m.HeaderReference.ValidateWithPath(path + "/HeaderReference"); err != nil {
			return err
		}
	}
	if m.FooterReference != nil {
		if err := m.FooterReference.ValidateWithPath(path + "/FooterReference"); err != nil {
			return err
		}
	}
	return nil
}
