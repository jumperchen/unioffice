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

type CT_TLOleBuildChart struct {
	// Build
	BldAttr ST_TLOleChartBuildType
	// Animate Background
	AnimBgAttr   *bool
	SpidAttr     *uint32
	GrpIdAttr    *uint32
	UiExpandAttr *bool
}

func NewCT_TLOleBuildChart() *CT_TLOleBuildChart {
	ret := &CT_TLOleBuildChart{}
	return ret
}
func (m *CT_TLOleBuildChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BldAttr != ST_TLOleChartBuildTypeUnset {
		attr, err := m.BldAttr.MarshalXMLAttr(xml.Name{Local: "bld"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnimBgAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "animBg"},
			Value: fmt.Sprintf("%v", *m.AnimBgAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.UiExpandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiExpand"},
			Value: fmt.Sprintf("%v", *m.UiExpandAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLOleBuildChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bld" {
			m.BldAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "animBg" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AnimBgAttr = &parsed
		}
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.SpidAttr = &pt
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
		}
		if attr.Name.Local == "uiExpand" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiExpandAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLOleBuildChart: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TLOleBuildChart) Validate() error {
	return m.ValidateWithPath("CT_TLOleBuildChart")
}
func (m *CT_TLOleBuildChart) ValidateWithPath(path string) error {
	if err := m.BldAttr.ValidateWithPath(path + "/BldAttr"); err != nil {
		return err
	}
	return nil
}
