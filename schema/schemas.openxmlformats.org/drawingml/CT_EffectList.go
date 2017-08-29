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

type CT_EffectList struct {
	Blur        *CT_BlurEffect
	FillOverlay *CT_FillOverlayEffect
	Glow        *CT_GlowEffect
	InnerShdw   *CT_InnerShadowEffect
	OuterShdw   *CT_OuterShadowEffect
	PrstShdw    *CT_PresetShadowEffect
	Reflection  *CT_ReflectionEffect
	SoftEdge    *CT_SoftEdgesEffect
}

func NewCT_EffectList() *CT_EffectList {
	ret := &CT_EffectList{}
	return ret
}
func (m *CT_EffectList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		e.EncodeElement(m.Blur, seblur)
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		e.EncodeElement(m.FillOverlay, sefillOverlay)
	}
	if m.Glow != nil {
		seglow := xml.StartElement{Name: xml.Name{Local: "a:glow"}}
		e.EncodeElement(m.Glow, seglow)
	}
	if m.InnerShdw != nil {
		seinnerShdw := xml.StartElement{Name: xml.Name{Local: "a:innerShdw"}}
		e.EncodeElement(m.InnerShdw, seinnerShdw)
	}
	if m.OuterShdw != nil {
		seouterShdw := xml.StartElement{Name: xml.Name{Local: "a:outerShdw"}}
		e.EncodeElement(m.OuterShdw, seouterShdw)
	}
	if m.PrstShdw != nil {
		seprstShdw := xml.StartElement{Name: xml.Name{Local: "a:prstShdw"}}
		e.EncodeElement(m.PrstShdw, seprstShdw)
	}
	if m.Reflection != nil {
		sereflection := xml.StartElement{Name: xml.Name{Local: "a:reflection"}}
		e.EncodeElement(m.Reflection, sereflection)
	}
	if m.SoftEdge != nil {
		sesoftEdge := xml.StartElement{Name: xml.Name{Local: "a:softEdge"}}
		e.EncodeElement(m.SoftEdge, sesoftEdge)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_EffectList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EffectList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "blur":
				m.Blur = NewCT_BlurEffect()
				if err := d.DecodeElement(m.Blur, &el); err != nil {
					return err
				}
			case "fillOverlay":
				m.FillOverlay = NewCT_FillOverlayEffect()
				if err := d.DecodeElement(m.FillOverlay, &el); err != nil {
					return err
				}
			case "glow":
				m.Glow = NewCT_GlowEffect()
				if err := d.DecodeElement(m.Glow, &el); err != nil {
					return err
				}
			case "innerShdw":
				m.InnerShdw = NewCT_InnerShadowEffect()
				if err := d.DecodeElement(m.InnerShdw, &el); err != nil {
					return err
				}
			case "outerShdw":
				m.OuterShdw = NewCT_OuterShadowEffect()
				if err := d.DecodeElement(m.OuterShdw, &el); err != nil {
					return err
				}
			case "prstShdw":
				m.PrstShdw = NewCT_PresetShadowEffect()
				if err := d.DecodeElement(m.PrstShdw, &el); err != nil {
					return err
				}
			case "reflection":
				m.Reflection = NewCT_ReflectionEffect()
				if err := d.DecodeElement(m.Reflection, &el); err != nil {
					return err
				}
			case "softEdge":
				m.SoftEdge = NewCT_SoftEdgesEffect()
				if err := d.DecodeElement(m.SoftEdge, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EffectList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_EffectList) Validate() error {
	return m.ValidateWithPath("CT_EffectList")
}
func (m *CT_EffectList) ValidateWithPath(path string) error {
	if m.Blur != nil {
		if err := m.Blur.ValidateWithPath(path + "/Blur"); err != nil {
			return err
		}
	}
	if m.FillOverlay != nil {
		if err := m.FillOverlay.ValidateWithPath(path + "/FillOverlay"); err != nil {
			return err
		}
	}
	if m.Glow != nil {
		if err := m.Glow.ValidateWithPath(path + "/Glow"); err != nil {
			return err
		}
	}
	if m.InnerShdw != nil {
		if err := m.InnerShdw.ValidateWithPath(path + "/InnerShdw"); err != nil {
			return err
		}
	}
	if m.OuterShdw != nil {
		if err := m.OuterShdw.ValidateWithPath(path + "/OuterShdw"); err != nil {
			return err
		}
	}
	if m.PrstShdw != nil {
		if err := m.PrstShdw.ValidateWithPath(path + "/PrstShdw"); err != nil {
			return err
		}
	}
	if m.Reflection != nil {
		if err := m.Reflection.ValidateWithPath(path + "/Reflection"); err != nil {
			return err
		}
	}
	if m.SoftEdge != nil {
		if err := m.SoftEdge.ValidateWithPath(path + "/SoftEdge"); err != nil {
			return err
		}
	}
	return nil
}
