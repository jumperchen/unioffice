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
)

type CT_DashStop struct {
	DAttr  ST_PositivePercentage
	SpAttr ST_PositivePercentage
}

func NewCT_DashStop() *CT_DashStop {
	ret := &CT_DashStop{}
	return ret
}
func (m *CT_DashStop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "d"},
		Value: fmt.Sprintf("%v", m.DAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sp"},
		Value: fmt.Sprintf("%v", m.SpAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DashStop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "d" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.DAttr = parsed
		}
		if attr.Name.Local == "sp" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.SpAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DashStop: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_DashStop) Validate() error {
	return m.ValidateWithPath("CT_DashStop")
}
func (m *CT_DashStop) ValidateWithPath(path string) error {
	if err := m.DAttr.ValidateWithPath(path + "/DAttr"); err != nil {
		return err
	}
	if err := m.SpAttr.ValidateWithPath(path + "/SpAttr"); err != nil {
		return err
	}
	return nil
}
