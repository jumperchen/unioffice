// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_Fonts struct {
	// Font Content Type
	HintAttr ST_Hint
	// ASCII Font
	AsciiAttr *string
	// High ANSI Font
	HAnsiAttr *string
	// East Asian Font
	EastAsiaAttr *string
	// Complex Script Font
	CsAttr *string
	// ASCII Theme Font
	AsciiThemeAttr ST_Theme
	// High ANSI Theme Font
	HAnsiThemeAttr ST_Theme
	// East Asian Theme Font
	EastAsiaThemeAttr ST_Theme
	// Complex Script Theme Font
	CsthemeAttr ST_Theme
}

func NewCT_Fonts() *CT_Fonts {
	ret := &CT_Fonts{}
	return ret
}
func (m *CT_Fonts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.HintAttr != ST_HintUnset {
		attr, err := m.HintAttr.MarshalXMLAttr(xml.Name{Local: "w:hint"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AsciiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:ascii"},
			Value: fmt.Sprintf("%v", *m.AsciiAttr)})
	}
	if m.HAnsiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hAnsi"},
			Value: fmt.Sprintf("%v", *m.HAnsiAttr)})
	}
	if m.EastAsiaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:eastAsia"},
			Value: fmt.Sprintf("%v", *m.EastAsiaAttr)})
	}
	if m.CsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cs"},
			Value: fmt.Sprintf("%v", *m.CsAttr)})
	}
	if m.AsciiThemeAttr != ST_ThemeUnset {
		attr, err := m.AsciiThemeAttr.MarshalXMLAttr(xml.Name{Local: "w:asciiTheme"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HAnsiThemeAttr != ST_ThemeUnset {
		attr, err := m.HAnsiThemeAttr.MarshalXMLAttr(xml.Name{Local: "w:hAnsiTheme"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EastAsiaThemeAttr != ST_ThemeUnset {
		attr, err := m.EastAsiaThemeAttr.MarshalXMLAttr(xml.Name{Local: "w:eastAsiaTheme"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CsthemeAttr != ST_ThemeUnset {
		attr, err := m.CsthemeAttr.MarshalXMLAttr(xml.Name{Local: "w:cstheme"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Fonts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "hint" {
			m.HintAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ascii" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AsciiAttr = &parsed
		}
		if attr.Name.Local == "hAnsi" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HAnsiAttr = &parsed
		}
		if attr.Name.Local == "eastAsia" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EastAsiaAttr = &parsed
		}
		if attr.Name.Local == "cs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsAttr = &parsed
		}
		if attr.Name.Local == "asciiTheme" {
			m.AsciiThemeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hAnsiTheme" {
			m.HAnsiThemeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "eastAsiaTheme" {
			m.EastAsiaThemeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cstheme" {
			m.CsthemeAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Fonts: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Fonts) Validate() error {
	return m.ValidateWithPath("CT_Fonts")
}
func (m *CT_Fonts) ValidateWithPath(path string) error {
	if err := m.HintAttr.ValidateWithPath(path + "/HintAttr"); err != nil {
		return err
	}
	if err := m.AsciiThemeAttr.ValidateWithPath(path + "/AsciiThemeAttr"); err != nil {
		return err
	}
	if err := m.HAnsiThemeAttr.ValidateWithPath(path + "/HAnsiThemeAttr"); err != nil {
		return err
	}
	if err := m.EastAsiaThemeAttr.ValidateWithPath(path + "/EastAsiaThemeAttr"); err != nil {
		return err
	}
	if err := m.CsthemeAttr.ValidateWithPath(path + "/CsthemeAttr"); err != nil {
		return err
	}
	return nil
}
