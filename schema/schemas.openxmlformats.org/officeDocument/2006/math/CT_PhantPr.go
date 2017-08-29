// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_PhantPr struct {
	Show     *CT_OnOff
	ZeroWid  *CT_OnOff
	ZeroAsc  *CT_OnOff
	ZeroDesc *CT_OnOff
	Transp   *CT_OnOff
	CtrlPr   *CT_CtrlPr
}

func NewCT_PhantPr() *CT_PhantPr {
	ret := &CT_PhantPr{}
	return ret
}
func (m *CT_PhantPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Show != nil {
		seshow := xml.StartElement{Name: xml.Name{Local: "m:show"}}
		e.EncodeElement(m.Show, seshow)
	}
	if m.ZeroWid != nil {
		sezeroWid := xml.StartElement{Name: xml.Name{Local: "m:zeroWid"}}
		e.EncodeElement(m.ZeroWid, sezeroWid)
	}
	if m.ZeroAsc != nil {
		sezeroAsc := xml.StartElement{Name: xml.Name{Local: "m:zeroAsc"}}
		e.EncodeElement(m.ZeroAsc, sezeroAsc)
	}
	if m.ZeroDesc != nil {
		sezeroDesc := xml.StartElement{Name: xml.Name{Local: "m:zeroDesc"}}
		e.EncodeElement(m.ZeroDesc, sezeroDesc)
	}
	if m.Transp != nil {
		setransp := xml.StartElement{Name: xml.Name{Local: "m:transp"}}
		e.EncodeElement(m.Transp, setransp)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PhantPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PhantPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "show":
				m.Show = NewCT_OnOff()
				if err := d.DecodeElement(m.Show, &el); err != nil {
					return err
				}
			case "zeroWid":
				m.ZeroWid = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroWid, &el); err != nil {
					return err
				}
			case "zeroAsc":
				m.ZeroAsc = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroAsc, &el); err != nil {
					return err
				}
			case "zeroDesc":
				m.ZeroDesc = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroDesc, &el); err != nil {
					return err
				}
			case "transp":
				m.Transp = NewCT_OnOff()
				if err := d.DecodeElement(m.Transp, &el); err != nil {
					return err
				}
			case "ctrlPr":
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PhantPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PhantPr) Validate() error {
	return m.ValidateWithPath("CT_PhantPr")
}
func (m *CT_PhantPr) ValidateWithPath(path string) error {
	if m.Show != nil {
		if err := m.Show.ValidateWithPath(path + "/Show"); err != nil {
			return err
		}
	}
	if m.ZeroWid != nil {
		if err := m.ZeroWid.ValidateWithPath(path + "/ZeroWid"); err != nil {
			return err
		}
	}
	if m.ZeroAsc != nil {
		if err := m.ZeroAsc.ValidateWithPath(path + "/ZeroAsc"); err != nil {
			return err
		}
	}
	if m.ZeroDesc != nil {
		if err := m.ZeroDesc.ValidateWithPath(path + "/ZeroDesc"); err != nil {
			return err
		}
	}
	if m.Transp != nil {
		if err := m.Transp.ValidateWithPath(path + "/Transp"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
