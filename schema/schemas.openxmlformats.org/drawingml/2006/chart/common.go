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
	"regexp"

	"baliance.com/gooxml"
)

const ST_HPercentWithSymbolPattern = `0*(([5-9])|([1-9][0-9])|([1-4][0-9][0-9])|500)%`

var ST_HPercentWithSymbolPatternRe = regexp.MustCompile(ST_HPercentWithSymbolPattern)

const ST_DepthPercentWithSymbolPattern = `0*(([2-9][0-9])|([1-9][0-9][0-9])|(1[0-9][0-9][0-9])|2000)%`

var ST_DepthPercentWithSymbolPatternRe = regexp.MustCompile(ST_DepthPercentWithSymbolPattern)

const ST_ThicknessPercentPattern = `([0-9]+)%`

var ST_ThicknessPercentPatternRe = regexp.MustCompile(ST_ThicknessPercentPattern)

const ST_GapAmountPercentPattern = `0*(([0-9])|([1-9][0-9])|([1-4][0-9][0-9])|500)%`

var ST_GapAmountPercentPatternRe = regexp.MustCompile(ST_GapAmountPercentPattern)

const ST_OverlapPercentPattern = `(-?0*(([0-9])|([1-9][0-9])|100))%`

var ST_OverlapPercentPatternRe = regexp.MustCompile(ST_OverlapPercentPattern)

const ST_BubbleScalePercentPattern = `0*(([0-9])|([1-9][0-9])|([1-2][0-9][0-9])|300)%`

var ST_BubbleScalePercentPatternRe = regexp.MustCompile(ST_BubbleScalePercentPattern)

const ST_HoleSizePercentPattern = `0*([1-9]|([1-8][0-9])|90)%`

var ST_HoleSizePercentPatternRe = regexp.MustCompile(ST_HoleSizePercentPattern)

const ST_SecondPieSizePercentPattern = `0*(([5-9])|([1-9][0-9])|(1[0-9][0-9])|200)%`

var ST_SecondPieSizePercentPatternRe = regexp.MustCompile(ST_SecondPieSizePercentPattern)

const ST_LblOffsetPercentPattern = `0*(([0-9])|([1-9][0-9])|([1-9][0-9][0-9])|1000)%`

var ST_LblOffsetPercentPatternRe = regexp.MustCompile(ST_LblOffsetPercentPattern)

func ParseUnionST_HPercent(s string) (ST_HPercent, error) {
	return ST_HPercent{}, nil
}
func ParseUnionST_DepthPercent(s string) (ST_DepthPercent, error) {
	return ST_DepthPercent{}, nil
}
func ParseUnionST_Thickness(s string) (ST_Thickness, error) {
	return ST_Thickness{}, nil
}
func ParseUnionST_GapAmount(s string) (ST_GapAmount, error) {
	return ST_GapAmount{}, nil
}
func ParseUnionST_Overlap(s string) (ST_Overlap, error) {
	return ST_Overlap{}, nil
}
func ParseUnionST_BubbleScale(s string) (ST_BubbleScale, error) {
	return ST_BubbleScale{}, nil
}
func ParseUnionST_HoleSize(s string) (ST_HoleSize, error) {
	return ST_HoleSize{}, nil
}
func ParseUnionST_SecondPieSize(s string) (ST_SecondPieSize, error) {
	return ST_SecondPieSize{}, nil
}
func ParseUnionST_LblOffset(s string) (ST_LblOffset, error) {
	return ST_LblOffset{}, nil
}

type Any interface {
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}

type ST_LayoutTarget byte

const (
	ST_LayoutTargetUnset ST_LayoutTarget = 0
	ST_LayoutTargetInner ST_LayoutTarget = 1
	ST_LayoutTargetOuter ST_LayoutTarget = 2
)

func (e ST_LayoutTarget) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LayoutTargetUnset:
		attr.Value = ""
	case ST_LayoutTargetInner:
		attr.Value = "inner"
	case ST_LayoutTargetOuter:
		attr.Value = "outer"
	}
	return attr, nil
}
func (e *ST_LayoutTarget) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "inner":
		*e = 1
	case "outer":
		*e = 2
	}
	return nil
}
func (m ST_LayoutTarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LayoutTarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "inner":
			*m = 1
		case "outer":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_LayoutTarget) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "inner"
	case 2:
		return "outer"
	}
	return ""
}
func (m ST_LayoutTarget) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LayoutTarget) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LayoutMode byte

const (
	ST_LayoutModeUnset  ST_LayoutMode = 0
	ST_LayoutModeEdge   ST_LayoutMode = 1
	ST_LayoutModeFactor ST_LayoutMode = 2
)

func (e ST_LayoutMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LayoutModeUnset:
		attr.Value = ""
	case ST_LayoutModeEdge:
		attr.Value = "edge"
	case ST_LayoutModeFactor:
		attr.Value = "factor"
	}
	return attr, nil
}
func (e *ST_LayoutMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "edge":
		*e = 1
	case "factor":
		*e = 2
	}
	return nil
}
func (m ST_LayoutMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LayoutMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "edge":
			*m = 1
		case "factor":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_LayoutMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "edge"
	case 2:
		return "factor"
	}
	return ""
}
func (m ST_LayoutMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LayoutMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SizeRepresents byte

const (
	ST_SizeRepresentsUnset ST_SizeRepresents = 0
	ST_SizeRepresentsArea  ST_SizeRepresents = 1
	ST_SizeRepresentsW     ST_SizeRepresents = 2
)

func (e ST_SizeRepresents) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SizeRepresentsUnset:
		attr.Value = ""
	case ST_SizeRepresentsArea:
		attr.Value = "area"
	case ST_SizeRepresentsW:
		attr.Value = "w"
	}
	return attr, nil
}
func (e *ST_SizeRepresents) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "area":
		*e = 1
	case "w":
		*e = 2
	}
	return nil
}
func (m ST_SizeRepresents) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SizeRepresents) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "area":
			*m = 1
		case "w":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_SizeRepresents) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "area"
	case 2:
		return "w"
	}
	return ""
}
func (m ST_SizeRepresents) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SizeRepresents) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SplitType byte

const (
	ST_SplitTypeUnset   ST_SplitType = 0
	ST_SplitTypeAuto    ST_SplitType = 1
	ST_SplitTypeCust    ST_SplitType = 2
	ST_SplitTypePercent ST_SplitType = 3
	ST_SplitTypePos     ST_SplitType = 4
	ST_SplitTypeVal     ST_SplitType = 5
)

func (e ST_SplitType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SplitTypeUnset:
		attr.Value = ""
	case ST_SplitTypeAuto:
		attr.Value = "auto"
	case ST_SplitTypeCust:
		attr.Value = "cust"
	case ST_SplitTypePercent:
		attr.Value = "percent"
	case ST_SplitTypePos:
		attr.Value = "pos"
	case ST_SplitTypeVal:
		attr.Value = "val"
	}
	return attr, nil
}
func (e *ST_SplitType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "cust":
		*e = 2
	case "percent":
		*e = 3
	case "pos":
		*e = 4
	case "val":
		*e = 5
	}
	return nil
}
func (m ST_SplitType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SplitType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "cust":
			*m = 2
		case "percent":
			*m = 3
		case "pos":
			*m = 4
		case "val":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_SplitType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "cust"
	case 3:
		return "percent"
	case 4:
		return "pos"
	case 5:
		return "val"
	}
	return ""
}
func (m ST_SplitType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SplitType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LblAlgn byte

const (
	ST_LblAlgnUnset ST_LblAlgn = 0
	ST_LblAlgnCtr   ST_LblAlgn = 1
	ST_LblAlgnL     ST_LblAlgn = 2
	ST_LblAlgnR     ST_LblAlgn = 3
)

func (e ST_LblAlgn) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LblAlgnUnset:
		attr.Value = ""
	case ST_LblAlgnCtr:
		attr.Value = "ctr"
	case ST_LblAlgnL:
		attr.Value = "l"
	case ST_LblAlgnR:
		attr.Value = "r"
	}
	return attr, nil
}
func (e *ST_LblAlgn) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ctr":
		*e = 1
	case "l":
		*e = 2
	case "r":
		*e = 3
	}
	return nil
}
func (m ST_LblAlgn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LblAlgn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "ctr":
			*m = 1
		case "l":
			*m = 2
		case "r":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_LblAlgn) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ctr"
	case 2:
		return "l"
	case 3:
		return "r"
	}
	return ""
}
func (m ST_LblAlgn) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LblAlgn) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DLblPos byte

const (
	ST_DLblPosUnset   ST_DLblPos = 0
	ST_DLblPosBestFit ST_DLblPos = 1
	ST_DLblPosB       ST_DLblPos = 2
	ST_DLblPosCtr     ST_DLblPos = 3
	ST_DLblPosInBase  ST_DLblPos = 4
	ST_DLblPosInEnd   ST_DLblPos = 5
	ST_DLblPosL       ST_DLblPos = 6
	ST_DLblPosOutEnd  ST_DLblPos = 7
	ST_DLblPosR       ST_DLblPos = 8
	ST_DLblPosT       ST_DLblPos = 9
)

func (e ST_DLblPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DLblPosUnset:
		attr.Value = ""
	case ST_DLblPosBestFit:
		attr.Value = "bestFit"
	case ST_DLblPosB:
		attr.Value = "b"
	case ST_DLblPosCtr:
		attr.Value = "ctr"
	case ST_DLblPosInBase:
		attr.Value = "inBase"
	case ST_DLblPosInEnd:
		attr.Value = "inEnd"
	case ST_DLblPosL:
		attr.Value = "l"
	case ST_DLblPosOutEnd:
		attr.Value = "outEnd"
	case ST_DLblPosR:
		attr.Value = "r"
	case ST_DLblPosT:
		attr.Value = "t"
	}
	return attr, nil
}
func (e *ST_DLblPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bestFit":
		*e = 1
	case "b":
		*e = 2
	case "ctr":
		*e = 3
	case "inBase":
		*e = 4
	case "inEnd":
		*e = 5
	case "l":
		*e = 6
	case "outEnd":
		*e = 7
	case "r":
		*e = 8
	case "t":
		*e = 9
	}
	return nil
}
func (m ST_DLblPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_DLblPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bestFit":
			*m = 1
		case "b":
			*m = 2
		case "ctr":
			*m = 3
		case "inBase":
			*m = 4
		case "inEnd":
			*m = 5
		case "l":
			*m = 6
		case "outEnd":
			*m = 7
		case "r":
			*m = 8
		case "t":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_DLblPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bestFit"
	case 2:
		return "b"
	case 3:
		return "ctr"
	case 4:
		return "inBase"
	case 5:
		return "inEnd"
	case 6:
		return "l"
	case 7:
		return "outEnd"
	case 8:
		return "r"
	case 9:
		return "t"
	}
	return ""
}
func (m ST_DLblPos) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_DLblPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MarkerStyle byte

const (
	ST_MarkerStyleUnset    ST_MarkerStyle = 0
	ST_MarkerStyleCircle   ST_MarkerStyle = 1
	ST_MarkerStyleDash     ST_MarkerStyle = 2
	ST_MarkerStyleDiamond  ST_MarkerStyle = 3
	ST_MarkerStyleDot      ST_MarkerStyle = 4
	ST_MarkerStyleNone     ST_MarkerStyle = 5
	ST_MarkerStylePicture  ST_MarkerStyle = 6
	ST_MarkerStylePlus     ST_MarkerStyle = 7
	ST_MarkerStyleSquare   ST_MarkerStyle = 8
	ST_MarkerStyleStar     ST_MarkerStyle = 9
	ST_MarkerStyleTriangle ST_MarkerStyle = 10
	ST_MarkerStyleX        ST_MarkerStyle = 11
	ST_MarkerStyleAuto     ST_MarkerStyle = 12
)

func (e ST_MarkerStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MarkerStyleUnset:
		attr.Value = ""
	case ST_MarkerStyleCircle:
		attr.Value = "circle"
	case ST_MarkerStyleDash:
		attr.Value = "dash"
	case ST_MarkerStyleDiamond:
		attr.Value = "diamond"
	case ST_MarkerStyleDot:
		attr.Value = "dot"
	case ST_MarkerStyleNone:
		attr.Value = "none"
	case ST_MarkerStylePicture:
		attr.Value = "picture"
	case ST_MarkerStylePlus:
		attr.Value = "plus"
	case ST_MarkerStyleSquare:
		attr.Value = "square"
	case ST_MarkerStyleStar:
		attr.Value = "star"
	case ST_MarkerStyleTriangle:
		attr.Value = "triangle"
	case ST_MarkerStyleX:
		attr.Value = "x"
	case ST_MarkerStyleAuto:
		attr.Value = "auto"
	}
	return attr, nil
}
func (e *ST_MarkerStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "circle":
		*e = 1
	case "dash":
		*e = 2
	case "diamond":
		*e = 3
	case "dot":
		*e = 4
	case "none":
		*e = 5
	case "picture":
		*e = 6
	case "plus":
		*e = 7
	case "square":
		*e = 8
	case "star":
		*e = 9
	case "triangle":
		*e = 10
	case "x":
		*e = 11
	case "auto":
		*e = 12
	}
	return nil
}
func (m ST_MarkerStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_MarkerStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "circle":
			*m = 1
		case "dash":
			*m = 2
		case "diamond":
			*m = 3
		case "dot":
			*m = 4
		case "none":
			*m = 5
		case "picture":
			*m = 6
		case "plus":
			*m = 7
		case "square":
			*m = 8
		case "star":
			*m = 9
		case "triangle":
			*m = 10
		case "x":
			*m = 11
		case "auto":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_MarkerStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "circle"
	case 2:
		return "dash"
	case 3:
		return "diamond"
	case 4:
		return "dot"
	case 5:
		return "none"
	case 6:
		return "picture"
	case 7:
		return "plus"
	case 8:
		return "square"
	case 9:
		return "star"
	case 10:
		return "triangle"
	case 11:
		return "x"
	case 12:
		return "auto"
	}
	return ""
}
func (m ST_MarkerStyle) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_MarkerStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TrendlineType byte

const (
	ST_TrendlineTypeUnset     ST_TrendlineType = 0
	ST_TrendlineTypeExp       ST_TrendlineType = 1
	ST_TrendlineTypeLinear    ST_TrendlineType = 2
	ST_TrendlineTypeLog       ST_TrendlineType = 3
	ST_TrendlineTypeMovingAvg ST_TrendlineType = 4
	ST_TrendlineTypePoly      ST_TrendlineType = 5
	ST_TrendlineTypePower     ST_TrendlineType = 6
)

func (e ST_TrendlineType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TrendlineTypeUnset:
		attr.Value = ""
	case ST_TrendlineTypeExp:
		attr.Value = "exp"
	case ST_TrendlineTypeLinear:
		attr.Value = "linear"
	case ST_TrendlineTypeLog:
		attr.Value = "log"
	case ST_TrendlineTypeMovingAvg:
		attr.Value = "movingAvg"
	case ST_TrendlineTypePoly:
		attr.Value = "poly"
	case ST_TrendlineTypePower:
		attr.Value = "power"
	}
	return attr, nil
}
func (e *ST_TrendlineType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "exp":
		*e = 1
	case "linear":
		*e = 2
	case "log":
		*e = 3
	case "movingAvg":
		*e = 4
	case "poly":
		*e = 5
	case "power":
		*e = 6
	}
	return nil
}
func (m ST_TrendlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TrendlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "exp":
			*m = 1
		case "linear":
			*m = 2
		case "log":
			*m = 3
		case "movingAvg":
			*m = 4
		case "poly":
			*m = 5
		case "power":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TrendlineType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "exp"
	case 2:
		return "linear"
	case 3:
		return "log"
	case 4:
		return "movingAvg"
	case 5:
		return "poly"
	case 6:
		return "power"
	}
	return ""
}
func (m ST_TrendlineType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TrendlineType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ErrDir byte

const (
	ST_ErrDirUnset ST_ErrDir = 0
	ST_ErrDirX     ST_ErrDir = 1
	ST_ErrDirY     ST_ErrDir = 2
)

func (e ST_ErrDir) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ErrDirUnset:
		attr.Value = ""
	case ST_ErrDirX:
		attr.Value = "x"
	case ST_ErrDirY:
		attr.Value = "y"
	}
	return attr, nil
}
func (e *ST_ErrDir) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "x":
		*e = 1
	case "y":
		*e = 2
	}
	return nil
}
func (m ST_ErrDir) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ErrDir) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "x":
			*m = 1
		case "y":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_ErrDir) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "x"
	case 2:
		return "y"
	}
	return ""
}
func (m ST_ErrDir) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ErrDir) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ErrBarType byte

const (
	ST_ErrBarTypeUnset ST_ErrBarType = 0
	ST_ErrBarTypeBoth  ST_ErrBarType = 1
	ST_ErrBarTypeMinus ST_ErrBarType = 2
	ST_ErrBarTypePlus  ST_ErrBarType = 3
)

func (e ST_ErrBarType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ErrBarTypeUnset:
		attr.Value = ""
	case ST_ErrBarTypeBoth:
		attr.Value = "both"
	case ST_ErrBarTypeMinus:
		attr.Value = "minus"
	case ST_ErrBarTypePlus:
		attr.Value = "plus"
	}
	return attr, nil
}
func (e *ST_ErrBarType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "both":
		*e = 1
	case "minus":
		*e = 2
	case "plus":
		*e = 3
	}
	return nil
}
func (m ST_ErrBarType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ErrBarType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "both":
			*m = 1
		case "minus":
			*m = 2
		case "plus":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_ErrBarType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "both"
	case 2:
		return "minus"
	case 3:
		return "plus"
	}
	return ""
}
func (m ST_ErrBarType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ErrBarType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ErrValType byte

const (
	ST_ErrValTypeUnset      ST_ErrValType = 0
	ST_ErrValTypeCust       ST_ErrValType = 1
	ST_ErrValTypeFixedVal   ST_ErrValType = 2
	ST_ErrValTypePercentage ST_ErrValType = 3
	ST_ErrValTypeStdDev     ST_ErrValType = 4
	ST_ErrValTypeStdErr     ST_ErrValType = 5
)

func (e ST_ErrValType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ErrValTypeUnset:
		attr.Value = ""
	case ST_ErrValTypeCust:
		attr.Value = "cust"
	case ST_ErrValTypeFixedVal:
		attr.Value = "fixedVal"
	case ST_ErrValTypePercentage:
		attr.Value = "percentage"
	case ST_ErrValTypeStdDev:
		attr.Value = "stdDev"
	case ST_ErrValTypeStdErr:
		attr.Value = "stdErr"
	}
	return attr, nil
}
func (e *ST_ErrValType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cust":
		*e = 1
	case "fixedVal":
		*e = 2
	case "percentage":
		*e = 3
	case "stdDev":
		*e = 4
	case "stdErr":
		*e = 5
	}
	return nil
}
func (m ST_ErrValType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ErrValType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cust":
			*m = 1
		case "fixedVal":
			*m = 2
		case "percentage":
			*m = 3
		case "stdDev":
			*m = 4
		case "stdErr":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_ErrValType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cust"
	case 2:
		return "fixedVal"
	case 3:
		return "percentage"
	case 4:
		return "stdDev"
	case 5:
		return "stdErr"
	}
	return ""
}
func (m ST_ErrValType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ErrValType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Grouping byte

const (
	ST_GroupingUnset          ST_Grouping = 0
	ST_GroupingPercentStacked ST_Grouping = 1
	ST_GroupingStandard       ST_Grouping = 2
	ST_GroupingStacked        ST_Grouping = 3
)

func (e ST_Grouping) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_GroupingUnset:
		attr.Value = ""
	case ST_GroupingPercentStacked:
		attr.Value = "percentStacked"
	case ST_GroupingStandard:
		attr.Value = "standard"
	case ST_GroupingStacked:
		attr.Value = "stacked"
	}
	return attr, nil
}
func (e *ST_Grouping) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "percentStacked":
		*e = 1
	case "standard":
		*e = 2
	case "stacked":
		*e = 3
	}
	return nil
}
func (m ST_Grouping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_Grouping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "percentStacked":
			*m = 1
		case "standard":
			*m = 2
		case "stacked":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_Grouping) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "percentStacked"
	case 2:
		return "standard"
	case 3:
		return "stacked"
	}
	return ""
}
func (m ST_Grouping) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_Grouping) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ScatterStyle byte

const (
	ST_ScatterStyleUnset        ST_ScatterStyle = 0
	ST_ScatterStyleNone         ST_ScatterStyle = 1
	ST_ScatterStyleLine         ST_ScatterStyle = 2
	ST_ScatterStyleLineMarker   ST_ScatterStyle = 3
	ST_ScatterStyleMarker       ST_ScatterStyle = 4
	ST_ScatterStyleSmooth       ST_ScatterStyle = 5
	ST_ScatterStyleSmoothMarker ST_ScatterStyle = 6
)

func (e ST_ScatterStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ScatterStyleUnset:
		attr.Value = ""
	case ST_ScatterStyleNone:
		attr.Value = "none"
	case ST_ScatterStyleLine:
		attr.Value = "line"
	case ST_ScatterStyleLineMarker:
		attr.Value = "lineMarker"
	case ST_ScatterStyleMarker:
		attr.Value = "marker"
	case ST_ScatterStyleSmooth:
		attr.Value = "smooth"
	case ST_ScatterStyleSmoothMarker:
		attr.Value = "smoothMarker"
	}
	return attr, nil
}
func (e *ST_ScatterStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "line":
		*e = 2
	case "lineMarker":
		*e = 3
	case "marker":
		*e = 4
	case "smooth":
		*e = 5
	case "smoothMarker":
		*e = 6
	}
	return nil
}
func (m ST_ScatterStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ScatterStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "line":
			*m = 2
		case "lineMarker":
			*m = 3
		case "marker":
			*m = 4
		case "smooth":
			*m = 5
		case "smoothMarker":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_ScatterStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "line"
	case 3:
		return "lineMarker"
	case 4:
		return "marker"
	case 5:
		return "smooth"
	case 6:
		return "smoothMarker"
	}
	return ""
}
func (m ST_ScatterStyle) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ScatterStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RadarStyle byte

const (
	ST_RadarStyleUnset    ST_RadarStyle = 0
	ST_RadarStyleStandard ST_RadarStyle = 1
	ST_RadarStyleMarker   ST_RadarStyle = 2
	ST_RadarStyleFilled   ST_RadarStyle = 3
)

func (e ST_RadarStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RadarStyleUnset:
		attr.Value = ""
	case ST_RadarStyleStandard:
		attr.Value = "standard"
	case ST_RadarStyleMarker:
		attr.Value = "marker"
	case ST_RadarStyleFilled:
		attr.Value = "filled"
	}
	return attr, nil
}
func (e *ST_RadarStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "standard":
		*e = 1
	case "marker":
		*e = 2
	case "filled":
		*e = 3
	}
	return nil
}
func (m ST_RadarStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_RadarStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "standard":
			*m = 1
		case "marker":
			*m = 2
		case "filled":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_RadarStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "standard"
	case 2:
		return "marker"
	case 3:
		return "filled"
	}
	return ""
}
func (m ST_RadarStyle) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_RadarStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BarGrouping byte

const (
	ST_BarGroupingUnset          ST_BarGrouping = 0
	ST_BarGroupingPercentStacked ST_BarGrouping = 1
	ST_BarGroupingClustered      ST_BarGrouping = 2
	ST_BarGroupingStandard       ST_BarGrouping = 3
	ST_BarGroupingStacked        ST_BarGrouping = 4
)

func (e ST_BarGrouping) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BarGroupingUnset:
		attr.Value = ""
	case ST_BarGroupingPercentStacked:
		attr.Value = "percentStacked"
	case ST_BarGroupingClustered:
		attr.Value = "clustered"
	case ST_BarGroupingStandard:
		attr.Value = "standard"
	case ST_BarGroupingStacked:
		attr.Value = "stacked"
	}
	return attr, nil
}
func (e *ST_BarGrouping) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "percentStacked":
		*e = 1
	case "clustered":
		*e = 2
	case "standard":
		*e = 3
	case "stacked":
		*e = 4
	}
	return nil
}
func (m ST_BarGrouping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BarGrouping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "percentStacked":
			*m = 1
		case "clustered":
			*m = 2
		case "standard":
			*m = 3
		case "stacked":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_BarGrouping) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "percentStacked"
	case 2:
		return "clustered"
	case 3:
		return "standard"
	case 4:
		return "stacked"
	}
	return ""
}
func (m ST_BarGrouping) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BarGrouping) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BarDir byte

const (
	ST_BarDirUnset ST_BarDir = 0
	ST_BarDirBar   ST_BarDir = 1
	ST_BarDirCol   ST_BarDir = 2
)

func (e ST_BarDir) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BarDirUnset:
		attr.Value = ""
	case ST_BarDirBar:
		attr.Value = "bar"
	case ST_BarDirCol:
		attr.Value = "col"
	}
	return attr, nil
}
func (e *ST_BarDir) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bar":
		*e = 1
	case "col":
		*e = 2
	}
	return nil
}
func (m ST_BarDir) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BarDir) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bar":
			*m = 1
		case "col":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_BarDir) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bar"
	case 2:
		return "col"
	}
	return ""
}
func (m ST_BarDir) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BarDir) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Shape byte

const (
	ST_ShapeUnset        ST_Shape = 0
	ST_ShapeCone         ST_Shape = 1
	ST_ShapeConeToMax    ST_Shape = 2
	ST_ShapeBox          ST_Shape = 3
	ST_ShapeCylinder     ST_Shape = 4
	ST_ShapePyramid      ST_Shape = 5
	ST_ShapePyramidToMax ST_Shape = 6
)

func (e ST_Shape) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShapeUnset:
		attr.Value = ""
	case ST_ShapeCone:
		attr.Value = "cone"
	case ST_ShapeConeToMax:
		attr.Value = "coneToMax"
	case ST_ShapeBox:
		attr.Value = "box"
	case ST_ShapeCylinder:
		attr.Value = "cylinder"
	case ST_ShapePyramid:
		attr.Value = "pyramid"
	case ST_ShapePyramidToMax:
		attr.Value = "pyramidToMax"
	}
	return attr, nil
}
func (e *ST_Shape) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cone":
		*e = 1
	case "coneToMax":
		*e = 2
	case "box":
		*e = 3
	case "cylinder":
		*e = 4
	case "pyramid":
		*e = 5
	case "pyramidToMax":
		*e = 6
	}
	return nil
}
func (m ST_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cone":
			*m = 1
		case "coneToMax":
			*m = 2
		case "box":
			*m = 3
		case "cylinder":
			*m = 4
		case "pyramid":
			*m = 5
		case "pyramidToMax":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_Shape) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cone"
	case 2:
		return "coneToMax"
	case 3:
		return "box"
	case 4:
		return "cylinder"
	case 5:
		return "pyramid"
	case 6:
		return "pyramidToMax"
	}
	return ""
}
func (m ST_Shape) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_Shape) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OfPieType byte

const (
	ST_OfPieTypeUnset ST_OfPieType = 0
	ST_OfPieTypePie   ST_OfPieType = 1
	ST_OfPieTypeBar   ST_OfPieType = 2
)

func (e ST_OfPieType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OfPieTypeUnset:
		attr.Value = ""
	case ST_OfPieTypePie:
		attr.Value = "pie"
	case ST_OfPieTypeBar:
		attr.Value = "bar"
	}
	return attr, nil
}
func (e *ST_OfPieType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "pie":
		*e = 1
	case "bar":
		*e = 2
	}
	return nil
}
func (m ST_OfPieType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_OfPieType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "pie":
			*m = 1
		case "bar":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_OfPieType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "pie"
	case 2:
		return "bar"
	}
	return ""
}
func (m ST_OfPieType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_OfPieType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AxPos byte

const (
	ST_AxPosUnset ST_AxPos = 0
	ST_AxPosB     ST_AxPos = 1
	ST_AxPosL     ST_AxPos = 2
	ST_AxPosR     ST_AxPos = 3
	ST_AxPosT     ST_AxPos = 4
)

func (e ST_AxPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AxPosUnset:
		attr.Value = ""
	case ST_AxPosB:
		attr.Value = "b"
	case ST_AxPosL:
		attr.Value = "l"
	case ST_AxPosR:
		attr.Value = "r"
	case ST_AxPosT:
		attr.Value = "t"
	}
	return attr, nil
}
func (e *ST_AxPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "b":
		*e = 1
	case "l":
		*e = 2
	case "r":
		*e = 3
	case "t":
		*e = 4
	}
	return nil
}
func (m ST_AxPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AxPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "b":
			*m = 1
		case "l":
			*m = 2
		case "r":
			*m = 3
		case "t":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_AxPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "b"
	case 2:
		return "l"
	case 3:
		return "r"
	case 4:
		return "t"
	}
	return ""
}
func (m ST_AxPos) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AxPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Crosses byte

const (
	ST_CrossesUnset    ST_Crosses = 0
	ST_CrossesAutoZero ST_Crosses = 1
	ST_CrossesMax      ST_Crosses = 2
	ST_CrossesMin      ST_Crosses = 3
)

func (e ST_Crosses) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CrossesUnset:
		attr.Value = ""
	case ST_CrossesAutoZero:
		attr.Value = "autoZero"
	case ST_CrossesMax:
		attr.Value = "max"
	case ST_CrossesMin:
		attr.Value = "min"
	}
	return attr, nil
}
func (e *ST_Crosses) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "autoZero":
		*e = 1
	case "max":
		*e = 2
	case "min":
		*e = 3
	}
	return nil
}
func (m ST_Crosses) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_Crosses) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "autoZero":
			*m = 1
		case "max":
			*m = 2
		case "min":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_Crosses) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "autoZero"
	case 2:
		return "max"
	case 3:
		return "min"
	}
	return ""
}
func (m ST_Crosses) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_Crosses) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CrossBetween byte

const (
	ST_CrossBetweenUnset   ST_CrossBetween = 0
	ST_CrossBetweenBetween ST_CrossBetween = 1
	ST_CrossBetweenMidCat  ST_CrossBetween = 2
)

func (e ST_CrossBetween) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CrossBetweenUnset:
		attr.Value = ""
	case ST_CrossBetweenBetween:
		attr.Value = "between"
	case ST_CrossBetweenMidCat:
		attr.Value = "midCat"
	}
	return attr, nil
}
func (e *ST_CrossBetween) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "between":
		*e = 1
	case "midCat":
		*e = 2
	}
	return nil
}
func (m ST_CrossBetween) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_CrossBetween) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "between":
			*m = 1
		case "midCat":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_CrossBetween) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "between"
	case 2:
		return "midCat"
	}
	return ""
}
func (m ST_CrossBetween) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_CrossBetween) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TickMark byte

const (
	ST_TickMarkUnset ST_TickMark = 0
	ST_TickMarkCross ST_TickMark = 1
	ST_TickMarkIn    ST_TickMark = 2
	ST_TickMarkNone  ST_TickMark = 3
	ST_TickMarkOut   ST_TickMark = 4
)

func (e ST_TickMark) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TickMarkUnset:
		attr.Value = ""
	case ST_TickMarkCross:
		attr.Value = "cross"
	case ST_TickMarkIn:
		attr.Value = "in"
	case ST_TickMarkNone:
		attr.Value = "none"
	case ST_TickMarkOut:
		attr.Value = "out"
	}
	return attr, nil
}
func (e *ST_TickMark) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cross":
		*e = 1
	case "in":
		*e = 2
	case "none":
		*e = 3
	case "out":
		*e = 4
	}
	return nil
}
func (m ST_TickMark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TickMark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cross":
			*m = 1
		case "in":
			*m = 2
		case "none":
			*m = 3
		case "out":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TickMark) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cross"
	case 2:
		return "in"
	case 3:
		return "none"
	case 4:
		return "out"
	}
	return ""
}
func (m ST_TickMark) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TickMark) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TickLblPos byte

const (
	ST_TickLblPosUnset  ST_TickLblPos = 0
	ST_TickLblPosHigh   ST_TickLblPos = 1
	ST_TickLblPosLow    ST_TickLblPos = 2
	ST_TickLblPosNextTo ST_TickLblPos = 3
	ST_TickLblPosNone   ST_TickLblPos = 4
)

func (e ST_TickLblPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TickLblPosUnset:
		attr.Value = ""
	case ST_TickLblPosHigh:
		attr.Value = "high"
	case ST_TickLblPosLow:
		attr.Value = "low"
	case ST_TickLblPosNextTo:
		attr.Value = "nextTo"
	case ST_TickLblPosNone:
		attr.Value = "none"
	}
	return attr, nil
}
func (e *ST_TickLblPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "high":
		*e = 1
	case "low":
		*e = 2
	case "nextTo":
		*e = 3
	case "none":
		*e = 4
	}
	return nil
}
func (m ST_TickLblPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TickLblPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "high":
			*m = 1
		case "low":
			*m = 2
		case "nextTo":
			*m = 3
		case "none":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TickLblPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "high"
	case 2:
		return "low"
	case 3:
		return "nextTo"
	case 4:
		return "none"
	}
	return ""
}
func (m ST_TickLblPos) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TickLblPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TimeUnit byte

const (
	ST_TimeUnitUnset  ST_TimeUnit = 0
	ST_TimeUnitDays   ST_TimeUnit = 1
	ST_TimeUnitMonths ST_TimeUnit = 2
	ST_TimeUnitYears  ST_TimeUnit = 3
)

func (e ST_TimeUnit) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TimeUnitUnset:
		attr.Value = ""
	case ST_TimeUnitDays:
		attr.Value = "days"
	case ST_TimeUnitMonths:
		attr.Value = "months"
	case ST_TimeUnitYears:
		attr.Value = "years"
	}
	return attr, nil
}
func (e *ST_TimeUnit) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "days":
		*e = 1
	case "months":
		*e = 2
	case "years":
		*e = 3
	}
	return nil
}
func (m ST_TimeUnit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TimeUnit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "days":
			*m = 1
		case "months":
			*m = 2
		case "years":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TimeUnit) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "days"
	case 2:
		return "months"
	case 3:
		return "years"
	}
	return ""
}
func (m ST_TimeUnit) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TimeUnit) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BuiltInUnit byte

const (
	ST_BuiltInUnitUnset            ST_BuiltInUnit = 0
	ST_BuiltInUnitHundreds         ST_BuiltInUnit = 1
	ST_BuiltInUnitThousands        ST_BuiltInUnit = 2
	ST_BuiltInUnitTenThousands     ST_BuiltInUnit = 3
	ST_BuiltInUnitHundredThousands ST_BuiltInUnit = 4
	ST_BuiltInUnitMillions         ST_BuiltInUnit = 5
	ST_BuiltInUnitTenMillions      ST_BuiltInUnit = 6
	ST_BuiltInUnitHundredMillions  ST_BuiltInUnit = 7
	ST_BuiltInUnitBillions         ST_BuiltInUnit = 8
	ST_BuiltInUnitTrillions        ST_BuiltInUnit = 9
)

func (e ST_BuiltInUnit) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BuiltInUnitUnset:
		attr.Value = ""
	case ST_BuiltInUnitHundreds:
		attr.Value = "hundreds"
	case ST_BuiltInUnitThousands:
		attr.Value = "thousands"
	case ST_BuiltInUnitTenThousands:
		attr.Value = "tenThousands"
	case ST_BuiltInUnitHundredThousands:
		attr.Value = "hundredThousands"
	case ST_BuiltInUnitMillions:
		attr.Value = "millions"
	case ST_BuiltInUnitTenMillions:
		attr.Value = "tenMillions"
	case ST_BuiltInUnitHundredMillions:
		attr.Value = "hundredMillions"
	case ST_BuiltInUnitBillions:
		attr.Value = "billions"
	case ST_BuiltInUnitTrillions:
		attr.Value = "trillions"
	}
	return attr, nil
}
func (e *ST_BuiltInUnit) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "hundreds":
		*e = 1
	case "thousands":
		*e = 2
	case "tenThousands":
		*e = 3
	case "hundredThousands":
		*e = 4
	case "millions":
		*e = 5
	case "tenMillions":
		*e = 6
	case "hundredMillions":
		*e = 7
	case "billions":
		*e = 8
	case "trillions":
		*e = 9
	}
	return nil
}
func (m ST_BuiltInUnit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BuiltInUnit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "hundreds":
			*m = 1
		case "thousands":
			*m = 2
		case "tenThousands":
			*m = 3
		case "hundredThousands":
			*m = 4
		case "millions":
			*m = 5
		case "tenMillions":
			*m = 6
		case "hundredMillions":
			*m = 7
		case "billions":
			*m = 8
		case "trillions":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_BuiltInUnit) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "hundreds"
	case 2:
		return "thousands"
	case 3:
		return "tenThousands"
	case 4:
		return "hundredThousands"
	case 5:
		return "millions"
	case 6:
		return "tenMillions"
	case 7:
		return "hundredMillions"
	case 8:
		return "billions"
	case 9:
		return "trillions"
	}
	return ""
}
func (m ST_BuiltInUnit) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BuiltInUnit) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PictureFormat byte

const (
	ST_PictureFormatUnset      ST_PictureFormat = 0
	ST_PictureFormatStretch    ST_PictureFormat = 1
	ST_PictureFormatStack      ST_PictureFormat = 2
	ST_PictureFormatStackScale ST_PictureFormat = 3
)

func (e ST_PictureFormat) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PictureFormatUnset:
		attr.Value = ""
	case ST_PictureFormatStretch:
		attr.Value = "stretch"
	case ST_PictureFormatStack:
		attr.Value = "stack"
	case ST_PictureFormatStackScale:
		attr.Value = "stackScale"
	}
	return attr, nil
}
func (e *ST_PictureFormat) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "stretch":
		*e = 1
	case "stack":
		*e = 2
	case "stackScale":
		*e = 3
	}
	return nil
}
func (m ST_PictureFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PictureFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "stretch":
			*m = 1
		case "stack":
			*m = 2
		case "stackScale":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PictureFormat) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "stretch"
	case 2:
		return "stack"
	case 3:
		return "stackScale"
	}
	return ""
}
func (m ST_PictureFormat) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PictureFormat) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Orientation byte

const (
	ST_OrientationUnset  ST_Orientation = 0
	ST_OrientationMaxMin ST_Orientation = 1
	ST_OrientationMinMax ST_Orientation = 2
)

func (e ST_Orientation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OrientationUnset:
		attr.Value = ""
	case ST_OrientationMaxMin:
		attr.Value = "maxMin"
	case ST_OrientationMinMax:
		attr.Value = "minMax"
	}
	return attr, nil
}
func (e *ST_Orientation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "maxMin":
		*e = 1
	case "minMax":
		*e = 2
	}
	return nil
}
func (m ST_Orientation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_Orientation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "maxMin":
			*m = 1
		case "minMax":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_Orientation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "maxMin"
	case 2:
		return "minMax"
	}
	return ""
}
func (m ST_Orientation) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_Orientation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LegendPos byte

const (
	ST_LegendPosUnset ST_LegendPos = 0
	ST_LegendPosB     ST_LegendPos = 1
	ST_LegendPosTr    ST_LegendPos = 2
	ST_LegendPosL     ST_LegendPos = 3
	ST_LegendPosR     ST_LegendPos = 4
	ST_LegendPosT     ST_LegendPos = 5
)

func (e ST_LegendPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LegendPosUnset:
		attr.Value = ""
	case ST_LegendPosB:
		attr.Value = "b"
	case ST_LegendPosTr:
		attr.Value = "tr"
	case ST_LegendPosL:
		attr.Value = "l"
	case ST_LegendPosR:
		attr.Value = "r"
	case ST_LegendPosT:
		attr.Value = "t"
	}
	return attr, nil
}
func (e *ST_LegendPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "b":
		*e = 1
	case "tr":
		*e = 2
	case "l":
		*e = 3
	case "r":
		*e = 4
	case "t":
		*e = 5
	}
	return nil
}
func (m ST_LegendPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LegendPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "b":
			*m = 1
		case "tr":
			*m = 2
		case "l":
			*m = 3
		case "r":
			*m = 4
		case "t":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_LegendPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "b"
	case 2:
		return "tr"
	case 3:
		return "l"
	case 4:
		return "r"
	case 5:
		return "t"
	}
	return ""
}
func (m ST_LegendPos) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LegendPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DispBlanksAs byte

const (
	ST_DispBlanksAsUnset ST_DispBlanksAs = 0
	ST_DispBlanksAsSpan  ST_DispBlanksAs = 1
	ST_DispBlanksAsGap   ST_DispBlanksAs = 2
	ST_DispBlanksAsZero  ST_DispBlanksAs = 3
)

func (e ST_DispBlanksAs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DispBlanksAsUnset:
		attr.Value = ""
	case ST_DispBlanksAsSpan:
		attr.Value = "span"
	case ST_DispBlanksAsGap:
		attr.Value = "gap"
	case ST_DispBlanksAsZero:
		attr.Value = "zero"
	}
	return attr, nil
}
func (e *ST_DispBlanksAs) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "span":
		*e = 1
	case "gap":
		*e = 2
	case "zero":
		*e = 3
	}
	return nil
}
func (m ST_DispBlanksAs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_DispBlanksAs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "span":
			*m = 1
		case "gap":
			*m = 2
		case "zero":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_DispBlanksAs) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "span"
	case 2:
		return "gap"
	case 3:
		return "zero"
	}
	return ""
}
func (m ST_DispBlanksAs) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_DispBlanksAs) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageSetupOrientation byte

const (
	ST_PageSetupOrientationUnset     ST_PageSetupOrientation = 0
	ST_PageSetupOrientationDefault   ST_PageSetupOrientation = 1
	ST_PageSetupOrientationPortrait  ST_PageSetupOrientation = 2
	ST_PageSetupOrientationLandscape ST_PageSetupOrientation = 3
)

func (e ST_PageSetupOrientation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageSetupOrientationUnset:
		attr.Value = ""
	case ST_PageSetupOrientationDefault:
		attr.Value = "default"
	case ST_PageSetupOrientationPortrait:
		attr.Value = "portrait"
	case ST_PageSetupOrientationLandscape:
		attr.Value = "landscape"
	}
	return attr, nil
}
func (e *ST_PageSetupOrientation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "default":
		*e = 1
	case "portrait":
		*e = 2
	case "landscape":
		*e = 3
	}
	return nil
}
func (m ST_PageSetupOrientation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PageSetupOrientation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "default":
			*m = 1
		case "portrait":
			*m = 2
		case "landscape":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PageSetupOrientation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "default"
	case 2:
		return "portrait"
	case 3:
		return "landscape"
	}
	return ""
}
func (m ST_PageSetupOrientation) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PageSetupOrientation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Boolean", NewCT_Boolean)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Double", NewCT_Double)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_UnsignedInt", NewCT_UnsignedInt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RelId", NewCT_RelId)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Extension", NewCT_Extension)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ExtensionList", NewCT_ExtensionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_NumVal", NewCT_NumVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_NumData", NewCT_NumData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_NumRef", NewCT_NumRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_NumDataSource", NewCT_NumDataSource)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_StrVal", NewCT_StrVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_StrData", NewCT_StrData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_StrRef", NewCT_StrRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Tx", NewCT_Tx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TextLanguageID", NewCT_TextLanguageID)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Lvl", NewCT_Lvl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_MultiLvlStrData", NewCT_MultiLvlStrData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_MultiLvlStrRef", NewCT_MultiLvlStrRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_AxDataSource", NewCT_AxDataSource)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SerTx", NewCT_SerTx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LayoutTarget", NewCT_LayoutTarget)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LayoutMode", NewCT_LayoutMode)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ManualLayout", NewCT_ManualLayout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Layout", NewCT_Layout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Title", NewCT_Title)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RotX", NewCT_RotX)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_HPercent", NewCT_HPercent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RotY", NewCT_RotY)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DepthPercent", NewCT_DepthPercent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Perspective", NewCT_Perspective)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_View3D", NewCT_View3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Surface", NewCT_Surface)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Thickness", NewCT_Thickness)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DTable", NewCT_DTable)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_GapAmount", NewCT_GapAmount)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Overlap", NewCT_Overlap)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BubbleScale", NewCT_BubbleScale)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SizeRepresents", NewCT_SizeRepresents)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_FirstSliceAng", NewCT_FirstSliceAng)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_HoleSize", NewCT_HoleSize)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SplitType", NewCT_SplitType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_CustSplit", NewCT_CustSplit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SecondPieSize", NewCT_SecondPieSize)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_NumFmt", NewCT_NumFmt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LblAlgn", NewCT_LblAlgn)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DLblPos", NewCT_DLblPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DLbl", NewCT_DLbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DLbls", NewCT_DLbls)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_MarkerStyle", NewCT_MarkerStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_MarkerSize", NewCT_MarkerSize)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Marker", NewCT_Marker)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DPt", NewCT_DPt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TrendlineType", NewCT_TrendlineType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Order", NewCT_Order)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Period", NewCT_Period)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TrendlineLbl", NewCT_TrendlineLbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Trendline", NewCT_Trendline)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ErrDir", NewCT_ErrDir)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ErrBarType", NewCT_ErrBarType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ErrValType", NewCT_ErrValType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ErrBars", NewCT_ErrBars)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_UpDownBar", NewCT_UpDownBar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_UpDownBars", NewCT_UpDownBars)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LineSer", NewCT_LineSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ScatterSer", NewCT_ScatterSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RadarSer", NewCT_RadarSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BarSer", NewCT_BarSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_AreaSer", NewCT_AreaSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PieSer", NewCT_PieSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BubbleSer", NewCT_BubbleSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SurfaceSer", NewCT_SurfaceSer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Grouping", NewCT_Grouping)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ChartLines", NewCT_ChartLines)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LineChart", NewCT_LineChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Line3DChart", NewCT_Line3DChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_StockChart", NewCT_StockChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ScatterStyle", NewCT_ScatterStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ScatterChart", NewCT_ScatterChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RadarStyle", NewCT_RadarStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_RadarChart", NewCT_RadarChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BarGrouping", NewCT_BarGrouping)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BarDir", NewCT_BarDir)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Shape", NewCT_Shape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BarChart", NewCT_BarChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Bar3DChart", NewCT_Bar3DChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_AreaChart", NewCT_AreaChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Area3DChart", NewCT_Area3DChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PieChart", NewCT_PieChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Pie3DChart", NewCT_Pie3DChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DoughnutChart", NewCT_DoughnutChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_OfPieType", NewCT_OfPieType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_OfPieChart", NewCT_OfPieChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BubbleChart", NewCT_BubbleChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BandFmt", NewCT_BandFmt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BandFmts", NewCT_BandFmts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SurfaceChart", NewCT_SurfaceChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Surface3DChart", NewCT_Surface3DChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_AxPos", NewCT_AxPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Crosses", NewCT_Crosses)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_CrossBetween", NewCT_CrossBetween)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TickMark", NewCT_TickMark)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TickLblPos", NewCT_TickLblPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Skip", NewCT_Skip)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_TimeUnit", NewCT_TimeUnit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_AxisUnit", NewCT_AxisUnit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_BuiltInUnit", NewCT_BuiltInUnit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PictureFormat", NewCT_PictureFormat)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PictureStackUnit", NewCT_PictureStackUnit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PictureOptions", NewCT_PictureOptions)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DispUnitsLbl", NewCT_DispUnitsLbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DispUnits", NewCT_DispUnits)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Orientation", NewCT_Orientation)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LogBase", NewCT_LogBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Scaling", NewCT_Scaling)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LblOffset", NewCT_LblOffset)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_CatAx", NewCT_CatAx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DateAx", NewCT_DateAx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_SerAx", NewCT_SerAx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ValAx", NewCT_ValAx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PlotArea", NewCT_PlotArea)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PivotFmt", NewCT_PivotFmt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PivotFmts", NewCT_PivotFmts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LegendPos", NewCT_LegendPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_LegendEntry", NewCT_LegendEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Legend", NewCT_Legend)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_DispBlanksAs", NewCT_DispBlanksAs)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Chart", NewCT_Chart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Style", NewCT_Style)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PivotSource", NewCT_PivotSource)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_Protection", NewCT_Protection)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_HeaderFooter", NewCT_HeaderFooter)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PageMargins", NewCT_PageMargins)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ExternalData", NewCT_ExternalData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PageSetup", NewCT_PageSetup)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_PrintSettings", NewCT_PrintSettings)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "CT_ChartSpace", NewCT_ChartSpace)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "chartSpace", NewChartSpace)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "userShapes", NewUserShapes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "chart", NewChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_DLblShared", NewEG_DLblShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "Group_DLbl", NewGroup_DLbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "Group_DLbls", NewGroup_DLbls)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_SerShared", NewEG_SerShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_LineChartShared", NewEG_LineChartShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_BarChartShared", NewEG_BarChartShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_AreaChartShared", NewEG_AreaChartShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_PieChartShared", NewEG_PieChartShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_SurfaceChartShared", NewEG_SurfaceChartShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_AxShared", NewEG_AxShared)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chart", "EG_LegendEntryData", NewEG_LegendEntryData)
}
