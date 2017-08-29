// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
)

type RelIds struct {
	CT_RelIds
}

func NewRelIds() *RelIds {
	ret := &RelIds{}
	ret.CT_RelIds = *NewCT_RelIds()
	return ret
}
func (m *RelIds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "relIds"
	return m.CT_RelIds.MarshalXML(e, start)
}
func (m *RelIds) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_RelIds = *NewCT_RelIds()
	for _, attr := range start.Attr {
		if attr.Name.Local == "dm" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DmAttr = parsed
		}
		if attr.Name.Local == "lo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoAttr = parsed
		}
		if attr.Name.Local == "qs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsAttr = parsed
		}
		if attr.Name.Local == "cs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing RelIds: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *RelIds) Validate() error {
	return m.ValidateWithPath("RelIds")
}
func (m *RelIds) ValidateWithPath(path string) error {
	if err := m.CT_RelIds.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
