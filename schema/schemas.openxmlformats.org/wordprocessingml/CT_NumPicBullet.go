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
	"strconv"
)

type CT_NumPicBullet struct {
	// Picture Numbering Symbol ID
	NumPicBulletIdAttr int32
	// Picture Numbering Symbol Properties
	Pict    *CT_Picture
	Drawing *CT_Drawing
}

func NewCT_NumPicBullet() *CT_NumPicBullet {
	ret := &CT_NumPicBullet{}
	return ret
}
func (m *CT_NumPicBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:numPicBulletId"},
		Value: fmt.Sprintf("%v", m.NumPicBulletIdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.Pict != nil {
		sepict := xml.StartElement{Name: xml.Name{Local: "w:pict"}}
		e.EncodeElement(m.Pict, sepict)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumPicBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "numPicBulletId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NumPicBulletIdAttr = int32(parsed)
		}
	}
lCT_NumPicBullet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pict":
				m.Pict = NewCT_Picture()
				if err := d.DecodeElement(m.Pict, &el); err != nil {
					return err
				}
			case "drawing":
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumPicBullet
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumPicBullet) Validate() error {
	return m.ValidateWithPath("CT_NumPicBullet")
}
func (m *CT_NumPicBullet) ValidateWithPath(path string) error {
	if m.Pict != nil {
		if err := m.Pict.ValidateWithPath(path + "/Pict"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	return nil
}
