// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_DPt struct {
	Idx              *CT_UnsignedInt
	InvertIfNegative *CT_Boolean
	Marker           *CT_Marker
	Bubble3D         *CT_Boolean
	Explosion        *CT_UnsignedInt
	SpPr             *drawingml.CT_ShapeProperties
	PictureOptions   *CT_PictureOptions
	ExtLst           *CT_ExtensionList
}

func NewCT_DPt() *CT_DPt {
	ret := &CT_DPt{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}
func (m *CT_DPt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.InvertIfNegative != nil {
		seinvertIfNegative := xml.StartElement{Name: xml.Name{Local: "invertIfNegative"}}
		e.EncodeElement(m.InvertIfNegative, seinvertIfNegative)
	}
	if m.Marker != nil {
		semarker := xml.StartElement{Name: xml.Name{Local: "marker"}}
		e.EncodeElement(m.Marker, semarker)
	}
	if m.Bubble3D != nil {
		sebubble3D := xml.StartElement{Name: xml.Name{Local: "bubble3D"}}
		e.EncodeElement(m.Bubble3D, sebubble3D)
	}
	if m.Explosion != nil {
		seexplosion := xml.StartElement{Name: xml.Name{Local: "explosion"}}
		e.EncodeElement(m.Explosion, seexplosion)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.PictureOptions != nil {
		sepictureOptions := xml.StartElement{Name: xml.Name{Local: "pictureOptions"}}
		e.EncodeElement(m.PictureOptions, sepictureOptions)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DPt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_DPt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "invertIfNegative":
				m.InvertIfNegative = NewCT_Boolean()
				if err := d.DecodeElement(m.InvertIfNegative, &el); err != nil {
					return err
				}
			case "marker":
				m.Marker = NewCT_Marker()
				if err := d.DecodeElement(m.Marker, &el); err != nil {
					return err
				}
			case "bubble3D":
				m.Bubble3D = NewCT_Boolean()
				if err := d.DecodeElement(m.Bubble3D, &el); err != nil {
					return err
				}
			case "explosion":
				m.Explosion = NewCT_UnsignedInt()
				if err := d.DecodeElement(m.Explosion, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "pictureOptions":
				m.PictureOptions = NewCT_PictureOptions()
				if err := d.DecodeElement(m.PictureOptions, &el); err != nil {
					return err
				}
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
			break lCT_DPt
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DPt) Validate() error {
	return m.ValidateWithPath("CT_DPt")
}
func (m *CT_DPt) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.InvertIfNegative != nil {
		if err := m.InvertIfNegative.ValidateWithPath(path + "/InvertIfNegative"); err != nil {
			return err
		}
	}
	if m.Marker != nil {
		if err := m.Marker.ValidateWithPath(path + "/Marker"); err != nil {
			return err
		}
	}
	if m.Bubble3D != nil {
		if err := m.Bubble3D.ValidateWithPath(path + "/Bubble3D"); err != nil {
			return err
		}
	}
	if m.Explosion != nil {
		if err := m.Explosion.ValidateWithPath(path + "/Explosion"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.PictureOptions != nil {
		if err := m.PictureOptions.ValidateWithPath(path + "/PictureOptions"); err != nil {
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
