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
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type CT_CustomXmlRun struct {
	// Custom XML Markup Namespace
	UriAttr *string
	// Element name
	ElementAttr string
	// Custom XML Element Properties
	CustomXmlPr *CT_CustomXmlPr
	EG_PContent []*EG_PContent
}

func NewCT_CustomXmlRun() *CT_CustomXmlRun {
	ret := &CT_CustomXmlRun{}
	return ret
}
func (m *CT_CustomXmlRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:element"},
		Value: fmt.Sprintf("%v", m.ElementAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.CustomXmlPr != nil {
		secustomXmlPr := xml.StartElement{Name: xml.Name{Local: "w:customXmlPr"}}
		e.EncodeElement(m.CustomXmlPr, secustomXmlPr)
	}
	if m.EG_PContent != nil {
		for _, c := range m.EG_PContent {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomXmlRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
		}
		if attr.Name.Local == "element" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ElementAttr = parsed
		}
	}
lCT_CustomXmlRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customXmlPr":
				m.CustomXmlPr = NewCT_CustomXmlPr()
				if err := d.DecodeElement(m.CustomXmlPr, &el); err != nil {
					return err
				}
			case "fldSimple":
				tmppcontent := NewEG_PContent()
				tmp := NewCT_SimpleField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmppcontent.FldSimple = append(tmppcontent.FldSimple, tmp)
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case "hyperlink":
				tmppcontent := NewEG_PContent()
				tmppcontent.Hyperlink = NewCT_Hyperlink()
				if err := d.DecodeElement(tmppcontent.Hyperlink, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case "subDoc":
				tmppcontent := NewEG_PContent()
				tmppcontent.SubDoc = NewCT_Rel()
				if err := d.DecodeElement(tmppcontent.SubDoc, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case "customXml":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.CustomXml = NewCT_CustomXmlRun()
				if err := d.DecodeElement(tmpcontentruncontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "smartTag":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.SmartTag = NewCT_SmartTagRun()
				if err := d.DecodeElement(tmpcontentruncontent.SmartTag, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "sdt":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Sdt = NewCT_SdtRun()
				if err := d.DecodeElement(tmpcontentruncontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "dir":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Dir = NewCT_DirContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Dir, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "bdo":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Bdo = NewCT_BdoContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Bdo, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "r":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.R = NewCT_R()
				if err := d.DecodeElement(tmpcontentruncontent.R, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case "proofErr":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "permStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "permEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "ins":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "del":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveFrom":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveTo":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case "bookmarkStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomXmlRun
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomXmlRun) Validate() error {
	return m.ValidateWithPath("CT_CustomXmlRun")
}
func (m *CT_CustomXmlRun) ValidateWithPath(path string) error {
	if m.CustomXmlPr != nil {
		if err := m.CustomXmlPr.ValidateWithPath(path + "/CustomXmlPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_PContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_PContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
