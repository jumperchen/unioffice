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
	"strconv"
)

type CT_BlurEffect struct {
	RadAttr  *int64
	GrowAttr *bool
}

func NewCT_BlurEffect() *CT_BlurEffect {
	ret := &CT_BlurEffect{}
	return ret
}
func (m *CT_BlurEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rad"},
			Value: fmt.Sprintf("%v", *m.RadAttr)})
	}
	if m.GrowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grow"},
			Value: fmt.Sprintf("%v", *m.GrowAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BlurEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rad" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RadAttr = &parsed
		}
		if attr.Name.Local == "grow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GrowAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BlurEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_BlurEffect) Validate() error {
	return m.ValidateWithPath("CT_BlurEffect")
}
func (m *CT_BlurEffect) ValidateWithPath(path string) error {
	if m.RadAttr != nil {
		if *m.RadAttr < 0 {
			return fmt.Errorf("%s/m.RadAttr must be >= 0 (have %v)", path, *m.RadAttr)
		}
		if *m.RadAttr > 27273042316900 {
			return fmt.Errorf("%s/m.RadAttr must be <= 27273042316900 (have %v)", path, *m.RadAttr)
		}
	}
	return nil
}
