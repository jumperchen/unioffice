// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Trendline struct {
	Name          *string
	SpPr          *drawingml.CT_ShapeProperties
	TrendlineType *CT_TrendlineType
	Order         *CT_Order
	Period        *CT_Period
	Forward       *CT_Double
	Backward      *CT_Double
	Intercept     *CT_Double
	DispRSqr      *CT_Boolean
	DispEq        *CT_Boolean
	TrendlineLbl  *CT_TrendlineLbl
	ExtLst        *CT_ExtensionList
}

func NewCT_Trendline() *CT_Trendline {
	ret := &CT_Trendline{}
	ret.TrendlineType = NewCT_TrendlineType()
	return ret
}
func (m *CT_Trendline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "name"}}
		gooxml.AddPreserveSpaceAttr(&sename, *m.Name)
		e.EncodeElement(m.Name, sename)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	setrendlineType := xml.StartElement{Name: xml.Name{Local: "trendlineType"}}
	e.EncodeElement(m.TrendlineType, setrendlineType)
	if m.Order != nil {
		seorder := xml.StartElement{Name: xml.Name{Local: "order"}}
		e.EncodeElement(m.Order, seorder)
	}
	if m.Period != nil {
		seperiod := xml.StartElement{Name: xml.Name{Local: "period"}}
		e.EncodeElement(m.Period, seperiod)
	}
	if m.Forward != nil {
		seforward := xml.StartElement{Name: xml.Name{Local: "forward"}}
		e.EncodeElement(m.Forward, seforward)
	}
	if m.Backward != nil {
		sebackward := xml.StartElement{Name: xml.Name{Local: "backward"}}
		e.EncodeElement(m.Backward, sebackward)
	}
	if m.Intercept != nil {
		seintercept := xml.StartElement{Name: xml.Name{Local: "intercept"}}
		e.EncodeElement(m.Intercept, seintercept)
	}
	if m.DispRSqr != nil {
		sedispRSqr := xml.StartElement{Name: xml.Name{Local: "dispRSqr"}}
		e.EncodeElement(m.DispRSqr, sedispRSqr)
	}
	if m.DispEq != nil {
		sedispEq := xml.StartElement{Name: xml.Name{Local: "dispEq"}}
		e.EncodeElement(m.DispEq, sedispEq)
	}
	if m.TrendlineLbl != nil {
		setrendlineLbl := xml.StartElement{Name: xml.Name{Local: "trendlineLbl"}}
		e.EncodeElement(m.TrendlineLbl, setrendlineLbl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Trendline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TrendlineType = NewCT_TrendlineType()
lCT_Trendline:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "name":
				m.Name = new(string)
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "trendlineType":
				if err := d.DecodeElement(m.TrendlineType, &el); err != nil {
					return err
				}
			case "order":
				m.Order = NewCT_Order()
				if err := d.DecodeElement(m.Order, &el); err != nil {
					return err
				}
			case "period":
				m.Period = NewCT_Period()
				if err := d.DecodeElement(m.Period, &el); err != nil {
					return err
				}
			case "forward":
				m.Forward = NewCT_Double()
				if err := d.DecodeElement(m.Forward, &el); err != nil {
					return err
				}
			case "backward":
				m.Backward = NewCT_Double()
				if err := d.DecodeElement(m.Backward, &el); err != nil {
					return err
				}
			case "intercept":
				m.Intercept = NewCT_Double()
				if err := d.DecodeElement(m.Intercept, &el); err != nil {
					return err
				}
			case "dispRSqr":
				m.DispRSqr = NewCT_Boolean()
				if err := d.DecodeElement(m.DispRSqr, &el); err != nil {
					return err
				}
			case "dispEq":
				m.DispEq = NewCT_Boolean()
				if err := d.DecodeElement(m.DispEq, &el); err != nil {
					return err
				}
			case "trendlineLbl":
				m.TrendlineLbl = NewCT_TrendlineLbl()
				if err := d.DecodeElement(m.TrendlineLbl, &el); err != nil {
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
			break lCT_Trendline
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Trendline) Validate() error {
	return m.ValidateWithPath("CT_Trendline")
}
func (m *CT_Trendline) ValidateWithPath(path string) error {
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if err := m.TrendlineType.ValidateWithPath(path + "/TrendlineType"); err != nil {
		return err
	}
	if m.Order != nil {
		if err := m.Order.ValidateWithPath(path + "/Order"); err != nil {
			return err
		}
	}
	if m.Period != nil {
		if err := m.Period.ValidateWithPath(path + "/Period"); err != nil {
			return err
		}
	}
	if m.Forward != nil {
		if err := m.Forward.ValidateWithPath(path + "/Forward"); err != nil {
			return err
		}
	}
	if m.Backward != nil {
		if err := m.Backward.ValidateWithPath(path + "/Backward"); err != nil {
			return err
		}
	}
	if m.Intercept != nil {
		if err := m.Intercept.ValidateWithPath(path + "/Intercept"); err != nil {
			return err
		}
	}
	if m.DispRSqr != nil {
		if err := m.DispRSqr.ValidateWithPath(path + "/DispRSqr"); err != nil {
			return err
		}
	}
	if m.DispEq != nil {
		if err := m.DispEq.ValidateWithPath(path + "/DispEq"); err != nil {
			return err
		}
	}
	if m.TrendlineLbl != nil {
		if err := m.TrendlineLbl.ValidateWithPath(path + "/TrendlineLbl"); err != nil {
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
