// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type ColorsDefHdr struct {
	CT_ColorTransformHeader
}

func NewColorsDefHdr() *ColorsDefHdr {
	ret := &ColorsDefHdr{}
	ret.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	return ret
}
func (m *ColorsDefHdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "colorsDefHdr"
	return m.CT_ColorTransformHeader.MarshalXML(e, start)
}
func (m *ColorsDefHdr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = parsed
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
		}
		if attr.Name.Local == "resId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.ResIdAttr = &pt
		}
	}
lColorsDefHdr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "title":
				tmp := NewCT_CTName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case "desc":
				tmp := NewCT_CTDescription()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case "catLst":
				m.CatLst = NewCT_CTCategories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
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
			break lColorsDefHdr
		case xml.CharData:
		}
	}
	return nil
}
func (m *ColorsDefHdr) Validate() error {
	return m.ValidateWithPath("ColorsDefHdr")
}
func (m *ColorsDefHdr) ValidateWithPath(path string) error {
	if err := m.CT_ColorTransformHeader.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
