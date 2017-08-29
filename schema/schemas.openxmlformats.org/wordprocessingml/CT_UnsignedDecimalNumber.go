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
)

type CT_UnsignedDecimalNumber struct {
	// Positive Decimal Number Value
	ValAttr uint64
}

func NewCT_UnsignedDecimalNumber() *CT_UnsignedDecimalNumber {
	ret := &CT_UnsignedDecimalNumber{}
	return ret
}
func (m *CT_UnsignedDecimalNumber) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_UnsignedDecimalNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ValAttr = uint64(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_UnsignedDecimalNumber: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_UnsignedDecimalNumber) Validate() error {
	return m.ValidateWithPath("CT_UnsignedDecimalNumber")
}
func (m *CT_UnsignedDecimalNumber) ValidateWithPath(path string) error {
	return nil
}
