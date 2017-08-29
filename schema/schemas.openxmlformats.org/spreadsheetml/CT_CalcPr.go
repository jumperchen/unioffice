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
)

type CT_CalcPr struct {
	// Calculation Id
	CalcIdAttr *uint32
	// Calculation Mode
	CalcModeAttr ST_CalcMode
	// Full Calculation On Load
	FullCalcOnLoadAttr *bool
	// Reference Mode
	RefModeAttr ST_RefMode
	// Calculation Iteration
	IterateAttr *bool
	// Iteration Count
	IterateCountAttr *uint32
	// Iterative Calculation Delta
	IterateDeltaAttr *float64
	// Full Precision Calculation
	FullPrecisionAttr *bool
	// Calc Completed
	CalcCompletedAttr *bool
	// Calculate On Save
	CalcOnSaveAttr *bool
	// Concurrent Calculations
	ConcurrentCalcAttr *bool
	// Concurrent Thread Manual Count
	ConcurrentManualCountAttr *uint32
	// Force Full Calculation
	ForceFullCalcAttr *bool
}

func NewCT_CalcPr() *CT_CalcPr {
	ret := &CT_CalcPr{}
	return ret
}
func (m *CT_CalcPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CalcIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "calcId"},
			Value: fmt.Sprintf("%v", *m.CalcIdAttr)})
	}
	if m.CalcModeAttr != ST_CalcModeUnset {
		attr, err := m.CalcModeAttr.MarshalXMLAttr(xml.Name{Local: "calcMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FullCalcOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullCalcOnLoad"},
			Value: fmt.Sprintf("%v", *m.FullCalcOnLoadAttr)})
	}
	if m.RefModeAttr != ST_RefModeUnset {
		attr, err := m.RefModeAttr.MarshalXMLAttr(xml.Name{Local: "refMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IterateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iterate"},
			Value: fmt.Sprintf("%v", *m.IterateAttr)})
	}
	if m.IterateCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iterateCount"},
			Value: fmt.Sprintf("%v", *m.IterateCountAttr)})
	}
	if m.IterateDeltaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iterateDelta"},
			Value: fmt.Sprintf("%v", *m.IterateDeltaAttr)})
	}
	if m.FullPrecisionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullPrecision"},
			Value: fmt.Sprintf("%v", *m.FullPrecisionAttr)})
	}
	if m.CalcCompletedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "calcCompleted"},
			Value: fmt.Sprintf("%v", *m.CalcCompletedAttr)})
	}
	if m.CalcOnSaveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "calcOnSave"},
			Value: fmt.Sprintf("%v", *m.CalcOnSaveAttr)})
	}
	if m.ConcurrentCalcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "concurrentCalc"},
			Value: fmt.Sprintf("%v", *m.ConcurrentCalcAttr)})
	}
	if m.ConcurrentManualCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "concurrentManualCount"},
			Value: fmt.Sprintf("%v", *m.ConcurrentManualCountAttr)})
	}
	if m.ForceFullCalcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "forceFullCalc"},
			Value: fmt.Sprintf("%v", *m.ForceFullCalcAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CalcPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "calcId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CalcIdAttr = &pt
		}
		if attr.Name.Local == "calcMode" {
			m.CalcModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "fullCalcOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullCalcOnLoadAttr = &parsed
		}
		if attr.Name.Local == "refMode" {
			m.RefModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "iterate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IterateAttr = &parsed
		}
		if attr.Name.Local == "iterateCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.IterateCountAttr = &pt
		}
		if attr.Name.Local == "iterateDelta" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.IterateDeltaAttr = &parsed
		}
		if attr.Name.Local == "fullPrecision" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullPrecisionAttr = &parsed
		}
		if attr.Name.Local == "calcCompleted" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CalcCompletedAttr = &parsed
		}
		if attr.Name.Local == "calcOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CalcOnSaveAttr = &parsed
		}
		if attr.Name.Local == "concurrentCalc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ConcurrentCalcAttr = &parsed
		}
		if attr.Name.Local == "concurrentManualCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.ConcurrentManualCountAttr = &pt
		}
		if attr.Name.Local == "forceFullCalc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ForceFullCalcAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CalcPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_CalcPr) Validate() error {
	return m.ValidateWithPath("CT_CalcPr")
}
func (m *CT_CalcPr) ValidateWithPath(path string) error {
	if err := m.CalcModeAttr.ValidateWithPath(path + "/CalcModeAttr"); err != nil {
		return err
	}
	if err := m.RefModeAttr.ValidateWithPath(path + "/RefModeAttr"); err != nil {
		return err
	}
	return nil
}
