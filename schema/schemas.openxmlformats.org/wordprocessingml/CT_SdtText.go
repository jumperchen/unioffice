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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_SdtText struct {
	// Allow Soft Line Breaks
	MultiLineAttr *sharedTypes.ST_OnOff
}

func NewCT_SdtText() *CT_SdtText {
	ret := &CT_SdtText{}
	return ret
}
func (m *CT_SdtText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MultiLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:multiLine"},
			Value: fmt.Sprintf("%v", *m.MultiLineAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SdtText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "multiLine" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.MultiLineAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SdtText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SdtText) Validate() error {
	return m.ValidateWithPath("CT_SdtText")
}
func (m *CT_SdtText) ValidateWithPath(path string) error {
	if m.MultiLineAttr != nil {
		if err := m.MultiLineAttr.ValidateWithPath(path + "/MultiLineAttr"); err != nil {
			return err
		}
	}
	return nil
}
