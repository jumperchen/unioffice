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

type CT_Xf struct {
	// Number Format Id
	NumFmtIdAttr *uint32
	// Font Id
	FontIdAttr *uint32
	// Fill Id
	FillIdAttr *uint32
	// Border Id
	BorderIdAttr *uint32
	// Format Id
	XfIdAttr *uint32
	// Quote Prefix
	QuotePrefixAttr *bool
	// Pivot Button
	PivotButtonAttr *bool
	// Apply Number Format
	ApplyNumberFormatAttr *bool
	// Apply Font
	ApplyFontAttr *bool
	// Apply Fill
	ApplyFillAttr *bool
	// Apply Border
	ApplyBorderAttr *bool
	// Apply Alignment
	ApplyAlignmentAttr *bool
	// Apply Protection
	ApplyProtectionAttr *bool
	// Alignment
	Alignment *CT_CellAlignment
	// Protection
	Protection *CT_CellProtection
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Xf() *CT_Xf {
	ret := &CT_Xf{}
	return ret
}
func (m *CT_Xf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	if m.FontIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fontId"},
			Value: fmt.Sprintf("%v", *m.FontIdAttr)})
	}
	if m.FillIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillId"},
			Value: fmt.Sprintf("%v", *m.FillIdAttr)})
	}
	if m.BorderIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "borderId"},
			Value: fmt.Sprintf("%v", *m.BorderIdAttr)})
	}
	if m.XfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xfId"},
			Value: fmt.Sprintf("%v", *m.XfIdAttr)})
	}
	if m.QuotePrefixAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "quotePrefix"},
			Value: fmt.Sprintf("%v", *m.QuotePrefixAttr)})
	}
	if m.PivotButtonAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivotButton"},
			Value: fmt.Sprintf("%v", *m.PivotButtonAttr)})
	}
	if m.ApplyNumberFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyNumberFormat"},
			Value: fmt.Sprintf("%v", *m.ApplyNumberFormatAttr)})
	}
	if m.ApplyFontAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFont"},
			Value: fmt.Sprintf("%v", *m.ApplyFontAttr)})
	}
	if m.ApplyFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFill"},
			Value: fmt.Sprintf("%v", *m.ApplyFillAttr)})
	}
	if m.ApplyBorderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyBorder"},
			Value: fmt.Sprintf("%v", *m.ApplyBorderAttr)})
	}
	if m.ApplyAlignmentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyAlignment"},
			Value: fmt.Sprintf("%v", *m.ApplyAlignmentAttr)})
	}
	if m.ApplyProtectionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyProtection"},
			Value: fmt.Sprintf("%v", *m.ApplyProtectionAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Alignment != nil {
		sealignment := xml.StartElement{Name: xml.Name{Local: "x:alignment"}}
		e.EncodeElement(m.Alignment, sealignment)
	}
	if m.Protection != nil {
		seprotection := xml.StartElement{Name: xml.Name{Local: "x:protection"}}
		e.EncodeElement(m.Protection, seprotection)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Xf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
		}
		if attr.Name.Local == "fontId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.FontIdAttr = &pt
		}
		if attr.Name.Local == "fillId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.FillIdAttr = &pt
		}
		if attr.Name.Local == "borderId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.BorderIdAttr = &pt
		}
		if attr.Name.Local == "xfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.XfIdAttr = &pt
		}
		if attr.Name.Local == "quotePrefix" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.QuotePrefixAttr = &parsed
		}
		if attr.Name.Local == "pivotButton" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PivotButtonAttr = &parsed
		}
		if attr.Name.Local == "applyNumberFormat" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatAttr = &parsed
		}
		if attr.Name.Local == "applyFont" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontAttr = &parsed
		}
		if attr.Name.Local == "applyFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFillAttr = &parsed
		}
		if attr.Name.Local == "applyBorder" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderAttr = &parsed
		}
		if attr.Name.Local == "applyAlignment" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentAttr = &parsed
		}
		if attr.Name.Local == "applyProtection" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyProtectionAttr = &parsed
		}
	}
lCT_Xf:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "alignment":
				m.Alignment = NewCT_CellAlignment()
				if err := d.DecodeElement(m.Alignment, &el); err != nil {
					return err
				}
			case "protection":
				m.Protection = NewCT_CellProtection()
				if err := d.DecodeElement(m.Protection, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Xf
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Xf) Validate() error {
	return m.ValidateWithPath("CT_Xf")
}
func (m *CT_Xf) ValidateWithPath(path string) error {
	if m.Alignment != nil {
		if err := m.Alignment.ValidateWithPath(path + "/Alignment"); err != nil {
			return err
		}
	}
	if m.Protection != nil {
		if err := m.Protection.ValidateWithPath(path + "/Protection"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
