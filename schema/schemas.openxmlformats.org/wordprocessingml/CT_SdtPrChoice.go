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

type CT_SdtPrChoice struct {
	Equation     *CT_Empty
	ComboBox     *CT_SdtComboBox
	Date         *CT_SdtDate
	DocPartObj   *CT_SdtDocPart
	DocPartList  *CT_SdtDocPart
	DropDownList *CT_SdtDropDownList
	Picture      *CT_Empty
	RichText     *CT_Empty
	Text         *CT_SdtText
	Citation     *CT_Empty
	Group        *CT_Empty
	Bibliography *CT_Empty
}

func NewCT_SdtPrChoice() *CT_SdtPrChoice {
	ret := &CT_SdtPrChoice{}
	return ret
}
func (m *CT_SdtPrChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Equation != nil {
		seequation := xml.StartElement{Name: xml.Name{Local: "w:equation"}}
		e.EncodeElement(m.Equation, seequation)
	}
	if m.ComboBox != nil {
		secomboBox := xml.StartElement{Name: xml.Name{Local: "w:comboBox"}}
		e.EncodeElement(m.ComboBox, secomboBox)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "w:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.DocPartObj != nil {
		sedocPartObj := xml.StartElement{Name: xml.Name{Local: "w:docPartObj"}}
		e.EncodeElement(m.DocPartObj, sedocPartObj)
	}
	if m.DocPartList != nil {
		sedocPartList := xml.StartElement{Name: xml.Name{Local: "w:docPartList"}}
		e.EncodeElement(m.DocPartList, sedocPartList)
	}
	if m.DropDownList != nil {
		sedropDownList := xml.StartElement{Name: xml.Name{Local: "w:dropDownList"}}
		e.EncodeElement(m.DropDownList, sedropDownList)
	}
	if m.Picture != nil {
		sepicture := xml.StartElement{Name: xml.Name{Local: "w:picture"}}
		e.EncodeElement(m.Picture, sepicture)
	}
	if m.RichText != nil {
		serichText := xml.StartElement{Name: xml.Name{Local: "w:richText"}}
		e.EncodeElement(m.RichText, serichText)
	}
	if m.Text != nil {
		setext := xml.StartElement{Name: xml.Name{Local: "w:text"}}
		e.EncodeElement(m.Text, setext)
	}
	if m.Citation != nil {
		secitation := xml.StartElement{Name: xml.Name{Local: "w:citation"}}
		e.EncodeElement(m.Citation, secitation)
	}
	if m.Group != nil {
		segroup := xml.StartElement{Name: xml.Name{Local: "w:group"}}
		e.EncodeElement(m.Group, segroup)
	}
	if m.Bibliography != nil {
		sebibliography := xml.StartElement{Name: xml.Name{Local: "w:bibliography"}}
		e.EncodeElement(m.Bibliography, sebibliography)
	}
	return nil
}
func (m *CT_SdtPrChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtPrChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "equation":
				m.Equation = NewCT_Empty()
				if err := d.DecodeElement(m.Equation, &el); err != nil {
					return err
				}
			case "comboBox":
				m.ComboBox = NewCT_SdtComboBox()
				if err := d.DecodeElement(m.ComboBox, &el); err != nil {
					return err
				}
			case "date":
				m.Date = NewCT_SdtDate()
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case "docPartObj":
				m.DocPartObj = NewCT_SdtDocPart()
				if err := d.DecodeElement(m.DocPartObj, &el); err != nil {
					return err
				}
			case "docPartList":
				m.DocPartList = NewCT_SdtDocPart()
				if err := d.DecodeElement(m.DocPartList, &el); err != nil {
					return err
				}
			case "dropDownList":
				m.DropDownList = NewCT_SdtDropDownList()
				if err := d.DecodeElement(m.DropDownList, &el); err != nil {
					return err
				}
			case "picture":
				m.Picture = NewCT_Empty()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case "richText":
				m.RichText = NewCT_Empty()
				if err := d.DecodeElement(m.RichText, &el); err != nil {
					return err
				}
			case "text":
				m.Text = NewCT_SdtText()
				if err := d.DecodeElement(m.Text, &el); err != nil {
					return err
				}
			case "citation":
				m.Citation = NewCT_Empty()
				if err := d.DecodeElement(m.Citation, &el); err != nil {
					return err
				}
			case "group":
				m.Group = NewCT_Empty()
				if err := d.DecodeElement(m.Group, &el); err != nil {
					return err
				}
			case "bibliography":
				m.Bibliography = NewCT_Empty()
				if err := d.DecodeElement(m.Bibliography, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtPrChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SdtPrChoice) Validate() error {
	return m.ValidateWithPath("CT_SdtPrChoice")
}
func (m *CT_SdtPrChoice) ValidateWithPath(path string) error {
	if m.Equation != nil {
		if err := m.Equation.ValidateWithPath(path + "/Equation"); err != nil {
			return err
		}
	}
	if m.ComboBox != nil {
		if err := m.ComboBox.ValidateWithPath(path + "/ComboBox"); err != nil {
			return err
		}
	}
	if m.Date != nil {
		if err := m.Date.ValidateWithPath(path + "/Date"); err != nil {
			return err
		}
	}
	if m.DocPartObj != nil {
		if err := m.DocPartObj.ValidateWithPath(path + "/DocPartObj"); err != nil {
			return err
		}
	}
	if m.DocPartList != nil {
		if err := m.DocPartList.ValidateWithPath(path + "/DocPartList"); err != nil {
			return err
		}
	}
	if m.DropDownList != nil {
		if err := m.DropDownList.ValidateWithPath(path + "/DropDownList"); err != nil {
			return err
		}
	}
	if m.Picture != nil {
		if err := m.Picture.ValidateWithPath(path + "/Picture"); err != nil {
			return err
		}
	}
	if m.RichText != nil {
		if err := m.RichText.ValidateWithPath(path + "/RichText"); err != nil {
			return err
		}
	}
	if m.Text != nil {
		if err := m.Text.ValidateWithPath(path + "/Text"); err != nil {
			return err
		}
	}
	if m.Citation != nil {
		if err := m.Citation.ValidateWithPath(path + "/Citation"); err != nil {
			return err
		}
	}
	if m.Group != nil {
		if err := m.Group.ValidateWithPath(path + "/Group"); err != nil {
			return err
		}
	}
	if m.Bibliography != nil {
		if err := m.Bibliography.ValidateWithPath(path + "/Bibliography"); err != nil {
			return err
		}
	}
	return nil
}
