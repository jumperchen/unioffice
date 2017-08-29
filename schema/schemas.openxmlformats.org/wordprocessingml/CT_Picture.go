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

type CT_Picture struct {
	// Embedded Video
	Movie *CT_Rel
	// Floating Embedded Control
	Control *CT_Control
}

func NewCT_Picture() *CT_Picture {
	ret := &CT_Picture{}
	return ret
}
func (m *CT_Picture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Movie != nil {
		semovie := xml.StartElement{Name: xml.Name{Local: "w:movie"}}
		e.EncodeElement(m.Movie, semovie)
	}
	if m.Control != nil {
		secontrol := xml.StartElement{Name: xml.Name{Local: "w:control"}}
		e.EncodeElement(m.Control, secontrol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Picture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Picture:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "movie":
				m.Movie = NewCT_Rel()
				if err := d.DecodeElement(m.Movie, &el); err != nil {
					return err
				}
			case "control":
				m.Control = NewCT_Control()
				if err := d.DecodeElement(m.Control, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Picture
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Picture) Validate() error {
	return m.ValidateWithPath("CT_Picture")
}
func (m *CT_Picture) ValidateWithPath(path string) error {
	if m.Movie != nil {
		if err := m.Movie.ValidateWithPath(path + "/Movie"); err != nil {
			return err
		}
	}
	if m.Control != nil {
		if err := m.Control.ValidateWithPath(path + "/Control"); err != nil {
			return err
		}
	}
	return nil
}
