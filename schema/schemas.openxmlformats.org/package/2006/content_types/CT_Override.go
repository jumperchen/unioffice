// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"encoding/xml"
	"fmt"
)

type CT_Override struct {
	ContentTypeAttr string
	PartNameAttr    string
}

func NewCT_Override() *CT_Override {
	ret := &CT_Override{}
	return ret
}
func (m *CT_Override) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ContentType"},
		Value: fmt.Sprintf("%v", m.ContentTypeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "PartName"},
		Value: fmt.Sprintf("%v", m.PartNameAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Override) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = parsed
		}
		if attr.Name.Local == "PartName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PartNameAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Override: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Override) Validate() error {
	return m.ValidateWithPath("CT_Override")
}
func (m *CT_Override) ValidateWithPath(path string) error {
	if !ST_ContentTypePatternRe.MatchString(m.ContentTypeAttr) {
		return fmt.Errorf(`%s/m.ContentTypeAttr must match '%s' (have %v)`, path, ST_ContentTypePatternRe, m.ContentTypeAttr)
	}
	return nil
}
