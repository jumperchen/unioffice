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
	"log"
	"strconv"
)

type CT_DataConsolidate struct {
	// Function Index
	FunctionAttr ST_DataConsolidateFunction
	// Use Starting Column Labels
	StartLabelsAttr *bool
	// Starting Column Labels
	LeftLabelsAttr *bool
	// Labels In Top Row
	TopLabelsAttr *bool
	// Link
	LinkAttr *bool
	// Data Consolidation References
	DataRefs *CT_DataRefs
}

func NewCT_DataConsolidate() *CT_DataConsolidate {
	ret := &CT_DataConsolidate{}
	return ret
}
func (m *CT_DataConsolidate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.FunctionAttr != ST_DataConsolidateFunctionUnset {
		attr, err := m.FunctionAttr.MarshalXMLAttr(xml.Name{Local: "function"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StartLabelsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "startLabels"},
			Value: fmt.Sprintf("%v", *m.StartLabelsAttr)})
	}
	if m.LeftLabelsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "leftLabels"},
			Value: fmt.Sprintf("%v", *m.LeftLabelsAttr)})
	}
	if m.TopLabelsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "topLabels"},
			Value: fmt.Sprintf("%v", *m.TopLabelsAttr)})
	}
	if m.LinkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "link"},
			Value: fmt.Sprintf("%v", *m.LinkAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.DataRefs != nil {
		sedataRefs := xml.StartElement{Name: xml.Name{Local: "x:dataRefs"}}
		e.EncodeElement(m.DataRefs, sedataRefs)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DataConsolidate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "function" {
			m.FunctionAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "startLabels" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StartLabelsAttr = &parsed
		}
		if attr.Name.Local == "leftLabels" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LeftLabelsAttr = &parsed
		}
		if attr.Name.Local == "topLabels" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TopLabelsAttr = &parsed
		}
		if attr.Name.Local == "link" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LinkAttr = &parsed
		}
	}
lCT_DataConsolidate:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dataRefs":
				m.DataRefs = NewCT_DataRefs()
				if err := d.DecodeElement(m.DataRefs, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataConsolidate
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DataConsolidate) Validate() error {
	return m.ValidateWithPath("CT_DataConsolidate")
}
func (m *CT_DataConsolidate) ValidateWithPath(path string) error {
	if err := m.FunctionAttr.ValidateWithPath(path + "/FunctionAttr"); err != nil {
		return err
	}
	if m.DataRefs != nil {
		if err := m.DataRefs.ValidateWithPath(path + "/DataRefs"); err != nil {
			return err
		}
	}
	return nil
}
