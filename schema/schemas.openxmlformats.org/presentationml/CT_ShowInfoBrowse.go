// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ShowInfoBrowse struct {
	// Show Scroll Bar in Window
	ShowScrollbarAttr *bool
}

func NewCT_ShowInfoBrowse() *CT_ShowInfoBrowse {
	ret := &CT_ShowInfoBrowse{}
	return ret
}
func (m *CT_ShowInfoBrowse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ShowScrollbarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showScrollbar"},
			Value: fmt.Sprintf("%v", *m.ShowScrollbarAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ShowInfoBrowse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "showScrollbar" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowScrollbarAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ShowInfoBrowse: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ShowInfoBrowse) Validate() error {
	return m.ValidateWithPath("CT_ShowInfoBrowse")
}
func (m *CT_ShowInfoBrowse) ValidateWithPath(path string) error {
	return nil
}
