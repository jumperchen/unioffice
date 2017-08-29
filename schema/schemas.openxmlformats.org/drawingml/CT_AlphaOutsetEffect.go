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

type CT_AlphaOutsetEffect struct {
	RadAttr *ST_Coordinate
}

func NewCT_AlphaOutsetEffect() *CT_AlphaOutsetEffect {
	ret := &CT_AlphaOutsetEffect{}
	return ret
}
func (m *CT_AlphaOutsetEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rad"},
			Value: fmt.Sprintf("%v", *m.RadAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AlphaOutsetEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rad" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.RadAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AlphaOutsetEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AlphaOutsetEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaOutsetEffect")
}
func (m *CT_AlphaOutsetEffect) ValidateWithPath(path string) error {
	if m.RadAttr != nil {
		if err := m.RadAttr.ValidateWithPath(path + "/RadAttr"); err != nil {
			return err
		}
	}
	return nil
}
