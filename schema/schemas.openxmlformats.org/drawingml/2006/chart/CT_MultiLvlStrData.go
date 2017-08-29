// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_MultiLvlStrData struct {
	PtCount *CT_UnsignedInt
	Lvl     []*CT_Lvl
	ExtLst  *CT_ExtensionList
}

func NewCT_MultiLvlStrData() *CT_MultiLvlStrData {
	ret := &CT_MultiLvlStrData{}
	return ret
}
func (m *CT_MultiLvlStrData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.PtCount != nil {
		septCount := xml.StartElement{Name: xml.Name{Local: "ptCount"}}
		e.EncodeElement(m.PtCount, septCount)
	}
	if m.Lvl != nil {
		selvl := xml.StartElement{Name: xml.Name{Local: "lvl"}}
		e.EncodeElement(m.Lvl, selvl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MultiLvlStrData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MultiLvlStrData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ptCount":
				m.PtCount = NewCT_UnsignedInt()
				if err := d.DecodeElement(m.PtCount, &el); err != nil {
					return err
				}
			case "lvl":
				tmp := NewCT_Lvl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Lvl = append(m.Lvl, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MultiLvlStrData
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MultiLvlStrData) Validate() error {
	return m.ValidateWithPath("CT_MultiLvlStrData")
}
func (m *CT_MultiLvlStrData) ValidateWithPath(path string) error {
	if m.PtCount != nil {
		if err := m.PtCount.ValidateWithPath(path + "/PtCount"); err != nil {
			return err
		}
	}
	for i, v := range m.Lvl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Lvl[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
