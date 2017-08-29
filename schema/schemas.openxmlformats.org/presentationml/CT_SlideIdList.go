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
	"log"
)

type CT_SlideIdList struct {
	// Slide ID
	SldId []*CT_SlideIdListEntry
}

func NewCT_SlideIdList() *CT_SlideIdList {
	ret := &CT_SlideIdList{}
	return ret
}
func (m *CT_SlideIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.SldId != nil {
		sesldId := xml.StartElement{Name: xml.Name{Local: "p:sldId"}}
		e.EncodeElement(m.SldId, sesldId)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldId":
				tmp := NewCT_SlideIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldId = append(m.SldId, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideIdList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SlideIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideIdList")
}
func (m *CT_SlideIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
