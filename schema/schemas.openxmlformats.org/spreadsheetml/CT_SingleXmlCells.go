// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_SingleXmlCells struct {
	// Table Properties
	SingleXmlCell []*CT_SingleXmlCell
}

func NewCT_SingleXmlCells() *CT_SingleXmlCells {
	ret := &CT_SingleXmlCells{}
	return ret
}
func (m *CT_SingleXmlCells) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sesingleXmlCell := xml.StartElement{Name: xml.Name{Local: "x:singleXmlCell"}}
	e.EncodeElement(m.SingleXmlCell, sesingleXmlCell)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SingleXmlCells) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SingleXmlCells:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "singleXmlCell":
				tmp := NewCT_SingleXmlCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SingleXmlCell = append(m.SingleXmlCell, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SingleXmlCells
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SingleXmlCells) Validate() error {
	return m.ValidateWithPath("CT_SingleXmlCells")
}
func (m *CT_SingleXmlCells) ValidateWithPath(path string) error {
	for i, v := range m.SingleXmlCell {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SingleXmlCell[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
