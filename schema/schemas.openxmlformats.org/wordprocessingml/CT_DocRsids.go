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

type CT_DocRsids struct {
	// Original Document Revision Save ID
	RsidRoot *CT_LongHexNumber
	// Single Session Revision Save ID
	Rsid []*CT_LongHexNumber
}

func NewCT_DocRsids() *CT_DocRsids {
	ret := &CT_DocRsids{}
	return ret
}
func (m *CT_DocRsids) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.RsidRoot != nil {
		sersidRoot := xml.StartElement{Name: xml.Name{Local: "w:rsidRoot"}}
		e.EncodeElement(m.RsidRoot, sersidRoot)
	}
	if m.Rsid != nil {
		sersid := xml.StartElement{Name: xml.Name{Local: "w:rsid"}}
		e.EncodeElement(m.Rsid, sersid)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DocRsids) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocRsids:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rsidRoot":
				m.RsidRoot = NewCT_LongHexNumber()
				if err := d.DecodeElement(m.RsidRoot, &el); err != nil {
					return err
				}
			case "rsid":
				tmp := NewCT_LongHexNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rsid = append(m.Rsid, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocRsids
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DocRsids) Validate() error {
	return m.ValidateWithPath("CT_DocRsids")
}
func (m *CT_DocRsids) ValidateWithPath(path string) error {
	if m.RsidRoot != nil {
		if err := m.RsidRoot.ValidateWithPath(path + "/RsidRoot"); err != nil {
			return err
		}
	}
	for i, v := range m.Rsid {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rsid[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
