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
)

type CT_RgbColor struct {
	// Alpha Red Green Blue
	RgbAttr *string
}

func NewCT_RgbColor() *CT_RgbColor {
	ret := &CT_RgbColor{}
	return ret
}
func (m *CT_RgbColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RgbAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rgb"},
			Value: fmt.Sprintf("%v", *m.RgbAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RgbColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rgb" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RgbAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RgbColor: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RgbColor) Validate() error {
	return m.ValidateWithPath("CT_RgbColor")
}
func (m *CT_RgbColor) ValidateWithPath(path string) error {
	return nil
}
