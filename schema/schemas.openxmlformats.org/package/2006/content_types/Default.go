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

type Default struct {
	CT_Default
}

func NewDefault() *Default {
	ret := &Default{}
	ret.CT_Default = *NewCT_Default()
	return ret
}
func (m *Default) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_Default.MarshalXML(e, start)
}
func (m *Default) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Default = *NewCT_Default()
	for _, attr := range start.Attr {
		if attr.Name.Local == "Extension" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ExtensionAttr = parsed
		}
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Default: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *Default) Validate() error {
	return m.ValidateWithPath("Default")
}
func (m *Default) ValidateWithPath(path string) error {
	if err := m.CT_Default.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
