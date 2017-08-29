// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_SdtDropDownList struct {
	// Drop-down List Last Saved Value
	LastValueAttr *string
	// Drop-Down List Item
	ListItem []*CT_SdtListItem
}

func NewCT_SdtDropDownList() *CT_SdtDropDownList {
	ret := &CT_SdtDropDownList{}
	return ret
}
func (m *CT_SdtDropDownList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LastValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastValue"},
			Value: fmt.Sprintf("%v", *m.LastValueAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ListItem != nil {
		selistItem := xml.StartElement{Name: xml.Name{Local: "w:listItem"}}
		e.EncodeElement(m.ListItem, selistItem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SdtDropDownList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastValueAttr = &parsed
		}
	}
lCT_SdtDropDownList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "listItem":
				tmp := NewCT_SdtListItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ListItem = append(m.ListItem, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtDropDownList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SdtDropDownList) Validate() error {
	return m.ValidateWithPath("CT_SdtDropDownList")
}
func (m *CT_SdtDropDownList) ValidateWithPath(path string) error {
	for i, v := range m.ListItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ListItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
