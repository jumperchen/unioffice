// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CustomColor struct {
	NameAttr  *string
	ScrgbClr  *CT_ScRgbColor
	SrgbClr   *CT_SRgbColor
	HslClr    *CT_HslColor
	SysClr    *CT_SystemColor
	SchemeClr *CT_SchemeColor
	PrstClr   *CT_PresetColor
}

func NewCT_CustomColor() *CT_CustomColor {
	ret := &CT_CustomColor{}
	return ret
}
func (m *CT_CustomColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ScrgbClr != nil {
		sescrgbClr := xml.StartElement{Name: xml.Name{Local: "a:scrgbClr"}}
		e.EncodeElement(m.ScrgbClr, sescrgbClr)
	}
	if m.SrgbClr != nil {
		sesrgbClr := xml.StartElement{Name: xml.Name{Local: "a:srgbClr"}}
		e.EncodeElement(m.SrgbClr, sesrgbClr)
	}
	if m.HslClr != nil {
		sehslClr := xml.StartElement{Name: xml.Name{Local: "a:hslClr"}}
		e.EncodeElement(m.HslClr, sehslClr)
	}
	if m.SysClr != nil {
		sesysClr := xml.StartElement{Name: xml.Name{Local: "a:sysClr"}}
		e.EncodeElement(m.SysClr, sesysClr)
	}
	if m.SchemeClr != nil {
		seschemeClr := xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}}
		e.EncodeElement(m.SchemeClr, seschemeClr)
	}
	if m.PrstClr != nil {
		seprstClr := xml.StartElement{Name: xml.Name{Local: "a:prstClr"}}
		e.EncodeElement(m.PrstClr, seprstClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
	}
lCT_CustomColor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "scrgbClr":
				m.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(m.ScrgbClr, &el); err != nil {
					return err
				}
			case "srgbClr":
				m.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(m.SrgbClr, &el); err != nil {
					return err
				}
			case "hslClr":
				m.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(m.HslClr, &el); err != nil {
					return err
				}
			case "sysClr":
				m.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(m.SysClr, &el); err != nil {
					return err
				}
			case "schemeClr":
				m.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(m.SchemeClr, &el); err != nil {
					return err
				}
			case "prstClr":
				m.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(m.PrstClr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomColor
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomColor) Validate() error {
	return m.ValidateWithPath("CT_CustomColor")
}
func (m *CT_CustomColor) ValidateWithPath(path string) error {
	if m.ScrgbClr != nil {
		if err := m.ScrgbClr.ValidateWithPath(path + "/ScrgbClr"); err != nil {
			return err
		}
	}
	if m.SrgbClr != nil {
		if err := m.SrgbClr.ValidateWithPath(path + "/SrgbClr"); err != nil {
			return err
		}
	}
	if m.HslClr != nil {
		if err := m.HslClr.ValidateWithPath(path + "/HslClr"); err != nil {
			return err
		}
	}
	if m.SysClr != nil {
		if err := m.SysClr.ValidateWithPath(path + "/SysClr"); err != nil {
			return err
		}
	}
	if m.SchemeClr != nil {
		if err := m.SchemeClr.ValidateWithPath(path + "/SchemeClr"); err != nil {
			return err
		}
	}
	if m.PrstClr != nil {
		if err := m.PrstClr.ValidateWithPath(path + "/PrstClr"); err != nil {
			return err
		}
	}
	return nil
}
