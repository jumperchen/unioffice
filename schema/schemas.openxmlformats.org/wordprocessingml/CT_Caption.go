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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Caption struct {
	// Caption Type Name
	NameAttr string
	// Automatic Caption Placement
	PosAttr ST_CaptionPos
	// Include Chapter Number in Field for Caption
	ChapNumAttr *sharedTypes.ST_OnOff
	// Style for Chapter Headings
	HeadingAttr *int32
	// Do Not Include Name In Caption
	NoLabelAttr *sharedTypes.ST_OnOff
	// Caption Numbering Format
	NumFmtAttr ST_NumberFormat
	// Chapter Number/Item Index Separator
	SepAttr ST_ChapterSep
}

func NewCT_Caption() *CT_Caption {
	ret := &CT_Caption{}
	return ret
}
func (m *CT_Caption) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.PosAttr != ST_CaptionPosUnset {
		attr, err := m.PosAttr.MarshalXMLAttr(xml.Name{Local: "w:pos"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ChapNumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:chapNum"},
			Value: fmt.Sprintf("%v", *m.ChapNumAttr)})
	}
	if m.HeadingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:heading"},
			Value: fmt.Sprintf("%v", *m.HeadingAttr)})
	}
	if m.NoLabelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:noLabel"},
			Value: fmt.Sprintf("%v", *m.NoLabelAttr)})
	}
	if m.NumFmtAttr != ST_NumberFormatUnset {
		attr, err := m.NumFmtAttr.MarshalXMLAttr(xml.Name{Local: "w:numFmt"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SepAttr != ST_ChapterSepUnset {
		attr, err := m.SepAttr.MarshalXMLAttr(xml.Name{Local: "w:sep"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Caption) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "pos" {
			m.PosAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "chapNum" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ChapNumAttr = &parsed
		}
		if attr.Name.Local == "heading" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.HeadingAttr = &pt
		}
		if attr.Name.Local == "noLabel" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NoLabelAttr = &parsed
		}
		if attr.Name.Local == "numFmt" {
			m.NumFmtAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "sep" {
			m.SepAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Caption: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Caption) Validate() error {
	return m.ValidateWithPath("CT_Caption")
}
func (m *CT_Caption) ValidateWithPath(path string) error {
	if err := m.PosAttr.ValidateWithPath(path + "/PosAttr"); err != nil {
		return err
	}
	if m.ChapNumAttr != nil {
		if err := m.ChapNumAttr.ValidateWithPath(path + "/ChapNumAttr"); err != nil {
			return err
		}
	}
	if m.NoLabelAttr != nil {
		if err := m.NoLabelAttr.ValidateWithPath(path + "/NoLabelAttr"); err != nil {
			return err
		}
	}
	if err := m.NumFmtAttr.ValidateWithPath(path + "/NumFmtAttr"); err != nil {
		return err
	}
	if err := m.SepAttr.ValidateWithPath(path + "/SepAttr"); err != nil {
		return err
	}
	return nil
}
