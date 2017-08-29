// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_AnimationGraphicalObjectBuildProperties struct {
	BldDgm   *CT_AnimationDgmBuildProperties
	BldChart *CT_AnimationChartBuildProperties
}

func NewCT_AnimationGraphicalObjectBuildProperties() *CT_AnimationGraphicalObjectBuildProperties {
	ret := &CT_AnimationGraphicalObjectBuildProperties{}
	return ret
}
func (m *CT_AnimationGraphicalObjectBuildProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "a:CT_AnimationGraphicalObjectBuildProperties"
	e.EncodeToken(start)
	start.Attr = nil
	if m.BldDgm != nil {
		sebldDgm := xml.StartElement{Name: xml.Name{Local: "a:bldDgm"}}
		e.EncodeElement(m.BldDgm, sebldDgm)
	}
	if m.BldChart != nil {
		sebldChart := xml.StartElement{Name: xml.Name{Local: "a:bldChart"}}
		e.EncodeElement(m.BldChart, sebldChart)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AnimationGraphicalObjectBuildProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AnimationGraphicalObjectBuildProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bldDgm":
				m.BldDgm = NewCT_AnimationDgmBuildProperties()
				if err := d.DecodeElement(m.BldDgm, &el); err != nil {
					return err
				}
			case "bldChart":
				m.BldChart = NewCT_AnimationChartBuildProperties()
				if err := d.DecodeElement(m.BldChart, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AnimationGraphicalObjectBuildProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_AnimationGraphicalObjectBuildProperties) Validate() error {
	return m.ValidateWithPath("CT_AnimationGraphicalObjectBuildProperties")
}
func (m *CT_AnimationGraphicalObjectBuildProperties) ValidateWithPath(path string) error {
	if m.BldDgm != nil {
		if err := m.BldDgm.ValidateWithPath(path + "/BldDgm"); err != nil {
			return err
		}
	}
	if m.BldChart != nil {
		if err := m.BldChart.ValidateWithPath(path + "/BldChart"); err != nil {
			return err
		}
	}
	return nil
}
