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
)

type CT_PhotoAlbum struct {
	// Black and White
	BwAttr *bool
	// Show/Hide Captions
	ShowCaptionsAttr *bool
	// Photo Album Layout
	LayoutAttr ST_PhotoAlbumLayout
	// Frame Type
	FrameAttr ST_PhotoAlbumFrameShape
	ExtLst    *CT_ExtensionList
}

func NewCT_PhotoAlbum() *CT_PhotoAlbum {
	ret := &CT_PhotoAlbum{}
	return ret
}
func (m *CT_PhotoAlbum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BwAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bw"},
			Value: fmt.Sprintf("%v", *m.BwAttr)})
	}
	if m.ShowCaptionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showCaptions"},
			Value: fmt.Sprintf("%v", *m.ShowCaptionsAttr)})
	}
	if m.LayoutAttr != ST_PhotoAlbumLayoutUnset {
		attr, err := m.LayoutAttr.MarshalXMLAttr(xml.Name{Local: "layout"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FrameAttr != ST_PhotoAlbumFrameShapeUnset {
		attr, err := m.FrameAttr.MarshalXMLAttr(xml.Name{Local: "frame"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PhotoAlbum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bw" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BwAttr = &parsed
		}
		if attr.Name.Local == "showCaptions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCaptionsAttr = &parsed
		}
		if attr.Name.Local == "layout" {
			m.LayoutAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "frame" {
			m.FrameAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PhotoAlbum:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
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
			break lCT_PhotoAlbum
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PhotoAlbum) Validate() error {
	return m.ValidateWithPath("CT_PhotoAlbum")
}
func (m *CT_PhotoAlbum) ValidateWithPath(path string) error {
	if err := m.LayoutAttr.ValidateWithPath(path + "/LayoutAttr"); err != nil {
		return err
	}
	if err := m.FrameAttr.ValidateWithPath(path + "/FrameAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
