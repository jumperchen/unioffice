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
)

type CT_TblWidth struct {
	// Table Width Value
	WAttr *ST_MeasurementOrPercent
	// Table Width Type
	TypeAttr ST_TblWidth
}

func NewCT_TblWidth() *CT_TblWidth {
	ret := &CT_TblWidth{}
	return ret
}
func (m *CT_TblWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.TypeAttr != ST_TblWidthUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
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
func (m *CT_TblWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_MeasurementOrPercent(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TblWidth: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TblWidth) Validate() error {
	return m.ValidateWithPath("CT_TblWidth")
}
func (m *CT_TblWidth) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
			return err
		}
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
