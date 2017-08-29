// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type CT_DynamicFilter struct {
	// Dynamic filter type
	TypeAttr ST_DynamicFilterType
	// Value
	ValAttr *float64
	// ISO Value
	ValIsoAttr *time.Time
	// Max Value
	MaxValAttr *float64
	// Max ISO Value
	MaxValIsoAttr *time.Time
}

func NewCT_DynamicFilter() *CT_DynamicFilter {
	ret := &CT_DynamicFilter{}
	return ret
}
func (m *CT_DynamicFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.ValIsoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "valIso"},
			Value: fmt.Sprintf("%v", *m.ValIsoAttr)})
	}
	if m.MaxValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxVal"},
			Value: fmt.Sprintf("%v", *m.MaxValAttr)})
	}
	if m.MaxValIsoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxValIso"},
			Value: fmt.Sprintf("%v", *m.MaxValIsoAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DynamicFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "valIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ValIsoAttr = &parsed
		}
		if attr.Name.Local == "maxVal" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MaxValAttr = &parsed
		}
		if attr.Name.Local == "maxValIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MaxValIsoAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DynamicFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_DynamicFilter) Validate() error {
	return m.ValidateWithPath("CT_DynamicFilter")
}
func (m *CT_DynamicFilter) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_DynamicFilterTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
