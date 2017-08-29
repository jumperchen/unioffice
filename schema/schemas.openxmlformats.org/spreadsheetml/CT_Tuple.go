// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Tuple struct {
	// Field Index
	FldAttr *uint32
	// Hierarchy Index
	HierAttr *uint32
	// Item Index
	ItemAttr uint32
}

func NewCT_Tuple() *CT_Tuple {
	ret := &CT_Tuple{}
	return ret
}
func (m *CT_Tuple) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.FldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fld"},
			Value: fmt.Sprintf("%v", *m.FldAttr)})
	}
	if m.HierAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hier"},
			Value: fmt.Sprintf("%v", *m.HierAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "item"},
		Value: fmt.Sprintf("%v", m.ItemAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Tuple) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.FldAttr = &pt
		}
		if attr.Name.Local == "hier" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.HierAttr = &pt
		}
		if attr.Name.Local == "item" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ItemAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Tuple: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Tuple) Validate() error {
	return m.ValidateWithPath("CT_Tuple")
}
func (m *CT_Tuple) ValidateWithPath(path string) error {
	return nil
}
