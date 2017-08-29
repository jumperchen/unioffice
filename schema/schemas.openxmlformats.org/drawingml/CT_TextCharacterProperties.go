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
	"log"
	"strconv"
)

type CT_TextCharacterProperties struct {
	KumimojiAttr   *bool
	LangAttr       *string
	AltLangAttr    *string
	SzAttr         *int32
	BAttr          *bool
	IAttr          *bool
	UAttr          ST_TextUnderlineType
	StrikeAttr     ST_TextStrikeType
	KernAttr       *int32
	CapAttr        ST_TextCapsType
	SpcAttr        *ST_TextPoint
	NormalizeHAttr *bool
	BaselineAttr   *ST_Percentage
	NoProofAttr    *bool
	DirtyAttr      *bool
	ErrAttr        *bool
	SmtCleanAttr   *bool
	SmtIdAttr      *uint32
	BmkAttr        *string
	Ln             *CT_LineProperties
	NoFill         *CT_NoFillProperties
	SolidFill      *CT_SolidColorFillProperties
	GradFill       *CT_GradientFillProperties
	BlipFill       *CT_BlipFillProperties
	PattFill       *CT_PatternFillProperties
	GrpFill        *CT_GroupFillProperties
	EffectLst      *CT_EffectList
	EffectDag      *CT_EffectContainer
	Highlight      *CT_Color
	ULnTx          *CT_TextUnderlineLineFollowText
	ULn            *CT_LineProperties
	UFillTx        *CT_TextUnderlineFillFollowText
	UFill          *CT_TextUnderlineFillGroupWrapper
	Latin          *CT_TextFont
	Ea             *CT_TextFont
	Cs             *CT_TextFont
	Sym            *CT_TextFont
	HlinkClick     *CT_Hyperlink
	HlinkMouseOver *CT_Hyperlink
	Rtl            *CT_Boolean
	ExtLst         *CT_OfficeArtExtensionList
}

func NewCT_TextCharacterProperties() *CT_TextCharacterProperties {
	ret := &CT_TextCharacterProperties{}
	return ret
}
func (m *CT_TextCharacterProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.KumimojiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "kumimoji"},
			Value: fmt.Sprintf("%v", *m.KumimojiAttr)})
	}
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	if m.AltLangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "altLang"},
			Value: fmt.Sprintf("%v", *m.AltLangAttr)})
	}
	if m.SzAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sz"},
			Value: fmt.Sprintf("%v", *m.SzAttr)})
	}
	if m.BAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
			Value: fmt.Sprintf("%v", *m.BAttr)})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%v", *m.IAttr)})
	}
	if m.UAttr != ST_TextUnderlineTypeUnset {
		attr, err := m.UAttr.MarshalXMLAttr(xml.Name{Local: "u"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrikeAttr != ST_TextStrikeTypeUnset {
		attr, err := m.StrikeAttr.MarshalXMLAttr(xml.Name{Local: "strike"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.KernAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "kern"},
			Value: fmt.Sprintf("%v", *m.KernAttr)})
	}
	if m.CapAttr != ST_TextCapsTypeUnset {
		attr, err := m.CapAttr.MarshalXMLAttr(xml.Name{Local: "cap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SpcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spc"},
			Value: fmt.Sprintf("%v", *m.SpcAttr)})
	}
	if m.NormalizeHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "normalizeH"},
			Value: fmt.Sprintf("%v", *m.NormalizeHAttr)})
	}
	if m.BaselineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "baseline"},
			Value: fmt.Sprintf("%v", *m.BaselineAttr)})
	}
	if m.NoProofAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noProof"},
			Value: fmt.Sprintf("%v", *m.NoProofAttr)})
	}
	if m.DirtyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dirty"},
			Value: fmt.Sprintf("%v", *m.DirtyAttr)})
	}
	if m.ErrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "err"},
			Value: fmt.Sprintf("%v", *m.ErrAttr)})
	}
	if m.SmtCleanAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "smtClean"},
			Value: fmt.Sprintf("%v", *m.SmtCleanAttr)})
	}
	if m.SmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "smtId"},
			Value: fmt.Sprintf("%v", *m.SmtIdAttr)})
	}
	if m.BmkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bmk"},
			Value: fmt.Sprintf("%v", *m.BmkAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Ln != nil {
		seln := xml.StartElement{Name: xml.Name{Local: "a:ln"}}
		e.EncodeElement(m.Ln, seln)
	}
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.BlipFill != nil {
		seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
		e.EncodeElement(m.BlipFill, seblipFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.GrpFill != nil {
		segrpFill := xml.StartElement{Name: xml.Name{Local: "a:grpFill"}}
		e.EncodeElement(m.GrpFill, segrpFill)
	}
	if m.EffectLst != nil {
		seeffectLst := xml.StartElement{Name: xml.Name{Local: "a:effectLst"}}
		e.EncodeElement(m.EffectLst, seeffectLst)
	}
	if m.EffectDag != nil {
		seeffectDag := xml.StartElement{Name: xml.Name{Local: "a:effectDag"}}
		e.EncodeElement(m.EffectDag, seeffectDag)
	}
	if m.Highlight != nil {
		sehighlight := xml.StartElement{Name: xml.Name{Local: "a:highlight"}}
		e.EncodeElement(m.Highlight, sehighlight)
	}
	if m.ULnTx != nil {
		seuLnTx := xml.StartElement{Name: xml.Name{Local: "a:uLnTx"}}
		e.EncodeElement(m.ULnTx, seuLnTx)
	}
	if m.ULn != nil {
		seuLn := xml.StartElement{Name: xml.Name{Local: "a:uLn"}}
		e.EncodeElement(m.ULn, seuLn)
	}
	if m.UFillTx != nil {
		seuFillTx := xml.StartElement{Name: xml.Name{Local: "a:uFillTx"}}
		e.EncodeElement(m.UFillTx, seuFillTx)
	}
	if m.UFill != nil {
		seuFill := xml.StartElement{Name: xml.Name{Local: "a:uFill"}}
		e.EncodeElement(m.UFill, seuFill)
	}
	if m.Latin != nil {
		selatin := xml.StartElement{Name: xml.Name{Local: "a:latin"}}
		e.EncodeElement(m.Latin, selatin)
	}
	if m.Ea != nil {
		seea := xml.StartElement{Name: xml.Name{Local: "a:ea"}}
		e.EncodeElement(m.Ea, seea)
	}
	if m.Cs != nil {
		secs := xml.StartElement{Name: xml.Name{Local: "a:cs"}}
		e.EncodeElement(m.Cs, secs)
	}
	if m.Sym != nil {
		sesym := xml.StartElement{Name: xml.Name{Local: "a:sym"}}
		e.EncodeElement(m.Sym, sesym)
	}
	if m.HlinkClick != nil {
		sehlinkClick := xml.StartElement{Name: xml.Name{Local: "a:hlinkClick"}}
		e.EncodeElement(m.HlinkClick, sehlinkClick)
	}
	if m.HlinkMouseOver != nil {
		sehlinkMouseOver := xml.StartElement{Name: xml.Name{Local: "a:hlinkMouseOver"}}
		e.EncodeElement(m.HlinkMouseOver, sehlinkMouseOver)
	}
	if m.Rtl != nil {
		sertl := xml.StartElement{Name: xml.Name{Local: "a:rtl"}}
		e.EncodeElement(m.Rtl, sertl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextCharacterProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "kumimoji" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.KumimojiAttr = &parsed
		}
		if attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
		}
		if attr.Name.Local == "altLang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltLangAttr = &parsed
		}
		if attr.Name.Local == "sz" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.SzAttr = &pt
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = &parsed
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
		}
		if attr.Name.Local == "u" {
			m.UAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "strike" {
			m.StrikeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "kern" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.KernAttr = &pt
		}
		if attr.Name.Local == "cap" {
			m.CapAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "spc" {
			parsed, err := ParseUnionST_TextPoint(attr.Value)
			if err != nil {
				return err
			}
			m.SpcAttr = &parsed
		}
		if attr.Name.Local == "normalizeH" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NormalizeHAttr = &parsed
		}
		if attr.Name.Local == "baseline" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.BaselineAttr = &parsed
		}
		if attr.Name.Local == "noProof" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoProofAttr = &parsed
		}
		if attr.Name.Local == "dirty" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DirtyAttr = &parsed
		}
		if attr.Name.Local == "err" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ErrAttr = &parsed
		}
		if attr.Name.Local == "smtClean" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SmtCleanAttr = &parsed
		}
		if attr.Name.Local == "smtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.SmtIdAttr = &pt
		}
		if attr.Name.Local == "bmk" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BmkAttr = &parsed
		}
	}
lCT_TextCharacterProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ln":
				m.Ln = NewCT_LineProperties()
				if err := d.DecodeElement(m.Ln, &el); err != nil {
					return err
				}
			case "noFill":
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case "solidFill":
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case "gradFill":
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case "blipFill":
				m.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case "pattFill":
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case "grpFill":
				m.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(m.GrpFill, &el); err != nil {
					return err
				}
			case "effectLst":
				m.EffectLst = NewCT_EffectList()
				if err := d.DecodeElement(m.EffectLst, &el); err != nil {
					return err
				}
			case "effectDag":
				m.EffectDag = NewCT_EffectContainer()
				if err := d.DecodeElement(m.EffectDag, &el); err != nil {
					return err
				}
			case "highlight":
				m.Highlight = NewCT_Color()
				if err := d.DecodeElement(m.Highlight, &el); err != nil {
					return err
				}
			case "uLnTx":
				m.ULnTx = NewCT_TextUnderlineLineFollowText()
				if err := d.DecodeElement(m.ULnTx, &el); err != nil {
					return err
				}
			case "uLn":
				m.ULn = NewCT_LineProperties()
				if err := d.DecodeElement(m.ULn, &el); err != nil {
					return err
				}
			case "uFillTx":
				m.UFillTx = NewCT_TextUnderlineFillFollowText()
				if err := d.DecodeElement(m.UFillTx, &el); err != nil {
					return err
				}
			case "uFill":
				m.UFill = NewCT_TextUnderlineFillGroupWrapper()
				if err := d.DecodeElement(m.UFill, &el); err != nil {
					return err
				}
			case "latin":
				m.Latin = NewCT_TextFont()
				if err := d.DecodeElement(m.Latin, &el); err != nil {
					return err
				}
			case "ea":
				m.Ea = NewCT_TextFont()
				if err := d.DecodeElement(m.Ea, &el); err != nil {
					return err
				}
			case "cs":
				m.Cs = NewCT_TextFont()
				if err := d.DecodeElement(m.Cs, &el); err != nil {
					return err
				}
			case "sym":
				m.Sym = NewCT_TextFont()
				if err := d.DecodeElement(m.Sym, &el); err != nil {
					return err
				}
			case "hlinkClick":
				m.HlinkClick = NewCT_Hyperlink()
				if err := d.DecodeElement(m.HlinkClick, &el); err != nil {
					return err
				}
			case "hlinkMouseOver":
				m.HlinkMouseOver = NewCT_Hyperlink()
				if err := d.DecodeElement(m.HlinkMouseOver, &el); err != nil {
					return err
				}
			case "rtl":
				m.Rtl = NewCT_Boolean()
				if err := d.DecodeElement(m.Rtl, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_TextCharacterProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TextCharacterProperties) Validate() error {
	return m.ValidateWithPath("CT_TextCharacterProperties")
}
func (m *CT_TextCharacterProperties) ValidateWithPath(path string) error {
	if m.SzAttr != nil {
		if *m.SzAttr < 100 {
			return fmt.Errorf("%s/m.SzAttr must be >= 100 (have %v)", path, *m.SzAttr)
		}
		if *m.SzAttr > 400000 {
			return fmt.Errorf("%s/m.SzAttr must be <= 400000 (have %v)", path, *m.SzAttr)
		}
	}
	if err := m.UAttr.ValidateWithPath(path + "/UAttr"); err != nil {
		return err
	}
	if err := m.StrikeAttr.ValidateWithPath(path + "/StrikeAttr"); err != nil {
		return err
	}
	if m.KernAttr != nil {
		if *m.KernAttr < 0 {
			return fmt.Errorf("%s/m.KernAttr must be >= 0 (have %v)", path, *m.KernAttr)
		}
		if *m.KernAttr > 400000 {
			return fmt.Errorf("%s/m.KernAttr must be <= 400000 (have %v)", path, *m.KernAttr)
		}
	}
	if err := m.CapAttr.ValidateWithPath(path + "/CapAttr"); err != nil {
		return err
	}
	if m.SpcAttr != nil {
		if err := m.SpcAttr.ValidateWithPath(path + "/SpcAttr"); err != nil {
			return err
		}
	}
	if m.BaselineAttr != nil {
		if err := m.BaselineAttr.ValidateWithPath(path + "/BaselineAttr"); err != nil {
			return err
		}
	}
	if m.Ln != nil {
		if err := m.Ln.ValidateWithPath(path + "/Ln"); err != nil {
			return err
		}
	}
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.BlipFill != nil {
		if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.GrpFill != nil {
		if err := m.GrpFill.ValidateWithPath(path + "/GrpFill"); err != nil {
			return err
		}
	}
	if m.EffectLst != nil {
		if err := m.EffectLst.ValidateWithPath(path + "/EffectLst"); err != nil {
			return err
		}
	}
	if m.EffectDag != nil {
		if err := m.EffectDag.ValidateWithPath(path + "/EffectDag"); err != nil {
			return err
		}
	}
	if m.Highlight != nil {
		if err := m.Highlight.ValidateWithPath(path + "/Highlight"); err != nil {
			return err
		}
	}
	if m.ULnTx != nil {
		if err := m.ULnTx.ValidateWithPath(path + "/ULnTx"); err != nil {
			return err
		}
	}
	if m.ULn != nil {
		if err := m.ULn.ValidateWithPath(path + "/ULn"); err != nil {
			return err
		}
	}
	if m.UFillTx != nil {
		if err := m.UFillTx.ValidateWithPath(path + "/UFillTx"); err != nil {
			return err
		}
	}
	if m.UFill != nil {
		if err := m.UFill.ValidateWithPath(path + "/UFill"); err != nil {
			return err
		}
	}
	if m.Latin != nil {
		if err := m.Latin.ValidateWithPath(path + "/Latin"); err != nil {
			return err
		}
	}
	if m.Ea != nil {
		if err := m.Ea.ValidateWithPath(path + "/Ea"); err != nil {
			return err
		}
	}
	if m.Cs != nil {
		if err := m.Cs.ValidateWithPath(path + "/Cs"); err != nil {
			return err
		}
	}
	if m.Sym != nil {
		if err := m.Sym.ValidateWithPath(path + "/Sym"); err != nil {
			return err
		}
	}
	if m.HlinkClick != nil {
		if err := m.HlinkClick.ValidateWithPath(path + "/HlinkClick"); err != nil {
			return err
		}
	}
	if m.HlinkMouseOver != nil {
		if err := m.HlinkMouseOver.ValidateWithPath(path + "/HlinkMouseOver"); err != nil {
			return err
		}
	}
	if m.Rtl != nil {
		if err := m.Rtl.ValidateWithPath(path + "/Rtl"); err != nil {
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
