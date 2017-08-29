// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Comment struct {
	// Comment Author ID
	AuthorIdAttr uint32
	// Comment Date/Time
	DtAttr *time.Time
	// Comment Index
	IdxAttr uint32
	// Comment Position
	Pos *drawingml.CT_Point2D
	// Comment's Text Content
	Text   string
	ExtLst *CT_ExtensionListModify
}

func NewCT_Comment() *CT_Comment {
	ret := &CT_Comment{}
	ret.Pos = drawingml.NewCT_Point2D()
	return ret
}
func (m *CT_Comment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "authorId"},
		Value: fmt.Sprintf("%v", m.AuthorIdAttr)})
	if m.DtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dt"},
			Value: fmt.Sprintf("%v", *m.DtAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	sepos := xml.StartElement{Name: xml.Name{Local: "p:pos"}}
	e.EncodeElement(m.Pos, sepos)
	setext := xml.StartElement{Name: xml.Name{Local: "p:text"}}
	gooxml.AddPreserveSpaceAttr(&setext, m.Text)
	e.EncodeElement(m.Text, setext)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Comment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = drawingml.NewCT_Point2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "authorId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.AuthorIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "dt" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DtAttr = &parsed
		}
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
		}
	}
lCT_Comment:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case "text":
				if err := d.DecodeElement(m.Text, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
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
			break lCT_Comment
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Comment) Validate() error {
	return m.ValidateWithPath("CT_Comment")
}
func (m *CT_Comment) ValidateWithPath(path string) error {
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
