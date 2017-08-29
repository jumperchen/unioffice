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

type CT_Surface3DChart struct {
	Wireframe *CT_Boolean
	Ser       []*CT_SurfaceSer
	BandFmts  *CT_BandFmts
	AxId      []*CT_UnsignedInt
	ExtLst    *CT_ExtensionList
}

func NewCT_Surface3DChart() *CT_Surface3DChart {
	ret := &CT_Surface3DChart{}
	return ret
}
func (m *CT_Surface3DChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Wireframe != nil {
		sewireframe := xml.StartElement{Name: xml.Name{Local: "wireframe"}}
		e.EncodeElement(m.Wireframe, sewireframe)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "ser"}}
		e.EncodeElement(m.Ser, seser)
	}
	if m.BandFmts != nil {
		sebandFmts := xml.StartElement{Name: xml.Name{Local: "bandFmts"}}
		e.EncodeElement(m.BandFmts, sebandFmts)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "axId"}}
	e.EncodeElement(m.AxId, seaxId)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Surface3DChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Surface3DChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wireframe":
				m.Wireframe = NewCT_Boolean()
				if err := d.DecodeElement(m.Wireframe, &el); err != nil {
					return err
				}
			case "ser":
				tmp := NewCT_SurfaceSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case "bandFmts":
				m.BandFmts = NewCT_BandFmts()
				if err := d.DecodeElement(m.BandFmts, &el); err != nil {
					return err
				}
			case "axId":
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
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
			break lCT_Surface3DChart
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Surface3DChart) Validate() error {
	return m.ValidateWithPath("CT_Surface3DChart")
}
func (m *CT_Surface3DChart) ValidateWithPath(path string) error {
	if m.Wireframe != nil {
		if err := m.Wireframe.ValidateWithPath(path + "/Wireframe"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.BandFmts != nil {
		if err := m.BandFmts.ValidateWithPath(path + "/BandFmts"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
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
