// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type Dialogsheet struct {
	CT_Dialogsheet
}

func NewDialogsheet() *Dialogsheet {
	ret := &Dialogsheet{}
	ret.CT_Dialogsheet = *NewCT_Dialogsheet()
	return ret
}
func (m *Dialogsheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:dialogsheet"
	return m.CT_Dialogsheet.MarshalXML(e, start)
}
func (m *Dialogsheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Dialogsheet = *NewCT_Dialogsheet()
lDialogsheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetPr":
				m.SheetPr = NewCT_SheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case "sheetViews":
				m.SheetViews = NewCT_SheetViews()
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case "sheetFormatPr":
				m.SheetFormatPr = NewCT_SheetFormatPr()
				if err := d.DecodeElement(m.SheetFormatPr, &el); err != nil {
					return err
				}
			case "sheetProtection":
				m.SheetProtection = NewCT_SheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case "customSheetViews":
				m.CustomSheetViews = NewCT_CustomSheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case "printOptions":
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case "drawing":
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case "legacyDrawing":
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case "legacyDrawingHF":
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case "drawingHF":
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case "oleObjects":
				m.OleObjects = NewCT_OleObjects()
				if err := d.DecodeElement(m.OleObjects, &el); err != nil {
					return err
				}
			case "controls":
				m.Controls = NewCT_Controls()
				if err := d.DecodeElement(m.Controls, &el); err != nil {
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
			break lDialogsheet
		case xml.CharData:
		}
	}
	return nil
}
func (m *Dialogsheet) Validate() error {
	return m.ValidateWithPath("Dialogsheet")
}
func (m *Dialogsheet) ValidateWithPath(path string) error {
	if err := m.CT_Dialogsheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
