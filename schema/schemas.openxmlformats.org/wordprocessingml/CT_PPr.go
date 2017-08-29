// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_PPr struct {
	// Referenced Paragraph Style
	PStyle *CT_String
	// Keep Paragraph With Next Paragraph
	KeepNext *CT_OnOff
	// Keep All Lines On One Page
	KeepLines *CT_OnOff
	// Start Paragraph on Next Page
	PageBreakBefore *CT_OnOff
	// Text Frame Properties
	FramePr *CT_FramePr
	// Allow First/Last Line to Display on a Separate Page
	WidowControl *CT_OnOff
	// Numbering Definition Instance Reference
	NumPr *CT_NumPr
	// Suppress Line Numbers for Paragraph
	SuppressLineNumbers *CT_OnOff
	// Paragraph Borders
	PBdr *CT_PBdr
	// Paragraph Shading
	Shd *CT_Shd
	// Set of Custom Tab Stops
	Tabs *CT_Tabs
	// Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *CT_OnOff
	// Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *CT_OnOff
	// Allow Line Breaking At Character Level
	WordWrap *CT_OnOff
	// Allow Punctuation to Extend Past Text Extents
	OverflowPunct *CT_OnOff
	// Compress Punctuation at Start of a Line
	TopLinePunct *CT_OnOff
	// Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *CT_OnOff
	// Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *CT_OnOff
	// Right to Left Paragraph Layout
	Bidi *CT_OnOff
	// Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *CT_OnOff
	// Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *CT_OnOff
	// Spacing Between Lines and Above/Below Paragraph
	Spacing *CT_Spacing
	// Paragraph Indentation
	Ind *CT_Ind
	// Ignore Spacing Above and Below When Using Identical Styles
	ContextualSpacing *CT_OnOff
	// Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *CT_OnOff
	// Prevent Text Frames From Overlapping
	SuppressOverlap *CT_OnOff
	// Paragraph Alignment
	Jc *CT_Jc
	// Paragraph Text Flow Direction
	TextDirection *CT_TextDirection
	// Vertical Character Alignment on Line
	TextAlignment *CT_TextAlignment
	// Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *CT_TextboxTightWrap
	// Associated Outline Level
	OutlineLvl *CT_DecimalNumber
	// Associated HTML div ID
	DivId *CT_DecimalNumber
	// Paragraph Conditional Formatting
	CnfStyle  *CT_Cnf
	RPr       *CT_ParaRPr
	SectPr    *CT_SectPr
	PPrChange *CT_PPrChange
}

func NewCT_PPr() *CT_PPr {
	ret := &CT_PPr{}
	return ret
}
func (m *CT_PPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.PStyle != nil {
		sepStyle := xml.StartElement{Name: xml.Name{Local: "w:pStyle"}}
		e.EncodeElement(m.PStyle, sepStyle)
	}
	if m.KeepNext != nil {
		sekeepNext := xml.StartElement{Name: xml.Name{Local: "w:keepNext"}}
		e.EncodeElement(m.KeepNext, sekeepNext)
	}
	if m.KeepLines != nil {
		sekeepLines := xml.StartElement{Name: xml.Name{Local: "w:keepLines"}}
		e.EncodeElement(m.KeepLines, sekeepLines)
	}
	if m.PageBreakBefore != nil {
		sepageBreakBefore := xml.StartElement{Name: xml.Name{Local: "w:pageBreakBefore"}}
		e.EncodeElement(m.PageBreakBefore, sepageBreakBefore)
	}
	if m.FramePr != nil {
		seframePr := xml.StartElement{Name: xml.Name{Local: "w:framePr"}}
		e.EncodeElement(m.FramePr, seframePr)
	}
	if m.WidowControl != nil {
		sewidowControl := xml.StartElement{Name: xml.Name{Local: "w:widowControl"}}
		e.EncodeElement(m.WidowControl, sewidowControl)
	}
	if m.NumPr != nil {
		senumPr := xml.StartElement{Name: xml.Name{Local: "w:numPr"}}
		e.EncodeElement(m.NumPr, senumPr)
	}
	if m.SuppressLineNumbers != nil {
		sesuppressLineNumbers := xml.StartElement{Name: xml.Name{Local: "w:suppressLineNumbers"}}
		e.EncodeElement(m.SuppressLineNumbers, sesuppressLineNumbers)
	}
	if m.PBdr != nil {
		sepBdr := xml.StartElement{Name: xml.Name{Local: "w:pBdr"}}
		e.EncodeElement(m.PBdr, sepBdr)
	}
	if m.Shd != nil {
		seshd := xml.StartElement{Name: xml.Name{Local: "w:shd"}}
		e.EncodeElement(m.Shd, seshd)
	}
	if m.Tabs != nil {
		setabs := xml.StartElement{Name: xml.Name{Local: "w:tabs"}}
		e.EncodeElement(m.Tabs, setabs)
	}
	if m.SuppressAutoHyphens != nil {
		sesuppressAutoHyphens := xml.StartElement{Name: xml.Name{Local: "w:suppressAutoHyphens"}}
		e.EncodeElement(m.SuppressAutoHyphens, sesuppressAutoHyphens)
	}
	if m.Kinsoku != nil {
		sekinsoku := xml.StartElement{Name: xml.Name{Local: "w:kinsoku"}}
		e.EncodeElement(m.Kinsoku, sekinsoku)
	}
	if m.WordWrap != nil {
		sewordWrap := xml.StartElement{Name: xml.Name{Local: "w:wordWrap"}}
		e.EncodeElement(m.WordWrap, sewordWrap)
	}
	if m.OverflowPunct != nil {
		seoverflowPunct := xml.StartElement{Name: xml.Name{Local: "w:overflowPunct"}}
		e.EncodeElement(m.OverflowPunct, seoverflowPunct)
	}
	if m.TopLinePunct != nil {
		setopLinePunct := xml.StartElement{Name: xml.Name{Local: "w:topLinePunct"}}
		e.EncodeElement(m.TopLinePunct, setopLinePunct)
	}
	if m.AutoSpaceDE != nil {
		seautoSpaceDE := xml.StartElement{Name: xml.Name{Local: "w:autoSpaceDE"}}
		e.EncodeElement(m.AutoSpaceDE, seautoSpaceDE)
	}
	if m.AutoSpaceDN != nil {
		seautoSpaceDN := xml.StartElement{Name: xml.Name{Local: "w:autoSpaceDN"}}
		e.EncodeElement(m.AutoSpaceDN, seautoSpaceDN)
	}
	if m.Bidi != nil {
		sebidi := xml.StartElement{Name: xml.Name{Local: "w:bidi"}}
		e.EncodeElement(m.Bidi, sebidi)
	}
	if m.AdjustRightInd != nil {
		seadjustRightInd := xml.StartElement{Name: xml.Name{Local: "w:adjustRightInd"}}
		e.EncodeElement(m.AdjustRightInd, seadjustRightInd)
	}
	if m.SnapToGrid != nil {
		sesnapToGrid := xml.StartElement{Name: xml.Name{Local: "w:snapToGrid"}}
		e.EncodeElement(m.SnapToGrid, sesnapToGrid)
	}
	if m.Spacing != nil {
		sespacing := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}
		e.EncodeElement(m.Spacing, sespacing)
	}
	if m.Ind != nil {
		seind := xml.StartElement{Name: xml.Name{Local: "w:ind"}}
		e.EncodeElement(m.Ind, seind)
	}
	if m.ContextualSpacing != nil {
		secontextualSpacing := xml.StartElement{Name: xml.Name{Local: "w:contextualSpacing"}}
		e.EncodeElement(m.ContextualSpacing, secontextualSpacing)
	}
	if m.MirrorIndents != nil {
		semirrorIndents := xml.StartElement{Name: xml.Name{Local: "w:mirrorIndents"}}
		e.EncodeElement(m.MirrorIndents, semirrorIndents)
	}
	if m.SuppressOverlap != nil {
		sesuppressOverlap := xml.StartElement{Name: xml.Name{Local: "w:suppressOverlap"}}
		e.EncodeElement(m.SuppressOverlap, sesuppressOverlap)
	}
	if m.Jc != nil {
		sejc := xml.StartElement{Name: xml.Name{Local: "w:jc"}}
		e.EncodeElement(m.Jc, sejc)
	}
	if m.TextDirection != nil {
		setextDirection := xml.StartElement{Name: xml.Name{Local: "w:textDirection"}}
		e.EncodeElement(m.TextDirection, setextDirection)
	}
	if m.TextAlignment != nil {
		setextAlignment := xml.StartElement{Name: xml.Name{Local: "w:textAlignment"}}
		e.EncodeElement(m.TextAlignment, setextAlignment)
	}
	if m.TextboxTightWrap != nil {
		setextboxTightWrap := xml.StartElement{Name: xml.Name{Local: "w:textboxTightWrap"}}
		e.EncodeElement(m.TextboxTightWrap, setextboxTightWrap)
	}
	if m.OutlineLvl != nil {
		seoutlineLvl := xml.StartElement{Name: xml.Name{Local: "w:outlineLvl"}}
		e.EncodeElement(m.OutlineLvl, seoutlineLvl)
	}
	if m.DivId != nil {
		sedivId := xml.StartElement{Name: xml.Name{Local: "w:divId"}}
		e.EncodeElement(m.DivId, sedivId)
	}
	if m.CnfStyle != nil {
		secnfStyle := xml.StartElement{Name: xml.Name{Local: "w:cnfStyle"}}
		e.EncodeElement(m.CnfStyle, secnfStyle)
	}
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.SectPr != nil {
		sesectPr := xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}
		e.EncodeElement(m.SectPr, sesectPr)
	}
	if m.PPrChange != nil {
		sepPrChange := xml.StartElement{Name: xml.Name{Local: "w:pPrChange"}}
		e.EncodeElement(m.PPrChange, sepPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pStyle":
				m.PStyle = NewCT_String()
				if err := d.DecodeElement(m.PStyle, &el); err != nil {
					return err
				}
			case "keepNext":
				m.KeepNext = NewCT_OnOff()
				if err := d.DecodeElement(m.KeepNext, &el); err != nil {
					return err
				}
			case "keepLines":
				m.KeepLines = NewCT_OnOff()
				if err := d.DecodeElement(m.KeepLines, &el); err != nil {
					return err
				}
			case "pageBreakBefore":
				m.PageBreakBefore = NewCT_OnOff()
				if err := d.DecodeElement(m.PageBreakBefore, &el); err != nil {
					return err
				}
			case "framePr":
				m.FramePr = NewCT_FramePr()
				if err := d.DecodeElement(m.FramePr, &el); err != nil {
					return err
				}
			case "widowControl":
				m.WidowControl = NewCT_OnOff()
				if err := d.DecodeElement(m.WidowControl, &el); err != nil {
					return err
				}
			case "numPr":
				m.NumPr = NewCT_NumPr()
				if err := d.DecodeElement(m.NumPr, &el); err != nil {
					return err
				}
			case "suppressLineNumbers":
				m.SuppressLineNumbers = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressLineNumbers, &el); err != nil {
					return err
				}
			case "pBdr":
				m.PBdr = NewCT_PBdr()
				if err := d.DecodeElement(m.PBdr, &el); err != nil {
					return err
				}
			case "shd":
				m.Shd = NewCT_Shd()
				if err := d.DecodeElement(m.Shd, &el); err != nil {
					return err
				}
			case "tabs":
				m.Tabs = NewCT_Tabs()
				if err := d.DecodeElement(m.Tabs, &el); err != nil {
					return err
				}
			case "suppressAutoHyphens":
				m.SuppressAutoHyphens = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressAutoHyphens, &el); err != nil {
					return err
				}
			case "kinsoku":
				m.Kinsoku = NewCT_OnOff()
				if err := d.DecodeElement(m.Kinsoku, &el); err != nil {
					return err
				}
			case "wordWrap":
				m.WordWrap = NewCT_OnOff()
				if err := d.DecodeElement(m.WordWrap, &el); err != nil {
					return err
				}
			case "overflowPunct":
				m.OverflowPunct = NewCT_OnOff()
				if err := d.DecodeElement(m.OverflowPunct, &el); err != nil {
					return err
				}
			case "topLinePunct":
				m.TopLinePunct = NewCT_OnOff()
				if err := d.DecodeElement(m.TopLinePunct, &el); err != nil {
					return err
				}
			case "autoSpaceDE":
				m.AutoSpaceDE = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoSpaceDE, &el); err != nil {
					return err
				}
			case "autoSpaceDN":
				m.AutoSpaceDN = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoSpaceDN, &el); err != nil {
					return err
				}
			case "bidi":
				m.Bidi = NewCT_OnOff()
				if err := d.DecodeElement(m.Bidi, &el); err != nil {
					return err
				}
			case "adjustRightInd":
				m.AdjustRightInd = NewCT_OnOff()
				if err := d.DecodeElement(m.AdjustRightInd, &el); err != nil {
					return err
				}
			case "snapToGrid":
				m.SnapToGrid = NewCT_OnOff()
				if err := d.DecodeElement(m.SnapToGrid, &el); err != nil {
					return err
				}
			case "spacing":
				m.Spacing = NewCT_Spacing()
				if err := d.DecodeElement(m.Spacing, &el); err != nil {
					return err
				}
			case "ind":
				m.Ind = NewCT_Ind()
				if err := d.DecodeElement(m.Ind, &el); err != nil {
					return err
				}
			case "contextualSpacing":
				m.ContextualSpacing = NewCT_OnOff()
				if err := d.DecodeElement(m.ContextualSpacing, &el); err != nil {
					return err
				}
			case "mirrorIndents":
				m.MirrorIndents = NewCT_OnOff()
				if err := d.DecodeElement(m.MirrorIndents, &el); err != nil {
					return err
				}
			case "suppressOverlap":
				m.SuppressOverlap = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressOverlap, &el); err != nil {
					return err
				}
			case "jc":
				m.Jc = NewCT_Jc()
				if err := d.DecodeElement(m.Jc, &el); err != nil {
					return err
				}
			case "textDirection":
				m.TextDirection = NewCT_TextDirection()
				if err := d.DecodeElement(m.TextDirection, &el); err != nil {
					return err
				}
			case "textAlignment":
				m.TextAlignment = NewCT_TextAlignment()
				if err := d.DecodeElement(m.TextAlignment, &el); err != nil {
					return err
				}
			case "textboxTightWrap":
				m.TextboxTightWrap = NewCT_TextboxTightWrap()
				if err := d.DecodeElement(m.TextboxTightWrap, &el); err != nil {
					return err
				}
			case "outlineLvl":
				m.OutlineLvl = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.OutlineLvl, &el); err != nil {
					return err
				}
			case "divId":
				m.DivId = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DivId, &el); err != nil {
					return err
				}
			case "cnfStyle":
				m.CnfStyle = NewCT_Cnf()
				if err := d.DecodeElement(m.CnfStyle, &el); err != nil {
					return err
				}
			case "rPr":
				m.RPr = NewCT_ParaRPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case "sectPr":
				m.SectPr = NewCT_SectPr()
				if err := d.DecodeElement(m.SectPr, &el); err != nil {
					return err
				}
			case "pPrChange":
				m.PPrChange = NewCT_PPrChange()
				if err := d.DecodeElement(m.PPrChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PPr) Validate() error {
	return m.ValidateWithPath("CT_PPr")
}
func (m *CT_PPr) ValidateWithPath(path string) error {
	if m.PStyle != nil {
		if err := m.PStyle.ValidateWithPath(path + "/PStyle"); err != nil {
			return err
		}
	}
	if m.KeepNext != nil {
		if err := m.KeepNext.ValidateWithPath(path + "/KeepNext"); err != nil {
			return err
		}
	}
	if m.KeepLines != nil {
		if err := m.KeepLines.ValidateWithPath(path + "/KeepLines"); err != nil {
			return err
		}
	}
	if m.PageBreakBefore != nil {
		if err := m.PageBreakBefore.ValidateWithPath(path + "/PageBreakBefore"); err != nil {
			return err
		}
	}
	if m.FramePr != nil {
		if err := m.FramePr.ValidateWithPath(path + "/FramePr"); err != nil {
			return err
		}
	}
	if m.WidowControl != nil {
		if err := m.WidowControl.ValidateWithPath(path + "/WidowControl"); err != nil {
			return err
		}
	}
	if m.NumPr != nil {
		if err := m.NumPr.ValidateWithPath(path + "/NumPr"); err != nil {
			return err
		}
	}
	if m.SuppressLineNumbers != nil {
		if err := m.SuppressLineNumbers.ValidateWithPath(path + "/SuppressLineNumbers"); err != nil {
			return err
		}
	}
	if m.PBdr != nil {
		if err := m.PBdr.ValidateWithPath(path + "/PBdr"); err != nil {
			return err
		}
	}
	if m.Shd != nil {
		if err := m.Shd.ValidateWithPath(path + "/Shd"); err != nil {
			return err
		}
	}
	if m.Tabs != nil {
		if err := m.Tabs.ValidateWithPath(path + "/Tabs"); err != nil {
			return err
		}
	}
	if m.SuppressAutoHyphens != nil {
		if err := m.SuppressAutoHyphens.ValidateWithPath(path + "/SuppressAutoHyphens"); err != nil {
			return err
		}
	}
	if m.Kinsoku != nil {
		if err := m.Kinsoku.ValidateWithPath(path + "/Kinsoku"); err != nil {
			return err
		}
	}
	if m.WordWrap != nil {
		if err := m.WordWrap.ValidateWithPath(path + "/WordWrap"); err != nil {
			return err
		}
	}
	if m.OverflowPunct != nil {
		if err := m.OverflowPunct.ValidateWithPath(path + "/OverflowPunct"); err != nil {
			return err
		}
	}
	if m.TopLinePunct != nil {
		if err := m.TopLinePunct.ValidateWithPath(path + "/TopLinePunct"); err != nil {
			return err
		}
	}
	if m.AutoSpaceDE != nil {
		if err := m.AutoSpaceDE.ValidateWithPath(path + "/AutoSpaceDE"); err != nil {
			return err
		}
	}
	if m.AutoSpaceDN != nil {
		if err := m.AutoSpaceDN.ValidateWithPath(path + "/AutoSpaceDN"); err != nil {
			return err
		}
	}
	if m.Bidi != nil {
		if err := m.Bidi.ValidateWithPath(path + "/Bidi"); err != nil {
			return err
		}
	}
	if m.AdjustRightInd != nil {
		if err := m.AdjustRightInd.ValidateWithPath(path + "/AdjustRightInd"); err != nil {
			return err
		}
	}
	if m.SnapToGrid != nil {
		if err := m.SnapToGrid.ValidateWithPath(path + "/SnapToGrid"); err != nil {
			return err
		}
	}
	if m.Spacing != nil {
		if err := m.Spacing.ValidateWithPath(path + "/Spacing"); err != nil {
			return err
		}
	}
	if m.Ind != nil {
		if err := m.Ind.ValidateWithPath(path + "/Ind"); err != nil {
			return err
		}
	}
	if m.ContextualSpacing != nil {
		if err := m.ContextualSpacing.ValidateWithPath(path + "/ContextualSpacing"); err != nil {
			return err
		}
	}
	if m.MirrorIndents != nil {
		if err := m.MirrorIndents.ValidateWithPath(path + "/MirrorIndents"); err != nil {
			return err
		}
	}
	if m.SuppressOverlap != nil {
		if err := m.SuppressOverlap.ValidateWithPath(path + "/SuppressOverlap"); err != nil {
			return err
		}
	}
	if m.Jc != nil {
		if err := m.Jc.ValidateWithPath(path + "/Jc"); err != nil {
			return err
		}
	}
	if m.TextDirection != nil {
		if err := m.TextDirection.ValidateWithPath(path + "/TextDirection"); err != nil {
			return err
		}
	}
	if m.TextAlignment != nil {
		if err := m.TextAlignment.ValidateWithPath(path + "/TextAlignment"); err != nil {
			return err
		}
	}
	if m.TextboxTightWrap != nil {
		if err := m.TextboxTightWrap.ValidateWithPath(path + "/TextboxTightWrap"); err != nil {
			return err
		}
	}
	if m.OutlineLvl != nil {
		if err := m.OutlineLvl.ValidateWithPath(path + "/OutlineLvl"); err != nil {
			return err
		}
	}
	if m.DivId != nil {
		if err := m.DivId.ValidateWithPath(path + "/DivId"); err != nil {
			return err
		}
	}
	if m.CnfStyle != nil {
		if err := m.CnfStyle.ValidateWithPath(path + "/CnfStyle"); err != nil {
			return err
		}
	}
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	if m.SectPr != nil {
		if err := m.SectPr.ValidateWithPath(path + "/SectPr"); err != nil {
			return err
		}
	}
	if m.PPrChange != nil {
		if err := m.PPrChange.ValidateWithPath(path + "/PPrChange"); err != nil {
			return err
		}
	}
	return nil
}
