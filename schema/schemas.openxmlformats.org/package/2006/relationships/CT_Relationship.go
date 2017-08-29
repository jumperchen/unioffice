// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package relationships

import (
	"encoding/xml"
	"fmt"
)

type CT_Relationship struct {
	TargetModeAttr ST_TargetMode
	TargetAttr     string
	TypeAttr       string
	IdAttr         string
	Content        string
}

func NewCT_Relationship() *CT_Relationship {
	ret := &CT_Relationship{}
	return ret
}
func (m *CT_Relationship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TargetModeAttr != ST_TargetModeUnset {
		attr, err := m.TargetModeAttr.MarshalXMLAttr(xml.Name{Local: "TargetMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Target"},
		Value: fmt.Sprintf("%v", m.TargetAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Type"},
		Value: fmt.Sprintf("%v", m.TypeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Relationship) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "TargetMode" {
			m.TargetModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "Target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = parsed
		}
		if attr.Name.Local == "Type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = parsed
		}
		if attr.Name.Local == "Id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Relationship: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Relationship) Validate() error {
	return m.ValidateWithPath("CT_Relationship")
}
func (m *CT_Relationship) ValidateWithPath(path string) error {
	if err := m.TargetModeAttr.ValidateWithPath(path + "/TargetModeAttr"); err != nil {
		return err
	}
	return nil
}
