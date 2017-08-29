// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
)

type AG_ConstraintRefAttributes struct {
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func NewAG_ConstraintRefAttributes() *AG_ConstraintRefAttributes {
	ret := &AG_ConstraintRefAttributes{}
	return ret
}
func (m *AG_ConstraintRefAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RefTypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.RefTypeAttr.MarshalXMLAttr(xml.Name{Local: "refType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.RefForAttr.MarshalXMLAttr(xml.Name{Local: "refFor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refForName"},
			Value: fmt.Sprintf("%v", *m.RefForNameAttr)})
	}
	if m.RefPtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.RefPtTypeAttr.MarshalXMLAttr(xml.Name{Local: "refPtType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}
func (m *AG_ConstraintRefAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "refType" {
			m.RefTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "refFor" {
			m.RefForAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "refForName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefForNameAttr = &parsed
		}
		if attr.Name.Local == "refPtType" {
			m.RefPtTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_ConstraintRefAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *AG_ConstraintRefAttributes) Validate() error {
	return m.ValidateWithPath("AG_ConstraintRefAttributes")
}
func (m *AG_ConstraintRefAttributes) ValidateWithPath(path string) error {
	if err := m.RefTypeAttr.ValidateWithPath(path + "/RefTypeAttr"); err != nil {
		return err
	}
	if err := m.RefForAttr.ValidateWithPath(path + "/RefForAttr"); err != nil {
		return err
	}
	if err := m.RefPtTypeAttr.ValidateWithPath(path + "/RefPtTypeAttr"); err != nil {
		return err
	}
	return nil
}
