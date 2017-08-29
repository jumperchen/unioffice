// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ReflectionEffect struct {
	BlurRadAttr      *int64
	StAAttr          *ST_PositiveFixedPercentage
	StPosAttr        *ST_PositiveFixedPercentage
	EndAAttr         *ST_PositiveFixedPercentage
	EndPosAttr       *ST_PositiveFixedPercentage
	DistAttr         *int64
	DirAttr          *int32
	FadeDirAttr      *int32
	SxAttr           *ST_Percentage
	SyAttr           *ST_Percentage
	KxAttr           *int32
	KyAttr           *int32
	AlgnAttr         ST_RectAlignment
	RotWithShapeAttr *bool
}

func NewCT_ReflectionEffect() *CT_ReflectionEffect {
	ret := &CT_ReflectionEffect{}
	return ret
}
func (m *CT_ReflectionEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BlurRadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blurRad"},
			Value: fmt.Sprintf("%v", *m.BlurRadAttr)})
	}
	if m.StAAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stA"},
			Value: fmt.Sprintf("%v", *m.StAAttr)})
	}
	if m.StPosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stPos"},
			Value: fmt.Sprintf("%v", *m.StPosAttr)})
	}
	if m.EndAAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endA"},
			Value: fmt.Sprintf("%v", *m.EndAAttr)})
	}
	if m.EndPosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endPos"},
			Value: fmt.Sprintf("%v", *m.EndPosAttr)})
	}
	if m.DistAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dist"},
			Value: fmt.Sprintf("%v", *m.DistAttr)})
	}
	if m.DirAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dir"},
			Value: fmt.Sprintf("%v", *m.DirAttr)})
	}
	if m.FadeDirAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fadeDir"},
			Value: fmt.Sprintf("%v", *m.FadeDirAttr)})
	}
	if m.SxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sx"},
			Value: fmt.Sprintf("%v", *m.SxAttr)})
	}
	if m.SyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sy"},
			Value: fmt.Sprintf("%v", *m.SyAttr)})
	}
	if m.KxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "kx"},
			Value: fmt.Sprintf("%v", *m.KxAttr)})
	}
	if m.KyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ky"},
			Value: fmt.Sprintf("%v", *m.KyAttr)})
	}
	if m.AlgnAttr != ST_RectAlignmentUnset {
		attr, err := m.AlgnAttr.MarshalXMLAttr(xml.Name{Local: "algn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RotWithShapeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rotWithShape"},
			Value: fmt.Sprintf("%v", *m.RotWithShapeAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ReflectionEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "blurRad" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.BlurRadAttr = &parsed
		}
		if attr.Name.Local == "stA" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.StAAttr = &parsed
		}
		if attr.Name.Local == "stPos" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.StPosAttr = &parsed
		}
		if attr.Name.Local == "endA" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.EndAAttr = &parsed
		}
		if attr.Name.Local == "endPos" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.EndPosAttr = &parsed
		}
		if attr.Name.Local == "dist" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DistAttr = &parsed
		}
		if attr.Name.Local == "dir" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.DirAttr = &pt
		}
		if attr.Name.Local == "fadeDir" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.FadeDirAttr = &pt
		}
		if attr.Name.Local == "sx" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SxAttr = &parsed
		}
		if attr.Name.Local == "sy" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SyAttr = &parsed
		}
		if attr.Name.Local == "kx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.KxAttr = &pt
		}
		if attr.Name.Local == "ky" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.KyAttr = &pt
		}
		if attr.Name.Local == "algn" {
			m.AlgnAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "rotWithShape" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RotWithShapeAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ReflectionEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ReflectionEffect) Validate() error {
	return m.ValidateWithPath("CT_ReflectionEffect")
}
func (m *CT_ReflectionEffect) ValidateWithPath(path string) error {
	if m.BlurRadAttr != nil {
		if *m.BlurRadAttr < 0 {
			return fmt.Errorf("%s/m.BlurRadAttr must be >= 0 (have %v)", path, *m.BlurRadAttr)
		}
		if *m.BlurRadAttr > 27273042316900 {
			return fmt.Errorf("%s/m.BlurRadAttr must be <= 27273042316900 (have %v)", path, *m.BlurRadAttr)
		}
	}
	if m.StAAttr != nil {
		if err := m.StAAttr.ValidateWithPath(path + "/StAAttr"); err != nil {
			return err
		}
	}
	if m.StPosAttr != nil {
		if err := m.StPosAttr.ValidateWithPath(path + "/StPosAttr"); err != nil {
			return err
		}
	}
	if m.EndAAttr != nil {
		if err := m.EndAAttr.ValidateWithPath(path + "/EndAAttr"); err != nil {
			return err
		}
	}
	if m.EndPosAttr != nil {
		if err := m.EndPosAttr.ValidateWithPath(path + "/EndPosAttr"); err != nil {
			return err
		}
	}
	if m.DistAttr != nil {
		if *m.DistAttr < 0 {
			return fmt.Errorf("%s/m.DistAttr must be >= 0 (have %v)", path, *m.DistAttr)
		}
		if *m.DistAttr > 27273042316900 {
			return fmt.Errorf("%s/m.DistAttr must be <= 27273042316900 (have %v)", path, *m.DistAttr)
		}
	}
	if m.DirAttr != nil {
		if *m.DirAttr < 0 {
			return fmt.Errorf("%s/m.DirAttr must be >= 0 (have %v)", path, *m.DirAttr)
		}
		if *m.DirAttr >= 21600000 {
			return fmt.Errorf("%s/m.DirAttr must be < 21600000 (have %v)", path, *m.DirAttr)
		}
	}
	if m.FadeDirAttr != nil {
		if *m.FadeDirAttr < 0 {
			return fmt.Errorf("%s/m.FadeDirAttr must be >= 0 (have %v)", path, *m.FadeDirAttr)
		}
		if *m.FadeDirAttr >= 21600000 {
			return fmt.Errorf("%s/m.FadeDirAttr must be < 21600000 (have %v)", path, *m.FadeDirAttr)
		}
	}
	if m.SxAttr != nil {
		if err := m.SxAttr.ValidateWithPath(path + "/SxAttr"); err != nil {
			return err
		}
	}
	if m.SyAttr != nil {
		if err := m.SyAttr.ValidateWithPath(path + "/SyAttr"); err != nil {
			return err
		}
	}
	if m.KxAttr != nil {
		if *m.KxAttr <= -5400000 {
			return fmt.Errorf("%s/m.KxAttr must be > -5400000 (have %v)", path, *m.KxAttr)
		}
		if *m.KxAttr >= 5400000 {
			return fmt.Errorf("%s/m.KxAttr must be < 5400000 (have %v)", path, *m.KxAttr)
		}
	}
	if m.KyAttr != nil {
		if *m.KyAttr <= -5400000 {
			return fmt.Errorf("%s/m.KyAttr must be > -5400000 (have %v)", path, *m.KyAttr)
		}
		if *m.KyAttr >= 5400000 {
			return fmt.Errorf("%s/m.KyAttr must be < 5400000 (have %v)", path, *m.KyAttr)
		}
	}
	if err := m.AlgnAttr.ValidateWithPath(path + "/AlgnAttr"); err != nil {
		return err
	}
	return nil
}
