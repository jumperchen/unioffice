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

type CT_TabStop struct {
	// Tab Stop Type
	ValAttr ST_TabJc
	// Tab Leader Character
	LeaderAttr ST_TabTlc
	// Tab Stop Position
	PosAttr ST_SignedTwipsMeasure
}

func NewCT_TabStop() *CT_TabStop {
	ret := &CT_TabStop{}
	return ret
}
func (m *CT_TabStop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LeaderAttr != ST_TabTlcUnset {
		attr, err := m.LeaderAttr.MarshalXMLAttr(xml.Name{Local: "w:leader"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:pos"},
		Value: fmt.Sprintf("%v", m.PosAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TabStop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "leader" {
			m.LeaderAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "pos" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.PosAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TabStop: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TabStop) Validate() error {
	return m.ValidateWithPath("CT_TabStop")
}
func (m *CT_TabStop) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TabJcUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if err := m.LeaderAttr.ValidateWithPath(path + "/LeaderAttr"); err != nil {
		return err
	}
	if err := m.PosAttr.ValidateWithPath(path + "/PosAttr"); err != nil {
		return err
	}
	return nil
}
