// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_NumFmt struct {
	FormatCodeAttr   string
	SourceLinkedAttr *bool
}

func NewCT_NumFmt() *CT_NumFmt {
	ret := &CT_NumFmt{}
	return ret
}
func (m *CT_NumFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCode"},
		Value: fmt.Sprintf("%v", m.FormatCodeAttr)})
	if m.SourceLinkedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceLinked"},
			Value: fmt.Sprintf("%v", *m.SourceLinkedAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "formatCode" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatCodeAttr = parsed
		}
		if attr.Name.Local == "sourceLinked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SourceLinkedAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NumFmt: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_NumFmt) Validate() error {
	return m.ValidateWithPath("CT_NumFmt")
}
func (m *CT_NumFmt) ValidateWithPath(path string) error {
	return nil
}
