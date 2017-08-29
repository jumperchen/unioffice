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

type CT_UndoInfo struct {
	// Index
	IndexAttr uint32
	// Expression
	ExpAttr ST_FormulaExpression
	// Reference 3D
	Ref3DAttr *bool
	// Array Formula
	ArrayAttr *bool
	// Value Needed
	VAttr *bool
	// Defined Name Formula
	NfAttr *bool
	// Cross Sheet Move
	CsAttr *bool
	// Range
	DrAttr string
	// Defined Name
	DnAttr *string
	// Cell Reference
	RAttr *string
	// Sheet Id
	SIdAttr *uint32
}

func NewCT_UndoInfo() *CT_UndoInfo {
	ret := &CT_UndoInfo{}
	return ret
}
func (m *CT_UndoInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "index"},
		Value: fmt.Sprintf("%v", m.IndexAttr)})
	attr, err := m.ExpAttr.MarshalXMLAttr(xml.Name{Local: "exp"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.Ref3DAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref3D"},
			Value: fmt.Sprintf("%v", *m.Ref3DAttr)})
	}
	if m.ArrayAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "array"},
			Value: fmt.Sprintf("%v", *m.ArrayAttr)})
	}
	if m.VAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
			Value: fmt.Sprintf("%v", *m.VAttr)})
	}
	if m.NfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nf"},
			Value: fmt.Sprintf("%v", *m.NfAttr)})
	}
	if m.CsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cs"},
			Value: fmt.Sprintf("%v", *m.CsAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dr"},
		Value: fmt.Sprintf("%v", m.DrAttr)})
	if m.DnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dn"},
			Value: fmt.Sprintf("%v", *m.DnAttr)})
	}
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.SIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sId"},
			Value: fmt.Sprintf("%v", *m.SIdAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_UndoInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "index" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IndexAttr = uint32(parsed)
		}
		if attr.Name.Local == "exp" {
			m.ExpAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ref3D" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Ref3DAttr = &parsed
		}
		if attr.Name.Local == "array" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ArrayAttr = &parsed
		}
		if attr.Name.Local == "v" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VAttr = &parsed
		}
		if attr.Name.Local == "nf" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NfAttr = &parsed
		}
		if attr.Name.Local == "cs" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CsAttr = &parsed
		}
		if attr.Name.Local == "dr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DrAttr = parsed
		}
		if attr.Name.Local == "dn" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DnAttr = &parsed
		}
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = &parsed
		}
		if attr.Name.Local == "sId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.SIdAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_UndoInfo: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_UndoInfo) Validate() error {
	return m.ValidateWithPath("CT_UndoInfo")
}
func (m *CT_UndoInfo) ValidateWithPath(path string) error {
	if m.ExpAttr == ST_FormulaExpressionUnset {
		return fmt.Errorf("%s/ExpAttr is a mandatory field", path)
	}
	if err := m.ExpAttr.ValidateWithPath(path + "/ExpAttr"); err != nil {
		return err
	}
	return nil
}
