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

type CT_AdjPoint2D struct {
	XAttr ST_AdjCoordinate
	YAttr ST_AdjCoordinate
}

func NewCT_AdjPoint2D() *CT_AdjPoint2D {
	ret := &CT_AdjPoint2D{}
	return ret
}
func (m *CT_AdjPoint2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "y"},
		Value: fmt.Sprintf("%v", m.YAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AdjPoint2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.XAttr = parsed
		}
		if attr.Name.Local == "y" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.YAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AdjPoint2D: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AdjPoint2D) Validate() error {
	return m.ValidateWithPath("CT_AdjPoint2D")
}
func (m *CT_AdjPoint2D) ValidateWithPath(path string) error {
	if err := m.XAttr.ValidateWithPath(path + "/XAttr"); err != nil {
		return err
	}
	if err := m.YAttr.ValidateWithPath(path + "/YAttr"); err != nil {
		return err
	}
	return nil
}
