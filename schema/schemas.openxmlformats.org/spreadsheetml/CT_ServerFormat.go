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

type CT_ServerFormat struct {
	// Culture
	CultureAttr *string
	// Format
	FormatAttr *string
}

func NewCT_ServerFormat() *CT_ServerFormat {
	ret := &CT_ServerFormat{}
	return ret
}
func (m *CT_ServerFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CultureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "culture"},
			Value: fmt.Sprintf("%v", *m.CultureAttr)})
	}
	if m.FormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "format"},
			Value: fmt.Sprintf("%v", *m.FormatAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ServerFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "culture" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CultureAttr = &parsed
		}
		if attr.Name.Local == "format" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ServerFormat: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ServerFormat) Validate() error {
	return m.ValidateWithPath("CT_ServerFormat")
}
func (m *CT_ServerFormat) ValidateWithPath(path string) error {
	return nil
}
