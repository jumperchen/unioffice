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

type CT_AlphaModulateFixedEffect struct {
	AmtAttr *ST_PositivePercentage
}

func NewCT_AlphaModulateFixedEffect() *CT_AlphaModulateFixedEffect {
	ret := &CT_AlphaModulateFixedEffect{}
	return ret
}
func (m *CT_AlphaModulateFixedEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AmtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "amt"},
			Value: fmt.Sprintf("%v", *m.AmtAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AlphaModulateFixedEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "amt" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.AmtAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AlphaModulateFixedEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AlphaModulateFixedEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaModulateFixedEffect")
}
func (m *CT_AlphaModulateFixedEffect) ValidateWithPath(path string) error {
	if m.AmtAttr != nil {
		if err := m.AmtAttr.ValidateWithPath(path + "/AmtAttr"); err != nil {
			return err
		}
	}
	return nil
}
