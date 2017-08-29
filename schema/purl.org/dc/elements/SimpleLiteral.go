// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import (
	"encoding/xml"
	"fmt"
)

type SimpleLiteral struct {
}

func NewSimpleLiteral() *SimpleLiteral {
	ret := &SimpleLiteral{}
	return ret
}
func (m *SimpleLiteral) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *SimpleLiteral) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing SimpleLiteral: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *SimpleLiteral) Validate() error {
	return m.ValidateWithPath("SimpleLiteral")
}
func (m *SimpleLiteral) ValidateWithPath(path string) error {
	return nil
}
