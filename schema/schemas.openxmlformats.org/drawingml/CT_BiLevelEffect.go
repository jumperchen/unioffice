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

type CT_BiLevelEffect struct {
	ThreshAttr ST_PositiveFixedPercentage
}

func NewCT_BiLevelEffect() *CT_BiLevelEffect {
	ret := &CT_BiLevelEffect{}
	return ret
}
func (m *CT_BiLevelEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thresh"},
		Value: fmt.Sprintf("%v", m.ThreshAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BiLevelEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "thresh" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ThreshAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BiLevelEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_BiLevelEffect) Validate() error {
	return m.ValidateWithPath("CT_BiLevelEffect")
}
func (m *CT_BiLevelEffect) ValidateWithPath(path string) error {
	if err := m.ThreshAttr.ValidateWithPath(path + "/ThreshAttr"); err != nil {
		return err
	}
	return nil
}
