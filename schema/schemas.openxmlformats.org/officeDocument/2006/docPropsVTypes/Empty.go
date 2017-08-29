// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
)

type Empty struct {
	CT_Empty
}

func NewEmpty() *Empty {
	ret := &Empty{}
	ret.CT_Empty = *NewCT_Empty()
	return ret
}
func (m *Empty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_Empty.MarshalXML(e, start)
}
func (m *Empty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Empty = *NewCT_Empty()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Empty: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *Empty) Validate() error {
	return m.ValidateWithPath("Empty")
}
func (m *Empty) ValidateWithPath(path string) error {
	if err := m.CT_Empty.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
