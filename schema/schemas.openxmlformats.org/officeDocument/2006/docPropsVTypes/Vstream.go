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

type Vstream struct {
	CT_Vstream
}

func NewVstream() *Vstream {
	ret := &Vstream{}
	ret.CT_Vstream = *NewCT_Vstream()
	return ret
}
func (m *Vstream) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	return m.CT_Vstream.MarshalXML(e, start)
}
func (m *Vstream) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Vstream = *NewCT_Vstream()
	for _, attr := range start.Attr {
		if attr.Name.Local == "version" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VersionAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Vstream: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *Vstream) Validate() error {
	return m.ValidateWithPath("Vstream")
}
func (m *Vstream) ValidateWithPath(path string) error {
	if err := m.CT_Vstream.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
