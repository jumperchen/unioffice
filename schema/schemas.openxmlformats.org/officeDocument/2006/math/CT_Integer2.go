// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Integer2 struct {
	ValAttr int32
}

func NewCT_Integer2() *CT_Integer2 {
	ret := &CT_Integer2{}
	return ret
}
func (m *CT_Integer2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "m:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Integer2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Integer2: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Integer2) Validate() error {
	return m.ValidateWithPath("CT_Integer2")
}
func (m *CT_Integer2) ValidateWithPath(path string) error {
	if m.ValAttr < -2 {
		return fmt.Errorf("%s/m.ValAttr must be >= -2 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 2 {
		return fmt.Errorf("%s/m.ValAttr must be <= 2 (have %v)", path, m.ValAttr)
	}
	return nil
}
