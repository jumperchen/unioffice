// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_OfPieChart struct {
	OfPieType     *CT_OfPieType
	VaryColors    *CT_Boolean
	Ser           []*CT_PieSer
	DLbls         *CT_DLbls
	GapWidth      *CT_GapAmount
	SplitType     *CT_SplitType
	SplitPos      *CT_Double
	CustSplit     *CT_CustSplit
	SecondPieSize *CT_SecondPieSize
	SerLines      []*CT_ChartLines
	ExtLst        *CT_ExtensionList
}

func NewCT_OfPieChart() *CT_OfPieChart {
	ret := &CT_OfPieChart{}
	ret.OfPieType = NewCT_OfPieType()
	return ret
}
func (m *CT_OfPieChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seofPieType := xml.StartElement{Name: xml.Name{Local: "ofPieType"}}
	e.EncodeElement(m.OfPieType, seofPieType)
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "ser"}}
		e.EncodeElement(m.Ser, seser)
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.GapWidth != nil {
		segapWidth := xml.StartElement{Name: xml.Name{Local: "gapWidth"}}
		e.EncodeElement(m.GapWidth, segapWidth)
	}
	if m.SplitType != nil {
		sesplitType := xml.StartElement{Name: xml.Name{Local: "splitType"}}
		e.EncodeElement(m.SplitType, sesplitType)
	}
	if m.SplitPos != nil {
		sesplitPos := xml.StartElement{Name: xml.Name{Local: "splitPos"}}
		e.EncodeElement(m.SplitPos, sesplitPos)
	}
	if m.CustSplit != nil {
		secustSplit := xml.StartElement{Name: xml.Name{Local: "custSplit"}}
		e.EncodeElement(m.CustSplit, secustSplit)
	}
	if m.SecondPieSize != nil {
		sesecondPieSize := xml.StartElement{Name: xml.Name{Local: "secondPieSize"}}
		e.EncodeElement(m.SecondPieSize, sesecondPieSize)
	}
	if m.SerLines != nil {
		seserLines := xml.StartElement{Name: xml.Name{Local: "serLines"}}
		e.EncodeElement(m.SerLines, seserLines)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OfPieChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfPieType = NewCT_OfPieType()
lCT_OfPieChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ofPieType":
				if err := d.DecodeElement(m.OfPieType, &el); err != nil {
					return err
				}
			case "varyColors":
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case "ser":
				tmp := NewCT_PieSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case "dLbls":
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case "gapWidth":
				m.GapWidth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapWidth, &el); err != nil {
					return err
				}
			case "splitType":
				m.SplitType = NewCT_SplitType()
				if err := d.DecodeElement(m.SplitType, &el); err != nil {
					return err
				}
			case "splitPos":
				m.SplitPos = NewCT_Double()
				if err := d.DecodeElement(m.SplitPos, &el); err != nil {
					return err
				}
			case "custSplit":
				m.CustSplit = NewCT_CustSplit()
				if err := d.DecodeElement(m.CustSplit, &el); err != nil {
					return err
				}
			case "secondPieSize":
				m.SecondPieSize = NewCT_SecondPieSize()
				if err := d.DecodeElement(m.SecondPieSize, &el); err != nil {
					return err
				}
			case "serLines":
				tmp := NewCT_ChartLines()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SerLines = append(m.SerLines, tmp)
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
			break lCT_OfPieChart
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OfPieChart) Validate() error {
	return m.ValidateWithPath("CT_OfPieChart")
}
func (m *CT_OfPieChart) ValidateWithPath(path string) error {
	if err := m.OfPieType.ValidateWithPath(path + "/OfPieType"); err != nil {
		return err
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.GapWidth != nil {
		if err := m.GapWidth.ValidateWithPath(path + "/GapWidth"); err != nil {
			return err
		}
	}
	if m.SplitType != nil {
		if err := m.SplitType.ValidateWithPath(path + "/SplitType"); err != nil {
			return err
		}
	}
	if m.SplitPos != nil {
		if err := m.SplitPos.ValidateWithPath(path + "/SplitPos"); err != nil {
			return err
		}
	}
	if m.CustSplit != nil {
		if err := m.CustSplit.ValidateWithPath(path + "/CustSplit"); err != nil {
			return err
		}
	}
	if m.SecondPieSize != nil {
		if err := m.SecondPieSize.ValidateWithPath(path + "/SecondPieSize"); err != nil {
			return err
		}
	}
	for i, v := range m.SerLines {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SerLines[%d]", path, i)); err != nil {
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
