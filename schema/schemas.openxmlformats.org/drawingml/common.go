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
	"regexp"
	"strconv"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

const ST_TextBulletSizePercentPattern = `0*((2[5-9])|([3-9][0-9])|([1-3][0-9][0-9])|400)%`

var ST_TextBulletSizePercentPatternRe = regexp.MustCompile(ST_TextBulletSizePercentPattern)

func ParseUnionST_Percentage(s string) (ST_Percentage, error) {
	r := ST_Percentage{}
	if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
		r.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_PercentageDecimal = &v32
	}
	return r, nil
}

func ParseUnionST_PositivePercentage(s string) (ST_PositivePercentage, error) {
	r := ST_PositivePercentage{}
	if sharedTypes.ST_PositivePercentagePatternRe.MatchString(s) {
		r.ST_PositivePercentage = &ST_Percentage{}
		r.ST_PositivePercentage.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_PositivePercentageDecimal = &v32
	}
	return r, nil
}

func ParseUnionST_PositiveFixedPercentage(s string) (ST_PositiveFixedPercentage, error) {
	r := ST_PositiveFixedPercentage{}
	if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
		r.ST_PositiveFixedPercentage = &ST_Percentage{}
		r.ST_PositiveFixedPercentage.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_PositiveFixedPercentageDecimal = &v32
	}
	return r, nil
}

func ParseUnionST_FixedPercentage(s string) (ST_FixedPercentage, error) {
	r := ST_FixedPercentage{}
	if sharedTypes.ST_FixedPercentagePatternRe.MatchString(s) {
		r.ST_FixedPercentage = &ST_Percentage{}
		r.ST_FixedPercentage.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_FixedPercentageDecimal = &v32
	}
	return r, nil
}

func ParseUnionST_Coordinate(s string) (ST_Coordinate, error) {
	r := ST_Coordinate{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		r.ST_CoordinateUnqualified = &v
	}
	return r, nil
}
func ParseUnionST_Coordinate32(s string) (ST_Coordinate32, error) {
	r := ST_Coordinate32{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_Coordinate32Unqualified = &v32
	}
	return r, nil
}
func ParseUnionST_AdjCoordinate(s string) (ST_AdjCoordinate, error) {
	r := ST_AdjCoordinate{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_Coordinate = &ST_Coordinate{}
		r.ST_Coordinate.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			// geom guide name can be anything
			r.ST_GeomGuideName = &s
		} else {
			r.ST_Coordinate.ST_CoordinateUnqualified = &v
		}
	}
	return r, nil
}

func ParseUnionST_AdjAngle(s string) (ST_AdjAngle, error) {
	r := ST_AdjAngle{}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// geom guide name can be anything
		r.ST_GeomGuideName = &s
	} else {
		v32 := int32(v)
		r.ST_Angle = &v32
	}
	return r, nil
}

func ParseUnionST_OnOff(s string) (sharedTypes.ST_OnOff, error) {
	return sharedTypes.ST_OnOff{}, nil
}

func ParseUnionST_TextPoint(s string) (ST_TextPoint, error) {
	r := ST_TextPoint{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_TextPointUnqualified = &v32

	}
	return r, nil
}

func ParseUnionST_AnimationDgmBuildType(s string) (ST_AnimationDgmBuildType, error) {
	r := ST_AnimationDgmBuildType{}
	switch s {
	case "allAtOnce":
		r.ST_AnimationBuildType = ST_AnimationBuildTypeAllAtOnce
	case "one":
		r.ST_AnimationDgmOnlyBuildType = ST_AnimationDgmOnlyBuildTypeOne
	case "lvlOne":
		r.ST_AnimationDgmOnlyBuildType = ST_AnimationDgmOnlyBuildTypeLvlOne
	case "lvlAtOnce":
		r.ST_AnimationDgmOnlyBuildType = ST_AnimationDgmOnlyBuildTypeLvlAtOnce
	}
	return r, nil
}

func ParseUnionST_AnimationChartBuildType(s string) (ST_AnimationChartBuildType, error) {
	r := ST_AnimationChartBuildType{}
	switch s {
	case "allAtOnce":
		r.ST_AnimationBuildType = ST_AnimationBuildTypeAllAtOnce
	case "series":
		r.ST_AnimationChartOnlyBuildType = ST_AnimationChartOnlyBuildTypeSeries
	case "category":
		r.ST_AnimationChartOnlyBuildType = ST_AnimationChartOnlyBuildTypeCategory
	case "seriesEl":
		r.ST_AnimationChartOnlyBuildType = ST_AnimationChartOnlyBuildTypeSeriesEl
	case "categoryEl":
		r.ST_AnimationChartOnlyBuildType = ST_AnimationChartOnlyBuildTypeCategoryEl
	}
	return r, nil
}

func ParseUnionST_TextSpacingPercentOrPercentString(s string) (ST_TextSpacingPercentOrPercentString, error) {
	r := ST_TextSpacingPercentOrPercentString{}
	if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
		r.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_TextSpacingPercent = &v32
	}
	return r, nil
}
func ParseUnionST_TextFontScalePercentOrPercentString(s string) (ST_TextFontScalePercentOrPercentString, error) {
	r := ST_TextFontScalePercentOrPercentString{}
	if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
		r.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int", err)
		}
		v32 := int32(v)
		r.ST_TextFontScalePercent = &v32
	}
	return r, nil
}

func ParseStdlibTime(s string) (time.Time, error) {
	// TODO: implement this
	return time.Time{}, nil
}

type Any interface {
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}

type ST_FontCollectionIndex byte

const (
	ST_FontCollectionIndexUnset ST_FontCollectionIndex = 0
	ST_FontCollectionIndexMajor ST_FontCollectionIndex = 1
	ST_FontCollectionIndexMinor ST_FontCollectionIndex = 2
	ST_FontCollectionIndexNone  ST_FontCollectionIndex = 3
)

func (e ST_FontCollectionIndex) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FontCollectionIndexUnset:
		attr.Value = ""
	case ST_FontCollectionIndexMajor:
		attr.Value = "major"
	case ST_FontCollectionIndexMinor:
		attr.Value = "minor"
	case ST_FontCollectionIndexNone:
		attr.Value = "none"
	}
	return attr, nil
}
func (e *ST_FontCollectionIndex) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "major":
		*e = 1
	case "minor":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}
func (m ST_FontCollectionIndex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_FontCollectionIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "major":
			*m = 1
		case "minor":
			*m = 2
		case "none":
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
func (m ST_FontCollectionIndex) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "major"
	case 2:
		return "minor"
	case 3:
		return "none"
	}
	return ""
}
func (m ST_FontCollectionIndex) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_FontCollectionIndex) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ColorSchemeIndex byte

const (
	ST_ColorSchemeIndexUnset    ST_ColorSchemeIndex = 0
	ST_ColorSchemeIndexDk1      ST_ColorSchemeIndex = 1
	ST_ColorSchemeIndexLt1      ST_ColorSchemeIndex = 2
	ST_ColorSchemeIndexDk2      ST_ColorSchemeIndex = 3
	ST_ColorSchemeIndexLt2      ST_ColorSchemeIndex = 4
	ST_ColorSchemeIndexAccent1  ST_ColorSchemeIndex = 5
	ST_ColorSchemeIndexAccent2  ST_ColorSchemeIndex = 6
	ST_ColorSchemeIndexAccent3  ST_ColorSchemeIndex = 7
	ST_ColorSchemeIndexAccent4  ST_ColorSchemeIndex = 8
	ST_ColorSchemeIndexAccent5  ST_ColorSchemeIndex = 9
	ST_ColorSchemeIndexAccent6  ST_ColorSchemeIndex = 10
	ST_ColorSchemeIndexHlink    ST_ColorSchemeIndex = 11
	ST_ColorSchemeIndexFolHlink ST_ColorSchemeIndex = 12
)

func (e ST_ColorSchemeIndex) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ColorSchemeIndexUnset:
		attr.Value = ""
	case ST_ColorSchemeIndexDk1:
		attr.Value = "dk1"
	case ST_ColorSchemeIndexLt1:
		attr.Value = "lt1"
	case ST_ColorSchemeIndexDk2:
		attr.Value = "dk2"
	case ST_ColorSchemeIndexLt2:
		attr.Value = "lt2"
	case ST_ColorSchemeIndexAccent1:
		attr.Value = "accent1"
	case ST_ColorSchemeIndexAccent2:
		attr.Value = "accent2"
	case ST_ColorSchemeIndexAccent3:
		attr.Value = "accent3"
	case ST_ColorSchemeIndexAccent4:
		attr.Value = "accent4"
	case ST_ColorSchemeIndexAccent5:
		attr.Value = "accent5"
	case ST_ColorSchemeIndexAccent6:
		attr.Value = "accent6"
	case ST_ColorSchemeIndexHlink:
		attr.Value = "hlink"
	case ST_ColorSchemeIndexFolHlink:
		attr.Value = "folHlink"
	}
	return attr, nil
}
func (e *ST_ColorSchemeIndex) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "dk1":
		*e = 1
	case "lt1":
		*e = 2
	case "dk2":
		*e = 3
	case "lt2":
		*e = 4
	case "accent1":
		*e = 5
	case "accent2":
		*e = 6
	case "accent3":
		*e = 7
	case "accent4":
		*e = 8
	case "accent5":
		*e = 9
	case "accent6":
		*e = 10
	case "hlink":
		*e = 11
	case "folHlink":
		*e = 12
	}
	return nil
}
func (m ST_ColorSchemeIndex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ColorSchemeIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "dk1":
			*m = 1
		case "lt1":
			*m = 2
		case "dk2":
			*m = 3
		case "lt2":
			*m = 4
		case "accent1":
			*m = 5
		case "accent2":
			*m = 6
		case "accent3":
			*m = 7
		case "accent4":
			*m = 8
		case "accent5":
			*m = 9
		case "accent6":
			*m = 10
		case "hlink":
			*m = 11
		case "folHlink":
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
func (m ST_ColorSchemeIndex) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "dk1"
	case 2:
		return "lt1"
	case 3:
		return "dk2"
	case 4:
		return "lt2"
	case 5:
		return "accent1"
	case 6:
		return "accent2"
	case 7:
		return "accent3"
	case 8:
		return "accent4"
	case 9:
		return "accent5"
	case 10:
		return "accent6"
	case 11:
		return "hlink"
	case 12:
		return "folHlink"
	}
	return ""
}
func (m ST_ColorSchemeIndex) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ColorSchemeIndex) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SystemColorVal byte

const (
	ST_SystemColorValUnset                   ST_SystemColorVal = 0
	ST_SystemColorValScrollBar               ST_SystemColorVal = 1
	ST_SystemColorValBackground              ST_SystemColorVal = 2
	ST_SystemColorValActiveCaption           ST_SystemColorVal = 3
	ST_SystemColorValInactiveCaption         ST_SystemColorVal = 4
	ST_SystemColorValMenu                    ST_SystemColorVal = 5
	ST_SystemColorValWindow                  ST_SystemColorVal = 6
	ST_SystemColorValWindowFrame             ST_SystemColorVal = 7
	ST_SystemColorValMenuText                ST_SystemColorVal = 8
	ST_SystemColorValWindowText              ST_SystemColorVal = 9
	ST_SystemColorValCaptionText             ST_SystemColorVal = 10
	ST_SystemColorValActiveBorder            ST_SystemColorVal = 11
	ST_SystemColorValInactiveBorder          ST_SystemColorVal = 12
	ST_SystemColorValAppWorkspace            ST_SystemColorVal = 13
	ST_SystemColorValHighlight               ST_SystemColorVal = 14
	ST_SystemColorValHighlightText           ST_SystemColorVal = 15
	ST_SystemColorValBtnFace                 ST_SystemColorVal = 16
	ST_SystemColorValBtnShadow               ST_SystemColorVal = 17
	ST_SystemColorValGrayText                ST_SystemColorVal = 18
	ST_SystemColorValBtnText                 ST_SystemColorVal = 19
	ST_SystemColorValInactiveCaptionText     ST_SystemColorVal = 20
	ST_SystemColorValBtnHighlight            ST_SystemColorVal = 21
	ST_SystemColorVal3dDkShadow              ST_SystemColorVal = 22
	ST_SystemColorVal3dLight                 ST_SystemColorVal = 23
	ST_SystemColorValInfoText                ST_SystemColorVal = 24
	ST_SystemColorValInfoBk                  ST_SystemColorVal = 25
	ST_SystemColorValHotLight                ST_SystemColorVal = 26
	ST_SystemColorValGradientActiveCaption   ST_SystemColorVal = 27
	ST_SystemColorValGradientInactiveCaption ST_SystemColorVal = 28
	ST_SystemColorValMenuHighlight           ST_SystemColorVal = 29
	ST_SystemColorValMenuBar                 ST_SystemColorVal = 30
)

func (e ST_SystemColorVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SystemColorValUnset:
		attr.Value = ""
	case ST_SystemColorValScrollBar:
		attr.Value = "scrollBar"
	case ST_SystemColorValBackground:
		attr.Value = "background"
	case ST_SystemColorValActiveCaption:
		attr.Value = "activeCaption"
	case ST_SystemColorValInactiveCaption:
		attr.Value = "inactiveCaption"
	case ST_SystemColorValMenu:
		attr.Value = "menu"
	case ST_SystemColorValWindow:
		attr.Value = "window"
	case ST_SystemColorValWindowFrame:
		attr.Value = "windowFrame"
	case ST_SystemColorValMenuText:
		attr.Value = "menuText"
	case ST_SystemColorValWindowText:
		attr.Value = "windowText"
	case ST_SystemColorValCaptionText:
		attr.Value = "captionText"
	case ST_SystemColorValActiveBorder:
		attr.Value = "activeBorder"
	case ST_SystemColorValInactiveBorder:
		attr.Value = "inactiveBorder"
	case ST_SystemColorValAppWorkspace:
		attr.Value = "appWorkspace"
	case ST_SystemColorValHighlight:
		attr.Value = "highlight"
	case ST_SystemColorValHighlightText:
		attr.Value = "highlightText"
	case ST_SystemColorValBtnFace:
		attr.Value = "btnFace"
	case ST_SystemColorValBtnShadow:
		attr.Value = "btnShadow"
	case ST_SystemColorValGrayText:
		attr.Value = "grayText"
	case ST_SystemColorValBtnText:
		attr.Value = "btnText"
	case ST_SystemColorValInactiveCaptionText:
		attr.Value = "inactiveCaptionText"
	case ST_SystemColorValBtnHighlight:
		attr.Value = "btnHighlight"
	case ST_SystemColorVal3dDkShadow:
		attr.Value = "3dDkShadow"
	case ST_SystemColorVal3dLight:
		attr.Value = "3dLight"
	case ST_SystemColorValInfoText:
		attr.Value = "infoText"
	case ST_SystemColorValInfoBk:
		attr.Value = "infoBk"
	case ST_SystemColorValHotLight:
		attr.Value = "hotLight"
	case ST_SystemColorValGradientActiveCaption:
		attr.Value = "gradientActiveCaption"
	case ST_SystemColorValGradientInactiveCaption:
		attr.Value = "gradientInactiveCaption"
	case ST_SystemColorValMenuHighlight:
		attr.Value = "menuHighlight"
	case ST_SystemColorValMenuBar:
		attr.Value = "menuBar"
	}
	return attr, nil
}
func (e *ST_SystemColorVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "scrollBar":
		*e = 1
	case "background":
		*e = 2
	case "activeCaption":
		*e = 3
	case "inactiveCaption":
		*e = 4
	case "menu":
		*e = 5
	case "window":
		*e = 6
	case "windowFrame":
		*e = 7
	case "menuText":
		*e = 8
	case "windowText":
		*e = 9
	case "captionText":
		*e = 10
	case "activeBorder":
		*e = 11
	case "inactiveBorder":
		*e = 12
	case "appWorkspace":
		*e = 13
	case "highlight":
		*e = 14
	case "highlightText":
		*e = 15
	case "btnFace":
		*e = 16
	case "btnShadow":
		*e = 17
	case "grayText":
		*e = 18
	case "btnText":
		*e = 19
	case "inactiveCaptionText":
		*e = 20
	case "btnHighlight":
		*e = 21
	case "3dDkShadow":
		*e = 22
	case "3dLight":
		*e = 23
	case "infoText":
		*e = 24
	case "infoBk":
		*e = 25
	case "hotLight":
		*e = 26
	case "gradientActiveCaption":
		*e = 27
	case "gradientInactiveCaption":
		*e = 28
	case "menuHighlight":
		*e = 29
	case "menuBar":
		*e = 30
	}
	return nil
}
func (m ST_SystemColorVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SystemColorVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "scrollBar":
			*m = 1
		case "background":
			*m = 2
		case "activeCaption":
			*m = 3
		case "inactiveCaption":
			*m = 4
		case "menu":
			*m = 5
		case "window":
			*m = 6
		case "windowFrame":
			*m = 7
		case "menuText":
			*m = 8
		case "windowText":
			*m = 9
		case "captionText":
			*m = 10
		case "activeBorder":
			*m = 11
		case "inactiveBorder":
			*m = 12
		case "appWorkspace":
			*m = 13
		case "highlight":
			*m = 14
		case "highlightText":
			*m = 15
		case "btnFace":
			*m = 16
		case "btnShadow":
			*m = 17
		case "grayText":
			*m = 18
		case "btnText":
			*m = 19
		case "inactiveCaptionText":
			*m = 20
		case "btnHighlight":
			*m = 21
		case "3dDkShadow":
			*m = 22
		case "3dLight":
			*m = 23
		case "infoText":
			*m = 24
		case "infoBk":
			*m = 25
		case "hotLight":
			*m = 26
		case "gradientActiveCaption":
			*m = 27
		case "gradientInactiveCaption":
			*m = 28
		case "menuHighlight":
			*m = 29
		case "menuBar":
			*m = 30
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
func (m ST_SystemColorVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "scrollBar"
	case 2:
		return "background"
	case 3:
		return "activeCaption"
	case 4:
		return "inactiveCaption"
	case 5:
		return "menu"
	case 6:
		return "window"
	case 7:
		return "windowFrame"
	case 8:
		return "menuText"
	case 9:
		return "windowText"
	case 10:
		return "captionText"
	case 11:
		return "activeBorder"
	case 12:
		return "inactiveBorder"
	case 13:
		return "appWorkspace"
	case 14:
		return "highlight"
	case 15:
		return "highlightText"
	case 16:
		return "btnFace"
	case 17:
		return "btnShadow"
	case 18:
		return "grayText"
	case 19:
		return "btnText"
	case 20:
		return "inactiveCaptionText"
	case 21:
		return "btnHighlight"
	case 22:
		return "3dDkShadow"
	case 23:
		return "3dLight"
	case 24:
		return "infoText"
	case 25:
		return "infoBk"
	case 26:
		return "hotLight"
	case 27:
		return "gradientActiveCaption"
	case 28:
		return "gradientInactiveCaption"
	case 29:
		return "menuHighlight"
	case 30:
		return "menuBar"
	}
	return ""
}
func (m ST_SystemColorVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SystemColorVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SchemeColorVal byte

const (
	ST_SchemeColorValUnset    ST_SchemeColorVal = 0
	ST_SchemeColorValBg1      ST_SchemeColorVal = 1
	ST_SchemeColorValTx1      ST_SchemeColorVal = 2
	ST_SchemeColorValBg2      ST_SchemeColorVal = 3
	ST_SchemeColorValTx2      ST_SchemeColorVal = 4
	ST_SchemeColorValAccent1  ST_SchemeColorVal = 5
	ST_SchemeColorValAccent2  ST_SchemeColorVal = 6
	ST_SchemeColorValAccent3  ST_SchemeColorVal = 7
	ST_SchemeColorValAccent4  ST_SchemeColorVal = 8
	ST_SchemeColorValAccent5  ST_SchemeColorVal = 9
	ST_SchemeColorValAccent6  ST_SchemeColorVal = 10
	ST_SchemeColorValHlink    ST_SchemeColorVal = 11
	ST_SchemeColorValFolHlink ST_SchemeColorVal = 12
	ST_SchemeColorValPhClr    ST_SchemeColorVal = 13
	ST_SchemeColorValDk1      ST_SchemeColorVal = 14
	ST_SchemeColorValLt1      ST_SchemeColorVal = 15
	ST_SchemeColorValDk2      ST_SchemeColorVal = 16
	ST_SchemeColorValLt2      ST_SchemeColorVal = 17
)

func (e ST_SchemeColorVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SchemeColorValUnset:
		attr.Value = ""
	case ST_SchemeColorValBg1:
		attr.Value = "bg1"
	case ST_SchemeColorValTx1:
		attr.Value = "tx1"
	case ST_SchemeColorValBg2:
		attr.Value = "bg2"
	case ST_SchemeColorValTx2:
		attr.Value = "tx2"
	case ST_SchemeColorValAccent1:
		attr.Value = "accent1"
	case ST_SchemeColorValAccent2:
		attr.Value = "accent2"
	case ST_SchemeColorValAccent3:
		attr.Value = "accent3"
	case ST_SchemeColorValAccent4:
		attr.Value = "accent4"
	case ST_SchemeColorValAccent5:
		attr.Value = "accent5"
	case ST_SchemeColorValAccent6:
		attr.Value = "accent6"
	case ST_SchemeColorValHlink:
		attr.Value = "hlink"
	case ST_SchemeColorValFolHlink:
		attr.Value = "folHlink"
	case ST_SchemeColorValPhClr:
		attr.Value = "phClr"
	case ST_SchemeColorValDk1:
		attr.Value = "dk1"
	case ST_SchemeColorValLt1:
		attr.Value = "lt1"
	case ST_SchemeColorValDk2:
		attr.Value = "dk2"
	case ST_SchemeColorValLt2:
		attr.Value = "lt2"
	}
	return attr, nil
}
func (e *ST_SchemeColorVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bg1":
		*e = 1
	case "tx1":
		*e = 2
	case "bg2":
		*e = 3
	case "tx2":
		*e = 4
	case "accent1":
		*e = 5
	case "accent2":
		*e = 6
	case "accent3":
		*e = 7
	case "accent4":
		*e = 8
	case "accent5":
		*e = 9
	case "accent6":
		*e = 10
	case "hlink":
		*e = 11
	case "folHlink":
		*e = 12
	case "phClr":
		*e = 13
	case "dk1":
		*e = 14
	case "lt1":
		*e = 15
	case "dk2":
		*e = 16
	case "lt2":
		*e = 17
	}
	return nil
}
func (m ST_SchemeColorVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SchemeColorVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "bg1":
			*m = 1
		case "tx1":
			*m = 2
		case "bg2":
			*m = 3
		case "tx2":
			*m = 4
		case "accent1":
			*m = 5
		case "accent2":
			*m = 6
		case "accent3":
			*m = 7
		case "accent4":
			*m = 8
		case "accent5":
			*m = 9
		case "accent6":
			*m = 10
		case "hlink":
			*m = 11
		case "folHlink":
			*m = 12
		case "phClr":
			*m = 13
		case "dk1":
			*m = 14
		case "lt1":
			*m = 15
		case "dk2":
			*m = 16
		case "lt2":
			*m = 17
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
func (m ST_SchemeColorVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bg1"
	case 2:
		return "tx1"
	case 3:
		return "bg2"
	case 4:
		return "tx2"
	case 5:
		return "accent1"
	case 6:
		return "accent2"
	case 7:
		return "accent3"
	case 8:
		return "accent4"
	case 9:
		return "accent5"
	case 10:
		return "accent6"
	case 11:
		return "hlink"
	case 12:
		return "folHlink"
	case 13:
		return "phClr"
	case 14:
		return "dk1"
	case 15:
		return "lt1"
	case 16:
		return "dk2"
	case 17:
		return "lt2"
	}
	return ""
}
func (m ST_SchemeColorVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SchemeColorVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetColorVal byte

const (
	ST_PresetColorValUnset                ST_PresetColorVal = 0
	ST_PresetColorValAliceBlue            ST_PresetColorVal = 1
	ST_PresetColorValAntiqueWhite         ST_PresetColorVal = 2
	ST_PresetColorValAqua                 ST_PresetColorVal = 3
	ST_PresetColorValAquamarine           ST_PresetColorVal = 4
	ST_PresetColorValAzure                ST_PresetColorVal = 5
	ST_PresetColorValBeige                ST_PresetColorVal = 6
	ST_PresetColorValBisque               ST_PresetColorVal = 7
	ST_PresetColorValBlack                ST_PresetColorVal = 8
	ST_PresetColorValBlanchedAlmond       ST_PresetColorVal = 9
	ST_PresetColorValBlue                 ST_PresetColorVal = 10
	ST_PresetColorValBlueViolet           ST_PresetColorVal = 11
	ST_PresetColorValBrown                ST_PresetColorVal = 12
	ST_PresetColorValBurlyWood            ST_PresetColorVal = 13
	ST_PresetColorValCadetBlue            ST_PresetColorVal = 14
	ST_PresetColorValChartreuse           ST_PresetColorVal = 15
	ST_PresetColorValChocolate            ST_PresetColorVal = 16
	ST_PresetColorValCoral                ST_PresetColorVal = 17
	ST_PresetColorValCornflowerBlue       ST_PresetColorVal = 18
	ST_PresetColorValCornsilk             ST_PresetColorVal = 19
	ST_PresetColorValCrimson              ST_PresetColorVal = 20
	ST_PresetColorValCyan                 ST_PresetColorVal = 21
	ST_PresetColorValDarkBlue             ST_PresetColorVal = 22
	ST_PresetColorValDarkCyan             ST_PresetColorVal = 23
	ST_PresetColorValDarkGoldenrod        ST_PresetColorVal = 24
	ST_PresetColorValDarkGray             ST_PresetColorVal = 25
	ST_PresetColorValDarkGrey             ST_PresetColorVal = 26
	ST_PresetColorValDarkGreen            ST_PresetColorVal = 27
	ST_PresetColorValDarkKhaki            ST_PresetColorVal = 28
	ST_PresetColorValDarkMagenta          ST_PresetColorVal = 29
	ST_PresetColorValDarkOliveGreen       ST_PresetColorVal = 30
	ST_PresetColorValDarkOrange           ST_PresetColorVal = 31
	ST_PresetColorValDarkOrchid           ST_PresetColorVal = 32
	ST_PresetColorValDarkRed              ST_PresetColorVal = 33
	ST_PresetColorValDarkSalmon           ST_PresetColorVal = 34
	ST_PresetColorValDarkSeaGreen         ST_PresetColorVal = 35
	ST_PresetColorValDarkSlateBlue        ST_PresetColorVal = 36
	ST_PresetColorValDarkSlateGray        ST_PresetColorVal = 37
	ST_PresetColorValDarkSlateGrey        ST_PresetColorVal = 38
	ST_PresetColorValDarkTurquoise        ST_PresetColorVal = 39
	ST_PresetColorValDarkViolet           ST_PresetColorVal = 40
	ST_PresetColorValDkBlue               ST_PresetColorVal = 41
	ST_PresetColorValDkCyan               ST_PresetColorVal = 42
	ST_PresetColorValDkGoldenrod          ST_PresetColorVal = 43
	ST_PresetColorValDkGray               ST_PresetColorVal = 44
	ST_PresetColorValDkGrey               ST_PresetColorVal = 45
	ST_PresetColorValDkGreen              ST_PresetColorVal = 46
	ST_PresetColorValDkKhaki              ST_PresetColorVal = 47
	ST_PresetColorValDkMagenta            ST_PresetColorVal = 48
	ST_PresetColorValDkOliveGreen         ST_PresetColorVal = 49
	ST_PresetColorValDkOrange             ST_PresetColorVal = 50
	ST_PresetColorValDkOrchid             ST_PresetColorVal = 51
	ST_PresetColorValDkRed                ST_PresetColorVal = 52
	ST_PresetColorValDkSalmon             ST_PresetColorVal = 53
	ST_PresetColorValDkSeaGreen           ST_PresetColorVal = 54
	ST_PresetColorValDkSlateBlue          ST_PresetColorVal = 55
	ST_PresetColorValDkSlateGray          ST_PresetColorVal = 56
	ST_PresetColorValDkSlateGrey          ST_PresetColorVal = 57
	ST_PresetColorValDkTurquoise          ST_PresetColorVal = 58
	ST_PresetColorValDkViolet             ST_PresetColorVal = 59
	ST_PresetColorValDeepPink             ST_PresetColorVal = 60
	ST_PresetColorValDeepSkyBlue          ST_PresetColorVal = 61
	ST_PresetColorValDimGray              ST_PresetColorVal = 62
	ST_PresetColorValDimGrey              ST_PresetColorVal = 63
	ST_PresetColorValDodgerBlue           ST_PresetColorVal = 64
	ST_PresetColorValFirebrick            ST_PresetColorVal = 65
	ST_PresetColorValFloralWhite          ST_PresetColorVal = 66
	ST_PresetColorValForestGreen          ST_PresetColorVal = 67
	ST_PresetColorValFuchsia              ST_PresetColorVal = 68
	ST_PresetColorValGainsboro            ST_PresetColorVal = 69
	ST_PresetColorValGhostWhite           ST_PresetColorVal = 70
	ST_PresetColorValGold                 ST_PresetColorVal = 71
	ST_PresetColorValGoldenrod            ST_PresetColorVal = 72
	ST_PresetColorValGray                 ST_PresetColorVal = 73
	ST_PresetColorValGrey                 ST_PresetColorVal = 74
	ST_PresetColorValGreen                ST_PresetColorVal = 75
	ST_PresetColorValGreenYellow          ST_PresetColorVal = 76
	ST_PresetColorValHoneydew             ST_PresetColorVal = 77
	ST_PresetColorValHotPink              ST_PresetColorVal = 78
	ST_PresetColorValIndianRed            ST_PresetColorVal = 79
	ST_PresetColorValIndigo               ST_PresetColorVal = 80
	ST_PresetColorValIvory                ST_PresetColorVal = 81
	ST_PresetColorValKhaki                ST_PresetColorVal = 82
	ST_PresetColorValLavender             ST_PresetColorVal = 83
	ST_PresetColorValLavenderBlush        ST_PresetColorVal = 84
	ST_PresetColorValLawnGreen            ST_PresetColorVal = 85
	ST_PresetColorValLemonChiffon         ST_PresetColorVal = 86
	ST_PresetColorValLightBlue            ST_PresetColorVal = 87
	ST_PresetColorValLightCoral           ST_PresetColorVal = 88
	ST_PresetColorValLightCyan            ST_PresetColorVal = 89
	ST_PresetColorValLightGoldenrodYellow ST_PresetColorVal = 90
	ST_PresetColorValLightGray            ST_PresetColorVal = 91
	ST_PresetColorValLightGrey            ST_PresetColorVal = 92
	ST_PresetColorValLightGreen           ST_PresetColorVal = 93
	ST_PresetColorValLightPink            ST_PresetColorVal = 94
	ST_PresetColorValLightSalmon          ST_PresetColorVal = 95
	ST_PresetColorValLightSeaGreen        ST_PresetColorVal = 96
	ST_PresetColorValLightSkyBlue         ST_PresetColorVal = 97
	ST_PresetColorValLightSlateGray       ST_PresetColorVal = 98
	ST_PresetColorValLightSlateGrey       ST_PresetColorVal = 99
	ST_PresetColorValLightSteelBlue       ST_PresetColorVal = 100
	ST_PresetColorValLightYellow          ST_PresetColorVal = 101
	ST_PresetColorValLtBlue               ST_PresetColorVal = 102
	ST_PresetColorValLtCoral              ST_PresetColorVal = 103
	ST_PresetColorValLtCyan               ST_PresetColorVal = 104
	ST_PresetColorValLtGoldenrodYellow    ST_PresetColorVal = 105
	ST_PresetColorValLtGray               ST_PresetColorVal = 106
	ST_PresetColorValLtGrey               ST_PresetColorVal = 107
	ST_PresetColorValLtGreen              ST_PresetColorVal = 108
	ST_PresetColorValLtPink               ST_PresetColorVal = 109
	ST_PresetColorValLtSalmon             ST_PresetColorVal = 110
	ST_PresetColorValLtSeaGreen           ST_PresetColorVal = 111
	ST_PresetColorValLtSkyBlue            ST_PresetColorVal = 112
	ST_PresetColorValLtSlateGray          ST_PresetColorVal = 113
	ST_PresetColorValLtSlateGrey          ST_PresetColorVal = 114
	ST_PresetColorValLtSteelBlue          ST_PresetColorVal = 115
	ST_PresetColorValLtYellow             ST_PresetColorVal = 116
	ST_PresetColorValLime                 ST_PresetColorVal = 117
	ST_PresetColorValLimeGreen            ST_PresetColorVal = 118
	ST_PresetColorValLinen                ST_PresetColorVal = 119
	ST_PresetColorValMagenta              ST_PresetColorVal = 120
	ST_PresetColorValMaroon               ST_PresetColorVal = 121
	ST_PresetColorValMedAquamarine        ST_PresetColorVal = 122
	ST_PresetColorValMedBlue              ST_PresetColorVal = 123
	ST_PresetColorValMedOrchid            ST_PresetColorVal = 124
	ST_PresetColorValMedPurple            ST_PresetColorVal = 125
	ST_PresetColorValMedSeaGreen          ST_PresetColorVal = 126
	ST_PresetColorValMedSlateBlue         ST_PresetColorVal = 127
	ST_PresetColorValMedSpringGreen       ST_PresetColorVal = 128
	ST_PresetColorValMedTurquoise         ST_PresetColorVal = 129
	ST_PresetColorValMedVioletRed         ST_PresetColorVal = 130
	ST_PresetColorValMediumAquamarine     ST_PresetColorVal = 131
	ST_PresetColorValMediumBlue           ST_PresetColorVal = 132
	ST_PresetColorValMediumOrchid         ST_PresetColorVal = 133
	ST_PresetColorValMediumPurple         ST_PresetColorVal = 134
	ST_PresetColorValMediumSeaGreen       ST_PresetColorVal = 135
	ST_PresetColorValMediumSlateBlue      ST_PresetColorVal = 136
	ST_PresetColorValMediumSpringGreen    ST_PresetColorVal = 137
	ST_PresetColorValMediumTurquoise      ST_PresetColorVal = 138
	ST_PresetColorValMediumVioletRed      ST_PresetColorVal = 139
	ST_PresetColorValMidnightBlue         ST_PresetColorVal = 140
	ST_PresetColorValMintCream            ST_PresetColorVal = 141
	ST_PresetColorValMistyRose            ST_PresetColorVal = 142
	ST_PresetColorValMoccasin             ST_PresetColorVal = 143
	ST_PresetColorValNavajoWhite          ST_PresetColorVal = 144
	ST_PresetColorValNavy                 ST_PresetColorVal = 145
	ST_PresetColorValOldLace              ST_PresetColorVal = 146
	ST_PresetColorValOlive                ST_PresetColorVal = 147
	ST_PresetColorValOliveDrab            ST_PresetColorVal = 148
	ST_PresetColorValOrange               ST_PresetColorVal = 149
	ST_PresetColorValOrangeRed            ST_PresetColorVal = 150
	ST_PresetColorValOrchid               ST_PresetColorVal = 151
	ST_PresetColorValPaleGoldenrod        ST_PresetColorVal = 152
	ST_PresetColorValPaleGreen            ST_PresetColorVal = 153
	ST_PresetColorValPaleTurquoise        ST_PresetColorVal = 154
	ST_PresetColorValPaleVioletRed        ST_PresetColorVal = 155
	ST_PresetColorValPapayaWhip           ST_PresetColorVal = 156
	ST_PresetColorValPeachPuff            ST_PresetColorVal = 157
	ST_PresetColorValPeru                 ST_PresetColorVal = 158
	ST_PresetColorValPink                 ST_PresetColorVal = 159
	ST_PresetColorValPlum                 ST_PresetColorVal = 160
	ST_PresetColorValPowderBlue           ST_PresetColorVal = 161
	ST_PresetColorValPurple               ST_PresetColorVal = 162
	ST_PresetColorValRed                  ST_PresetColorVal = 163
	ST_PresetColorValRosyBrown            ST_PresetColorVal = 164
	ST_PresetColorValRoyalBlue            ST_PresetColorVal = 165
	ST_PresetColorValSaddleBrown          ST_PresetColorVal = 166
	ST_PresetColorValSalmon               ST_PresetColorVal = 167
	ST_PresetColorValSandyBrown           ST_PresetColorVal = 168
	ST_PresetColorValSeaGreen             ST_PresetColorVal = 169
	ST_PresetColorValSeaShell             ST_PresetColorVal = 170
	ST_PresetColorValSienna               ST_PresetColorVal = 171
	ST_PresetColorValSilver               ST_PresetColorVal = 172
	ST_PresetColorValSkyBlue              ST_PresetColorVal = 173
	ST_PresetColorValSlateBlue            ST_PresetColorVal = 174
	ST_PresetColorValSlateGray            ST_PresetColorVal = 175
	ST_PresetColorValSlateGrey            ST_PresetColorVal = 176
	ST_PresetColorValSnow                 ST_PresetColorVal = 177
	ST_PresetColorValSpringGreen          ST_PresetColorVal = 178
	ST_PresetColorValSteelBlue            ST_PresetColorVal = 179
	ST_PresetColorValTan                  ST_PresetColorVal = 180
	ST_PresetColorValTeal                 ST_PresetColorVal = 181
	ST_PresetColorValThistle              ST_PresetColorVal = 182
	ST_PresetColorValTomato               ST_PresetColorVal = 183
	ST_PresetColorValTurquoise            ST_PresetColorVal = 184
	ST_PresetColorValViolet               ST_PresetColorVal = 185
	ST_PresetColorValWheat                ST_PresetColorVal = 186
	ST_PresetColorValWhite                ST_PresetColorVal = 187
	ST_PresetColorValWhiteSmoke           ST_PresetColorVal = 188
	ST_PresetColorValYellow               ST_PresetColorVal = 189
	ST_PresetColorValYellowGreen          ST_PresetColorVal = 190
)

func (e ST_PresetColorVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetColorValUnset:
		attr.Value = ""
	case ST_PresetColorValAliceBlue:
		attr.Value = "aliceBlue"
	case ST_PresetColorValAntiqueWhite:
		attr.Value = "antiqueWhite"
	case ST_PresetColorValAqua:
		attr.Value = "aqua"
	case ST_PresetColorValAquamarine:
		attr.Value = "aquamarine"
	case ST_PresetColorValAzure:
		attr.Value = "azure"
	case ST_PresetColorValBeige:
		attr.Value = "beige"
	case ST_PresetColorValBisque:
		attr.Value = "bisque"
	case ST_PresetColorValBlack:
		attr.Value = "black"
	case ST_PresetColorValBlanchedAlmond:
		attr.Value = "blanchedAlmond"
	case ST_PresetColorValBlue:
		attr.Value = "blue"
	case ST_PresetColorValBlueViolet:
		attr.Value = "blueViolet"
	case ST_PresetColorValBrown:
		attr.Value = "brown"
	case ST_PresetColorValBurlyWood:
		attr.Value = "burlyWood"
	case ST_PresetColorValCadetBlue:
		attr.Value = "cadetBlue"
	case ST_PresetColorValChartreuse:
		attr.Value = "chartreuse"
	case ST_PresetColorValChocolate:
		attr.Value = "chocolate"
	case ST_PresetColorValCoral:
		attr.Value = "coral"
	case ST_PresetColorValCornflowerBlue:
		attr.Value = "cornflowerBlue"
	case ST_PresetColorValCornsilk:
		attr.Value = "cornsilk"
	case ST_PresetColorValCrimson:
		attr.Value = "crimson"
	case ST_PresetColorValCyan:
		attr.Value = "cyan"
	case ST_PresetColorValDarkBlue:
		attr.Value = "darkBlue"
	case ST_PresetColorValDarkCyan:
		attr.Value = "darkCyan"
	case ST_PresetColorValDarkGoldenrod:
		attr.Value = "darkGoldenrod"
	case ST_PresetColorValDarkGray:
		attr.Value = "darkGray"
	case ST_PresetColorValDarkGrey:
		attr.Value = "darkGrey"
	case ST_PresetColorValDarkGreen:
		attr.Value = "darkGreen"
	case ST_PresetColorValDarkKhaki:
		attr.Value = "darkKhaki"
	case ST_PresetColorValDarkMagenta:
		attr.Value = "darkMagenta"
	case ST_PresetColorValDarkOliveGreen:
		attr.Value = "darkOliveGreen"
	case ST_PresetColorValDarkOrange:
		attr.Value = "darkOrange"
	case ST_PresetColorValDarkOrchid:
		attr.Value = "darkOrchid"
	case ST_PresetColorValDarkRed:
		attr.Value = "darkRed"
	case ST_PresetColorValDarkSalmon:
		attr.Value = "darkSalmon"
	case ST_PresetColorValDarkSeaGreen:
		attr.Value = "darkSeaGreen"
	case ST_PresetColorValDarkSlateBlue:
		attr.Value = "darkSlateBlue"
	case ST_PresetColorValDarkSlateGray:
		attr.Value = "darkSlateGray"
	case ST_PresetColorValDarkSlateGrey:
		attr.Value = "darkSlateGrey"
	case ST_PresetColorValDarkTurquoise:
		attr.Value = "darkTurquoise"
	case ST_PresetColorValDarkViolet:
		attr.Value = "darkViolet"
	case ST_PresetColorValDkBlue:
		attr.Value = "dkBlue"
	case ST_PresetColorValDkCyan:
		attr.Value = "dkCyan"
	case ST_PresetColorValDkGoldenrod:
		attr.Value = "dkGoldenrod"
	case ST_PresetColorValDkGray:
		attr.Value = "dkGray"
	case ST_PresetColorValDkGrey:
		attr.Value = "dkGrey"
	case ST_PresetColorValDkGreen:
		attr.Value = "dkGreen"
	case ST_PresetColorValDkKhaki:
		attr.Value = "dkKhaki"
	case ST_PresetColorValDkMagenta:
		attr.Value = "dkMagenta"
	case ST_PresetColorValDkOliveGreen:
		attr.Value = "dkOliveGreen"
	case ST_PresetColorValDkOrange:
		attr.Value = "dkOrange"
	case ST_PresetColorValDkOrchid:
		attr.Value = "dkOrchid"
	case ST_PresetColorValDkRed:
		attr.Value = "dkRed"
	case ST_PresetColorValDkSalmon:
		attr.Value = "dkSalmon"
	case ST_PresetColorValDkSeaGreen:
		attr.Value = "dkSeaGreen"
	case ST_PresetColorValDkSlateBlue:
		attr.Value = "dkSlateBlue"
	case ST_PresetColorValDkSlateGray:
		attr.Value = "dkSlateGray"
	case ST_PresetColorValDkSlateGrey:
		attr.Value = "dkSlateGrey"
	case ST_PresetColorValDkTurquoise:
		attr.Value = "dkTurquoise"
	case ST_PresetColorValDkViolet:
		attr.Value = "dkViolet"
	case ST_PresetColorValDeepPink:
		attr.Value = "deepPink"
	case ST_PresetColorValDeepSkyBlue:
		attr.Value = "deepSkyBlue"
	case ST_PresetColorValDimGray:
		attr.Value = "dimGray"
	case ST_PresetColorValDimGrey:
		attr.Value = "dimGrey"
	case ST_PresetColorValDodgerBlue:
		attr.Value = "dodgerBlue"
	case ST_PresetColorValFirebrick:
		attr.Value = "firebrick"
	case ST_PresetColorValFloralWhite:
		attr.Value = "floralWhite"
	case ST_PresetColorValForestGreen:
		attr.Value = "forestGreen"
	case ST_PresetColorValFuchsia:
		attr.Value = "fuchsia"
	case ST_PresetColorValGainsboro:
		attr.Value = "gainsboro"
	case ST_PresetColorValGhostWhite:
		attr.Value = "ghostWhite"
	case ST_PresetColorValGold:
		attr.Value = "gold"
	case ST_PresetColorValGoldenrod:
		attr.Value = "goldenrod"
	case ST_PresetColorValGray:
		attr.Value = "gray"
	case ST_PresetColorValGrey:
		attr.Value = "grey"
	case ST_PresetColorValGreen:
		attr.Value = "green"
	case ST_PresetColorValGreenYellow:
		attr.Value = "greenYellow"
	case ST_PresetColorValHoneydew:
		attr.Value = "honeydew"
	case ST_PresetColorValHotPink:
		attr.Value = "hotPink"
	case ST_PresetColorValIndianRed:
		attr.Value = "indianRed"
	case ST_PresetColorValIndigo:
		attr.Value = "indigo"
	case ST_PresetColorValIvory:
		attr.Value = "ivory"
	case ST_PresetColorValKhaki:
		attr.Value = "khaki"
	case ST_PresetColorValLavender:
		attr.Value = "lavender"
	case ST_PresetColorValLavenderBlush:
		attr.Value = "lavenderBlush"
	case ST_PresetColorValLawnGreen:
		attr.Value = "lawnGreen"
	case ST_PresetColorValLemonChiffon:
		attr.Value = "lemonChiffon"
	case ST_PresetColorValLightBlue:
		attr.Value = "lightBlue"
	case ST_PresetColorValLightCoral:
		attr.Value = "lightCoral"
	case ST_PresetColorValLightCyan:
		attr.Value = "lightCyan"
	case ST_PresetColorValLightGoldenrodYellow:
		attr.Value = "lightGoldenrodYellow"
	case ST_PresetColorValLightGray:
		attr.Value = "lightGray"
	case ST_PresetColorValLightGrey:
		attr.Value = "lightGrey"
	case ST_PresetColorValLightGreen:
		attr.Value = "lightGreen"
	case ST_PresetColorValLightPink:
		attr.Value = "lightPink"
	case ST_PresetColorValLightSalmon:
		attr.Value = "lightSalmon"
	case ST_PresetColorValLightSeaGreen:
		attr.Value = "lightSeaGreen"
	case ST_PresetColorValLightSkyBlue:
		attr.Value = "lightSkyBlue"
	case ST_PresetColorValLightSlateGray:
		attr.Value = "lightSlateGray"
	case ST_PresetColorValLightSlateGrey:
		attr.Value = "lightSlateGrey"
	case ST_PresetColorValLightSteelBlue:
		attr.Value = "lightSteelBlue"
	case ST_PresetColorValLightYellow:
		attr.Value = "lightYellow"
	case ST_PresetColorValLtBlue:
		attr.Value = "ltBlue"
	case ST_PresetColorValLtCoral:
		attr.Value = "ltCoral"
	case ST_PresetColorValLtCyan:
		attr.Value = "ltCyan"
	case ST_PresetColorValLtGoldenrodYellow:
		attr.Value = "ltGoldenrodYellow"
	case ST_PresetColorValLtGray:
		attr.Value = "ltGray"
	case ST_PresetColorValLtGrey:
		attr.Value = "ltGrey"
	case ST_PresetColorValLtGreen:
		attr.Value = "ltGreen"
	case ST_PresetColorValLtPink:
		attr.Value = "ltPink"
	case ST_PresetColorValLtSalmon:
		attr.Value = "ltSalmon"
	case ST_PresetColorValLtSeaGreen:
		attr.Value = "ltSeaGreen"
	case ST_PresetColorValLtSkyBlue:
		attr.Value = "ltSkyBlue"
	case ST_PresetColorValLtSlateGray:
		attr.Value = "ltSlateGray"
	case ST_PresetColorValLtSlateGrey:
		attr.Value = "ltSlateGrey"
	case ST_PresetColorValLtSteelBlue:
		attr.Value = "ltSteelBlue"
	case ST_PresetColorValLtYellow:
		attr.Value = "ltYellow"
	case ST_PresetColorValLime:
		attr.Value = "lime"
	case ST_PresetColorValLimeGreen:
		attr.Value = "limeGreen"
	case ST_PresetColorValLinen:
		attr.Value = "linen"
	case ST_PresetColorValMagenta:
		attr.Value = "magenta"
	case ST_PresetColorValMaroon:
		attr.Value = "maroon"
	case ST_PresetColorValMedAquamarine:
		attr.Value = "medAquamarine"
	case ST_PresetColorValMedBlue:
		attr.Value = "medBlue"
	case ST_PresetColorValMedOrchid:
		attr.Value = "medOrchid"
	case ST_PresetColorValMedPurple:
		attr.Value = "medPurple"
	case ST_PresetColorValMedSeaGreen:
		attr.Value = "medSeaGreen"
	case ST_PresetColorValMedSlateBlue:
		attr.Value = "medSlateBlue"
	case ST_PresetColorValMedSpringGreen:
		attr.Value = "medSpringGreen"
	case ST_PresetColorValMedTurquoise:
		attr.Value = "medTurquoise"
	case ST_PresetColorValMedVioletRed:
		attr.Value = "medVioletRed"
	case ST_PresetColorValMediumAquamarine:
		attr.Value = "mediumAquamarine"
	case ST_PresetColorValMediumBlue:
		attr.Value = "mediumBlue"
	case ST_PresetColorValMediumOrchid:
		attr.Value = "mediumOrchid"
	case ST_PresetColorValMediumPurple:
		attr.Value = "mediumPurple"
	case ST_PresetColorValMediumSeaGreen:
		attr.Value = "mediumSeaGreen"
	case ST_PresetColorValMediumSlateBlue:
		attr.Value = "mediumSlateBlue"
	case ST_PresetColorValMediumSpringGreen:
		attr.Value = "mediumSpringGreen"
	case ST_PresetColorValMediumTurquoise:
		attr.Value = "mediumTurquoise"
	case ST_PresetColorValMediumVioletRed:
		attr.Value = "mediumVioletRed"
	case ST_PresetColorValMidnightBlue:
		attr.Value = "midnightBlue"
	case ST_PresetColorValMintCream:
		attr.Value = "mintCream"
	case ST_PresetColorValMistyRose:
		attr.Value = "mistyRose"
	case ST_PresetColorValMoccasin:
		attr.Value = "moccasin"
	case ST_PresetColorValNavajoWhite:
		attr.Value = "navajoWhite"
	case ST_PresetColorValNavy:
		attr.Value = "navy"
	case ST_PresetColorValOldLace:
		attr.Value = "oldLace"
	case ST_PresetColorValOlive:
		attr.Value = "olive"
	case ST_PresetColorValOliveDrab:
		attr.Value = "oliveDrab"
	case ST_PresetColorValOrange:
		attr.Value = "orange"
	case ST_PresetColorValOrangeRed:
		attr.Value = "orangeRed"
	case ST_PresetColorValOrchid:
		attr.Value = "orchid"
	case ST_PresetColorValPaleGoldenrod:
		attr.Value = "paleGoldenrod"
	case ST_PresetColorValPaleGreen:
		attr.Value = "paleGreen"
	case ST_PresetColorValPaleTurquoise:
		attr.Value = "paleTurquoise"
	case ST_PresetColorValPaleVioletRed:
		attr.Value = "paleVioletRed"
	case ST_PresetColorValPapayaWhip:
		attr.Value = "papayaWhip"
	case ST_PresetColorValPeachPuff:
		attr.Value = "peachPuff"
	case ST_PresetColorValPeru:
		attr.Value = "peru"
	case ST_PresetColorValPink:
		attr.Value = "pink"
	case ST_PresetColorValPlum:
		attr.Value = "plum"
	case ST_PresetColorValPowderBlue:
		attr.Value = "powderBlue"
	case ST_PresetColorValPurple:
		attr.Value = "purple"
	case ST_PresetColorValRed:
		attr.Value = "red"
	case ST_PresetColorValRosyBrown:
		attr.Value = "rosyBrown"
	case ST_PresetColorValRoyalBlue:
		attr.Value = "royalBlue"
	case ST_PresetColorValSaddleBrown:
		attr.Value = "saddleBrown"
	case ST_PresetColorValSalmon:
		attr.Value = "salmon"
	case ST_PresetColorValSandyBrown:
		attr.Value = "sandyBrown"
	case ST_PresetColorValSeaGreen:
		attr.Value = "seaGreen"
	case ST_PresetColorValSeaShell:
		attr.Value = "seaShell"
	case ST_PresetColorValSienna:
		attr.Value = "sienna"
	case ST_PresetColorValSilver:
		attr.Value = "silver"
	case ST_PresetColorValSkyBlue:
		attr.Value = "skyBlue"
	case ST_PresetColorValSlateBlue:
		attr.Value = "slateBlue"
	case ST_PresetColorValSlateGray:
		attr.Value = "slateGray"
	case ST_PresetColorValSlateGrey:
		attr.Value = "slateGrey"
	case ST_PresetColorValSnow:
		attr.Value = "snow"
	case ST_PresetColorValSpringGreen:
		attr.Value = "springGreen"
	case ST_PresetColorValSteelBlue:
		attr.Value = "steelBlue"
	case ST_PresetColorValTan:
		attr.Value = "tan"
	case ST_PresetColorValTeal:
		attr.Value = "teal"
	case ST_PresetColorValThistle:
		attr.Value = "thistle"
	case ST_PresetColorValTomato:
		attr.Value = "tomato"
	case ST_PresetColorValTurquoise:
		attr.Value = "turquoise"
	case ST_PresetColorValViolet:
		attr.Value = "violet"
	case ST_PresetColorValWheat:
		attr.Value = "wheat"
	case ST_PresetColorValWhite:
		attr.Value = "white"
	case ST_PresetColorValWhiteSmoke:
		attr.Value = "whiteSmoke"
	case ST_PresetColorValYellow:
		attr.Value = "yellow"
	case ST_PresetColorValYellowGreen:
		attr.Value = "yellowGreen"
	}
	return attr, nil
}
func (e *ST_PresetColorVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "aliceBlue":
		*e = 1
	case "antiqueWhite":
		*e = 2
	case "aqua":
		*e = 3
	case "aquamarine":
		*e = 4
	case "azure":
		*e = 5
	case "beige":
		*e = 6
	case "bisque":
		*e = 7
	case "black":
		*e = 8
	case "blanchedAlmond":
		*e = 9
	case "blue":
		*e = 10
	case "blueViolet":
		*e = 11
	case "brown":
		*e = 12
	case "burlyWood":
		*e = 13
	case "cadetBlue":
		*e = 14
	case "chartreuse":
		*e = 15
	case "chocolate":
		*e = 16
	case "coral":
		*e = 17
	case "cornflowerBlue":
		*e = 18
	case "cornsilk":
		*e = 19
	case "crimson":
		*e = 20
	case "cyan":
		*e = 21
	case "darkBlue":
		*e = 22
	case "darkCyan":
		*e = 23
	case "darkGoldenrod":
		*e = 24
	case "darkGray":
		*e = 25
	case "darkGrey":
		*e = 26
	case "darkGreen":
		*e = 27
	case "darkKhaki":
		*e = 28
	case "darkMagenta":
		*e = 29
	case "darkOliveGreen":
		*e = 30
	case "darkOrange":
		*e = 31
	case "darkOrchid":
		*e = 32
	case "darkRed":
		*e = 33
	case "darkSalmon":
		*e = 34
	case "darkSeaGreen":
		*e = 35
	case "darkSlateBlue":
		*e = 36
	case "darkSlateGray":
		*e = 37
	case "darkSlateGrey":
		*e = 38
	case "darkTurquoise":
		*e = 39
	case "darkViolet":
		*e = 40
	case "dkBlue":
		*e = 41
	case "dkCyan":
		*e = 42
	case "dkGoldenrod":
		*e = 43
	case "dkGray":
		*e = 44
	case "dkGrey":
		*e = 45
	case "dkGreen":
		*e = 46
	case "dkKhaki":
		*e = 47
	case "dkMagenta":
		*e = 48
	case "dkOliveGreen":
		*e = 49
	case "dkOrange":
		*e = 50
	case "dkOrchid":
		*e = 51
	case "dkRed":
		*e = 52
	case "dkSalmon":
		*e = 53
	case "dkSeaGreen":
		*e = 54
	case "dkSlateBlue":
		*e = 55
	case "dkSlateGray":
		*e = 56
	case "dkSlateGrey":
		*e = 57
	case "dkTurquoise":
		*e = 58
	case "dkViolet":
		*e = 59
	case "deepPink":
		*e = 60
	case "deepSkyBlue":
		*e = 61
	case "dimGray":
		*e = 62
	case "dimGrey":
		*e = 63
	case "dodgerBlue":
		*e = 64
	case "firebrick":
		*e = 65
	case "floralWhite":
		*e = 66
	case "forestGreen":
		*e = 67
	case "fuchsia":
		*e = 68
	case "gainsboro":
		*e = 69
	case "ghostWhite":
		*e = 70
	case "gold":
		*e = 71
	case "goldenrod":
		*e = 72
	case "gray":
		*e = 73
	case "grey":
		*e = 74
	case "green":
		*e = 75
	case "greenYellow":
		*e = 76
	case "honeydew":
		*e = 77
	case "hotPink":
		*e = 78
	case "indianRed":
		*e = 79
	case "indigo":
		*e = 80
	case "ivory":
		*e = 81
	case "khaki":
		*e = 82
	case "lavender":
		*e = 83
	case "lavenderBlush":
		*e = 84
	case "lawnGreen":
		*e = 85
	case "lemonChiffon":
		*e = 86
	case "lightBlue":
		*e = 87
	case "lightCoral":
		*e = 88
	case "lightCyan":
		*e = 89
	case "lightGoldenrodYellow":
		*e = 90
	case "lightGray":
		*e = 91
	case "lightGrey":
		*e = 92
	case "lightGreen":
		*e = 93
	case "lightPink":
		*e = 94
	case "lightSalmon":
		*e = 95
	case "lightSeaGreen":
		*e = 96
	case "lightSkyBlue":
		*e = 97
	case "lightSlateGray":
		*e = 98
	case "lightSlateGrey":
		*e = 99
	case "lightSteelBlue":
		*e = 100
	case "lightYellow":
		*e = 101
	case "ltBlue":
		*e = 102
	case "ltCoral":
		*e = 103
	case "ltCyan":
		*e = 104
	case "ltGoldenrodYellow":
		*e = 105
	case "ltGray":
		*e = 106
	case "ltGrey":
		*e = 107
	case "ltGreen":
		*e = 108
	case "ltPink":
		*e = 109
	case "ltSalmon":
		*e = 110
	case "ltSeaGreen":
		*e = 111
	case "ltSkyBlue":
		*e = 112
	case "ltSlateGray":
		*e = 113
	case "ltSlateGrey":
		*e = 114
	case "ltSteelBlue":
		*e = 115
	case "ltYellow":
		*e = 116
	case "lime":
		*e = 117
	case "limeGreen":
		*e = 118
	case "linen":
		*e = 119
	case "magenta":
		*e = 120
	case "maroon":
		*e = 121
	case "medAquamarine":
		*e = 122
	case "medBlue":
		*e = 123
	case "medOrchid":
		*e = 124
	case "medPurple":
		*e = 125
	case "medSeaGreen":
		*e = 126
	case "medSlateBlue":
		*e = 127
	case "medSpringGreen":
		*e = 128
	case "medTurquoise":
		*e = 129
	case "medVioletRed":
		*e = 130
	case "mediumAquamarine":
		*e = 131
	case "mediumBlue":
		*e = 132
	case "mediumOrchid":
		*e = 133
	case "mediumPurple":
		*e = 134
	case "mediumSeaGreen":
		*e = 135
	case "mediumSlateBlue":
		*e = 136
	case "mediumSpringGreen":
		*e = 137
	case "mediumTurquoise":
		*e = 138
	case "mediumVioletRed":
		*e = 139
	case "midnightBlue":
		*e = 140
	case "mintCream":
		*e = 141
	case "mistyRose":
		*e = 142
	case "moccasin":
		*e = 143
	case "navajoWhite":
		*e = 144
	case "navy":
		*e = 145
	case "oldLace":
		*e = 146
	case "olive":
		*e = 147
	case "oliveDrab":
		*e = 148
	case "orange":
		*e = 149
	case "orangeRed":
		*e = 150
	case "orchid":
		*e = 151
	case "paleGoldenrod":
		*e = 152
	case "paleGreen":
		*e = 153
	case "paleTurquoise":
		*e = 154
	case "paleVioletRed":
		*e = 155
	case "papayaWhip":
		*e = 156
	case "peachPuff":
		*e = 157
	case "peru":
		*e = 158
	case "pink":
		*e = 159
	case "plum":
		*e = 160
	case "powderBlue":
		*e = 161
	case "purple":
		*e = 162
	case "red":
		*e = 163
	case "rosyBrown":
		*e = 164
	case "royalBlue":
		*e = 165
	case "saddleBrown":
		*e = 166
	case "salmon":
		*e = 167
	case "sandyBrown":
		*e = 168
	case "seaGreen":
		*e = 169
	case "seaShell":
		*e = 170
	case "sienna":
		*e = 171
	case "silver":
		*e = 172
	case "skyBlue":
		*e = 173
	case "slateBlue":
		*e = 174
	case "slateGray":
		*e = 175
	case "slateGrey":
		*e = 176
	case "snow":
		*e = 177
	case "springGreen":
		*e = 178
	case "steelBlue":
		*e = 179
	case "tan":
		*e = 180
	case "teal":
		*e = 181
	case "thistle":
		*e = 182
	case "tomato":
		*e = 183
	case "turquoise":
		*e = 184
	case "violet":
		*e = 185
	case "wheat":
		*e = 186
	case "white":
		*e = 187
	case "whiteSmoke":
		*e = 188
	case "yellow":
		*e = 189
	case "yellowGreen":
		*e = 190
	}
	return nil
}
func (m ST_PresetColorVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetColorVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "aliceBlue":
			*m = 1
		case "antiqueWhite":
			*m = 2
		case "aqua":
			*m = 3
		case "aquamarine":
			*m = 4
		case "azure":
			*m = 5
		case "beige":
			*m = 6
		case "bisque":
			*m = 7
		case "black":
			*m = 8
		case "blanchedAlmond":
			*m = 9
		case "blue":
			*m = 10
		case "blueViolet":
			*m = 11
		case "brown":
			*m = 12
		case "burlyWood":
			*m = 13
		case "cadetBlue":
			*m = 14
		case "chartreuse":
			*m = 15
		case "chocolate":
			*m = 16
		case "coral":
			*m = 17
		case "cornflowerBlue":
			*m = 18
		case "cornsilk":
			*m = 19
		case "crimson":
			*m = 20
		case "cyan":
			*m = 21
		case "darkBlue":
			*m = 22
		case "darkCyan":
			*m = 23
		case "darkGoldenrod":
			*m = 24
		case "darkGray":
			*m = 25
		case "darkGrey":
			*m = 26
		case "darkGreen":
			*m = 27
		case "darkKhaki":
			*m = 28
		case "darkMagenta":
			*m = 29
		case "darkOliveGreen":
			*m = 30
		case "darkOrange":
			*m = 31
		case "darkOrchid":
			*m = 32
		case "darkRed":
			*m = 33
		case "darkSalmon":
			*m = 34
		case "darkSeaGreen":
			*m = 35
		case "darkSlateBlue":
			*m = 36
		case "darkSlateGray":
			*m = 37
		case "darkSlateGrey":
			*m = 38
		case "darkTurquoise":
			*m = 39
		case "darkViolet":
			*m = 40
		case "dkBlue":
			*m = 41
		case "dkCyan":
			*m = 42
		case "dkGoldenrod":
			*m = 43
		case "dkGray":
			*m = 44
		case "dkGrey":
			*m = 45
		case "dkGreen":
			*m = 46
		case "dkKhaki":
			*m = 47
		case "dkMagenta":
			*m = 48
		case "dkOliveGreen":
			*m = 49
		case "dkOrange":
			*m = 50
		case "dkOrchid":
			*m = 51
		case "dkRed":
			*m = 52
		case "dkSalmon":
			*m = 53
		case "dkSeaGreen":
			*m = 54
		case "dkSlateBlue":
			*m = 55
		case "dkSlateGray":
			*m = 56
		case "dkSlateGrey":
			*m = 57
		case "dkTurquoise":
			*m = 58
		case "dkViolet":
			*m = 59
		case "deepPink":
			*m = 60
		case "deepSkyBlue":
			*m = 61
		case "dimGray":
			*m = 62
		case "dimGrey":
			*m = 63
		case "dodgerBlue":
			*m = 64
		case "firebrick":
			*m = 65
		case "floralWhite":
			*m = 66
		case "forestGreen":
			*m = 67
		case "fuchsia":
			*m = 68
		case "gainsboro":
			*m = 69
		case "ghostWhite":
			*m = 70
		case "gold":
			*m = 71
		case "goldenrod":
			*m = 72
		case "gray":
			*m = 73
		case "grey":
			*m = 74
		case "green":
			*m = 75
		case "greenYellow":
			*m = 76
		case "honeydew":
			*m = 77
		case "hotPink":
			*m = 78
		case "indianRed":
			*m = 79
		case "indigo":
			*m = 80
		case "ivory":
			*m = 81
		case "khaki":
			*m = 82
		case "lavender":
			*m = 83
		case "lavenderBlush":
			*m = 84
		case "lawnGreen":
			*m = 85
		case "lemonChiffon":
			*m = 86
		case "lightBlue":
			*m = 87
		case "lightCoral":
			*m = 88
		case "lightCyan":
			*m = 89
		case "lightGoldenrodYellow":
			*m = 90
		case "lightGray":
			*m = 91
		case "lightGrey":
			*m = 92
		case "lightGreen":
			*m = 93
		case "lightPink":
			*m = 94
		case "lightSalmon":
			*m = 95
		case "lightSeaGreen":
			*m = 96
		case "lightSkyBlue":
			*m = 97
		case "lightSlateGray":
			*m = 98
		case "lightSlateGrey":
			*m = 99
		case "lightSteelBlue":
			*m = 100
		case "lightYellow":
			*m = 101
		case "ltBlue":
			*m = 102
		case "ltCoral":
			*m = 103
		case "ltCyan":
			*m = 104
		case "ltGoldenrodYellow":
			*m = 105
		case "ltGray":
			*m = 106
		case "ltGrey":
			*m = 107
		case "ltGreen":
			*m = 108
		case "ltPink":
			*m = 109
		case "ltSalmon":
			*m = 110
		case "ltSeaGreen":
			*m = 111
		case "ltSkyBlue":
			*m = 112
		case "ltSlateGray":
			*m = 113
		case "ltSlateGrey":
			*m = 114
		case "ltSteelBlue":
			*m = 115
		case "ltYellow":
			*m = 116
		case "lime":
			*m = 117
		case "limeGreen":
			*m = 118
		case "linen":
			*m = 119
		case "magenta":
			*m = 120
		case "maroon":
			*m = 121
		case "medAquamarine":
			*m = 122
		case "medBlue":
			*m = 123
		case "medOrchid":
			*m = 124
		case "medPurple":
			*m = 125
		case "medSeaGreen":
			*m = 126
		case "medSlateBlue":
			*m = 127
		case "medSpringGreen":
			*m = 128
		case "medTurquoise":
			*m = 129
		case "medVioletRed":
			*m = 130
		case "mediumAquamarine":
			*m = 131
		case "mediumBlue":
			*m = 132
		case "mediumOrchid":
			*m = 133
		case "mediumPurple":
			*m = 134
		case "mediumSeaGreen":
			*m = 135
		case "mediumSlateBlue":
			*m = 136
		case "mediumSpringGreen":
			*m = 137
		case "mediumTurquoise":
			*m = 138
		case "mediumVioletRed":
			*m = 139
		case "midnightBlue":
			*m = 140
		case "mintCream":
			*m = 141
		case "mistyRose":
			*m = 142
		case "moccasin":
			*m = 143
		case "navajoWhite":
			*m = 144
		case "navy":
			*m = 145
		case "oldLace":
			*m = 146
		case "olive":
			*m = 147
		case "oliveDrab":
			*m = 148
		case "orange":
			*m = 149
		case "orangeRed":
			*m = 150
		case "orchid":
			*m = 151
		case "paleGoldenrod":
			*m = 152
		case "paleGreen":
			*m = 153
		case "paleTurquoise":
			*m = 154
		case "paleVioletRed":
			*m = 155
		case "papayaWhip":
			*m = 156
		case "peachPuff":
			*m = 157
		case "peru":
			*m = 158
		case "pink":
			*m = 159
		case "plum":
			*m = 160
		case "powderBlue":
			*m = 161
		case "purple":
			*m = 162
		case "red":
			*m = 163
		case "rosyBrown":
			*m = 164
		case "royalBlue":
			*m = 165
		case "saddleBrown":
			*m = 166
		case "salmon":
			*m = 167
		case "sandyBrown":
			*m = 168
		case "seaGreen":
			*m = 169
		case "seaShell":
			*m = 170
		case "sienna":
			*m = 171
		case "silver":
			*m = 172
		case "skyBlue":
			*m = 173
		case "slateBlue":
			*m = 174
		case "slateGray":
			*m = 175
		case "slateGrey":
			*m = 176
		case "snow":
			*m = 177
		case "springGreen":
			*m = 178
		case "steelBlue":
			*m = 179
		case "tan":
			*m = 180
		case "teal":
			*m = 181
		case "thistle":
			*m = 182
		case "tomato":
			*m = 183
		case "turquoise":
			*m = 184
		case "violet":
			*m = 185
		case "wheat":
			*m = 186
		case "white":
			*m = 187
		case "whiteSmoke":
			*m = 188
		case "yellow":
			*m = 189
		case "yellowGreen":
			*m = 190
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
func (m ST_PresetColorVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "aliceBlue"
	case 2:
		return "antiqueWhite"
	case 3:
		return "aqua"
	case 4:
		return "aquamarine"
	case 5:
		return "azure"
	case 6:
		return "beige"
	case 7:
		return "bisque"
	case 8:
		return "black"
	case 9:
		return "blanchedAlmond"
	case 10:
		return "blue"
	case 11:
		return "blueViolet"
	case 12:
		return "brown"
	case 13:
		return "burlyWood"
	case 14:
		return "cadetBlue"
	case 15:
		return "chartreuse"
	case 16:
		return "chocolate"
	case 17:
		return "coral"
	case 18:
		return "cornflowerBlue"
	case 19:
		return "cornsilk"
	case 20:
		return "crimson"
	case 21:
		return "cyan"
	case 22:
		return "darkBlue"
	case 23:
		return "darkCyan"
	case 24:
		return "darkGoldenrod"
	case 25:
		return "darkGray"
	case 26:
		return "darkGrey"
	case 27:
		return "darkGreen"
	case 28:
		return "darkKhaki"
	case 29:
		return "darkMagenta"
	case 30:
		return "darkOliveGreen"
	case 31:
		return "darkOrange"
	case 32:
		return "darkOrchid"
	case 33:
		return "darkRed"
	case 34:
		return "darkSalmon"
	case 35:
		return "darkSeaGreen"
	case 36:
		return "darkSlateBlue"
	case 37:
		return "darkSlateGray"
	case 38:
		return "darkSlateGrey"
	case 39:
		return "darkTurquoise"
	case 40:
		return "darkViolet"
	case 41:
		return "dkBlue"
	case 42:
		return "dkCyan"
	case 43:
		return "dkGoldenrod"
	case 44:
		return "dkGray"
	case 45:
		return "dkGrey"
	case 46:
		return "dkGreen"
	case 47:
		return "dkKhaki"
	case 48:
		return "dkMagenta"
	case 49:
		return "dkOliveGreen"
	case 50:
		return "dkOrange"
	case 51:
		return "dkOrchid"
	case 52:
		return "dkRed"
	case 53:
		return "dkSalmon"
	case 54:
		return "dkSeaGreen"
	case 55:
		return "dkSlateBlue"
	case 56:
		return "dkSlateGray"
	case 57:
		return "dkSlateGrey"
	case 58:
		return "dkTurquoise"
	case 59:
		return "dkViolet"
	case 60:
		return "deepPink"
	case 61:
		return "deepSkyBlue"
	case 62:
		return "dimGray"
	case 63:
		return "dimGrey"
	case 64:
		return "dodgerBlue"
	case 65:
		return "firebrick"
	case 66:
		return "floralWhite"
	case 67:
		return "forestGreen"
	case 68:
		return "fuchsia"
	case 69:
		return "gainsboro"
	case 70:
		return "ghostWhite"
	case 71:
		return "gold"
	case 72:
		return "goldenrod"
	case 73:
		return "gray"
	case 74:
		return "grey"
	case 75:
		return "green"
	case 76:
		return "greenYellow"
	case 77:
		return "honeydew"
	case 78:
		return "hotPink"
	case 79:
		return "indianRed"
	case 80:
		return "indigo"
	case 81:
		return "ivory"
	case 82:
		return "khaki"
	case 83:
		return "lavender"
	case 84:
		return "lavenderBlush"
	case 85:
		return "lawnGreen"
	case 86:
		return "lemonChiffon"
	case 87:
		return "lightBlue"
	case 88:
		return "lightCoral"
	case 89:
		return "lightCyan"
	case 90:
		return "lightGoldenrodYellow"
	case 91:
		return "lightGray"
	case 92:
		return "lightGrey"
	case 93:
		return "lightGreen"
	case 94:
		return "lightPink"
	case 95:
		return "lightSalmon"
	case 96:
		return "lightSeaGreen"
	case 97:
		return "lightSkyBlue"
	case 98:
		return "lightSlateGray"
	case 99:
		return "lightSlateGrey"
	case 100:
		return "lightSteelBlue"
	case 101:
		return "lightYellow"
	case 102:
		return "ltBlue"
	case 103:
		return "ltCoral"
	case 104:
		return "ltCyan"
	case 105:
		return "ltGoldenrodYellow"
	case 106:
		return "ltGray"
	case 107:
		return "ltGrey"
	case 108:
		return "ltGreen"
	case 109:
		return "ltPink"
	case 110:
		return "ltSalmon"
	case 111:
		return "ltSeaGreen"
	case 112:
		return "ltSkyBlue"
	case 113:
		return "ltSlateGray"
	case 114:
		return "ltSlateGrey"
	case 115:
		return "ltSteelBlue"
	case 116:
		return "ltYellow"
	case 117:
		return "lime"
	case 118:
		return "limeGreen"
	case 119:
		return "linen"
	case 120:
		return "magenta"
	case 121:
		return "maroon"
	case 122:
		return "medAquamarine"
	case 123:
		return "medBlue"
	case 124:
		return "medOrchid"
	case 125:
		return "medPurple"
	case 126:
		return "medSeaGreen"
	case 127:
		return "medSlateBlue"
	case 128:
		return "medSpringGreen"
	case 129:
		return "medTurquoise"
	case 130:
		return "medVioletRed"
	case 131:
		return "mediumAquamarine"
	case 132:
		return "mediumBlue"
	case 133:
		return "mediumOrchid"
	case 134:
		return "mediumPurple"
	case 135:
		return "mediumSeaGreen"
	case 136:
		return "mediumSlateBlue"
	case 137:
		return "mediumSpringGreen"
	case 138:
		return "mediumTurquoise"
	case 139:
		return "mediumVioletRed"
	case 140:
		return "midnightBlue"
	case 141:
		return "mintCream"
	case 142:
		return "mistyRose"
	case 143:
		return "moccasin"
	case 144:
		return "navajoWhite"
	case 145:
		return "navy"
	case 146:
		return "oldLace"
	case 147:
		return "olive"
	case 148:
		return "oliveDrab"
	case 149:
		return "orange"
	case 150:
		return "orangeRed"
	case 151:
		return "orchid"
	case 152:
		return "paleGoldenrod"
	case 153:
		return "paleGreen"
	case 154:
		return "paleTurquoise"
	case 155:
		return "paleVioletRed"
	case 156:
		return "papayaWhip"
	case 157:
		return "peachPuff"
	case 158:
		return "peru"
	case 159:
		return "pink"
	case 160:
		return "plum"
	case 161:
		return "powderBlue"
	case 162:
		return "purple"
	case 163:
		return "red"
	case 164:
		return "rosyBrown"
	case 165:
		return "royalBlue"
	case 166:
		return "saddleBrown"
	case 167:
		return "salmon"
	case 168:
		return "sandyBrown"
	case 169:
		return "seaGreen"
	case 170:
		return "seaShell"
	case 171:
		return "sienna"
	case 172:
		return "silver"
	case 173:
		return "skyBlue"
	case 174:
		return "slateBlue"
	case 175:
		return "slateGray"
	case 176:
		return "slateGrey"
	case 177:
		return "snow"
	case 178:
		return "springGreen"
	case 179:
		return "steelBlue"
	case 180:
		return "tan"
	case 181:
		return "teal"
	case 182:
		return "thistle"
	case 183:
		return "tomato"
	case 184:
		return "turquoise"
	case 185:
		return "violet"
	case 186:
		return "wheat"
	case 187:
		return "white"
	case 188:
		return "whiteSmoke"
	case 189:
		return "yellow"
	case 190:
		return "yellowGreen"
	}
	return ""
}
func (m ST_PresetColorVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetColorVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RectAlignment byte

const (
	ST_RectAlignmentUnset ST_RectAlignment = 0
	ST_RectAlignmentTl    ST_RectAlignment = 1
	ST_RectAlignmentT     ST_RectAlignment = 2
	ST_RectAlignmentTr    ST_RectAlignment = 3
	ST_RectAlignmentL     ST_RectAlignment = 4
	ST_RectAlignmentCtr   ST_RectAlignment = 5
	ST_RectAlignmentR     ST_RectAlignment = 6
	ST_RectAlignmentBl    ST_RectAlignment = 7
	ST_RectAlignmentB     ST_RectAlignment = 8
	ST_RectAlignmentBr    ST_RectAlignment = 9
)

func (e ST_RectAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RectAlignmentUnset:
		attr.Value = ""
	case ST_RectAlignmentTl:
		attr.Value = "tl"
	case ST_RectAlignmentT:
		attr.Value = "t"
	case ST_RectAlignmentTr:
		attr.Value = "tr"
	case ST_RectAlignmentL:
		attr.Value = "l"
	case ST_RectAlignmentCtr:
		attr.Value = "ctr"
	case ST_RectAlignmentR:
		attr.Value = "r"
	case ST_RectAlignmentBl:
		attr.Value = "bl"
	case ST_RectAlignmentB:
		attr.Value = "b"
	case ST_RectAlignmentBr:
		attr.Value = "br"
	}
	return attr, nil
}
func (e *ST_RectAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tl":
		*e = 1
	case "t":
		*e = 2
	case "tr":
		*e = 3
	case "l":
		*e = 4
	case "ctr":
		*e = 5
	case "r":
		*e = 6
	case "bl":
		*e = 7
	case "b":
		*e = 8
	case "br":
		*e = 9
	}
	return nil
}
func (m ST_RectAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_RectAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "tl":
			*m = 1
		case "t":
			*m = 2
		case "tr":
			*m = 3
		case "l":
			*m = 4
		case "ctr":
			*m = 5
		case "r":
			*m = 6
		case "bl":
			*m = 7
		case "b":
			*m = 8
		case "br":
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
func (m ST_RectAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tl"
	case 2:
		return "t"
	case 3:
		return "tr"
	case 4:
		return "l"
	case 5:
		return "ctr"
	case 6:
		return "r"
	case 7:
		return "bl"
	case 8:
		return "b"
	case 9:
		return "br"
	}
	return ""
}
func (m ST_RectAlignment) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_RectAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BlackWhiteMode byte

const (
	ST_BlackWhiteModeUnset      ST_BlackWhiteMode = 0
	ST_BlackWhiteModeClr        ST_BlackWhiteMode = 1
	ST_BlackWhiteModeAuto       ST_BlackWhiteMode = 2
	ST_BlackWhiteModeGray       ST_BlackWhiteMode = 3
	ST_BlackWhiteModeLtGray     ST_BlackWhiteMode = 4
	ST_BlackWhiteModeInvGray    ST_BlackWhiteMode = 5
	ST_BlackWhiteModeGrayWhite  ST_BlackWhiteMode = 6
	ST_BlackWhiteModeBlackGray  ST_BlackWhiteMode = 7
	ST_BlackWhiteModeBlackWhite ST_BlackWhiteMode = 8
	ST_BlackWhiteModeBlack      ST_BlackWhiteMode = 9
	ST_BlackWhiteModeWhite      ST_BlackWhiteMode = 10
	ST_BlackWhiteModeHidden     ST_BlackWhiteMode = 11
)

func (e ST_BlackWhiteMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BlackWhiteModeUnset:
		attr.Value = ""
	case ST_BlackWhiteModeClr:
		attr.Value = "clr"
	case ST_BlackWhiteModeAuto:
		attr.Value = "auto"
	case ST_BlackWhiteModeGray:
		attr.Value = "gray"
	case ST_BlackWhiteModeLtGray:
		attr.Value = "ltGray"
	case ST_BlackWhiteModeInvGray:
		attr.Value = "invGray"
	case ST_BlackWhiteModeGrayWhite:
		attr.Value = "grayWhite"
	case ST_BlackWhiteModeBlackGray:
		attr.Value = "blackGray"
	case ST_BlackWhiteModeBlackWhite:
		attr.Value = "blackWhite"
	case ST_BlackWhiteModeBlack:
		attr.Value = "black"
	case ST_BlackWhiteModeWhite:
		attr.Value = "white"
	case ST_BlackWhiteModeHidden:
		attr.Value = "hidden"
	}
	return attr, nil
}
func (e *ST_BlackWhiteMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "clr":
		*e = 1
	case "auto":
		*e = 2
	case "gray":
		*e = 3
	case "ltGray":
		*e = 4
	case "invGray":
		*e = 5
	case "grayWhite":
		*e = 6
	case "blackGray":
		*e = 7
	case "blackWhite":
		*e = 8
	case "black":
		*e = 9
	case "white":
		*e = 10
	case "hidden":
		*e = 11
	}
	return nil
}
func (m ST_BlackWhiteMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BlackWhiteMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "clr":
			*m = 1
		case "auto":
			*m = 2
		case "gray":
			*m = 3
		case "ltGray":
			*m = 4
		case "invGray":
			*m = 5
		case "grayWhite":
			*m = 6
		case "blackGray":
			*m = 7
		case "blackWhite":
			*m = 8
		case "black":
			*m = 9
		case "white":
			*m = 10
		case "hidden":
			*m = 11
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
func (m ST_BlackWhiteMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "clr"
	case 2:
		return "auto"
	case 3:
		return "gray"
	case 4:
		return "ltGray"
	case 5:
		return "invGray"
	case 6:
		return "grayWhite"
	case 7:
		return "blackGray"
	case 8:
		return "blackWhite"
	case 9:
		return "black"
	case 10:
		return "white"
	case 11:
		return "hidden"
	}
	return ""
}
func (m ST_BlackWhiteMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BlackWhiteMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ChartBuildStep byte

const (
	ST_ChartBuildStepUnset        ST_ChartBuildStep = 0
	ST_ChartBuildStepCategory     ST_ChartBuildStep = 1
	ST_ChartBuildStepPtInCategory ST_ChartBuildStep = 2
	ST_ChartBuildStepSeries       ST_ChartBuildStep = 3
	ST_ChartBuildStepPtInSeries   ST_ChartBuildStep = 4
	ST_ChartBuildStepAllPts       ST_ChartBuildStep = 5
	ST_ChartBuildStepGridLegend   ST_ChartBuildStep = 6
)

func (e ST_ChartBuildStep) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ChartBuildStepUnset:
		attr.Value = ""
	case ST_ChartBuildStepCategory:
		attr.Value = "category"
	case ST_ChartBuildStepPtInCategory:
		attr.Value = "ptInCategory"
	case ST_ChartBuildStepSeries:
		attr.Value = "series"
	case ST_ChartBuildStepPtInSeries:
		attr.Value = "ptInSeries"
	case ST_ChartBuildStepAllPts:
		attr.Value = "allPts"
	case ST_ChartBuildStepGridLegend:
		attr.Value = "gridLegend"
	}
	return attr, nil
}
func (e *ST_ChartBuildStep) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "category":
		*e = 1
	case "ptInCategory":
		*e = 2
	case "series":
		*e = 3
	case "ptInSeries":
		*e = 4
	case "allPts":
		*e = 5
	case "gridLegend":
		*e = 6
	}
	return nil
}
func (m ST_ChartBuildStep) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ChartBuildStep) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "category":
			*m = 1
		case "ptInCategory":
			*m = 2
		case "series":
			*m = 3
		case "ptInSeries":
			*m = 4
		case "allPts":
			*m = 5
		case "gridLegend":
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
func (m ST_ChartBuildStep) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "category"
	case 2:
		return "ptInCategory"
	case 3:
		return "series"
	case 4:
		return "ptInSeries"
	case 5:
		return "allPts"
	case 6:
		return "gridLegend"
	}
	return ""
}
func (m ST_ChartBuildStep) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ChartBuildStep) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DgmBuildStep byte

const (
	ST_DgmBuildStepUnset ST_DgmBuildStep = 0
	ST_DgmBuildStepSp    ST_DgmBuildStep = 1
	ST_DgmBuildStepBg    ST_DgmBuildStep = 2
)

func (e ST_DgmBuildStep) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DgmBuildStepUnset:
		attr.Value = ""
	case ST_DgmBuildStepSp:
		attr.Value = "sp"
	case ST_DgmBuildStepBg:
		attr.Value = "bg"
	}
	return attr, nil
}
func (e *ST_DgmBuildStep) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sp":
		*e = 1
	case "bg":
		*e = 2
	}
	return nil
}
func (m ST_DgmBuildStep) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_DgmBuildStep) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sp":
			*m = 1
		case "bg":
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
func (m ST_DgmBuildStep) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sp"
	case 2:
		return "bg"
	}
	return ""
}
func (m ST_DgmBuildStep) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_DgmBuildStep) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnimationBuildType byte

const (
	ST_AnimationBuildTypeUnset     ST_AnimationBuildType = 0
	ST_AnimationBuildTypeAllAtOnce ST_AnimationBuildType = 1
)

func (e ST_AnimationBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnimationBuildTypeUnset:
		attr.Value = ""
	case ST_AnimationBuildTypeAllAtOnce:
		attr.Value = "allAtOnce"
	}
	return attr, nil
}
func (e *ST_AnimationBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "allAtOnce":
		*e = 1
	}
	return nil
}
func (m ST_AnimationBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AnimationBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "allAtOnce":
			*m = 1
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
func (m ST_AnimationBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "allAtOnce"
	}
	return ""
}
func (m ST_AnimationBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AnimationBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnimationDgmOnlyBuildType byte

const (
	ST_AnimationDgmOnlyBuildTypeUnset     ST_AnimationDgmOnlyBuildType = 0
	ST_AnimationDgmOnlyBuildTypeOne       ST_AnimationDgmOnlyBuildType = 1
	ST_AnimationDgmOnlyBuildTypeLvlOne    ST_AnimationDgmOnlyBuildType = 2
	ST_AnimationDgmOnlyBuildTypeLvlAtOnce ST_AnimationDgmOnlyBuildType = 3
)

func (e ST_AnimationDgmOnlyBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnimationDgmOnlyBuildTypeUnset:
		attr.Value = ""
	case ST_AnimationDgmOnlyBuildTypeOne:
		attr.Value = "one"
	case ST_AnimationDgmOnlyBuildTypeLvlOne:
		attr.Value = "lvlOne"
	case ST_AnimationDgmOnlyBuildTypeLvlAtOnce:
		attr.Value = "lvlAtOnce"
	}
	return attr, nil
}
func (e *ST_AnimationDgmOnlyBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "one":
		*e = 1
	case "lvlOne":
		*e = 2
	case "lvlAtOnce":
		*e = 3
	}
	return nil
}
func (m ST_AnimationDgmOnlyBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AnimationDgmOnlyBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "one":
			*m = 1
		case "lvlOne":
			*m = 2
		case "lvlAtOnce":
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
func (m ST_AnimationDgmOnlyBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "one"
	case 2:
		return "lvlOne"
	case 3:
		return "lvlAtOnce"
	}
	return ""
}
func (m ST_AnimationDgmOnlyBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AnimationDgmOnlyBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnimationChartOnlyBuildType byte

const (
	ST_AnimationChartOnlyBuildTypeUnset      ST_AnimationChartOnlyBuildType = 0
	ST_AnimationChartOnlyBuildTypeSeries     ST_AnimationChartOnlyBuildType = 1
	ST_AnimationChartOnlyBuildTypeCategory   ST_AnimationChartOnlyBuildType = 2
	ST_AnimationChartOnlyBuildTypeSeriesEl   ST_AnimationChartOnlyBuildType = 3
	ST_AnimationChartOnlyBuildTypeCategoryEl ST_AnimationChartOnlyBuildType = 4
)

func (e ST_AnimationChartOnlyBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnimationChartOnlyBuildTypeUnset:
		attr.Value = ""
	case ST_AnimationChartOnlyBuildTypeSeries:
		attr.Value = "series"
	case ST_AnimationChartOnlyBuildTypeCategory:
		attr.Value = "category"
	case ST_AnimationChartOnlyBuildTypeSeriesEl:
		attr.Value = "seriesEl"
	case ST_AnimationChartOnlyBuildTypeCategoryEl:
		attr.Value = "categoryEl"
	}
	return attr, nil
}
func (e *ST_AnimationChartOnlyBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "series":
		*e = 1
	case "category":
		*e = 2
	case "seriesEl":
		*e = 3
	case "categoryEl":
		*e = 4
	}
	return nil
}
func (m ST_AnimationChartOnlyBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_AnimationChartOnlyBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "series":
			*m = 1
		case "category":
			*m = 2
		case "seriesEl":
			*m = 3
		case "categoryEl":
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
func (m ST_AnimationChartOnlyBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "series"
	case 2:
		return "category"
	case 3:
		return "seriesEl"
	case 4:
		return "categoryEl"
	}
	return ""
}
func (m ST_AnimationChartOnlyBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_AnimationChartOnlyBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetCameraType byte

const (
	ST_PresetCameraTypeUnset                               ST_PresetCameraType = 0
	ST_PresetCameraTypeLegacyObliqueTopLeft                ST_PresetCameraType = 1
	ST_PresetCameraTypeLegacyObliqueTop                    ST_PresetCameraType = 2
	ST_PresetCameraTypeLegacyObliqueTopRight               ST_PresetCameraType = 3
	ST_PresetCameraTypeLegacyObliqueLeft                   ST_PresetCameraType = 4
	ST_PresetCameraTypeLegacyObliqueFront                  ST_PresetCameraType = 5
	ST_PresetCameraTypeLegacyObliqueRight                  ST_PresetCameraType = 6
	ST_PresetCameraTypeLegacyObliqueBottomLeft             ST_PresetCameraType = 7
	ST_PresetCameraTypeLegacyObliqueBottom                 ST_PresetCameraType = 8
	ST_PresetCameraTypeLegacyObliqueBottomRight            ST_PresetCameraType = 9
	ST_PresetCameraTypeLegacyPerspectiveTopLeft            ST_PresetCameraType = 10
	ST_PresetCameraTypeLegacyPerspectiveTop                ST_PresetCameraType = 11
	ST_PresetCameraTypeLegacyPerspectiveTopRight           ST_PresetCameraType = 12
	ST_PresetCameraTypeLegacyPerspectiveLeft               ST_PresetCameraType = 13
	ST_PresetCameraTypeLegacyPerspectiveFront              ST_PresetCameraType = 14
	ST_PresetCameraTypeLegacyPerspectiveRight              ST_PresetCameraType = 15
	ST_PresetCameraTypeLegacyPerspectiveBottomLeft         ST_PresetCameraType = 16
	ST_PresetCameraTypeLegacyPerspectiveBottom             ST_PresetCameraType = 17
	ST_PresetCameraTypeLegacyPerspectiveBottomRight        ST_PresetCameraType = 18
	ST_PresetCameraTypeOrthographicFront                   ST_PresetCameraType = 19
	ST_PresetCameraTypeIsometricTopUp                      ST_PresetCameraType = 20
	ST_PresetCameraTypeIsometricTopDown                    ST_PresetCameraType = 21
	ST_PresetCameraTypeIsometricBottomUp                   ST_PresetCameraType = 22
	ST_PresetCameraTypeIsometricBottomDown                 ST_PresetCameraType = 23
	ST_PresetCameraTypeIsometricLeftUp                     ST_PresetCameraType = 24
	ST_PresetCameraTypeIsometricLeftDown                   ST_PresetCameraType = 25
	ST_PresetCameraTypeIsometricRightUp                    ST_PresetCameraType = 26
	ST_PresetCameraTypeIsometricRightDown                  ST_PresetCameraType = 27
	ST_PresetCameraTypeIsometricOffAxis1Left               ST_PresetCameraType = 28
	ST_PresetCameraTypeIsometricOffAxis1Right              ST_PresetCameraType = 29
	ST_PresetCameraTypeIsometricOffAxis1Top                ST_PresetCameraType = 30
	ST_PresetCameraTypeIsometricOffAxis2Left               ST_PresetCameraType = 31
	ST_PresetCameraTypeIsometricOffAxis2Right              ST_PresetCameraType = 32
	ST_PresetCameraTypeIsometricOffAxis2Top                ST_PresetCameraType = 33
	ST_PresetCameraTypeIsometricOffAxis3Left               ST_PresetCameraType = 34
	ST_PresetCameraTypeIsometricOffAxis3Right              ST_PresetCameraType = 35
	ST_PresetCameraTypeIsometricOffAxis3Bottom             ST_PresetCameraType = 36
	ST_PresetCameraTypeIsometricOffAxis4Left               ST_PresetCameraType = 37
	ST_PresetCameraTypeIsometricOffAxis4Right              ST_PresetCameraType = 38
	ST_PresetCameraTypeIsometricOffAxis4Bottom             ST_PresetCameraType = 39
	ST_PresetCameraTypeObliqueTopLeft                      ST_PresetCameraType = 40
	ST_PresetCameraTypeObliqueTop                          ST_PresetCameraType = 41
	ST_PresetCameraTypeObliqueTopRight                     ST_PresetCameraType = 42
	ST_PresetCameraTypeObliqueLeft                         ST_PresetCameraType = 43
	ST_PresetCameraTypeObliqueRight                        ST_PresetCameraType = 44
	ST_PresetCameraTypeObliqueBottomLeft                   ST_PresetCameraType = 45
	ST_PresetCameraTypeObliqueBottom                       ST_PresetCameraType = 46
	ST_PresetCameraTypeObliqueBottomRight                  ST_PresetCameraType = 47
	ST_PresetCameraTypePerspectiveFront                    ST_PresetCameraType = 48
	ST_PresetCameraTypePerspectiveLeft                     ST_PresetCameraType = 49
	ST_PresetCameraTypePerspectiveRight                    ST_PresetCameraType = 50
	ST_PresetCameraTypePerspectiveAbove                    ST_PresetCameraType = 51
	ST_PresetCameraTypePerspectiveBelow                    ST_PresetCameraType = 52
	ST_PresetCameraTypePerspectiveAboveLeftFacing          ST_PresetCameraType = 53
	ST_PresetCameraTypePerspectiveAboveRightFacing         ST_PresetCameraType = 54
	ST_PresetCameraTypePerspectiveContrastingLeftFacing    ST_PresetCameraType = 55
	ST_PresetCameraTypePerspectiveContrastingRightFacing   ST_PresetCameraType = 56
	ST_PresetCameraTypePerspectiveHeroicLeftFacing         ST_PresetCameraType = 57
	ST_PresetCameraTypePerspectiveHeroicRightFacing        ST_PresetCameraType = 58
	ST_PresetCameraTypePerspectiveHeroicExtremeLeftFacing  ST_PresetCameraType = 59
	ST_PresetCameraTypePerspectiveHeroicExtremeRightFacing ST_PresetCameraType = 60
	ST_PresetCameraTypePerspectiveRelaxed                  ST_PresetCameraType = 61
	ST_PresetCameraTypePerspectiveRelaxedModerately        ST_PresetCameraType = 62
)

func (e ST_PresetCameraType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetCameraTypeUnset:
		attr.Value = ""
	case ST_PresetCameraTypeLegacyObliqueTopLeft:
		attr.Value = "legacyObliqueTopLeft"
	case ST_PresetCameraTypeLegacyObliqueTop:
		attr.Value = "legacyObliqueTop"
	case ST_PresetCameraTypeLegacyObliqueTopRight:
		attr.Value = "legacyObliqueTopRight"
	case ST_PresetCameraTypeLegacyObliqueLeft:
		attr.Value = "legacyObliqueLeft"
	case ST_PresetCameraTypeLegacyObliqueFront:
		attr.Value = "legacyObliqueFront"
	case ST_PresetCameraTypeLegacyObliqueRight:
		attr.Value = "legacyObliqueRight"
	case ST_PresetCameraTypeLegacyObliqueBottomLeft:
		attr.Value = "legacyObliqueBottomLeft"
	case ST_PresetCameraTypeLegacyObliqueBottom:
		attr.Value = "legacyObliqueBottom"
	case ST_PresetCameraTypeLegacyObliqueBottomRight:
		attr.Value = "legacyObliqueBottomRight"
	case ST_PresetCameraTypeLegacyPerspectiveTopLeft:
		attr.Value = "legacyPerspectiveTopLeft"
	case ST_PresetCameraTypeLegacyPerspectiveTop:
		attr.Value = "legacyPerspectiveTop"
	case ST_PresetCameraTypeLegacyPerspectiveTopRight:
		attr.Value = "legacyPerspectiveTopRight"
	case ST_PresetCameraTypeLegacyPerspectiveLeft:
		attr.Value = "legacyPerspectiveLeft"
	case ST_PresetCameraTypeLegacyPerspectiveFront:
		attr.Value = "legacyPerspectiveFront"
	case ST_PresetCameraTypeLegacyPerspectiveRight:
		attr.Value = "legacyPerspectiveRight"
	case ST_PresetCameraTypeLegacyPerspectiveBottomLeft:
		attr.Value = "legacyPerspectiveBottomLeft"
	case ST_PresetCameraTypeLegacyPerspectiveBottom:
		attr.Value = "legacyPerspectiveBottom"
	case ST_PresetCameraTypeLegacyPerspectiveBottomRight:
		attr.Value = "legacyPerspectiveBottomRight"
	case ST_PresetCameraTypeOrthographicFront:
		attr.Value = "orthographicFront"
	case ST_PresetCameraTypeIsometricTopUp:
		attr.Value = "isometricTopUp"
	case ST_PresetCameraTypeIsometricTopDown:
		attr.Value = "isometricTopDown"
	case ST_PresetCameraTypeIsometricBottomUp:
		attr.Value = "isometricBottomUp"
	case ST_PresetCameraTypeIsometricBottomDown:
		attr.Value = "isometricBottomDown"
	case ST_PresetCameraTypeIsometricLeftUp:
		attr.Value = "isometricLeftUp"
	case ST_PresetCameraTypeIsometricLeftDown:
		attr.Value = "isometricLeftDown"
	case ST_PresetCameraTypeIsometricRightUp:
		attr.Value = "isometricRightUp"
	case ST_PresetCameraTypeIsometricRightDown:
		attr.Value = "isometricRightDown"
	case ST_PresetCameraTypeIsometricOffAxis1Left:
		attr.Value = "isometricOffAxis1Left"
	case ST_PresetCameraTypeIsometricOffAxis1Right:
		attr.Value = "isometricOffAxis1Right"
	case ST_PresetCameraTypeIsometricOffAxis1Top:
		attr.Value = "isometricOffAxis1Top"
	case ST_PresetCameraTypeIsometricOffAxis2Left:
		attr.Value = "isometricOffAxis2Left"
	case ST_PresetCameraTypeIsometricOffAxis2Right:
		attr.Value = "isometricOffAxis2Right"
	case ST_PresetCameraTypeIsometricOffAxis2Top:
		attr.Value = "isometricOffAxis2Top"
	case ST_PresetCameraTypeIsometricOffAxis3Left:
		attr.Value = "isometricOffAxis3Left"
	case ST_PresetCameraTypeIsometricOffAxis3Right:
		attr.Value = "isometricOffAxis3Right"
	case ST_PresetCameraTypeIsometricOffAxis3Bottom:
		attr.Value = "isometricOffAxis3Bottom"
	case ST_PresetCameraTypeIsometricOffAxis4Left:
		attr.Value = "isometricOffAxis4Left"
	case ST_PresetCameraTypeIsometricOffAxis4Right:
		attr.Value = "isometricOffAxis4Right"
	case ST_PresetCameraTypeIsometricOffAxis4Bottom:
		attr.Value = "isometricOffAxis4Bottom"
	case ST_PresetCameraTypeObliqueTopLeft:
		attr.Value = "obliqueTopLeft"
	case ST_PresetCameraTypeObliqueTop:
		attr.Value = "obliqueTop"
	case ST_PresetCameraTypeObliqueTopRight:
		attr.Value = "obliqueTopRight"
	case ST_PresetCameraTypeObliqueLeft:
		attr.Value = "obliqueLeft"
	case ST_PresetCameraTypeObliqueRight:
		attr.Value = "obliqueRight"
	case ST_PresetCameraTypeObliqueBottomLeft:
		attr.Value = "obliqueBottomLeft"
	case ST_PresetCameraTypeObliqueBottom:
		attr.Value = "obliqueBottom"
	case ST_PresetCameraTypeObliqueBottomRight:
		attr.Value = "obliqueBottomRight"
	case ST_PresetCameraTypePerspectiveFront:
		attr.Value = "perspectiveFront"
	case ST_PresetCameraTypePerspectiveLeft:
		attr.Value = "perspectiveLeft"
	case ST_PresetCameraTypePerspectiveRight:
		attr.Value = "perspectiveRight"
	case ST_PresetCameraTypePerspectiveAbove:
		attr.Value = "perspectiveAbove"
	case ST_PresetCameraTypePerspectiveBelow:
		attr.Value = "perspectiveBelow"
	case ST_PresetCameraTypePerspectiveAboveLeftFacing:
		attr.Value = "perspectiveAboveLeftFacing"
	case ST_PresetCameraTypePerspectiveAboveRightFacing:
		attr.Value = "perspectiveAboveRightFacing"
	case ST_PresetCameraTypePerspectiveContrastingLeftFacing:
		attr.Value = "perspectiveContrastingLeftFacing"
	case ST_PresetCameraTypePerspectiveContrastingRightFacing:
		attr.Value = "perspectiveContrastingRightFacing"
	case ST_PresetCameraTypePerspectiveHeroicLeftFacing:
		attr.Value = "perspectiveHeroicLeftFacing"
	case ST_PresetCameraTypePerspectiveHeroicRightFacing:
		attr.Value = "perspectiveHeroicRightFacing"
	case ST_PresetCameraTypePerspectiveHeroicExtremeLeftFacing:
		attr.Value = "perspectiveHeroicExtremeLeftFacing"
	case ST_PresetCameraTypePerspectiveHeroicExtremeRightFacing:
		attr.Value = "perspectiveHeroicExtremeRightFacing"
	case ST_PresetCameraTypePerspectiveRelaxed:
		attr.Value = "perspectiveRelaxed"
	case ST_PresetCameraTypePerspectiveRelaxedModerately:
		attr.Value = "perspectiveRelaxedModerately"
	}
	return attr, nil
}
func (e *ST_PresetCameraType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "legacyObliqueTopLeft":
		*e = 1
	case "legacyObliqueTop":
		*e = 2
	case "legacyObliqueTopRight":
		*e = 3
	case "legacyObliqueLeft":
		*e = 4
	case "legacyObliqueFront":
		*e = 5
	case "legacyObliqueRight":
		*e = 6
	case "legacyObliqueBottomLeft":
		*e = 7
	case "legacyObliqueBottom":
		*e = 8
	case "legacyObliqueBottomRight":
		*e = 9
	case "legacyPerspectiveTopLeft":
		*e = 10
	case "legacyPerspectiveTop":
		*e = 11
	case "legacyPerspectiveTopRight":
		*e = 12
	case "legacyPerspectiveLeft":
		*e = 13
	case "legacyPerspectiveFront":
		*e = 14
	case "legacyPerspectiveRight":
		*e = 15
	case "legacyPerspectiveBottomLeft":
		*e = 16
	case "legacyPerspectiveBottom":
		*e = 17
	case "legacyPerspectiveBottomRight":
		*e = 18
	case "orthographicFront":
		*e = 19
	case "isometricTopUp":
		*e = 20
	case "isometricTopDown":
		*e = 21
	case "isometricBottomUp":
		*e = 22
	case "isometricBottomDown":
		*e = 23
	case "isometricLeftUp":
		*e = 24
	case "isometricLeftDown":
		*e = 25
	case "isometricRightUp":
		*e = 26
	case "isometricRightDown":
		*e = 27
	case "isometricOffAxis1Left":
		*e = 28
	case "isometricOffAxis1Right":
		*e = 29
	case "isometricOffAxis1Top":
		*e = 30
	case "isometricOffAxis2Left":
		*e = 31
	case "isometricOffAxis2Right":
		*e = 32
	case "isometricOffAxis2Top":
		*e = 33
	case "isometricOffAxis3Left":
		*e = 34
	case "isometricOffAxis3Right":
		*e = 35
	case "isometricOffAxis3Bottom":
		*e = 36
	case "isometricOffAxis4Left":
		*e = 37
	case "isometricOffAxis4Right":
		*e = 38
	case "isometricOffAxis4Bottom":
		*e = 39
	case "obliqueTopLeft":
		*e = 40
	case "obliqueTop":
		*e = 41
	case "obliqueTopRight":
		*e = 42
	case "obliqueLeft":
		*e = 43
	case "obliqueRight":
		*e = 44
	case "obliqueBottomLeft":
		*e = 45
	case "obliqueBottom":
		*e = 46
	case "obliqueBottomRight":
		*e = 47
	case "perspectiveFront":
		*e = 48
	case "perspectiveLeft":
		*e = 49
	case "perspectiveRight":
		*e = 50
	case "perspectiveAbove":
		*e = 51
	case "perspectiveBelow":
		*e = 52
	case "perspectiveAboveLeftFacing":
		*e = 53
	case "perspectiveAboveRightFacing":
		*e = 54
	case "perspectiveContrastingLeftFacing":
		*e = 55
	case "perspectiveContrastingRightFacing":
		*e = 56
	case "perspectiveHeroicLeftFacing":
		*e = 57
	case "perspectiveHeroicRightFacing":
		*e = 58
	case "perspectiveHeroicExtremeLeftFacing":
		*e = 59
	case "perspectiveHeroicExtremeRightFacing":
		*e = 60
	case "perspectiveRelaxed":
		*e = 61
	case "perspectiveRelaxedModerately":
		*e = 62
	}
	return nil
}
func (m ST_PresetCameraType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetCameraType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "legacyObliqueTopLeft":
			*m = 1
		case "legacyObliqueTop":
			*m = 2
		case "legacyObliqueTopRight":
			*m = 3
		case "legacyObliqueLeft":
			*m = 4
		case "legacyObliqueFront":
			*m = 5
		case "legacyObliqueRight":
			*m = 6
		case "legacyObliqueBottomLeft":
			*m = 7
		case "legacyObliqueBottom":
			*m = 8
		case "legacyObliqueBottomRight":
			*m = 9
		case "legacyPerspectiveTopLeft":
			*m = 10
		case "legacyPerspectiveTop":
			*m = 11
		case "legacyPerspectiveTopRight":
			*m = 12
		case "legacyPerspectiveLeft":
			*m = 13
		case "legacyPerspectiveFront":
			*m = 14
		case "legacyPerspectiveRight":
			*m = 15
		case "legacyPerspectiveBottomLeft":
			*m = 16
		case "legacyPerspectiveBottom":
			*m = 17
		case "legacyPerspectiveBottomRight":
			*m = 18
		case "orthographicFront":
			*m = 19
		case "isometricTopUp":
			*m = 20
		case "isometricTopDown":
			*m = 21
		case "isometricBottomUp":
			*m = 22
		case "isometricBottomDown":
			*m = 23
		case "isometricLeftUp":
			*m = 24
		case "isometricLeftDown":
			*m = 25
		case "isometricRightUp":
			*m = 26
		case "isometricRightDown":
			*m = 27
		case "isometricOffAxis1Left":
			*m = 28
		case "isometricOffAxis1Right":
			*m = 29
		case "isometricOffAxis1Top":
			*m = 30
		case "isometricOffAxis2Left":
			*m = 31
		case "isometricOffAxis2Right":
			*m = 32
		case "isometricOffAxis2Top":
			*m = 33
		case "isometricOffAxis3Left":
			*m = 34
		case "isometricOffAxis3Right":
			*m = 35
		case "isometricOffAxis3Bottom":
			*m = 36
		case "isometricOffAxis4Left":
			*m = 37
		case "isometricOffAxis4Right":
			*m = 38
		case "isometricOffAxis4Bottom":
			*m = 39
		case "obliqueTopLeft":
			*m = 40
		case "obliqueTop":
			*m = 41
		case "obliqueTopRight":
			*m = 42
		case "obliqueLeft":
			*m = 43
		case "obliqueRight":
			*m = 44
		case "obliqueBottomLeft":
			*m = 45
		case "obliqueBottom":
			*m = 46
		case "obliqueBottomRight":
			*m = 47
		case "perspectiveFront":
			*m = 48
		case "perspectiveLeft":
			*m = 49
		case "perspectiveRight":
			*m = 50
		case "perspectiveAbove":
			*m = 51
		case "perspectiveBelow":
			*m = 52
		case "perspectiveAboveLeftFacing":
			*m = 53
		case "perspectiveAboveRightFacing":
			*m = 54
		case "perspectiveContrastingLeftFacing":
			*m = 55
		case "perspectiveContrastingRightFacing":
			*m = 56
		case "perspectiveHeroicLeftFacing":
			*m = 57
		case "perspectiveHeroicRightFacing":
			*m = 58
		case "perspectiveHeroicExtremeLeftFacing":
			*m = 59
		case "perspectiveHeroicExtremeRightFacing":
			*m = 60
		case "perspectiveRelaxed":
			*m = 61
		case "perspectiveRelaxedModerately":
			*m = 62
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
func (m ST_PresetCameraType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "legacyObliqueTopLeft"
	case 2:
		return "legacyObliqueTop"
	case 3:
		return "legacyObliqueTopRight"
	case 4:
		return "legacyObliqueLeft"
	case 5:
		return "legacyObliqueFront"
	case 6:
		return "legacyObliqueRight"
	case 7:
		return "legacyObliqueBottomLeft"
	case 8:
		return "legacyObliqueBottom"
	case 9:
		return "legacyObliqueBottomRight"
	case 10:
		return "legacyPerspectiveTopLeft"
	case 11:
		return "legacyPerspectiveTop"
	case 12:
		return "legacyPerspectiveTopRight"
	case 13:
		return "legacyPerspectiveLeft"
	case 14:
		return "legacyPerspectiveFront"
	case 15:
		return "legacyPerspectiveRight"
	case 16:
		return "legacyPerspectiveBottomLeft"
	case 17:
		return "legacyPerspectiveBottom"
	case 18:
		return "legacyPerspectiveBottomRight"
	case 19:
		return "orthographicFront"
	case 20:
		return "isometricTopUp"
	case 21:
		return "isometricTopDown"
	case 22:
		return "isometricBottomUp"
	case 23:
		return "isometricBottomDown"
	case 24:
		return "isometricLeftUp"
	case 25:
		return "isometricLeftDown"
	case 26:
		return "isometricRightUp"
	case 27:
		return "isometricRightDown"
	case 28:
		return "isometricOffAxis1Left"
	case 29:
		return "isometricOffAxis1Right"
	case 30:
		return "isometricOffAxis1Top"
	case 31:
		return "isometricOffAxis2Left"
	case 32:
		return "isometricOffAxis2Right"
	case 33:
		return "isometricOffAxis2Top"
	case 34:
		return "isometricOffAxis3Left"
	case 35:
		return "isometricOffAxis3Right"
	case 36:
		return "isometricOffAxis3Bottom"
	case 37:
		return "isometricOffAxis4Left"
	case 38:
		return "isometricOffAxis4Right"
	case 39:
		return "isometricOffAxis4Bottom"
	case 40:
		return "obliqueTopLeft"
	case 41:
		return "obliqueTop"
	case 42:
		return "obliqueTopRight"
	case 43:
		return "obliqueLeft"
	case 44:
		return "obliqueRight"
	case 45:
		return "obliqueBottomLeft"
	case 46:
		return "obliqueBottom"
	case 47:
		return "obliqueBottomRight"
	case 48:
		return "perspectiveFront"
	case 49:
		return "perspectiveLeft"
	case 50:
		return "perspectiveRight"
	case 51:
		return "perspectiveAbove"
	case 52:
		return "perspectiveBelow"
	case 53:
		return "perspectiveAboveLeftFacing"
	case 54:
		return "perspectiveAboveRightFacing"
	case 55:
		return "perspectiveContrastingLeftFacing"
	case 56:
		return "perspectiveContrastingRightFacing"
	case 57:
		return "perspectiveHeroicLeftFacing"
	case 58:
		return "perspectiveHeroicRightFacing"
	case 59:
		return "perspectiveHeroicExtremeLeftFacing"
	case 60:
		return "perspectiveHeroicExtremeRightFacing"
	case 61:
		return "perspectiveRelaxed"
	case 62:
		return "perspectiveRelaxedModerately"
	}
	return ""
}
func (m ST_PresetCameraType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetCameraType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LightRigDirection byte

const (
	ST_LightRigDirectionUnset ST_LightRigDirection = 0
	ST_LightRigDirectionTl    ST_LightRigDirection = 1
	ST_LightRigDirectionT     ST_LightRigDirection = 2
	ST_LightRigDirectionTr    ST_LightRigDirection = 3
	ST_LightRigDirectionL     ST_LightRigDirection = 4
	ST_LightRigDirectionR     ST_LightRigDirection = 5
	ST_LightRigDirectionBl    ST_LightRigDirection = 6
	ST_LightRigDirectionB     ST_LightRigDirection = 7
	ST_LightRigDirectionBr    ST_LightRigDirection = 8
)

func (e ST_LightRigDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LightRigDirectionUnset:
		attr.Value = ""
	case ST_LightRigDirectionTl:
		attr.Value = "tl"
	case ST_LightRigDirectionT:
		attr.Value = "t"
	case ST_LightRigDirectionTr:
		attr.Value = "tr"
	case ST_LightRigDirectionL:
		attr.Value = "l"
	case ST_LightRigDirectionR:
		attr.Value = "r"
	case ST_LightRigDirectionBl:
		attr.Value = "bl"
	case ST_LightRigDirectionB:
		attr.Value = "b"
	case ST_LightRigDirectionBr:
		attr.Value = "br"
	}
	return attr, nil
}
func (e *ST_LightRigDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tl":
		*e = 1
	case "t":
		*e = 2
	case "tr":
		*e = 3
	case "l":
		*e = 4
	case "r":
		*e = 5
	case "bl":
		*e = 6
	case "b":
		*e = 7
	case "br":
		*e = 8
	}
	return nil
}
func (m ST_LightRigDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LightRigDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "tl":
			*m = 1
		case "t":
			*m = 2
		case "tr":
			*m = 3
		case "l":
			*m = 4
		case "r":
			*m = 5
		case "bl":
			*m = 6
		case "b":
			*m = 7
		case "br":
			*m = 8
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
func (m ST_LightRigDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tl"
	case 2:
		return "t"
	case 3:
		return "tr"
	case 4:
		return "l"
	case 5:
		return "r"
	case 6:
		return "bl"
	case 7:
		return "b"
	case 8:
		return "br"
	}
	return ""
}
func (m ST_LightRigDirection) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LightRigDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LightRigType byte

const (
	ST_LightRigTypeUnset         ST_LightRigType = 0
	ST_LightRigTypeLegacyFlat1   ST_LightRigType = 1
	ST_LightRigTypeLegacyFlat2   ST_LightRigType = 2
	ST_LightRigTypeLegacyFlat3   ST_LightRigType = 3
	ST_LightRigTypeLegacyFlat4   ST_LightRigType = 4
	ST_LightRigTypeLegacyNormal1 ST_LightRigType = 5
	ST_LightRigTypeLegacyNormal2 ST_LightRigType = 6
	ST_LightRigTypeLegacyNormal3 ST_LightRigType = 7
	ST_LightRigTypeLegacyNormal4 ST_LightRigType = 8
	ST_LightRigTypeLegacyHarsh1  ST_LightRigType = 9
	ST_LightRigTypeLegacyHarsh2  ST_LightRigType = 10
	ST_LightRigTypeLegacyHarsh3  ST_LightRigType = 11
	ST_LightRigTypeLegacyHarsh4  ST_LightRigType = 12
	ST_LightRigTypeThreePt       ST_LightRigType = 13
	ST_LightRigTypeBalanced      ST_LightRigType = 14
	ST_LightRigTypeSoft          ST_LightRigType = 15
	ST_LightRigTypeHarsh         ST_LightRigType = 16
	ST_LightRigTypeFlood         ST_LightRigType = 17
	ST_LightRigTypeContrasting   ST_LightRigType = 18
	ST_LightRigTypeMorning       ST_LightRigType = 19
	ST_LightRigTypeSunrise       ST_LightRigType = 20
	ST_LightRigTypeSunset        ST_LightRigType = 21
	ST_LightRigTypeChilly        ST_LightRigType = 22
	ST_LightRigTypeFreezing      ST_LightRigType = 23
	ST_LightRigTypeFlat          ST_LightRigType = 24
	ST_LightRigTypeTwoPt         ST_LightRigType = 25
	ST_LightRigTypeGlow          ST_LightRigType = 26
	ST_LightRigTypeBrightRoom    ST_LightRigType = 27
)

func (e ST_LightRigType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LightRigTypeUnset:
		attr.Value = ""
	case ST_LightRigTypeLegacyFlat1:
		attr.Value = "legacyFlat1"
	case ST_LightRigTypeLegacyFlat2:
		attr.Value = "legacyFlat2"
	case ST_LightRigTypeLegacyFlat3:
		attr.Value = "legacyFlat3"
	case ST_LightRigTypeLegacyFlat4:
		attr.Value = "legacyFlat4"
	case ST_LightRigTypeLegacyNormal1:
		attr.Value = "legacyNormal1"
	case ST_LightRigTypeLegacyNormal2:
		attr.Value = "legacyNormal2"
	case ST_LightRigTypeLegacyNormal3:
		attr.Value = "legacyNormal3"
	case ST_LightRigTypeLegacyNormal4:
		attr.Value = "legacyNormal4"
	case ST_LightRigTypeLegacyHarsh1:
		attr.Value = "legacyHarsh1"
	case ST_LightRigTypeLegacyHarsh2:
		attr.Value = "legacyHarsh2"
	case ST_LightRigTypeLegacyHarsh3:
		attr.Value = "legacyHarsh3"
	case ST_LightRigTypeLegacyHarsh4:
		attr.Value = "legacyHarsh4"
	case ST_LightRigTypeThreePt:
		attr.Value = "threePt"
	case ST_LightRigTypeBalanced:
		attr.Value = "balanced"
	case ST_LightRigTypeSoft:
		attr.Value = "soft"
	case ST_LightRigTypeHarsh:
		attr.Value = "harsh"
	case ST_LightRigTypeFlood:
		attr.Value = "flood"
	case ST_LightRigTypeContrasting:
		attr.Value = "contrasting"
	case ST_LightRigTypeMorning:
		attr.Value = "morning"
	case ST_LightRigTypeSunrise:
		attr.Value = "sunrise"
	case ST_LightRigTypeSunset:
		attr.Value = "sunset"
	case ST_LightRigTypeChilly:
		attr.Value = "chilly"
	case ST_LightRigTypeFreezing:
		attr.Value = "freezing"
	case ST_LightRigTypeFlat:
		attr.Value = "flat"
	case ST_LightRigTypeTwoPt:
		attr.Value = "twoPt"
	case ST_LightRigTypeGlow:
		attr.Value = "glow"
	case ST_LightRigTypeBrightRoom:
		attr.Value = "brightRoom"
	}
	return attr, nil
}
func (e *ST_LightRigType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "legacyFlat1":
		*e = 1
	case "legacyFlat2":
		*e = 2
	case "legacyFlat3":
		*e = 3
	case "legacyFlat4":
		*e = 4
	case "legacyNormal1":
		*e = 5
	case "legacyNormal2":
		*e = 6
	case "legacyNormal3":
		*e = 7
	case "legacyNormal4":
		*e = 8
	case "legacyHarsh1":
		*e = 9
	case "legacyHarsh2":
		*e = 10
	case "legacyHarsh3":
		*e = 11
	case "legacyHarsh4":
		*e = 12
	case "threePt":
		*e = 13
	case "balanced":
		*e = 14
	case "soft":
		*e = 15
	case "harsh":
		*e = 16
	case "flood":
		*e = 17
	case "contrasting":
		*e = 18
	case "morning":
		*e = 19
	case "sunrise":
		*e = 20
	case "sunset":
		*e = 21
	case "chilly":
		*e = 22
	case "freezing":
		*e = 23
	case "flat":
		*e = 24
	case "twoPt":
		*e = 25
	case "glow":
		*e = 26
	case "brightRoom":
		*e = 27
	}
	return nil
}
func (m ST_LightRigType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LightRigType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "legacyFlat1":
			*m = 1
		case "legacyFlat2":
			*m = 2
		case "legacyFlat3":
			*m = 3
		case "legacyFlat4":
			*m = 4
		case "legacyNormal1":
			*m = 5
		case "legacyNormal2":
			*m = 6
		case "legacyNormal3":
			*m = 7
		case "legacyNormal4":
			*m = 8
		case "legacyHarsh1":
			*m = 9
		case "legacyHarsh2":
			*m = 10
		case "legacyHarsh3":
			*m = 11
		case "legacyHarsh4":
			*m = 12
		case "threePt":
			*m = 13
		case "balanced":
			*m = 14
		case "soft":
			*m = 15
		case "harsh":
			*m = 16
		case "flood":
			*m = 17
		case "contrasting":
			*m = 18
		case "morning":
			*m = 19
		case "sunrise":
			*m = 20
		case "sunset":
			*m = 21
		case "chilly":
			*m = 22
		case "freezing":
			*m = 23
		case "flat":
			*m = 24
		case "twoPt":
			*m = 25
		case "glow":
			*m = 26
		case "brightRoom":
			*m = 27
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
func (m ST_LightRigType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "legacyFlat1"
	case 2:
		return "legacyFlat2"
	case 3:
		return "legacyFlat3"
	case 4:
		return "legacyFlat4"
	case 5:
		return "legacyNormal1"
	case 6:
		return "legacyNormal2"
	case 7:
		return "legacyNormal3"
	case 8:
		return "legacyNormal4"
	case 9:
		return "legacyHarsh1"
	case 10:
		return "legacyHarsh2"
	case 11:
		return "legacyHarsh3"
	case 12:
		return "legacyHarsh4"
	case 13:
		return "threePt"
	case 14:
		return "balanced"
	case 15:
		return "soft"
	case 16:
		return "harsh"
	case 17:
		return "flood"
	case 18:
		return "contrasting"
	case 19:
		return "morning"
	case 20:
		return "sunrise"
	case 21:
		return "sunset"
	case 22:
		return "chilly"
	case 23:
		return "freezing"
	case 24:
		return "flat"
	case 25:
		return "twoPt"
	case 26:
		return "glow"
	case 27:
		return "brightRoom"
	}
	return ""
}
func (m ST_LightRigType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LightRigType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BevelPresetType byte

const (
	ST_BevelPresetTypeUnset        ST_BevelPresetType = 0
	ST_BevelPresetTypeRelaxedInset ST_BevelPresetType = 1
	ST_BevelPresetTypeCircle       ST_BevelPresetType = 2
	ST_BevelPresetTypeSlope        ST_BevelPresetType = 3
	ST_BevelPresetTypeCross        ST_BevelPresetType = 4
	ST_BevelPresetTypeAngle        ST_BevelPresetType = 5
	ST_BevelPresetTypeSoftRound    ST_BevelPresetType = 6
	ST_BevelPresetTypeConvex       ST_BevelPresetType = 7
	ST_BevelPresetTypeCoolSlant    ST_BevelPresetType = 8
	ST_BevelPresetTypeDivot        ST_BevelPresetType = 9
	ST_BevelPresetTypeRiblet       ST_BevelPresetType = 10
	ST_BevelPresetTypeHardEdge     ST_BevelPresetType = 11
	ST_BevelPresetTypeArtDeco      ST_BevelPresetType = 12
)

func (e ST_BevelPresetType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BevelPresetTypeUnset:
		attr.Value = ""
	case ST_BevelPresetTypeRelaxedInset:
		attr.Value = "relaxedInset"
	case ST_BevelPresetTypeCircle:
		attr.Value = "circle"
	case ST_BevelPresetTypeSlope:
		attr.Value = "slope"
	case ST_BevelPresetTypeCross:
		attr.Value = "cross"
	case ST_BevelPresetTypeAngle:
		attr.Value = "angle"
	case ST_BevelPresetTypeSoftRound:
		attr.Value = "softRound"
	case ST_BevelPresetTypeConvex:
		attr.Value = "convex"
	case ST_BevelPresetTypeCoolSlant:
		attr.Value = "coolSlant"
	case ST_BevelPresetTypeDivot:
		attr.Value = "divot"
	case ST_BevelPresetTypeRiblet:
		attr.Value = "riblet"
	case ST_BevelPresetTypeHardEdge:
		attr.Value = "hardEdge"
	case ST_BevelPresetTypeArtDeco:
		attr.Value = "artDeco"
	}
	return attr, nil
}
func (e *ST_BevelPresetType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "relaxedInset":
		*e = 1
	case "circle":
		*e = 2
	case "slope":
		*e = 3
	case "cross":
		*e = 4
	case "angle":
		*e = 5
	case "softRound":
		*e = 6
	case "convex":
		*e = 7
	case "coolSlant":
		*e = 8
	case "divot":
		*e = 9
	case "riblet":
		*e = 10
	case "hardEdge":
		*e = 11
	case "artDeco":
		*e = 12
	}
	return nil
}
func (m ST_BevelPresetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BevelPresetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "relaxedInset":
			*m = 1
		case "circle":
			*m = 2
		case "slope":
			*m = 3
		case "cross":
			*m = 4
		case "angle":
			*m = 5
		case "softRound":
			*m = 6
		case "convex":
			*m = 7
		case "coolSlant":
			*m = 8
		case "divot":
			*m = 9
		case "riblet":
			*m = 10
		case "hardEdge":
			*m = 11
		case "artDeco":
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
func (m ST_BevelPresetType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "relaxedInset"
	case 2:
		return "circle"
	case 3:
		return "slope"
	case 4:
		return "cross"
	case 5:
		return "angle"
	case 6:
		return "softRound"
	case 7:
		return "convex"
	case 8:
		return "coolSlant"
	case 9:
		return "divot"
	case 10:
		return "riblet"
	case 11:
		return "hardEdge"
	case 12:
		return "artDeco"
	}
	return ""
}
func (m ST_BevelPresetType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BevelPresetType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetMaterialType byte

const (
	ST_PresetMaterialTypeUnset             ST_PresetMaterialType = 0
	ST_PresetMaterialTypeLegacyMatte       ST_PresetMaterialType = 1
	ST_PresetMaterialTypeLegacyPlastic     ST_PresetMaterialType = 2
	ST_PresetMaterialTypeLegacyMetal       ST_PresetMaterialType = 3
	ST_PresetMaterialTypeLegacyWireframe   ST_PresetMaterialType = 4
	ST_PresetMaterialTypeMatte             ST_PresetMaterialType = 5
	ST_PresetMaterialTypePlastic           ST_PresetMaterialType = 6
	ST_PresetMaterialTypeMetal             ST_PresetMaterialType = 7
	ST_PresetMaterialTypeWarmMatte         ST_PresetMaterialType = 8
	ST_PresetMaterialTypeTranslucentPowder ST_PresetMaterialType = 9
	ST_PresetMaterialTypePowder            ST_PresetMaterialType = 10
	ST_PresetMaterialTypeDkEdge            ST_PresetMaterialType = 11
	ST_PresetMaterialTypeSoftEdge          ST_PresetMaterialType = 12
	ST_PresetMaterialTypeClear             ST_PresetMaterialType = 13
	ST_PresetMaterialTypeFlat              ST_PresetMaterialType = 14
	ST_PresetMaterialTypeSoftmetal         ST_PresetMaterialType = 15
)

func (e ST_PresetMaterialType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetMaterialTypeUnset:
		attr.Value = ""
	case ST_PresetMaterialTypeLegacyMatte:
		attr.Value = "legacyMatte"
	case ST_PresetMaterialTypeLegacyPlastic:
		attr.Value = "legacyPlastic"
	case ST_PresetMaterialTypeLegacyMetal:
		attr.Value = "legacyMetal"
	case ST_PresetMaterialTypeLegacyWireframe:
		attr.Value = "legacyWireframe"
	case ST_PresetMaterialTypeMatte:
		attr.Value = "matte"
	case ST_PresetMaterialTypePlastic:
		attr.Value = "plastic"
	case ST_PresetMaterialTypeMetal:
		attr.Value = "metal"
	case ST_PresetMaterialTypeWarmMatte:
		attr.Value = "warmMatte"
	case ST_PresetMaterialTypeTranslucentPowder:
		attr.Value = "translucentPowder"
	case ST_PresetMaterialTypePowder:
		attr.Value = "powder"
	case ST_PresetMaterialTypeDkEdge:
		attr.Value = "dkEdge"
	case ST_PresetMaterialTypeSoftEdge:
		attr.Value = "softEdge"
	case ST_PresetMaterialTypeClear:
		attr.Value = "clear"
	case ST_PresetMaterialTypeFlat:
		attr.Value = "flat"
	case ST_PresetMaterialTypeSoftmetal:
		attr.Value = "softmetal"
	}
	return attr, nil
}
func (e *ST_PresetMaterialType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "legacyMatte":
		*e = 1
	case "legacyPlastic":
		*e = 2
	case "legacyMetal":
		*e = 3
	case "legacyWireframe":
		*e = 4
	case "matte":
		*e = 5
	case "plastic":
		*e = 6
	case "metal":
		*e = 7
	case "warmMatte":
		*e = 8
	case "translucentPowder":
		*e = 9
	case "powder":
		*e = 10
	case "dkEdge":
		*e = 11
	case "softEdge":
		*e = 12
	case "clear":
		*e = 13
	case "flat":
		*e = 14
	case "softmetal":
		*e = 15
	}
	return nil
}
func (m ST_PresetMaterialType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetMaterialType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "legacyMatte":
			*m = 1
		case "legacyPlastic":
			*m = 2
		case "legacyMetal":
			*m = 3
		case "legacyWireframe":
			*m = 4
		case "matte":
			*m = 5
		case "plastic":
			*m = 6
		case "metal":
			*m = 7
		case "warmMatte":
			*m = 8
		case "translucentPowder":
			*m = 9
		case "powder":
			*m = 10
		case "dkEdge":
			*m = 11
		case "softEdge":
			*m = 12
		case "clear":
			*m = 13
		case "flat":
			*m = 14
		case "softmetal":
			*m = 15
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
func (m ST_PresetMaterialType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "legacyMatte"
	case 2:
		return "legacyPlastic"
	case 3:
		return "legacyMetal"
	case 4:
		return "legacyWireframe"
	case 5:
		return "matte"
	case 6:
		return "plastic"
	case 7:
		return "metal"
	case 8:
		return "warmMatte"
	case 9:
		return "translucentPowder"
	case 10:
		return "powder"
	case 11:
		return "dkEdge"
	case 12:
		return "softEdge"
	case 13:
		return "clear"
	case 14:
		return "flat"
	case 15:
		return "softmetal"
	}
	return ""
}
func (m ST_PresetMaterialType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetMaterialType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetShadowVal byte

const (
	ST_PresetShadowValUnset  ST_PresetShadowVal = 0
	ST_PresetShadowValShdw1  ST_PresetShadowVal = 1
	ST_PresetShadowValShdw2  ST_PresetShadowVal = 2
	ST_PresetShadowValShdw3  ST_PresetShadowVal = 3
	ST_PresetShadowValShdw4  ST_PresetShadowVal = 4
	ST_PresetShadowValShdw5  ST_PresetShadowVal = 5
	ST_PresetShadowValShdw6  ST_PresetShadowVal = 6
	ST_PresetShadowValShdw7  ST_PresetShadowVal = 7
	ST_PresetShadowValShdw8  ST_PresetShadowVal = 8
	ST_PresetShadowValShdw9  ST_PresetShadowVal = 9
	ST_PresetShadowValShdw10 ST_PresetShadowVal = 10
	ST_PresetShadowValShdw11 ST_PresetShadowVal = 11
	ST_PresetShadowValShdw12 ST_PresetShadowVal = 12
	ST_PresetShadowValShdw13 ST_PresetShadowVal = 13
	ST_PresetShadowValShdw14 ST_PresetShadowVal = 14
	ST_PresetShadowValShdw15 ST_PresetShadowVal = 15
	ST_PresetShadowValShdw16 ST_PresetShadowVal = 16
	ST_PresetShadowValShdw17 ST_PresetShadowVal = 17
	ST_PresetShadowValShdw18 ST_PresetShadowVal = 18
	ST_PresetShadowValShdw19 ST_PresetShadowVal = 19
	ST_PresetShadowValShdw20 ST_PresetShadowVal = 20
)

func (e ST_PresetShadowVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetShadowValUnset:
		attr.Value = ""
	case ST_PresetShadowValShdw1:
		attr.Value = "shdw1"
	case ST_PresetShadowValShdw2:
		attr.Value = "shdw2"
	case ST_PresetShadowValShdw3:
		attr.Value = "shdw3"
	case ST_PresetShadowValShdw4:
		attr.Value = "shdw4"
	case ST_PresetShadowValShdw5:
		attr.Value = "shdw5"
	case ST_PresetShadowValShdw6:
		attr.Value = "shdw6"
	case ST_PresetShadowValShdw7:
		attr.Value = "shdw7"
	case ST_PresetShadowValShdw8:
		attr.Value = "shdw8"
	case ST_PresetShadowValShdw9:
		attr.Value = "shdw9"
	case ST_PresetShadowValShdw10:
		attr.Value = "shdw10"
	case ST_PresetShadowValShdw11:
		attr.Value = "shdw11"
	case ST_PresetShadowValShdw12:
		attr.Value = "shdw12"
	case ST_PresetShadowValShdw13:
		attr.Value = "shdw13"
	case ST_PresetShadowValShdw14:
		attr.Value = "shdw14"
	case ST_PresetShadowValShdw15:
		attr.Value = "shdw15"
	case ST_PresetShadowValShdw16:
		attr.Value = "shdw16"
	case ST_PresetShadowValShdw17:
		attr.Value = "shdw17"
	case ST_PresetShadowValShdw18:
		attr.Value = "shdw18"
	case ST_PresetShadowValShdw19:
		attr.Value = "shdw19"
	case ST_PresetShadowValShdw20:
		attr.Value = "shdw20"
	}
	return attr, nil
}
func (e *ST_PresetShadowVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "shdw1":
		*e = 1
	case "shdw2":
		*e = 2
	case "shdw3":
		*e = 3
	case "shdw4":
		*e = 4
	case "shdw5":
		*e = 5
	case "shdw6":
		*e = 6
	case "shdw7":
		*e = 7
	case "shdw8":
		*e = 8
	case "shdw9":
		*e = 9
	case "shdw10":
		*e = 10
	case "shdw11":
		*e = 11
	case "shdw12":
		*e = 12
	case "shdw13":
		*e = 13
	case "shdw14":
		*e = 14
	case "shdw15":
		*e = 15
	case "shdw16":
		*e = 16
	case "shdw17":
		*e = 17
	case "shdw18":
		*e = 18
	case "shdw19":
		*e = 19
	case "shdw20":
		*e = 20
	}
	return nil
}
func (m ST_PresetShadowVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetShadowVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "shdw1":
			*m = 1
		case "shdw2":
			*m = 2
		case "shdw3":
			*m = 3
		case "shdw4":
			*m = 4
		case "shdw5":
			*m = 5
		case "shdw6":
			*m = 6
		case "shdw7":
			*m = 7
		case "shdw8":
			*m = 8
		case "shdw9":
			*m = 9
		case "shdw10":
			*m = 10
		case "shdw11":
			*m = 11
		case "shdw12":
			*m = 12
		case "shdw13":
			*m = 13
		case "shdw14":
			*m = 14
		case "shdw15":
			*m = 15
		case "shdw16":
			*m = 16
		case "shdw17":
			*m = 17
		case "shdw18":
			*m = 18
		case "shdw19":
			*m = 19
		case "shdw20":
			*m = 20
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
func (m ST_PresetShadowVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "shdw1"
	case 2:
		return "shdw2"
	case 3:
		return "shdw3"
	case 4:
		return "shdw4"
	case 5:
		return "shdw5"
	case 6:
		return "shdw6"
	case 7:
		return "shdw7"
	case 8:
		return "shdw8"
	case 9:
		return "shdw9"
	case 10:
		return "shdw10"
	case 11:
		return "shdw11"
	case 12:
		return "shdw12"
	case 13:
		return "shdw13"
	case 14:
		return "shdw14"
	case 15:
		return "shdw15"
	case 16:
		return "shdw16"
	case 17:
		return "shdw17"
	case 18:
		return "shdw18"
	case 19:
		return "shdw19"
	case 20:
		return "shdw20"
	}
	return ""
}
func (m ST_PresetShadowVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetShadowVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PathShadeType byte

const (
	ST_PathShadeTypeUnset  ST_PathShadeType = 0
	ST_PathShadeTypeShape  ST_PathShadeType = 1
	ST_PathShadeTypeCircle ST_PathShadeType = 2
	ST_PathShadeTypeRect   ST_PathShadeType = 3
)

func (e ST_PathShadeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PathShadeTypeUnset:
		attr.Value = ""
	case ST_PathShadeTypeShape:
		attr.Value = "shape"
	case ST_PathShadeTypeCircle:
		attr.Value = "circle"
	case ST_PathShadeTypeRect:
		attr.Value = "rect"
	}
	return attr, nil
}
func (e *ST_PathShadeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "shape":
		*e = 1
	case "circle":
		*e = 2
	case "rect":
		*e = 3
	}
	return nil
}
func (m ST_PathShadeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PathShadeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "shape":
			*m = 1
		case "circle":
			*m = 2
		case "rect":
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
func (m ST_PathShadeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "shape"
	case 2:
		return "circle"
	case 3:
		return "rect"
	}
	return ""
}
func (m ST_PathShadeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PathShadeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TileFlipMode byte

const (
	ST_TileFlipModeUnset ST_TileFlipMode = 0
	ST_TileFlipModeNone  ST_TileFlipMode = 1
	ST_TileFlipModeX     ST_TileFlipMode = 2
	ST_TileFlipModeY     ST_TileFlipMode = 3
	ST_TileFlipModeXy    ST_TileFlipMode = 4
)

func (e ST_TileFlipMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TileFlipModeUnset:
		attr.Value = ""
	case ST_TileFlipModeNone:
		attr.Value = "none"
	case ST_TileFlipModeX:
		attr.Value = "x"
	case ST_TileFlipModeY:
		attr.Value = "y"
	case ST_TileFlipModeXy:
		attr.Value = "xy"
	}
	return attr, nil
}
func (e *ST_TileFlipMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "x":
		*e = 2
	case "y":
		*e = 3
	case "xy":
		*e = 4
	}
	return nil
}
func (m ST_TileFlipMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TileFlipMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "x":
			*m = 2
		case "y":
			*m = 3
		case "xy":
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
func (m ST_TileFlipMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "x"
	case 3:
		return "y"
	case 4:
		return "xy"
	}
	return ""
}
func (m ST_TileFlipMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TileFlipMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BlipCompression byte

const (
	ST_BlipCompressionUnset   ST_BlipCompression = 0
	ST_BlipCompressionEmail   ST_BlipCompression = 1
	ST_BlipCompressionScreen  ST_BlipCompression = 2
	ST_BlipCompressionPrint   ST_BlipCompression = 3
	ST_BlipCompressionHqprint ST_BlipCompression = 4
	ST_BlipCompressionNone    ST_BlipCompression = 5
)

func (e ST_BlipCompression) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BlipCompressionUnset:
		attr.Value = ""
	case ST_BlipCompressionEmail:
		attr.Value = "email"
	case ST_BlipCompressionScreen:
		attr.Value = "screen"
	case ST_BlipCompressionPrint:
		attr.Value = "print"
	case ST_BlipCompressionHqprint:
		attr.Value = "hqprint"
	case ST_BlipCompressionNone:
		attr.Value = "none"
	}
	return attr, nil
}
func (e *ST_BlipCompression) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "email":
		*e = 1
	case "screen":
		*e = 2
	case "print":
		*e = 3
	case "hqprint":
		*e = 4
	case "none":
		*e = 5
	}
	return nil
}
func (m ST_BlipCompression) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BlipCompression) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "email":
			*m = 1
		case "screen":
			*m = 2
		case "print":
			*m = 3
		case "hqprint":
			*m = 4
		case "none":
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
func (m ST_BlipCompression) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "email"
	case 2:
		return "screen"
	case 3:
		return "print"
	case 4:
		return "hqprint"
	case 5:
		return "none"
	}
	return ""
}
func (m ST_BlipCompression) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BlipCompression) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetPatternVal byte

const (
	ST_PresetPatternValUnset      ST_PresetPatternVal = 0
	ST_PresetPatternValPct5       ST_PresetPatternVal = 1
	ST_PresetPatternValPct10      ST_PresetPatternVal = 2
	ST_PresetPatternValPct20      ST_PresetPatternVal = 3
	ST_PresetPatternValPct25      ST_PresetPatternVal = 4
	ST_PresetPatternValPct30      ST_PresetPatternVal = 5
	ST_PresetPatternValPct40      ST_PresetPatternVal = 6
	ST_PresetPatternValPct50      ST_PresetPatternVal = 7
	ST_PresetPatternValPct60      ST_PresetPatternVal = 8
	ST_PresetPatternValPct70      ST_PresetPatternVal = 9
	ST_PresetPatternValPct75      ST_PresetPatternVal = 10
	ST_PresetPatternValPct80      ST_PresetPatternVal = 11
	ST_PresetPatternValPct90      ST_PresetPatternVal = 12
	ST_PresetPatternValHorz       ST_PresetPatternVal = 13
	ST_PresetPatternValVert       ST_PresetPatternVal = 14
	ST_PresetPatternValLtHorz     ST_PresetPatternVal = 15
	ST_PresetPatternValLtVert     ST_PresetPatternVal = 16
	ST_PresetPatternValDkHorz     ST_PresetPatternVal = 17
	ST_PresetPatternValDkVert     ST_PresetPatternVal = 18
	ST_PresetPatternValNarHorz    ST_PresetPatternVal = 19
	ST_PresetPatternValNarVert    ST_PresetPatternVal = 20
	ST_PresetPatternValDashHorz   ST_PresetPatternVal = 21
	ST_PresetPatternValDashVert   ST_PresetPatternVal = 22
	ST_PresetPatternValCross      ST_PresetPatternVal = 23
	ST_PresetPatternValDnDiag     ST_PresetPatternVal = 24
	ST_PresetPatternValUpDiag     ST_PresetPatternVal = 25
	ST_PresetPatternValLtDnDiag   ST_PresetPatternVal = 26
	ST_PresetPatternValLtUpDiag   ST_PresetPatternVal = 27
	ST_PresetPatternValDkDnDiag   ST_PresetPatternVal = 28
	ST_PresetPatternValDkUpDiag   ST_PresetPatternVal = 29
	ST_PresetPatternValWdDnDiag   ST_PresetPatternVal = 30
	ST_PresetPatternValWdUpDiag   ST_PresetPatternVal = 31
	ST_PresetPatternValDashDnDiag ST_PresetPatternVal = 32
	ST_PresetPatternValDashUpDiag ST_PresetPatternVal = 33
	ST_PresetPatternValDiagCross  ST_PresetPatternVal = 34
	ST_PresetPatternValSmCheck    ST_PresetPatternVal = 35
	ST_PresetPatternValLgCheck    ST_PresetPatternVal = 36
	ST_PresetPatternValSmGrid     ST_PresetPatternVal = 37
	ST_PresetPatternValLgGrid     ST_PresetPatternVal = 38
	ST_PresetPatternValDotGrid    ST_PresetPatternVal = 39
	ST_PresetPatternValSmConfetti ST_PresetPatternVal = 40
	ST_PresetPatternValLgConfetti ST_PresetPatternVal = 41
	ST_PresetPatternValHorzBrick  ST_PresetPatternVal = 42
	ST_PresetPatternValDiagBrick  ST_PresetPatternVal = 43
	ST_PresetPatternValSolidDmnd  ST_PresetPatternVal = 44
	ST_PresetPatternValOpenDmnd   ST_PresetPatternVal = 45
	ST_PresetPatternValDotDmnd    ST_PresetPatternVal = 46
	ST_PresetPatternValPlaid      ST_PresetPatternVal = 47
	ST_PresetPatternValSphere     ST_PresetPatternVal = 48
	ST_PresetPatternValWeave      ST_PresetPatternVal = 49
	ST_PresetPatternValDivot      ST_PresetPatternVal = 50
	ST_PresetPatternValShingle    ST_PresetPatternVal = 51
	ST_PresetPatternValWave       ST_PresetPatternVal = 52
	ST_PresetPatternValTrellis    ST_PresetPatternVal = 53
	ST_PresetPatternValZigZag     ST_PresetPatternVal = 54
)

func (e ST_PresetPatternVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetPatternValUnset:
		attr.Value = ""
	case ST_PresetPatternValPct5:
		attr.Value = "pct5"
	case ST_PresetPatternValPct10:
		attr.Value = "pct10"
	case ST_PresetPatternValPct20:
		attr.Value = "pct20"
	case ST_PresetPatternValPct25:
		attr.Value = "pct25"
	case ST_PresetPatternValPct30:
		attr.Value = "pct30"
	case ST_PresetPatternValPct40:
		attr.Value = "pct40"
	case ST_PresetPatternValPct50:
		attr.Value = "pct50"
	case ST_PresetPatternValPct60:
		attr.Value = "pct60"
	case ST_PresetPatternValPct70:
		attr.Value = "pct70"
	case ST_PresetPatternValPct75:
		attr.Value = "pct75"
	case ST_PresetPatternValPct80:
		attr.Value = "pct80"
	case ST_PresetPatternValPct90:
		attr.Value = "pct90"
	case ST_PresetPatternValHorz:
		attr.Value = "horz"
	case ST_PresetPatternValVert:
		attr.Value = "vert"
	case ST_PresetPatternValLtHorz:
		attr.Value = "ltHorz"
	case ST_PresetPatternValLtVert:
		attr.Value = "ltVert"
	case ST_PresetPatternValDkHorz:
		attr.Value = "dkHorz"
	case ST_PresetPatternValDkVert:
		attr.Value = "dkVert"
	case ST_PresetPatternValNarHorz:
		attr.Value = "narHorz"
	case ST_PresetPatternValNarVert:
		attr.Value = "narVert"
	case ST_PresetPatternValDashHorz:
		attr.Value = "dashHorz"
	case ST_PresetPatternValDashVert:
		attr.Value = "dashVert"
	case ST_PresetPatternValCross:
		attr.Value = "cross"
	case ST_PresetPatternValDnDiag:
		attr.Value = "dnDiag"
	case ST_PresetPatternValUpDiag:
		attr.Value = "upDiag"
	case ST_PresetPatternValLtDnDiag:
		attr.Value = "ltDnDiag"
	case ST_PresetPatternValLtUpDiag:
		attr.Value = "ltUpDiag"
	case ST_PresetPatternValDkDnDiag:
		attr.Value = "dkDnDiag"
	case ST_PresetPatternValDkUpDiag:
		attr.Value = "dkUpDiag"
	case ST_PresetPatternValWdDnDiag:
		attr.Value = "wdDnDiag"
	case ST_PresetPatternValWdUpDiag:
		attr.Value = "wdUpDiag"
	case ST_PresetPatternValDashDnDiag:
		attr.Value = "dashDnDiag"
	case ST_PresetPatternValDashUpDiag:
		attr.Value = "dashUpDiag"
	case ST_PresetPatternValDiagCross:
		attr.Value = "diagCross"
	case ST_PresetPatternValSmCheck:
		attr.Value = "smCheck"
	case ST_PresetPatternValLgCheck:
		attr.Value = "lgCheck"
	case ST_PresetPatternValSmGrid:
		attr.Value = "smGrid"
	case ST_PresetPatternValLgGrid:
		attr.Value = "lgGrid"
	case ST_PresetPatternValDotGrid:
		attr.Value = "dotGrid"
	case ST_PresetPatternValSmConfetti:
		attr.Value = "smConfetti"
	case ST_PresetPatternValLgConfetti:
		attr.Value = "lgConfetti"
	case ST_PresetPatternValHorzBrick:
		attr.Value = "horzBrick"
	case ST_PresetPatternValDiagBrick:
		attr.Value = "diagBrick"
	case ST_PresetPatternValSolidDmnd:
		attr.Value = "solidDmnd"
	case ST_PresetPatternValOpenDmnd:
		attr.Value = "openDmnd"
	case ST_PresetPatternValDotDmnd:
		attr.Value = "dotDmnd"
	case ST_PresetPatternValPlaid:
		attr.Value = "plaid"
	case ST_PresetPatternValSphere:
		attr.Value = "sphere"
	case ST_PresetPatternValWeave:
		attr.Value = "weave"
	case ST_PresetPatternValDivot:
		attr.Value = "divot"
	case ST_PresetPatternValShingle:
		attr.Value = "shingle"
	case ST_PresetPatternValWave:
		attr.Value = "wave"
	case ST_PresetPatternValTrellis:
		attr.Value = "trellis"
	case ST_PresetPatternValZigZag:
		attr.Value = "zigZag"
	}
	return attr, nil
}
func (e *ST_PresetPatternVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "pct5":
		*e = 1
	case "pct10":
		*e = 2
	case "pct20":
		*e = 3
	case "pct25":
		*e = 4
	case "pct30":
		*e = 5
	case "pct40":
		*e = 6
	case "pct50":
		*e = 7
	case "pct60":
		*e = 8
	case "pct70":
		*e = 9
	case "pct75":
		*e = 10
	case "pct80":
		*e = 11
	case "pct90":
		*e = 12
	case "horz":
		*e = 13
	case "vert":
		*e = 14
	case "ltHorz":
		*e = 15
	case "ltVert":
		*e = 16
	case "dkHorz":
		*e = 17
	case "dkVert":
		*e = 18
	case "narHorz":
		*e = 19
	case "narVert":
		*e = 20
	case "dashHorz":
		*e = 21
	case "dashVert":
		*e = 22
	case "cross":
		*e = 23
	case "dnDiag":
		*e = 24
	case "upDiag":
		*e = 25
	case "ltDnDiag":
		*e = 26
	case "ltUpDiag":
		*e = 27
	case "dkDnDiag":
		*e = 28
	case "dkUpDiag":
		*e = 29
	case "wdDnDiag":
		*e = 30
	case "wdUpDiag":
		*e = 31
	case "dashDnDiag":
		*e = 32
	case "dashUpDiag":
		*e = 33
	case "diagCross":
		*e = 34
	case "smCheck":
		*e = 35
	case "lgCheck":
		*e = 36
	case "smGrid":
		*e = 37
	case "lgGrid":
		*e = 38
	case "dotGrid":
		*e = 39
	case "smConfetti":
		*e = 40
	case "lgConfetti":
		*e = 41
	case "horzBrick":
		*e = 42
	case "diagBrick":
		*e = 43
	case "solidDmnd":
		*e = 44
	case "openDmnd":
		*e = 45
	case "dotDmnd":
		*e = 46
	case "plaid":
		*e = 47
	case "sphere":
		*e = 48
	case "weave":
		*e = 49
	case "divot":
		*e = 50
	case "shingle":
		*e = 51
	case "wave":
		*e = 52
	case "trellis":
		*e = 53
	case "zigZag":
		*e = 54
	}
	return nil
}
func (m ST_PresetPatternVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetPatternVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "pct5":
			*m = 1
		case "pct10":
			*m = 2
		case "pct20":
			*m = 3
		case "pct25":
			*m = 4
		case "pct30":
			*m = 5
		case "pct40":
			*m = 6
		case "pct50":
			*m = 7
		case "pct60":
			*m = 8
		case "pct70":
			*m = 9
		case "pct75":
			*m = 10
		case "pct80":
			*m = 11
		case "pct90":
			*m = 12
		case "horz":
			*m = 13
		case "vert":
			*m = 14
		case "ltHorz":
			*m = 15
		case "ltVert":
			*m = 16
		case "dkHorz":
			*m = 17
		case "dkVert":
			*m = 18
		case "narHorz":
			*m = 19
		case "narVert":
			*m = 20
		case "dashHorz":
			*m = 21
		case "dashVert":
			*m = 22
		case "cross":
			*m = 23
		case "dnDiag":
			*m = 24
		case "upDiag":
			*m = 25
		case "ltDnDiag":
			*m = 26
		case "ltUpDiag":
			*m = 27
		case "dkDnDiag":
			*m = 28
		case "dkUpDiag":
			*m = 29
		case "wdDnDiag":
			*m = 30
		case "wdUpDiag":
			*m = 31
		case "dashDnDiag":
			*m = 32
		case "dashUpDiag":
			*m = 33
		case "diagCross":
			*m = 34
		case "smCheck":
			*m = 35
		case "lgCheck":
			*m = 36
		case "smGrid":
			*m = 37
		case "lgGrid":
			*m = 38
		case "dotGrid":
			*m = 39
		case "smConfetti":
			*m = 40
		case "lgConfetti":
			*m = 41
		case "horzBrick":
			*m = 42
		case "diagBrick":
			*m = 43
		case "solidDmnd":
			*m = 44
		case "openDmnd":
			*m = 45
		case "dotDmnd":
			*m = 46
		case "plaid":
			*m = 47
		case "sphere":
			*m = 48
		case "weave":
			*m = 49
		case "divot":
			*m = 50
		case "shingle":
			*m = 51
		case "wave":
			*m = 52
		case "trellis":
			*m = 53
		case "zigZag":
			*m = 54
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
func (m ST_PresetPatternVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "pct5"
	case 2:
		return "pct10"
	case 3:
		return "pct20"
	case 4:
		return "pct25"
	case 5:
		return "pct30"
	case 6:
		return "pct40"
	case 7:
		return "pct50"
	case 8:
		return "pct60"
	case 9:
		return "pct70"
	case 10:
		return "pct75"
	case 11:
		return "pct80"
	case 12:
		return "pct90"
	case 13:
		return "horz"
	case 14:
		return "vert"
	case 15:
		return "ltHorz"
	case 16:
		return "ltVert"
	case 17:
		return "dkHorz"
	case 18:
		return "dkVert"
	case 19:
		return "narHorz"
	case 20:
		return "narVert"
	case 21:
		return "dashHorz"
	case 22:
		return "dashVert"
	case 23:
		return "cross"
	case 24:
		return "dnDiag"
	case 25:
		return "upDiag"
	case 26:
		return "ltDnDiag"
	case 27:
		return "ltUpDiag"
	case 28:
		return "dkDnDiag"
	case 29:
		return "dkUpDiag"
	case 30:
		return "wdDnDiag"
	case 31:
		return "wdUpDiag"
	case 32:
		return "dashDnDiag"
	case 33:
		return "dashUpDiag"
	case 34:
		return "diagCross"
	case 35:
		return "smCheck"
	case 36:
		return "lgCheck"
	case 37:
		return "smGrid"
	case 38:
		return "lgGrid"
	case 39:
		return "dotGrid"
	case 40:
		return "smConfetti"
	case 41:
		return "lgConfetti"
	case 42:
		return "horzBrick"
	case 43:
		return "diagBrick"
	case 44:
		return "solidDmnd"
	case 45:
		return "openDmnd"
	case 46:
		return "dotDmnd"
	case 47:
		return "plaid"
	case 48:
		return "sphere"
	case 49:
		return "weave"
	case 50:
		return "divot"
	case 51:
		return "shingle"
	case 52:
		return "wave"
	case 53:
		return "trellis"
	case 54:
		return "zigZag"
	}
	return ""
}
func (m ST_PresetPatternVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetPatternVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BlendMode byte

const (
	ST_BlendModeUnset   ST_BlendMode = 0
	ST_BlendModeOver    ST_BlendMode = 1
	ST_BlendModeMult    ST_BlendMode = 2
	ST_BlendModeScreen  ST_BlendMode = 3
	ST_BlendModeDarken  ST_BlendMode = 4
	ST_BlendModeLighten ST_BlendMode = 5
)

func (e ST_BlendMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BlendModeUnset:
		attr.Value = ""
	case ST_BlendModeOver:
		attr.Value = "over"
	case ST_BlendModeMult:
		attr.Value = "mult"
	case ST_BlendModeScreen:
		attr.Value = "screen"
	case ST_BlendModeDarken:
		attr.Value = "darken"
	case ST_BlendModeLighten:
		attr.Value = "lighten"
	}
	return attr, nil
}
func (e *ST_BlendMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "over":
		*e = 1
	case "mult":
		*e = 2
	case "screen":
		*e = 3
	case "darken":
		*e = 4
	case "lighten":
		*e = 5
	}
	return nil
}
func (m ST_BlendMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_BlendMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "over":
			*m = 1
		case "mult":
			*m = 2
		case "screen":
			*m = 3
		case "darken":
			*m = 4
		case "lighten":
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
func (m ST_BlendMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "over"
	case 2:
		return "mult"
	case 3:
		return "screen"
	case 4:
		return "darken"
	case 5:
		return "lighten"
	}
	return ""
}
func (m ST_BlendMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_BlendMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_EffectContainerType byte

const (
	ST_EffectContainerTypeUnset ST_EffectContainerType = 0
	ST_EffectContainerTypeSib   ST_EffectContainerType = 1
	ST_EffectContainerTypeTree  ST_EffectContainerType = 2
)

func (e ST_EffectContainerType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EffectContainerTypeUnset:
		attr.Value = ""
	case ST_EffectContainerTypeSib:
		attr.Value = "sib"
	case ST_EffectContainerTypeTree:
		attr.Value = "tree"
	}
	return attr, nil
}
func (e *ST_EffectContainerType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sib":
		*e = 1
	case "tree":
		*e = 2
	}
	return nil
}
func (m ST_EffectContainerType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_EffectContainerType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sib":
			*m = 1
		case "tree":
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
func (m ST_EffectContainerType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sib"
	case 2:
		return "tree"
	}
	return ""
}
func (m ST_EffectContainerType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_EffectContainerType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ShapeType byte

const (
	ST_ShapeTypeUnset                      ST_ShapeType = 0
	ST_ShapeTypeLine                       ST_ShapeType = 1
	ST_ShapeTypeLineInv                    ST_ShapeType = 2
	ST_ShapeTypeTriangle                   ST_ShapeType = 3
	ST_ShapeTypeRtTriangle                 ST_ShapeType = 4
	ST_ShapeTypeRect                       ST_ShapeType = 5
	ST_ShapeTypeDiamond                    ST_ShapeType = 6
	ST_ShapeTypeParallelogram              ST_ShapeType = 7
	ST_ShapeTypeTrapezoid                  ST_ShapeType = 8
	ST_ShapeTypeNonIsoscelesTrapezoid      ST_ShapeType = 9
	ST_ShapeTypePentagon                   ST_ShapeType = 10
	ST_ShapeTypeHexagon                    ST_ShapeType = 11
	ST_ShapeTypeHeptagon                   ST_ShapeType = 12
	ST_ShapeTypeOctagon                    ST_ShapeType = 13
	ST_ShapeTypeDecagon                    ST_ShapeType = 14
	ST_ShapeTypeDodecagon                  ST_ShapeType = 15
	ST_ShapeTypeStar4                      ST_ShapeType = 16
	ST_ShapeTypeStar5                      ST_ShapeType = 17
	ST_ShapeTypeStar6                      ST_ShapeType = 18
	ST_ShapeTypeStar7                      ST_ShapeType = 19
	ST_ShapeTypeStar8                      ST_ShapeType = 20
	ST_ShapeTypeStar10                     ST_ShapeType = 21
	ST_ShapeTypeStar12                     ST_ShapeType = 22
	ST_ShapeTypeStar16                     ST_ShapeType = 23
	ST_ShapeTypeStar24                     ST_ShapeType = 24
	ST_ShapeTypeStar32                     ST_ShapeType = 25
	ST_ShapeTypeRoundRect                  ST_ShapeType = 26
	ST_ShapeTypeRound1Rect                 ST_ShapeType = 27
	ST_ShapeTypeRound2SameRect             ST_ShapeType = 28
	ST_ShapeTypeRound2DiagRect             ST_ShapeType = 29
	ST_ShapeTypeSnipRoundRect              ST_ShapeType = 30
	ST_ShapeTypeSnip1Rect                  ST_ShapeType = 31
	ST_ShapeTypeSnip2SameRect              ST_ShapeType = 32
	ST_ShapeTypeSnip2DiagRect              ST_ShapeType = 33
	ST_ShapeTypePlaque                     ST_ShapeType = 34
	ST_ShapeTypeEllipse                    ST_ShapeType = 35
	ST_ShapeTypeTeardrop                   ST_ShapeType = 36
	ST_ShapeTypeHomePlate                  ST_ShapeType = 37
	ST_ShapeTypeChevron                    ST_ShapeType = 38
	ST_ShapeTypePieWedge                   ST_ShapeType = 39
	ST_ShapeTypePie                        ST_ShapeType = 40
	ST_ShapeTypeBlockArc                   ST_ShapeType = 41
	ST_ShapeTypeDonut                      ST_ShapeType = 42
	ST_ShapeTypeNoSmoking                  ST_ShapeType = 43
	ST_ShapeTypeRightArrow                 ST_ShapeType = 44
	ST_ShapeTypeLeftArrow                  ST_ShapeType = 45
	ST_ShapeTypeUpArrow                    ST_ShapeType = 46
	ST_ShapeTypeDownArrow                  ST_ShapeType = 47
	ST_ShapeTypeStripedRightArrow          ST_ShapeType = 48
	ST_ShapeTypeNotchedRightArrow          ST_ShapeType = 49
	ST_ShapeTypeBentUpArrow                ST_ShapeType = 50
	ST_ShapeTypeLeftRightArrow             ST_ShapeType = 51
	ST_ShapeTypeUpDownArrow                ST_ShapeType = 52
	ST_ShapeTypeLeftUpArrow                ST_ShapeType = 53
	ST_ShapeTypeLeftRightUpArrow           ST_ShapeType = 54
	ST_ShapeTypeQuadArrow                  ST_ShapeType = 55
	ST_ShapeTypeLeftArrowCallout           ST_ShapeType = 56
	ST_ShapeTypeRightArrowCallout          ST_ShapeType = 57
	ST_ShapeTypeUpArrowCallout             ST_ShapeType = 58
	ST_ShapeTypeDownArrowCallout           ST_ShapeType = 59
	ST_ShapeTypeLeftRightArrowCallout      ST_ShapeType = 60
	ST_ShapeTypeUpDownArrowCallout         ST_ShapeType = 61
	ST_ShapeTypeQuadArrowCallout           ST_ShapeType = 62
	ST_ShapeTypeBentArrow                  ST_ShapeType = 63
	ST_ShapeTypeUturnArrow                 ST_ShapeType = 64
	ST_ShapeTypeCircularArrow              ST_ShapeType = 65
	ST_ShapeTypeLeftCircularArrow          ST_ShapeType = 66
	ST_ShapeTypeLeftRightCircularArrow     ST_ShapeType = 67
	ST_ShapeTypeCurvedRightArrow           ST_ShapeType = 68
	ST_ShapeTypeCurvedLeftArrow            ST_ShapeType = 69
	ST_ShapeTypeCurvedUpArrow              ST_ShapeType = 70
	ST_ShapeTypeCurvedDownArrow            ST_ShapeType = 71
	ST_ShapeTypeSwooshArrow                ST_ShapeType = 72
	ST_ShapeTypeCube                       ST_ShapeType = 73
	ST_ShapeTypeCan                        ST_ShapeType = 74
	ST_ShapeTypeLightningBolt              ST_ShapeType = 75
	ST_ShapeTypeHeart                      ST_ShapeType = 76
	ST_ShapeTypeSun                        ST_ShapeType = 77
	ST_ShapeTypeMoon                       ST_ShapeType = 78
	ST_ShapeTypeSmileyFace                 ST_ShapeType = 79
	ST_ShapeTypeIrregularSeal1             ST_ShapeType = 80
	ST_ShapeTypeIrregularSeal2             ST_ShapeType = 81
	ST_ShapeTypeFoldedCorner               ST_ShapeType = 82
	ST_ShapeTypeBevel                      ST_ShapeType = 83
	ST_ShapeTypeFrame                      ST_ShapeType = 84
	ST_ShapeTypeHalfFrame                  ST_ShapeType = 85
	ST_ShapeTypeCorner                     ST_ShapeType = 86
	ST_ShapeTypeDiagStripe                 ST_ShapeType = 87
	ST_ShapeTypeChord                      ST_ShapeType = 88
	ST_ShapeTypeArc                        ST_ShapeType = 89
	ST_ShapeTypeLeftBracket                ST_ShapeType = 90
	ST_ShapeTypeRightBracket               ST_ShapeType = 91
	ST_ShapeTypeLeftBrace                  ST_ShapeType = 92
	ST_ShapeTypeRightBrace                 ST_ShapeType = 93
	ST_ShapeTypeBracketPair                ST_ShapeType = 94
	ST_ShapeTypeBracePair                  ST_ShapeType = 95
	ST_ShapeTypeStraightConnector1         ST_ShapeType = 96
	ST_ShapeTypeBentConnector2             ST_ShapeType = 97
	ST_ShapeTypeBentConnector3             ST_ShapeType = 98
	ST_ShapeTypeBentConnector4             ST_ShapeType = 99
	ST_ShapeTypeBentConnector5             ST_ShapeType = 100
	ST_ShapeTypeCurvedConnector2           ST_ShapeType = 101
	ST_ShapeTypeCurvedConnector3           ST_ShapeType = 102
	ST_ShapeTypeCurvedConnector4           ST_ShapeType = 103
	ST_ShapeTypeCurvedConnector5           ST_ShapeType = 104
	ST_ShapeTypeCallout1                   ST_ShapeType = 105
	ST_ShapeTypeCallout2                   ST_ShapeType = 106
	ST_ShapeTypeCallout3                   ST_ShapeType = 107
	ST_ShapeTypeAccentCallout1             ST_ShapeType = 108
	ST_ShapeTypeAccentCallout2             ST_ShapeType = 109
	ST_ShapeTypeAccentCallout3             ST_ShapeType = 110
	ST_ShapeTypeBorderCallout1             ST_ShapeType = 111
	ST_ShapeTypeBorderCallout2             ST_ShapeType = 112
	ST_ShapeTypeBorderCallout3             ST_ShapeType = 113
	ST_ShapeTypeAccentBorderCallout1       ST_ShapeType = 114
	ST_ShapeTypeAccentBorderCallout2       ST_ShapeType = 115
	ST_ShapeTypeAccentBorderCallout3       ST_ShapeType = 116
	ST_ShapeTypeWedgeRectCallout           ST_ShapeType = 117
	ST_ShapeTypeWedgeRoundRectCallout      ST_ShapeType = 118
	ST_ShapeTypeWedgeEllipseCallout        ST_ShapeType = 119
	ST_ShapeTypeCloudCallout               ST_ShapeType = 120
	ST_ShapeTypeCloud                      ST_ShapeType = 121
	ST_ShapeTypeRibbon                     ST_ShapeType = 122
	ST_ShapeTypeRibbon2                    ST_ShapeType = 123
	ST_ShapeTypeEllipseRibbon              ST_ShapeType = 124
	ST_ShapeTypeEllipseRibbon2             ST_ShapeType = 125
	ST_ShapeTypeLeftRightRibbon            ST_ShapeType = 126
	ST_ShapeTypeVerticalScroll             ST_ShapeType = 127
	ST_ShapeTypeHorizontalScroll           ST_ShapeType = 128
	ST_ShapeTypeWave                       ST_ShapeType = 129
	ST_ShapeTypeDoubleWave                 ST_ShapeType = 130
	ST_ShapeTypePlus                       ST_ShapeType = 131
	ST_ShapeTypeFlowChartProcess           ST_ShapeType = 132
	ST_ShapeTypeFlowChartDecision          ST_ShapeType = 133
	ST_ShapeTypeFlowChartInputOutput       ST_ShapeType = 134
	ST_ShapeTypeFlowChartPredefinedProcess ST_ShapeType = 135
	ST_ShapeTypeFlowChartInternalStorage   ST_ShapeType = 136
	ST_ShapeTypeFlowChartDocument          ST_ShapeType = 137
	ST_ShapeTypeFlowChartMultidocument     ST_ShapeType = 138
	ST_ShapeTypeFlowChartTerminator        ST_ShapeType = 139
	ST_ShapeTypeFlowChartPreparation       ST_ShapeType = 140
	ST_ShapeTypeFlowChartManualInput       ST_ShapeType = 141
	ST_ShapeTypeFlowChartManualOperation   ST_ShapeType = 142
	ST_ShapeTypeFlowChartConnector         ST_ShapeType = 143
	ST_ShapeTypeFlowChartPunchedCard       ST_ShapeType = 144
	ST_ShapeTypeFlowChartPunchedTape       ST_ShapeType = 145
	ST_ShapeTypeFlowChartSummingJunction   ST_ShapeType = 146
	ST_ShapeTypeFlowChartOr                ST_ShapeType = 147
	ST_ShapeTypeFlowChartCollate           ST_ShapeType = 148
	ST_ShapeTypeFlowChartSort              ST_ShapeType = 149
	ST_ShapeTypeFlowChartExtract           ST_ShapeType = 150
	ST_ShapeTypeFlowChartMerge             ST_ShapeType = 151
	ST_ShapeTypeFlowChartOfflineStorage    ST_ShapeType = 152
	ST_ShapeTypeFlowChartOnlineStorage     ST_ShapeType = 153
	ST_ShapeTypeFlowChartMagneticTape      ST_ShapeType = 154
	ST_ShapeTypeFlowChartMagneticDisk      ST_ShapeType = 155
	ST_ShapeTypeFlowChartMagneticDrum      ST_ShapeType = 156
	ST_ShapeTypeFlowChartDisplay           ST_ShapeType = 157
	ST_ShapeTypeFlowChartDelay             ST_ShapeType = 158
	ST_ShapeTypeFlowChartAlternateProcess  ST_ShapeType = 159
	ST_ShapeTypeFlowChartOffpageConnector  ST_ShapeType = 160
	ST_ShapeTypeActionButtonBlank          ST_ShapeType = 161
	ST_ShapeTypeActionButtonHome           ST_ShapeType = 162
	ST_ShapeTypeActionButtonHelp           ST_ShapeType = 163
	ST_ShapeTypeActionButtonInformation    ST_ShapeType = 164
	ST_ShapeTypeActionButtonForwardNext    ST_ShapeType = 165
	ST_ShapeTypeActionButtonBackPrevious   ST_ShapeType = 166
	ST_ShapeTypeActionButtonEnd            ST_ShapeType = 167
	ST_ShapeTypeActionButtonBeginning      ST_ShapeType = 168
	ST_ShapeTypeActionButtonReturn         ST_ShapeType = 169
	ST_ShapeTypeActionButtonDocument       ST_ShapeType = 170
	ST_ShapeTypeActionButtonSound          ST_ShapeType = 171
	ST_ShapeTypeActionButtonMovie          ST_ShapeType = 172
	ST_ShapeTypeGear6                      ST_ShapeType = 173
	ST_ShapeTypeGear9                      ST_ShapeType = 174
	ST_ShapeTypeFunnel                     ST_ShapeType = 175
	ST_ShapeTypeMathPlus                   ST_ShapeType = 176
	ST_ShapeTypeMathMinus                  ST_ShapeType = 177
	ST_ShapeTypeMathMultiply               ST_ShapeType = 178
	ST_ShapeTypeMathDivide                 ST_ShapeType = 179
	ST_ShapeTypeMathEqual                  ST_ShapeType = 180
	ST_ShapeTypeMathNotEqual               ST_ShapeType = 181
	ST_ShapeTypeCornerTabs                 ST_ShapeType = 182
	ST_ShapeTypeSquareTabs                 ST_ShapeType = 183
	ST_ShapeTypePlaqueTabs                 ST_ShapeType = 184
	ST_ShapeTypeChartX                     ST_ShapeType = 185
	ST_ShapeTypeChartStar                  ST_ShapeType = 186
	ST_ShapeTypeChartPlus                  ST_ShapeType = 187
)

func (e ST_ShapeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShapeTypeUnset:
		attr.Value = ""
	case ST_ShapeTypeLine:
		attr.Value = "line"
	case ST_ShapeTypeLineInv:
		attr.Value = "lineInv"
	case ST_ShapeTypeTriangle:
		attr.Value = "triangle"
	case ST_ShapeTypeRtTriangle:
		attr.Value = "rtTriangle"
	case ST_ShapeTypeRect:
		attr.Value = "rect"
	case ST_ShapeTypeDiamond:
		attr.Value = "diamond"
	case ST_ShapeTypeParallelogram:
		attr.Value = "parallelogram"
	case ST_ShapeTypeTrapezoid:
		attr.Value = "trapezoid"
	case ST_ShapeTypeNonIsoscelesTrapezoid:
		attr.Value = "nonIsoscelesTrapezoid"
	case ST_ShapeTypePentagon:
		attr.Value = "pentagon"
	case ST_ShapeTypeHexagon:
		attr.Value = "hexagon"
	case ST_ShapeTypeHeptagon:
		attr.Value = "heptagon"
	case ST_ShapeTypeOctagon:
		attr.Value = "octagon"
	case ST_ShapeTypeDecagon:
		attr.Value = "decagon"
	case ST_ShapeTypeDodecagon:
		attr.Value = "dodecagon"
	case ST_ShapeTypeStar4:
		attr.Value = "star4"
	case ST_ShapeTypeStar5:
		attr.Value = "star5"
	case ST_ShapeTypeStar6:
		attr.Value = "star6"
	case ST_ShapeTypeStar7:
		attr.Value = "star7"
	case ST_ShapeTypeStar8:
		attr.Value = "star8"
	case ST_ShapeTypeStar10:
		attr.Value = "star10"
	case ST_ShapeTypeStar12:
		attr.Value = "star12"
	case ST_ShapeTypeStar16:
		attr.Value = "star16"
	case ST_ShapeTypeStar24:
		attr.Value = "star24"
	case ST_ShapeTypeStar32:
		attr.Value = "star32"
	case ST_ShapeTypeRoundRect:
		attr.Value = "roundRect"
	case ST_ShapeTypeRound1Rect:
		attr.Value = "round1Rect"
	case ST_ShapeTypeRound2SameRect:
		attr.Value = "round2SameRect"
	case ST_ShapeTypeRound2DiagRect:
		attr.Value = "round2DiagRect"
	case ST_ShapeTypeSnipRoundRect:
		attr.Value = "snipRoundRect"
	case ST_ShapeTypeSnip1Rect:
		attr.Value = "snip1Rect"
	case ST_ShapeTypeSnip2SameRect:
		attr.Value = "snip2SameRect"
	case ST_ShapeTypeSnip2DiagRect:
		attr.Value = "snip2DiagRect"
	case ST_ShapeTypePlaque:
		attr.Value = "plaque"
	case ST_ShapeTypeEllipse:
		attr.Value = "ellipse"
	case ST_ShapeTypeTeardrop:
		attr.Value = "teardrop"
	case ST_ShapeTypeHomePlate:
		attr.Value = "homePlate"
	case ST_ShapeTypeChevron:
		attr.Value = "chevron"
	case ST_ShapeTypePieWedge:
		attr.Value = "pieWedge"
	case ST_ShapeTypePie:
		attr.Value = "pie"
	case ST_ShapeTypeBlockArc:
		attr.Value = "blockArc"
	case ST_ShapeTypeDonut:
		attr.Value = "donut"
	case ST_ShapeTypeNoSmoking:
		attr.Value = "noSmoking"
	case ST_ShapeTypeRightArrow:
		attr.Value = "rightArrow"
	case ST_ShapeTypeLeftArrow:
		attr.Value = "leftArrow"
	case ST_ShapeTypeUpArrow:
		attr.Value = "upArrow"
	case ST_ShapeTypeDownArrow:
		attr.Value = "downArrow"
	case ST_ShapeTypeStripedRightArrow:
		attr.Value = "stripedRightArrow"
	case ST_ShapeTypeNotchedRightArrow:
		attr.Value = "notchedRightArrow"
	case ST_ShapeTypeBentUpArrow:
		attr.Value = "bentUpArrow"
	case ST_ShapeTypeLeftRightArrow:
		attr.Value = "leftRightArrow"
	case ST_ShapeTypeUpDownArrow:
		attr.Value = "upDownArrow"
	case ST_ShapeTypeLeftUpArrow:
		attr.Value = "leftUpArrow"
	case ST_ShapeTypeLeftRightUpArrow:
		attr.Value = "leftRightUpArrow"
	case ST_ShapeTypeQuadArrow:
		attr.Value = "quadArrow"
	case ST_ShapeTypeLeftArrowCallout:
		attr.Value = "leftArrowCallout"
	case ST_ShapeTypeRightArrowCallout:
		attr.Value = "rightArrowCallout"
	case ST_ShapeTypeUpArrowCallout:
		attr.Value = "upArrowCallout"
	case ST_ShapeTypeDownArrowCallout:
		attr.Value = "downArrowCallout"
	case ST_ShapeTypeLeftRightArrowCallout:
		attr.Value = "leftRightArrowCallout"
	case ST_ShapeTypeUpDownArrowCallout:
		attr.Value = "upDownArrowCallout"
	case ST_ShapeTypeQuadArrowCallout:
		attr.Value = "quadArrowCallout"
	case ST_ShapeTypeBentArrow:
		attr.Value = "bentArrow"
	case ST_ShapeTypeUturnArrow:
		attr.Value = "uturnArrow"
	case ST_ShapeTypeCircularArrow:
		attr.Value = "circularArrow"
	case ST_ShapeTypeLeftCircularArrow:
		attr.Value = "leftCircularArrow"
	case ST_ShapeTypeLeftRightCircularArrow:
		attr.Value = "leftRightCircularArrow"
	case ST_ShapeTypeCurvedRightArrow:
		attr.Value = "curvedRightArrow"
	case ST_ShapeTypeCurvedLeftArrow:
		attr.Value = "curvedLeftArrow"
	case ST_ShapeTypeCurvedUpArrow:
		attr.Value = "curvedUpArrow"
	case ST_ShapeTypeCurvedDownArrow:
		attr.Value = "curvedDownArrow"
	case ST_ShapeTypeSwooshArrow:
		attr.Value = "swooshArrow"
	case ST_ShapeTypeCube:
		attr.Value = "cube"
	case ST_ShapeTypeCan:
		attr.Value = "can"
	case ST_ShapeTypeLightningBolt:
		attr.Value = "lightningBolt"
	case ST_ShapeTypeHeart:
		attr.Value = "heart"
	case ST_ShapeTypeSun:
		attr.Value = "sun"
	case ST_ShapeTypeMoon:
		attr.Value = "moon"
	case ST_ShapeTypeSmileyFace:
		attr.Value = "smileyFace"
	case ST_ShapeTypeIrregularSeal1:
		attr.Value = "irregularSeal1"
	case ST_ShapeTypeIrregularSeal2:
		attr.Value = "irregularSeal2"
	case ST_ShapeTypeFoldedCorner:
		attr.Value = "foldedCorner"
	case ST_ShapeTypeBevel:
		attr.Value = "bevel"
	case ST_ShapeTypeFrame:
		attr.Value = "frame"
	case ST_ShapeTypeHalfFrame:
		attr.Value = "halfFrame"
	case ST_ShapeTypeCorner:
		attr.Value = "corner"
	case ST_ShapeTypeDiagStripe:
		attr.Value = "diagStripe"
	case ST_ShapeTypeChord:
		attr.Value = "chord"
	case ST_ShapeTypeArc:
		attr.Value = "arc"
	case ST_ShapeTypeLeftBracket:
		attr.Value = "leftBracket"
	case ST_ShapeTypeRightBracket:
		attr.Value = "rightBracket"
	case ST_ShapeTypeLeftBrace:
		attr.Value = "leftBrace"
	case ST_ShapeTypeRightBrace:
		attr.Value = "rightBrace"
	case ST_ShapeTypeBracketPair:
		attr.Value = "bracketPair"
	case ST_ShapeTypeBracePair:
		attr.Value = "bracePair"
	case ST_ShapeTypeStraightConnector1:
		attr.Value = "straightConnector1"
	case ST_ShapeTypeBentConnector2:
		attr.Value = "bentConnector2"
	case ST_ShapeTypeBentConnector3:
		attr.Value = "bentConnector3"
	case ST_ShapeTypeBentConnector4:
		attr.Value = "bentConnector4"
	case ST_ShapeTypeBentConnector5:
		attr.Value = "bentConnector5"
	case ST_ShapeTypeCurvedConnector2:
		attr.Value = "curvedConnector2"
	case ST_ShapeTypeCurvedConnector3:
		attr.Value = "curvedConnector3"
	case ST_ShapeTypeCurvedConnector4:
		attr.Value = "curvedConnector4"
	case ST_ShapeTypeCurvedConnector5:
		attr.Value = "curvedConnector5"
	case ST_ShapeTypeCallout1:
		attr.Value = "callout1"
	case ST_ShapeTypeCallout2:
		attr.Value = "callout2"
	case ST_ShapeTypeCallout3:
		attr.Value = "callout3"
	case ST_ShapeTypeAccentCallout1:
		attr.Value = "accentCallout1"
	case ST_ShapeTypeAccentCallout2:
		attr.Value = "accentCallout2"
	case ST_ShapeTypeAccentCallout3:
		attr.Value = "accentCallout3"
	case ST_ShapeTypeBorderCallout1:
		attr.Value = "borderCallout1"
	case ST_ShapeTypeBorderCallout2:
		attr.Value = "borderCallout2"
	case ST_ShapeTypeBorderCallout3:
		attr.Value = "borderCallout3"
	case ST_ShapeTypeAccentBorderCallout1:
		attr.Value = "accentBorderCallout1"
	case ST_ShapeTypeAccentBorderCallout2:
		attr.Value = "accentBorderCallout2"
	case ST_ShapeTypeAccentBorderCallout3:
		attr.Value = "accentBorderCallout3"
	case ST_ShapeTypeWedgeRectCallout:
		attr.Value = "wedgeRectCallout"
	case ST_ShapeTypeWedgeRoundRectCallout:
		attr.Value = "wedgeRoundRectCallout"
	case ST_ShapeTypeWedgeEllipseCallout:
		attr.Value = "wedgeEllipseCallout"
	case ST_ShapeTypeCloudCallout:
		attr.Value = "cloudCallout"
	case ST_ShapeTypeCloud:
		attr.Value = "cloud"
	case ST_ShapeTypeRibbon:
		attr.Value = "ribbon"
	case ST_ShapeTypeRibbon2:
		attr.Value = "ribbon2"
	case ST_ShapeTypeEllipseRibbon:
		attr.Value = "ellipseRibbon"
	case ST_ShapeTypeEllipseRibbon2:
		attr.Value = "ellipseRibbon2"
	case ST_ShapeTypeLeftRightRibbon:
		attr.Value = "leftRightRibbon"
	case ST_ShapeTypeVerticalScroll:
		attr.Value = "verticalScroll"
	case ST_ShapeTypeHorizontalScroll:
		attr.Value = "horizontalScroll"
	case ST_ShapeTypeWave:
		attr.Value = "wave"
	case ST_ShapeTypeDoubleWave:
		attr.Value = "doubleWave"
	case ST_ShapeTypePlus:
		attr.Value = "plus"
	case ST_ShapeTypeFlowChartProcess:
		attr.Value = "flowChartProcess"
	case ST_ShapeTypeFlowChartDecision:
		attr.Value = "flowChartDecision"
	case ST_ShapeTypeFlowChartInputOutput:
		attr.Value = "flowChartInputOutput"
	case ST_ShapeTypeFlowChartPredefinedProcess:
		attr.Value = "flowChartPredefinedProcess"
	case ST_ShapeTypeFlowChartInternalStorage:
		attr.Value = "flowChartInternalStorage"
	case ST_ShapeTypeFlowChartDocument:
		attr.Value = "flowChartDocument"
	case ST_ShapeTypeFlowChartMultidocument:
		attr.Value = "flowChartMultidocument"
	case ST_ShapeTypeFlowChartTerminator:
		attr.Value = "flowChartTerminator"
	case ST_ShapeTypeFlowChartPreparation:
		attr.Value = "flowChartPreparation"
	case ST_ShapeTypeFlowChartManualInput:
		attr.Value = "flowChartManualInput"
	case ST_ShapeTypeFlowChartManualOperation:
		attr.Value = "flowChartManualOperation"
	case ST_ShapeTypeFlowChartConnector:
		attr.Value = "flowChartConnector"
	case ST_ShapeTypeFlowChartPunchedCard:
		attr.Value = "flowChartPunchedCard"
	case ST_ShapeTypeFlowChartPunchedTape:
		attr.Value = "flowChartPunchedTape"
	case ST_ShapeTypeFlowChartSummingJunction:
		attr.Value = "flowChartSummingJunction"
	case ST_ShapeTypeFlowChartOr:
		attr.Value = "flowChartOr"
	case ST_ShapeTypeFlowChartCollate:
		attr.Value = "flowChartCollate"
	case ST_ShapeTypeFlowChartSort:
		attr.Value = "flowChartSort"
	case ST_ShapeTypeFlowChartExtract:
		attr.Value = "flowChartExtract"
	case ST_ShapeTypeFlowChartMerge:
		attr.Value = "flowChartMerge"
	case ST_ShapeTypeFlowChartOfflineStorage:
		attr.Value = "flowChartOfflineStorage"
	case ST_ShapeTypeFlowChartOnlineStorage:
		attr.Value = "flowChartOnlineStorage"
	case ST_ShapeTypeFlowChartMagneticTape:
		attr.Value = "flowChartMagneticTape"
	case ST_ShapeTypeFlowChartMagneticDisk:
		attr.Value = "flowChartMagneticDisk"
	case ST_ShapeTypeFlowChartMagneticDrum:
		attr.Value = "flowChartMagneticDrum"
	case ST_ShapeTypeFlowChartDisplay:
		attr.Value = "flowChartDisplay"
	case ST_ShapeTypeFlowChartDelay:
		attr.Value = "flowChartDelay"
	case ST_ShapeTypeFlowChartAlternateProcess:
		attr.Value = "flowChartAlternateProcess"
	case ST_ShapeTypeFlowChartOffpageConnector:
		attr.Value = "flowChartOffpageConnector"
	case ST_ShapeTypeActionButtonBlank:
		attr.Value = "actionButtonBlank"
	case ST_ShapeTypeActionButtonHome:
		attr.Value = "actionButtonHome"
	case ST_ShapeTypeActionButtonHelp:
		attr.Value = "actionButtonHelp"
	case ST_ShapeTypeActionButtonInformation:
		attr.Value = "actionButtonInformation"
	case ST_ShapeTypeActionButtonForwardNext:
		attr.Value = "actionButtonForwardNext"
	case ST_ShapeTypeActionButtonBackPrevious:
		attr.Value = "actionButtonBackPrevious"
	case ST_ShapeTypeActionButtonEnd:
		attr.Value = "actionButtonEnd"
	case ST_ShapeTypeActionButtonBeginning:
		attr.Value = "actionButtonBeginning"
	case ST_ShapeTypeActionButtonReturn:
		attr.Value = "actionButtonReturn"
	case ST_ShapeTypeActionButtonDocument:
		attr.Value = "actionButtonDocument"
	case ST_ShapeTypeActionButtonSound:
		attr.Value = "actionButtonSound"
	case ST_ShapeTypeActionButtonMovie:
		attr.Value = "actionButtonMovie"
	case ST_ShapeTypeGear6:
		attr.Value = "gear6"
	case ST_ShapeTypeGear9:
		attr.Value = "gear9"
	case ST_ShapeTypeFunnel:
		attr.Value = "funnel"
	case ST_ShapeTypeMathPlus:
		attr.Value = "mathPlus"
	case ST_ShapeTypeMathMinus:
		attr.Value = "mathMinus"
	case ST_ShapeTypeMathMultiply:
		attr.Value = "mathMultiply"
	case ST_ShapeTypeMathDivide:
		attr.Value = "mathDivide"
	case ST_ShapeTypeMathEqual:
		attr.Value = "mathEqual"
	case ST_ShapeTypeMathNotEqual:
		attr.Value = "mathNotEqual"
	case ST_ShapeTypeCornerTabs:
		attr.Value = "cornerTabs"
	case ST_ShapeTypeSquareTabs:
		attr.Value = "squareTabs"
	case ST_ShapeTypePlaqueTabs:
		attr.Value = "plaqueTabs"
	case ST_ShapeTypeChartX:
		attr.Value = "chartX"
	case ST_ShapeTypeChartStar:
		attr.Value = "chartStar"
	case ST_ShapeTypeChartPlus:
		attr.Value = "chartPlus"
	}
	return attr, nil
}
func (e *ST_ShapeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "line":
		*e = 1
	case "lineInv":
		*e = 2
	case "triangle":
		*e = 3
	case "rtTriangle":
		*e = 4
	case "rect":
		*e = 5
	case "diamond":
		*e = 6
	case "parallelogram":
		*e = 7
	case "trapezoid":
		*e = 8
	case "nonIsoscelesTrapezoid":
		*e = 9
	case "pentagon":
		*e = 10
	case "hexagon":
		*e = 11
	case "heptagon":
		*e = 12
	case "octagon":
		*e = 13
	case "decagon":
		*e = 14
	case "dodecagon":
		*e = 15
	case "star4":
		*e = 16
	case "star5":
		*e = 17
	case "star6":
		*e = 18
	case "star7":
		*e = 19
	case "star8":
		*e = 20
	case "star10":
		*e = 21
	case "star12":
		*e = 22
	case "star16":
		*e = 23
	case "star24":
		*e = 24
	case "star32":
		*e = 25
	case "roundRect":
		*e = 26
	case "round1Rect":
		*e = 27
	case "round2SameRect":
		*e = 28
	case "round2DiagRect":
		*e = 29
	case "snipRoundRect":
		*e = 30
	case "snip1Rect":
		*e = 31
	case "snip2SameRect":
		*e = 32
	case "snip2DiagRect":
		*e = 33
	case "plaque":
		*e = 34
	case "ellipse":
		*e = 35
	case "teardrop":
		*e = 36
	case "homePlate":
		*e = 37
	case "chevron":
		*e = 38
	case "pieWedge":
		*e = 39
	case "pie":
		*e = 40
	case "blockArc":
		*e = 41
	case "donut":
		*e = 42
	case "noSmoking":
		*e = 43
	case "rightArrow":
		*e = 44
	case "leftArrow":
		*e = 45
	case "upArrow":
		*e = 46
	case "downArrow":
		*e = 47
	case "stripedRightArrow":
		*e = 48
	case "notchedRightArrow":
		*e = 49
	case "bentUpArrow":
		*e = 50
	case "leftRightArrow":
		*e = 51
	case "upDownArrow":
		*e = 52
	case "leftUpArrow":
		*e = 53
	case "leftRightUpArrow":
		*e = 54
	case "quadArrow":
		*e = 55
	case "leftArrowCallout":
		*e = 56
	case "rightArrowCallout":
		*e = 57
	case "upArrowCallout":
		*e = 58
	case "downArrowCallout":
		*e = 59
	case "leftRightArrowCallout":
		*e = 60
	case "upDownArrowCallout":
		*e = 61
	case "quadArrowCallout":
		*e = 62
	case "bentArrow":
		*e = 63
	case "uturnArrow":
		*e = 64
	case "circularArrow":
		*e = 65
	case "leftCircularArrow":
		*e = 66
	case "leftRightCircularArrow":
		*e = 67
	case "curvedRightArrow":
		*e = 68
	case "curvedLeftArrow":
		*e = 69
	case "curvedUpArrow":
		*e = 70
	case "curvedDownArrow":
		*e = 71
	case "swooshArrow":
		*e = 72
	case "cube":
		*e = 73
	case "can":
		*e = 74
	case "lightningBolt":
		*e = 75
	case "heart":
		*e = 76
	case "sun":
		*e = 77
	case "moon":
		*e = 78
	case "smileyFace":
		*e = 79
	case "irregularSeal1":
		*e = 80
	case "irregularSeal2":
		*e = 81
	case "foldedCorner":
		*e = 82
	case "bevel":
		*e = 83
	case "frame":
		*e = 84
	case "halfFrame":
		*e = 85
	case "corner":
		*e = 86
	case "diagStripe":
		*e = 87
	case "chord":
		*e = 88
	case "arc":
		*e = 89
	case "leftBracket":
		*e = 90
	case "rightBracket":
		*e = 91
	case "leftBrace":
		*e = 92
	case "rightBrace":
		*e = 93
	case "bracketPair":
		*e = 94
	case "bracePair":
		*e = 95
	case "straightConnector1":
		*e = 96
	case "bentConnector2":
		*e = 97
	case "bentConnector3":
		*e = 98
	case "bentConnector4":
		*e = 99
	case "bentConnector5":
		*e = 100
	case "curvedConnector2":
		*e = 101
	case "curvedConnector3":
		*e = 102
	case "curvedConnector4":
		*e = 103
	case "curvedConnector5":
		*e = 104
	case "callout1":
		*e = 105
	case "callout2":
		*e = 106
	case "callout3":
		*e = 107
	case "accentCallout1":
		*e = 108
	case "accentCallout2":
		*e = 109
	case "accentCallout3":
		*e = 110
	case "borderCallout1":
		*e = 111
	case "borderCallout2":
		*e = 112
	case "borderCallout3":
		*e = 113
	case "accentBorderCallout1":
		*e = 114
	case "accentBorderCallout2":
		*e = 115
	case "accentBorderCallout3":
		*e = 116
	case "wedgeRectCallout":
		*e = 117
	case "wedgeRoundRectCallout":
		*e = 118
	case "wedgeEllipseCallout":
		*e = 119
	case "cloudCallout":
		*e = 120
	case "cloud":
		*e = 121
	case "ribbon":
		*e = 122
	case "ribbon2":
		*e = 123
	case "ellipseRibbon":
		*e = 124
	case "ellipseRibbon2":
		*e = 125
	case "leftRightRibbon":
		*e = 126
	case "verticalScroll":
		*e = 127
	case "horizontalScroll":
		*e = 128
	case "wave":
		*e = 129
	case "doubleWave":
		*e = 130
	case "plus":
		*e = 131
	case "flowChartProcess":
		*e = 132
	case "flowChartDecision":
		*e = 133
	case "flowChartInputOutput":
		*e = 134
	case "flowChartPredefinedProcess":
		*e = 135
	case "flowChartInternalStorage":
		*e = 136
	case "flowChartDocument":
		*e = 137
	case "flowChartMultidocument":
		*e = 138
	case "flowChartTerminator":
		*e = 139
	case "flowChartPreparation":
		*e = 140
	case "flowChartManualInput":
		*e = 141
	case "flowChartManualOperation":
		*e = 142
	case "flowChartConnector":
		*e = 143
	case "flowChartPunchedCard":
		*e = 144
	case "flowChartPunchedTape":
		*e = 145
	case "flowChartSummingJunction":
		*e = 146
	case "flowChartOr":
		*e = 147
	case "flowChartCollate":
		*e = 148
	case "flowChartSort":
		*e = 149
	case "flowChartExtract":
		*e = 150
	case "flowChartMerge":
		*e = 151
	case "flowChartOfflineStorage":
		*e = 152
	case "flowChartOnlineStorage":
		*e = 153
	case "flowChartMagneticTape":
		*e = 154
	case "flowChartMagneticDisk":
		*e = 155
	case "flowChartMagneticDrum":
		*e = 156
	case "flowChartDisplay":
		*e = 157
	case "flowChartDelay":
		*e = 158
	case "flowChartAlternateProcess":
		*e = 159
	case "flowChartOffpageConnector":
		*e = 160
	case "actionButtonBlank":
		*e = 161
	case "actionButtonHome":
		*e = 162
	case "actionButtonHelp":
		*e = 163
	case "actionButtonInformation":
		*e = 164
	case "actionButtonForwardNext":
		*e = 165
	case "actionButtonBackPrevious":
		*e = 166
	case "actionButtonEnd":
		*e = 167
	case "actionButtonBeginning":
		*e = 168
	case "actionButtonReturn":
		*e = 169
	case "actionButtonDocument":
		*e = 170
	case "actionButtonSound":
		*e = 171
	case "actionButtonMovie":
		*e = 172
	case "gear6":
		*e = 173
	case "gear9":
		*e = 174
	case "funnel":
		*e = 175
	case "mathPlus":
		*e = 176
	case "mathMinus":
		*e = 177
	case "mathMultiply":
		*e = 178
	case "mathDivide":
		*e = 179
	case "mathEqual":
		*e = 180
	case "mathNotEqual":
		*e = 181
	case "cornerTabs":
		*e = 182
	case "squareTabs":
		*e = 183
	case "plaqueTabs":
		*e = 184
	case "chartX":
		*e = 185
	case "chartStar":
		*e = 186
	case "chartPlus":
		*e = 187
	}
	return nil
}
func (m ST_ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ShapeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "line":
			*m = 1
		case "lineInv":
			*m = 2
		case "triangle":
			*m = 3
		case "rtTriangle":
			*m = 4
		case "rect":
			*m = 5
		case "diamond":
			*m = 6
		case "parallelogram":
			*m = 7
		case "trapezoid":
			*m = 8
		case "nonIsoscelesTrapezoid":
			*m = 9
		case "pentagon":
			*m = 10
		case "hexagon":
			*m = 11
		case "heptagon":
			*m = 12
		case "octagon":
			*m = 13
		case "decagon":
			*m = 14
		case "dodecagon":
			*m = 15
		case "star4":
			*m = 16
		case "star5":
			*m = 17
		case "star6":
			*m = 18
		case "star7":
			*m = 19
		case "star8":
			*m = 20
		case "star10":
			*m = 21
		case "star12":
			*m = 22
		case "star16":
			*m = 23
		case "star24":
			*m = 24
		case "star32":
			*m = 25
		case "roundRect":
			*m = 26
		case "round1Rect":
			*m = 27
		case "round2SameRect":
			*m = 28
		case "round2DiagRect":
			*m = 29
		case "snipRoundRect":
			*m = 30
		case "snip1Rect":
			*m = 31
		case "snip2SameRect":
			*m = 32
		case "snip2DiagRect":
			*m = 33
		case "plaque":
			*m = 34
		case "ellipse":
			*m = 35
		case "teardrop":
			*m = 36
		case "homePlate":
			*m = 37
		case "chevron":
			*m = 38
		case "pieWedge":
			*m = 39
		case "pie":
			*m = 40
		case "blockArc":
			*m = 41
		case "donut":
			*m = 42
		case "noSmoking":
			*m = 43
		case "rightArrow":
			*m = 44
		case "leftArrow":
			*m = 45
		case "upArrow":
			*m = 46
		case "downArrow":
			*m = 47
		case "stripedRightArrow":
			*m = 48
		case "notchedRightArrow":
			*m = 49
		case "bentUpArrow":
			*m = 50
		case "leftRightArrow":
			*m = 51
		case "upDownArrow":
			*m = 52
		case "leftUpArrow":
			*m = 53
		case "leftRightUpArrow":
			*m = 54
		case "quadArrow":
			*m = 55
		case "leftArrowCallout":
			*m = 56
		case "rightArrowCallout":
			*m = 57
		case "upArrowCallout":
			*m = 58
		case "downArrowCallout":
			*m = 59
		case "leftRightArrowCallout":
			*m = 60
		case "upDownArrowCallout":
			*m = 61
		case "quadArrowCallout":
			*m = 62
		case "bentArrow":
			*m = 63
		case "uturnArrow":
			*m = 64
		case "circularArrow":
			*m = 65
		case "leftCircularArrow":
			*m = 66
		case "leftRightCircularArrow":
			*m = 67
		case "curvedRightArrow":
			*m = 68
		case "curvedLeftArrow":
			*m = 69
		case "curvedUpArrow":
			*m = 70
		case "curvedDownArrow":
			*m = 71
		case "swooshArrow":
			*m = 72
		case "cube":
			*m = 73
		case "can":
			*m = 74
		case "lightningBolt":
			*m = 75
		case "heart":
			*m = 76
		case "sun":
			*m = 77
		case "moon":
			*m = 78
		case "smileyFace":
			*m = 79
		case "irregularSeal1":
			*m = 80
		case "irregularSeal2":
			*m = 81
		case "foldedCorner":
			*m = 82
		case "bevel":
			*m = 83
		case "frame":
			*m = 84
		case "halfFrame":
			*m = 85
		case "corner":
			*m = 86
		case "diagStripe":
			*m = 87
		case "chord":
			*m = 88
		case "arc":
			*m = 89
		case "leftBracket":
			*m = 90
		case "rightBracket":
			*m = 91
		case "leftBrace":
			*m = 92
		case "rightBrace":
			*m = 93
		case "bracketPair":
			*m = 94
		case "bracePair":
			*m = 95
		case "straightConnector1":
			*m = 96
		case "bentConnector2":
			*m = 97
		case "bentConnector3":
			*m = 98
		case "bentConnector4":
			*m = 99
		case "bentConnector5":
			*m = 100
		case "curvedConnector2":
			*m = 101
		case "curvedConnector3":
			*m = 102
		case "curvedConnector4":
			*m = 103
		case "curvedConnector5":
			*m = 104
		case "callout1":
			*m = 105
		case "callout2":
			*m = 106
		case "callout3":
			*m = 107
		case "accentCallout1":
			*m = 108
		case "accentCallout2":
			*m = 109
		case "accentCallout3":
			*m = 110
		case "borderCallout1":
			*m = 111
		case "borderCallout2":
			*m = 112
		case "borderCallout3":
			*m = 113
		case "accentBorderCallout1":
			*m = 114
		case "accentBorderCallout2":
			*m = 115
		case "accentBorderCallout3":
			*m = 116
		case "wedgeRectCallout":
			*m = 117
		case "wedgeRoundRectCallout":
			*m = 118
		case "wedgeEllipseCallout":
			*m = 119
		case "cloudCallout":
			*m = 120
		case "cloud":
			*m = 121
		case "ribbon":
			*m = 122
		case "ribbon2":
			*m = 123
		case "ellipseRibbon":
			*m = 124
		case "ellipseRibbon2":
			*m = 125
		case "leftRightRibbon":
			*m = 126
		case "verticalScroll":
			*m = 127
		case "horizontalScroll":
			*m = 128
		case "wave":
			*m = 129
		case "doubleWave":
			*m = 130
		case "plus":
			*m = 131
		case "flowChartProcess":
			*m = 132
		case "flowChartDecision":
			*m = 133
		case "flowChartInputOutput":
			*m = 134
		case "flowChartPredefinedProcess":
			*m = 135
		case "flowChartInternalStorage":
			*m = 136
		case "flowChartDocument":
			*m = 137
		case "flowChartMultidocument":
			*m = 138
		case "flowChartTerminator":
			*m = 139
		case "flowChartPreparation":
			*m = 140
		case "flowChartManualInput":
			*m = 141
		case "flowChartManualOperation":
			*m = 142
		case "flowChartConnector":
			*m = 143
		case "flowChartPunchedCard":
			*m = 144
		case "flowChartPunchedTape":
			*m = 145
		case "flowChartSummingJunction":
			*m = 146
		case "flowChartOr":
			*m = 147
		case "flowChartCollate":
			*m = 148
		case "flowChartSort":
			*m = 149
		case "flowChartExtract":
			*m = 150
		case "flowChartMerge":
			*m = 151
		case "flowChartOfflineStorage":
			*m = 152
		case "flowChartOnlineStorage":
			*m = 153
		case "flowChartMagneticTape":
			*m = 154
		case "flowChartMagneticDisk":
			*m = 155
		case "flowChartMagneticDrum":
			*m = 156
		case "flowChartDisplay":
			*m = 157
		case "flowChartDelay":
			*m = 158
		case "flowChartAlternateProcess":
			*m = 159
		case "flowChartOffpageConnector":
			*m = 160
		case "actionButtonBlank":
			*m = 161
		case "actionButtonHome":
			*m = 162
		case "actionButtonHelp":
			*m = 163
		case "actionButtonInformation":
			*m = 164
		case "actionButtonForwardNext":
			*m = 165
		case "actionButtonBackPrevious":
			*m = 166
		case "actionButtonEnd":
			*m = 167
		case "actionButtonBeginning":
			*m = 168
		case "actionButtonReturn":
			*m = 169
		case "actionButtonDocument":
			*m = 170
		case "actionButtonSound":
			*m = 171
		case "actionButtonMovie":
			*m = 172
		case "gear6":
			*m = 173
		case "gear9":
			*m = 174
		case "funnel":
			*m = 175
		case "mathPlus":
			*m = 176
		case "mathMinus":
			*m = 177
		case "mathMultiply":
			*m = 178
		case "mathDivide":
			*m = 179
		case "mathEqual":
			*m = 180
		case "mathNotEqual":
			*m = 181
		case "cornerTabs":
			*m = 182
		case "squareTabs":
			*m = 183
		case "plaqueTabs":
			*m = 184
		case "chartX":
			*m = 185
		case "chartStar":
			*m = 186
		case "chartPlus":
			*m = 187
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
func (m ST_ShapeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "line"
	case 2:
		return "lineInv"
	case 3:
		return "triangle"
	case 4:
		return "rtTriangle"
	case 5:
		return "rect"
	case 6:
		return "diamond"
	case 7:
		return "parallelogram"
	case 8:
		return "trapezoid"
	case 9:
		return "nonIsoscelesTrapezoid"
	case 10:
		return "pentagon"
	case 11:
		return "hexagon"
	case 12:
		return "heptagon"
	case 13:
		return "octagon"
	case 14:
		return "decagon"
	case 15:
		return "dodecagon"
	case 16:
		return "star4"
	case 17:
		return "star5"
	case 18:
		return "star6"
	case 19:
		return "star7"
	case 20:
		return "star8"
	case 21:
		return "star10"
	case 22:
		return "star12"
	case 23:
		return "star16"
	case 24:
		return "star24"
	case 25:
		return "star32"
	case 26:
		return "roundRect"
	case 27:
		return "round1Rect"
	case 28:
		return "round2SameRect"
	case 29:
		return "round2DiagRect"
	case 30:
		return "snipRoundRect"
	case 31:
		return "snip1Rect"
	case 32:
		return "snip2SameRect"
	case 33:
		return "snip2DiagRect"
	case 34:
		return "plaque"
	case 35:
		return "ellipse"
	case 36:
		return "teardrop"
	case 37:
		return "homePlate"
	case 38:
		return "chevron"
	case 39:
		return "pieWedge"
	case 40:
		return "pie"
	case 41:
		return "blockArc"
	case 42:
		return "donut"
	case 43:
		return "noSmoking"
	case 44:
		return "rightArrow"
	case 45:
		return "leftArrow"
	case 46:
		return "upArrow"
	case 47:
		return "downArrow"
	case 48:
		return "stripedRightArrow"
	case 49:
		return "notchedRightArrow"
	case 50:
		return "bentUpArrow"
	case 51:
		return "leftRightArrow"
	case 52:
		return "upDownArrow"
	case 53:
		return "leftUpArrow"
	case 54:
		return "leftRightUpArrow"
	case 55:
		return "quadArrow"
	case 56:
		return "leftArrowCallout"
	case 57:
		return "rightArrowCallout"
	case 58:
		return "upArrowCallout"
	case 59:
		return "downArrowCallout"
	case 60:
		return "leftRightArrowCallout"
	case 61:
		return "upDownArrowCallout"
	case 62:
		return "quadArrowCallout"
	case 63:
		return "bentArrow"
	case 64:
		return "uturnArrow"
	case 65:
		return "circularArrow"
	case 66:
		return "leftCircularArrow"
	case 67:
		return "leftRightCircularArrow"
	case 68:
		return "curvedRightArrow"
	case 69:
		return "curvedLeftArrow"
	case 70:
		return "curvedUpArrow"
	case 71:
		return "curvedDownArrow"
	case 72:
		return "swooshArrow"
	case 73:
		return "cube"
	case 74:
		return "can"
	case 75:
		return "lightningBolt"
	case 76:
		return "heart"
	case 77:
		return "sun"
	case 78:
		return "moon"
	case 79:
		return "smileyFace"
	case 80:
		return "irregularSeal1"
	case 81:
		return "irregularSeal2"
	case 82:
		return "foldedCorner"
	case 83:
		return "bevel"
	case 84:
		return "frame"
	case 85:
		return "halfFrame"
	case 86:
		return "corner"
	case 87:
		return "diagStripe"
	case 88:
		return "chord"
	case 89:
		return "arc"
	case 90:
		return "leftBracket"
	case 91:
		return "rightBracket"
	case 92:
		return "leftBrace"
	case 93:
		return "rightBrace"
	case 94:
		return "bracketPair"
	case 95:
		return "bracePair"
	case 96:
		return "straightConnector1"
	case 97:
		return "bentConnector2"
	case 98:
		return "bentConnector3"
	case 99:
		return "bentConnector4"
	case 100:
		return "bentConnector5"
	case 101:
		return "curvedConnector2"
	case 102:
		return "curvedConnector3"
	case 103:
		return "curvedConnector4"
	case 104:
		return "curvedConnector5"
	case 105:
		return "callout1"
	case 106:
		return "callout2"
	case 107:
		return "callout3"
	case 108:
		return "accentCallout1"
	case 109:
		return "accentCallout2"
	case 110:
		return "accentCallout3"
	case 111:
		return "borderCallout1"
	case 112:
		return "borderCallout2"
	case 113:
		return "borderCallout3"
	case 114:
		return "accentBorderCallout1"
	case 115:
		return "accentBorderCallout2"
	case 116:
		return "accentBorderCallout3"
	case 117:
		return "wedgeRectCallout"
	case 118:
		return "wedgeRoundRectCallout"
	case 119:
		return "wedgeEllipseCallout"
	case 120:
		return "cloudCallout"
	case 121:
		return "cloud"
	case 122:
		return "ribbon"
	case 123:
		return "ribbon2"
	case 124:
		return "ellipseRibbon"
	case 125:
		return "ellipseRibbon2"
	case 126:
		return "leftRightRibbon"
	case 127:
		return "verticalScroll"
	case 128:
		return "horizontalScroll"
	case 129:
		return "wave"
	case 130:
		return "doubleWave"
	case 131:
		return "plus"
	case 132:
		return "flowChartProcess"
	case 133:
		return "flowChartDecision"
	case 134:
		return "flowChartInputOutput"
	case 135:
		return "flowChartPredefinedProcess"
	case 136:
		return "flowChartInternalStorage"
	case 137:
		return "flowChartDocument"
	case 138:
		return "flowChartMultidocument"
	case 139:
		return "flowChartTerminator"
	case 140:
		return "flowChartPreparation"
	case 141:
		return "flowChartManualInput"
	case 142:
		return "flowChartManualOperation"
	case 143:
		return "flowChartConnector"
	case 144:
		return "flowChartPunchedCard"
	case 145:
		return "flowChartPunchedTape"
	case 146:
		return "flowChartSummingJunction"
	case 147:
		return "flowChartOr"
	case 148:
		return "flowChartCollate"
	case 149:
		return "flowChartSort"
	case 150:
		return "flowChartExtract"
	case 151:
		return "flowChartMerge"
	case 152:
		return "flowChartOfflineStorage"
	case 153:
		return "flowChartOnlineStorage"
	case 154:
		return "flowChartMagneticTape"
	case 155:
		return "flowChartMagneticDisk"
	case 156:
		return "flowChartMagneticDrum"
	case 157:
		return "flowChartDisplay"
	case 158:
		return "flowChartDelay"
	case 159:
		return "flowChartAlternateProcess"
	case 160:
		return "flowChartOffpageConnector"
	case 161:
		return "actionButtonBlank"
	case 162:
		return "actionButtonHome"
	case 163:
		return "actionButtonHelp"
	case 164:
		return "actionButtonInformation"
	case 165:
		return "actionButtonForwardNext"
	case 166:
		return "actionButtonBackPrevious"
	case 167:
		return "actionButtonEnd"
	case 168:
		return "actionButtonBeginning"
	case 169:
		return "actionButtonReturn"
	case 170:
		return "actionButtonDocument"
	case 171:
		return "actionButtonSound"
	case 172:
		return "actionButtonMovie"
	case 173:
		return "gear6"
	case 174:
		return "gear9"
	case 175:
		return "funnel"
	case 176:
		return "mathPlus"
	case 177:
		return "mathMinus"
	case 178:
		return "mathMultiply"
	case 179:
		return "mathDivide"
	case 180:
		return "mathEqual"
	case 181:
		return "mathNotEqual"
	case 182:
		return "cornerTabs"
	case 183:
		return "squareTabs"
	case 184:
		return "plaqueTabs"
	case 185:
		return "chartX"
	case 186:
		return "chartStar"
	case 187:
		return "chartPlus"
	}
	return ""
}
func (m ST_ShapeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ShapeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextShapeType byte

const (
	ST_TextShapeTypeUnset                     ST_TextShapeType = 0
	ST_TextShapeTypeTextNoShape               ST_TextShapeType = 1
	ST_TextShapeTypeTextPlain                 ST_TextShapeType = 2
	ST_TextShapeTypeTextStop                  ST_TextShapeType = 3
	ST_TextShapeTypeTextTriangle              ST_TextShapeType = 4
	ST_TextShapeTypeTextTriangleInverted      ST_TextShapeType = 5
	ST_TextShapeTypeTextChevron               ST_TextShapeType = 6
	ST_TextShapeTypeTextChevronInverted       ST_TextShapeType = 7
	ST_TextShapeTypeTextRingInside            ST_TextShapeType = 8
	ST_TextShapeTypeTextRingOutside           ST_TextShapeType = 9
	ST_TextShapeTypeTextArchUp                ST_TextShapeType = 10
	ST_TextShapeTypeTextArchDown              ST_TextShapeType = 11
	ST_TextShapeTypeTextCircle                ST_TextShapeType = 12
	ST_TextShapeTypeTextButton                ST_TextShapeType = 13
	ST_TextShapeTypeTextArchUpPour            ST_TextShapeType = 14
	ST_TextShapeTypeTextArchDownPour          ST_TextShapeType = 15
	ST_TextShapeTypeTextCirclePour            ST_TextShapeType = 16
	ST_TextShapeTypeTextButtonPour            ST_TextShapeType = 17
	ST_TextShapeTypeTextCurveUp               ST_TextShapeType = 18
	ST_TextShapeTypeTextCurveDown             ST_TextShapeType = 19
	ST_TextShapeTypeTextCanUp                 ST_TextShapeType = 20
	ST_TextShapeTypeTextCanDown               ST_TextShapeType = 21
	ST_TextShapeTypeTextWave1                 ST_TextShapeType = 22
	ST_TextShapeTypeTextWave2                 ST_TextShapeType = 23
	ST_TextShapeTypeTextDoubleWave1           ST_TextShapeType = 24
	ST_TextShapeTypeTextWave4                 ST_TextShapeType = 25
	ST_TextShapeTypeTextInflate               ST_TextShapeType = 26
	ST_TextShapeTypeTextDeflate               ST_TextShapeType = 27
	ST_TextShapeTypeTextInflateBottom         ST_TextShapeType = 28
	ST_TextShapeTypeTextDeflateBottom         ST_TextShapeType = 29
	ST_TextShapeTypeTextInflateTop            ST_TextShapeType = 30
	ST_TextShapeTypeTextDeflateTop            ST_TextShapeType = 31
	ST_TextShapeTypeTextDeflateInflate        ST_TextShapeType = 32
	ST_TextShapeTypeTextDeflateInflateDeflate ST_TextShapeType = 33
	ST_TextShapeTypeTextFadeRight             ST_TextShapeType = 34
	ST_TextShapeTypeTextFadeLeft              ST_TextShapeType = 35
	ST_TextShapeTypeTextFadeUp                ST_TextShapeType = 36
	ST_TextShapeTypeTextFadeDown              ST_TextShapeType = 37
	ST_TextShapeTypeTextSlantUp               ST_TextShapeType = 38
	ST_TextShapeTypeTextSlantDown             ST_TextShapeType = 39
	ST_TextShapeTypeTextCascadeUp             ST_TextShapeType = 40
	ST_TextShapeTypeTextCascadeDown           ST_TextShapeType = 41
)

func (e ST_TextShapeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextShapeTypeUnset:
		attr.Value = ""
	case ST_TextShapeTypeTextNoShape:
		attr.Value = "textNoShape"
	case ST_TextShapeTypeTextPlain:
		attr.Value = "textPlain"
	case ST_TextShapeTypeTextStop:
		attr.Value = "textStop"
	case ST_TextShapeTypeTextTriangle:
		attr.Value = "textTriangle"
	case ST_TextShapeTypeTextTriangleInverted:
		attr.Value = "textTriangleInverted"
	case ST_TextShapeTypeTextChevron:
		attr.Value = "textChevron"
	case ST_TextShapeTypeTextChevronInverted:
		attr.Value = "textChevronInverted"
	case ST_TextShapeTypeTextRingInside:
		attr.Value = "textRingInside"
	case ST_TextShapeTypeTextRingOutside:
		attr.Value = "textRingOutside"
	case ST_TextShapeTypeTextArchUp:
		attr.Value = "textArchUp"
	case ST_TextShapeTypeTextArchDown:
		attr.Value = "textArchDown"
	case ST_TextShapeTypeTextCircle:
		attr.Value = "textCircle"
	case ST_TextShapeTypeTextButton:
		attr.Value = "textButton"
	case ST_TextShapeTypeTextArchUpPour:
		attr.Value = "textArchUpPour"
	case ST_TextShapeTypeTextArchDownPour:
		attr.Value = "textArchDownPour"
	case ST_TextShapeTypeTextCirclePour:
		attr.Value = "textCirclePour"
	case ST_TextShapeTypeTextButtonPour:
		attr.Value = "textButtonPour"
	case ST_TextShapeTypeTextCurveUp:
		attr.Value = "textCurveUp"
	case ST_TextShapeTypeTextCurveDown:
		attr.Value = "textCurveDown"
	case ST_TextShapeTypeTextCanUp:
		attr.Value = "textCanUp"
	case ST_TextShapeTypeTextCanDown:
		attr.Value = "textCanDown"
	case ST_TextShapeTypeTextWave1:
		attr.Value = "textWave1"
	case ST_TextShapeTypeTextWave2:
		attr.Value = "textWave2"
	case ST_TextShapeTypeTextDoubleWave1:
		attr.Value = "textDoubleWave1"
	case ST_TextShapeTypeTextWave4:
		attr.Value = "textWave4"
	case ST_TextShapeTypeTextInflate:
		attr.Value = "textInflate"
	case ST_TextShapeTypeTextDeflate:
		attr.Value = "textDeflate"
	case ST_TextShapeTypeTextInflateBottom:
		attr.Value = "textInflateBottom"
	case ST_TextShapeTypeTextDeflateBottom:
		attr.Value = "textDeflateBottom"
	case ST_TextShapeTypeTextInflateTop:
		attr.Value = "textInflateTop"
	case ST_TextShapeTypeTextDeflateTop:
		attr.Value = "textDeflateTop"
	case ST_TextShapeTypeTextDeflateInflate:
		attr.Value = "textDeflateInflate"
	case ST_TextShapeTypeTextDeflateInflateDeflate:
		attr.Value = "textDeflateInflateDeflate"
	case ST_TextShapeTypeTextFadeRight:
		attr.Value = "textFadeRight"
	case ST_TextShapeTypeTextFadeLeft:
		attr.Value = "textFadeLeft"
	case ST_TextShapeTypeTextFadeUp:
		attr.Value = "textFadeUp"
	case ST_TextShapeTypeTextFadeDown:
		attr.Value = "textFadeDown"
	case ST_TextShapeTypeTextSlantUp:
		attr.Value = "textSlantUp"
	case ST_TextShapeTypeTextSlantDown:
		attr.Value = "textSlantDown"
	case ST_TextShapeTypeTextCascadeUp:
		attr.Value = "textCascadeUp"
	case ST_TextShapeTypeTextCascadeDown:
		attr.Value = "textCascadeDown"
	}
	return attr, nil
}
func (e *ST_TextShapeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "textNoShape":
		*e = 1
	case "textPlain":
		*e = 2
	case "textStop":
		*e = 3
	case "textTriangle":
		*e = 4
	case "textTriangleInverted":
		*e = 5
	case "textChevron":
		*e = 6
	case "textChevronInverted":
		*e = 7
	case "textRingInside":
		*e = 8
	case "textRingOutside":
		*e = 9
	case "textArchUp":
		*e = 10
	case "textArchDown":
		*e = 11
	case "textCircle":
		*e = 12
	case "textButton":
		*e = 13
	case "textArchUpPour":
		*e = 14
	case "textArchDownPour":
		*e = 15
	case "textCirclePour":
		*e = 16
	case "textButtonPour":
		*e = 17
	case "textCurveUp":
		*e = 18
	case "textCurveDown":
		*e = 19
	case "textCanUp":
		*e = 20
	case "textCanDown":
		*e = 21
	case "textWave1":
		*e = 22
	case "textWave2":
		*e = 23
	case "textDoubleWave1":
		*e = 24
	case "textWave4":
		*e = 25
	case "textInflate":
		*e = 26
	case "textDeflate":
		*e = 27
	case "textInflateBottom":
		*e = 28
	case "textDeflateBottom":
		*e = 29
	case "textInflateTop":
		*e = 30
	case "textDeflateTop":
		*e = 31
	case "textDeflateInflate":
		*e = 32
	case "textDeflateInflateDeflate":
		*e = 33
	case "textFadeRight":
		*e = 34
	case "textFadeLeft":
		*e = 35
	case "textFadeUp":
		*e = 36
	case "textFadeDown":
		*e = 37
	case "textSlantUp":
		*e = 38
	case "textSlantDown":
		*e = 39
	case "textCascadeUp":
		*e = 40
	case "textCascadeDown":
		*e = 41
	}
	return nil
}
func (m ST_TextShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextShapeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "textNoShape":
			*m = 1
		case "textPlain":
			*m = 2
		case "textStop":
			*m = 3
		case "textTriangle":
			*m = 4
		case "textTriangleInverted":
			*m = 5
		case "textChevron":
			*m = 6
		case "textChevronInverted":
			*m = 7
		case "textRingInside":
			*m = 8
		case "textRingOutside":
			*m = 9
		case "textArchUp":
			*m = 10
		case "textArchDown":
			*m = 11
		case "textCircle":
			*m = 12
		case "textButton":
			*m = 13
		case "textArchUpPour":
			*m = 14
		case "textArchDownPour":
			*m = 15
		case "textCirclePour":
			*m = 16
		case "textButtonPour":
			*m = 17
		case "textCurveUp":
			*m = 18
		case "textCurveDown":
			*m = 19
		case "textCanUp":
			*m = 20
		case "textCanDown":
			*m = 21
		case "textWave1":
			*m = 22
		case "textWave2":
			*m = 23
		case "textDoubleWave1":
			*m = 24
		case "textWave4":
			*m = 25
		case "textInflate":
			*m = 26
		case "textDeflate":
			*m = 27
		case "textInflateBottom":
			*m = 28
		case "textDeflateBottom":
			*m = 29
		case "textInflateTop":
			*m = 30
		case "textDeflateTop":
			*m = 31
		case "textDeflateInflate":
			*m = 32
		case "textDeflateInflateDeflate":
			*m = 33
		case "textFadeRight":
			*m = 34
		case "textFadeLeft":
			*m = 35
		case "textFadeUp":
			*m = 36
		case "textFadeDown":
			*m = 37
		case "textSlantUp":
			*m = 38
		case "textSlantDown":
			*m = 39
		case "textCascadeUp":
			*m = 40
		case "textCascadeDown":
			*m = 41
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
func (m ST_TextShapeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "textNoShape"
	case 2:
		return "textPlain"
	case 3:
		return "textStop"
	case 4:
		return "textTriangle"
	case 5:
		return "textTriangleInverted"
	case 6:
		return "textChevron"
	case 7:
		return "textChevronInverted"
	case 8:
		return "textRingInside"
	case 9:
		return "textRingOutside"
	case 10:
		return "textArchUp"
	case 11:
		return "textArchDown"
	case 12:
		return "textCircle"
	case 13:
		return "textButton"
	case 14:
		return "textArchUpPour"
	case 15:
		return "textArchDownPour"
	case 16:
		return "textCirclePour"
	case 17:
		return "textButtonPour"
	case 18:
		return "textCurveUp"
	case 19:
		return "textCurveDown"
	case 20:
		return "textCanUp"
	case 21:
		return "textCanDown"
	case 22:
		return "textWave1"
	case 23:
		return "textWave2"
	case 24:
		return "textDoubleWave1"
	case 25:
		return "textWave4"
	case 26:
		return "textInflate"
	case 27:
		return "textDeflate"
	case 28:
		return "textInflateBottom"
	case 29:
		return "textDeflateBottom"
	case 30:
		return "textInflateTop"
	case 31:
		return "textDeflateTop"
	case 32:
		return "textDeflateInflate"
	case 33:
		return "textDeflateInflateDeflate"
	case 34:
		return "textFadeRight"
	case 35:
		return "textFadeLeft"
	case 36:
		return "textFadeUp"
	case 37:
		return "textFadeDown"
	case 38:
		return "textSlantUp"
	case 39:
		return "textSlantDown"
	case 40:
		return "textCascadeUp"
	case 41:
		return "textCascadeDown"
	}
	return ""
}
func (m ST_TextShapeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextShapeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PathFillMode byte

const (
	ST_PathFillModeUnset       ST_PathFillMode = 0
	ST_PathFillModeNone        ST_PathFillMode = 1
	ST_PathFillModeNorm        ST_PathFillMode = 2
	ST_PathFillModeLighten     ST_PathFillMode = 3
	ST_PathFillModeLightenLess ST_PathFillMode = 4
	ST_PathFillModeDarken      ST_PathFillMode = 5
	ST_PathFillModeDarkenLess  ST_PathFillMode = 6
)

func (e ST_PathFillMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PathFillModeUnset:
		attr.Value = ""
	case ST_PathFillModeNone:
		attr.Value = "none"
	case ST_PathFillModeNorm:
		attr.Value = "norm"
	case ST_PathFillModeLighten:
		attr.Value = "lighten"
	case ST_PathFillModeLightenLess:
		attr.Value = "lightenLess"
	case ST_PathFillModeDarken:
		attr.Value = "darken"
	case ST_PathFillModeDarkenLess:
		attr.Value = "darkenLess"
	}
	return attr, nil
}
func (e *ST_PathFillMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "norm":
		*e = 2
	case "lighten":
		*e = 3
	case "lightenLess":
		*e = 4
	case "darken":
		*e = 5
	case "darkenLess":
		*e = 6
	}
	return nil
}
func (m ST_PathFillMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PathFillMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "norm":
			*m = 2
		case "lighten":
			*m = 3
		case "lightenLess":
			*m = 4
		case "darken":
			*m = 5
		case "darkenLess":
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
func (m ST_PathFillMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "norm"
	case 3:
		return "lighten"
	case 4:
		return "lightenLess"
	case 5:
		return "darken"
	case 6:
		return "darkenLess"
	}
	return ""
}
func (m ST_PathFillMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PathFillMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineEndType byte

const (
	ST_LineEndTypeUnset    ST_LineEndType = 0
	ST_LineEndTypeNone     ST_LineEndType = 1
	ST_LineEndTypeTriangle ST_LineEndType = 2
	ST_LineEndTypeStealth  ST_LineEndType = 3
	ST_LineEndTypeDiamond  ST_LineEndType = 4
	ST_LineEndTypeOval     ST_LineEndType = 5
	ST_LineEndTypeArrow    ST_LineEndType = 6
)

func (e ST_LineEndType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineEndTypeUnset:
		attr.Value = ""
	case ST_LineEndTypeNone:
		attr.Value = "none"
	case ST_LineEndTypeTriangle:
		attr.Value = "triangle"
	case ST_LineEndTypeStealth:
		attr.Value = "stealth"
	case ST_LineEndTypeDiamond:
		attr.Value = "diamond"
	case ST_LineEndTypeOval:
		attr.Value = "oval"
	case ST_LineEndTypeArrow:
		attr.Value = "arrow"
	}
	return attr, nil
}
func (e *ST_LineEndType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "triangle":
		*e = 2
	case "stealth":
		*e = 3
	case "diamond":
		*e = 4
	case "oval":
		*e = 5
	case "arrow":
		*e = 6
	}
	return nil
}
func (m ST_LineEndType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LineEndType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "triangle":
			*m = 2
		case "stealth":
			*m = 3
		case "diamond":
			*m = 4
		case "oval":
			*m = 5
		case "arrow":
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
func (m ST_LineEndType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "triangle"
	case 3:
		return "stealth"
	case 4:
		return "diamond"
	case 5:
		return "oval"
	case 6:
		return "arrow"
	}
	return ""
}
func (m ST_LineEndType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LineEndType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineEndWidth byte

const (
	ST_LineEndWidthUnset ST_LineEndWidth = 0
	ST_LineEndWidthSm    ST_LineEndWidth = 1
	ST_LineEndWidthMed   ST_LineEndWidth = 2
	ST_LineEndWidthLg    ST_LineEndWidth = 3
)

func (e ST_LineEndWidth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineEndWidthUnset:
		attr.Value = ""
	case ST_LineEndWidthSm:
		attr.Value = "sm"
	case ST_LineEndWidthMed:
		attr.Value = "med"
	case ST_LineEndWidthLg:
		attr.Value = "lg"
	}
	return attr, nil
}
func (e *ST_LineEndWidth) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sm":
		*e = 1
	case "med":
		*e = 2
	case "lg":
		*e = 3
	}
	return nil
}
func (m ST_LineEndWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LineEndWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sm":
			*m = 1
		case "med":
			*m = 2
		case "lg":
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
func (m ST_LineEndWidth) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sm"
	case 2:
		return "med"
	case 3:
		return "lg"
	}
	return ""
}
func (m ST_LineEndWidth) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LineEndWidth) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineEndLength byte

const (
	ST_LineEndLengthUnset ST_LineEndLength = 0
	ST_LineEndLengthSm    ST_LineEndLength = 1
	ST_LineEndLengthMed   ST_LineEndLength = 2
	ST_LineEndLengthLg    ST_LineEndLength = 3
)

func (e ST_LineEndLength) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineEndLengthUnset:
		attr.Value = ""
	case ST_LineEndLengthSm:
		attr.Value = "sm"
	case ST_LineEndLengthMed:
		attr.Value = "med"
	case ST_LineEndLengthLg:
		attr.Value = "lg"
	}
	return attr, nil
}
func (e *ST_LineEndLength) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sm":
		*e = 1
	case "med":
		*e = 2
	case "lg":
		*e = 3
	}
	return nil
}
func (m ST_LineEndLength) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LineEndLength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sm":
			*m = 1
		case "med":
			*m = 2
		case "lg":
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
func (m ST_LineEndLength) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sm"
	case 2:
		return "med"
	case 3:
		return "lg"
	}
	return ""
}
func (m ST_LineEndLength) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LineEndLength) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PresetLineDashVal byte

const (
	ST_PresetLineDashValUnset         ST_PresetLineDashVal = 0
	ST_PresetLineDashValSolid         ST_PresetLineDashVal = 1
	ST_PresetLineDashValDot           ST_PresetLineDashVal = 2
	ST_PresetLineDashValDash          ST_PresetLineDashVal = 3
	ST_PresetLineDashValLgDash        ST_PresetLineDashVal = 4
	ST_PresetLineDashValDashDot       ST_PresetLineDashVal = 5
	ST_PresetLineDashValLgDashDot     ST_PresetLineDashVal = 6
	ST_PresetLineDashValLgDashDotDot  ST_PresetLineDashVal = 7
	ST_PresetLineDashValSysDash       ST_PresetLineDashVal = 8
	ST_PresetLineDashValSysDot        ST_PresetLineDashVal = 9
	ST_PresetLineDashValSysDashDot    ST_PresetLineDashVal = 10
	ST_PresetLineDashValSysDashDotDot ST_PresetLineDashVal = 11
)

func (e ST_PresetLineDashVal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PresetLineDashValUnset:
		attr.Value = ""
	case ST_PresetLineDashValSolid:
		attr.Value = "solid"
	case ST_PresetLineDashValDot:
		attr.Value = "dot"
	case ST_PresetLineDashValDash:
		attr.Value = "dash"
	case ST_PresetLineDashValLgDash:
		attr.Value = "lgDash"
	case ST_PresetLineDashValDashDot:
		attr.Value = "dashDot"
	case ST_PresetLineDashValLgDashDot:
		attr.Value = "lgDashDot"
	case ST_PresetLineDashValLgDashDotDot:
		attr.Value = "lgDashDotDot"
	case ST_PresetLineDashValSysDash:
		attr.Value = "sysDash"
	case ST_PresetLineDashValSysDot:
		attr.Value = "sysDot"
	case ST_PresetLineDashValSysDashDot:
		attr.Value = "sysDashDot"
	case ST_PresetLineDashValSysDashDotDot:
		attr.Value = "sysDashDotDot"
	}
	return attr, nil
}
func (e *ST_PresetLineDashVal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "solid":
		*e = 1
	case "dot":
		*e = 2
	case "dash":
		*e = 3
	case "lgDash":
		*e = 4
	case "dashDot":
		*e = 5
	case "lgDashDot":
		*e = 6
	case "lgDashDotDot":
		*e = 7
	case "sysDash":
		*e = 8
	case "sysDot":
		*e = 9
	case "sysDashDot":
		*e = 10
	case "sysDashDotDot":
		*e = 11
	}
	return nil
}
func (m ST_PresetLineDashVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PresetLineDashVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "solid":
			*m = 1
		case "dot":
			*m = 2
		case "dash":
			*m = 3
		case "lgDash":
			*m = 4
		case "dashDot":
			*m = 5
		case "lgDashDot":
			*m = 6
		case "lgDashDotDot":
			*m = 7
		case "sysDash":
			*m = 8
		case "sysDot":
			*m = 9
		case "sysDashDot":
			*m = 10
		case "sysDashDotDot":
			*m = 11
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
func (m ST_PresetLineDashVal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "solid"
	case 2:
		return "dot"
	case 3:
		return "dash"
	case 4:
		return "lgDash"
	case 5:
		return "dashDot"
	case 6:
		return "lgDashDot"
	case 7:
		return "lgDashDotDot"
	case 8:
		return "sysDash"
	case 9:
		return "sysDot"
	case 10:
		return "sysDashDot"
	case 11:
		return "sysDashDotDot"
	}
	return ""
}
func (m ST_PresetLineDashVal) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PresetLineDashVal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineCap byte

const (
	ST_LineCapUnset ST_LineCap = 0
	ST_LineCapRnd   ST_LineCap = 1
	ST_LineCapSq    ST_LineCap = 2
	ST_LineCapFlat  ST_LineCap = 3
)

func (e ST_LineCap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineCapUnset:
		attr.Value = ""
	case ST_LineCapRnd:
		attr.Value = "rnd"
	case ST_LineCapSq:
		attr.Value = "sq"
	case ST_LineCapFlat:
		attr.Value = "flat"
	}
	return attr, nil
}
func (e *ST_LineCap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "rnd":
		*e = 1
	case "sq":
		*e = 2
	case "flat":
		*e = 3
	}
	return nil
}
func (m ST_LineCap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_LineCap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "rnd":
			*m = 1
		case "sq":
			*m = 2
		case "flat":
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
func (m ST_LineCap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "rnd"
	case 2:
		return "sq"
	case 3:
		return "flat"
	}
	return ""
}
func (m ST_LineCap) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_LineCap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PenAlignment byte

const (
	ST_PenAlignmentUnset ST_PenAlignment = 0
	ST_PenAlignmentCtr   ST_PenAlignment = 1
	ST_PenAlignmentIn    ST_PenAlignment = 2
)

func (e ST_PenAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PenAlignmentUnset:
		attr.Value = ""
	case ST_PenAlignmentCtr:
		attr.Value = "ctr"
	case ST_PenAlignmentIn:
		attr.Value = "in"
	}
	return attr, nil
}
func (e *ST_PenAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ctr":
		*e = 1
	case "in":
		*e = 2
	}
	return nil
}
func (m ST_PenAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PenAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "in":
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
func (m ST_PenAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ctr"
	case 2:
		return "in"
	}
	return ""
}
func (m ST_PenAlignment) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PenAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CompoundLine byte

const (
	ST_CompoundLineUnset     ST_CompoundLine = 0
	ST_CompoundLineSng       ST_CompoundLine = 1
	ST_CompoundLineDbl       ST_CompoundLine = 2
	ST_CompoundLineThickThin ST_CompoundLine = 3
	ST_CompoundLineThinThick ST_CompoundLine = 4
	ST_CompoundLineTri       ST_CompoundLine = 5
)

func (e ST_CompoundLine) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CompoundLineUnset:
		attr.Value = ""
	case ST_CompoundLineSng:
		attr.Value = "sng"
	case ST_CompoundLineDbl:
		attr.Value = "dbl"
	case ST_CompoundLineThickThin:
		attr.Value = "thickThin"
	case ST_CompoundLineThinThick:
		attr.Value = "thinThick"
	case ST_CompoundLineTri:
		attr.Value = "tri"
	}
	return attr, nil
}
func (e *ST_CompoundLine) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sng":
		*e = 1
	case "dbl":
		*e = 2
	case "thickThin":
		*e = 3
	case "thinThick":
		*e = 4
	case "tri":
		*e = 5
	}
	return nil
}
func (m ST_CompoundLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_CompoundLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sng":
			*m = 1
		case "dbl":
			*m = 2
		case "thickThin":
			*m = 3
		case "thinThick":
			*m = 4
		case "tri":
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
func (m ST_CompoundLine) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sng"
	case 2:
		return "dbl"
	case 3:
		return "thickThin"
	case 4:
		return "thinThick"
	case 5:
		return "tri"
	}
	return ""
}
func (m ST_CompoundLine) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_CompoundLine) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OnOffStyleType byte

const (
	ST_OnOffStyleTypeUnset ST_OnOffStyleType = 0
	ST_OnOffStyleTypeOn    ST_OnOffStyleType = 1
	ST_OnOffStyleTypeOff   ST_OnOffStyleType = 2
	ST_OnOffStyleTypeDef   ST_OnOffStyleType = 3
)

func (e ST_OnOffStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OnOffStyleTypeUnset:
		attr.Value = ""
	case ST_OnOffStyleTypeOn:
		attr.Value = "on"
	case ST_OnOffStyleTypeOff:
		attr.Value = "off"
	case ST_OnOffStyleTypeDef:
		attr.Value = "def"
	}
	return attr, nil
}
func (e *ST_OnOffStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "on":
		*e = 1
	case "off":
		*e = 2
	case "def":
		*e = 3
	}
	return nil
}
func (m ST_OnOffStyleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_OnOffStyleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "on":
			*m = 1
		case "off":
			*m = 2
		case "def":
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
func (m ST_OnOffStyleType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "on"
	case 2:
		return "off"
	case 3:
		return "def"
	}
	return ""
}
func (m ST_OnOffStyleType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_OnOffStyleType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAnchoringType byte

const (
	ST_TextAnchoringTypeUnset ST_TextAnchoringType = 0
	ST_TextAnchoringTypeT     ST_TextAnchoringType = 1
	ST_TextAnchoringTypeCtr   ST_TextAnchoringType = 2
	ST_TextAnchoringTypeB     ST_TextAnchoringType = 3
	ST_TextAnchoringTypeJust  ST_TextAnchoringType = 4
	ST_TextAnchoringTypeDist  ST_TextAnchoringType = 5
)

func (e ST_TextAnchoringType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAnchoringTypeUnset:
		attr.Value = ""
	case ST_TextAnchoringTypeT:
		attr.Value = "t"
	case ST_TextAnchoringTypeCtr:
		attr.Value = "ctr"
	case ST_TextAnchoringTypeB:
		attr.Value = "b"
	case ST_TextAnchoringTypeJust:
		attr.Value = "just"
	case ST_TextAnchoringTypeDist:
		attr.Value = "dist"
	}
	return attr, nil
}
func (e *ST_TextAnchoringType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "ctr":
		*e = 2
	case "b":
		*e = 3
	case "just":
		*e = 4
	case "dist":
		*e = 5
	}
	return nil
}
func (m ST_TextAnchoringType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextAnchoringType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "t":
			*m = 1
		case "ctr":
			*m = 2
		case "b":
			*m = 3
		case "just":
			*m = 4
		case "dist":
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
func (m ST_TextAnchoringType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "ctr"
	case 3:
		return "b"
	case 4:
		return "just"
	case 5:
		return "dist"
	}
	return ""
}
func (m ST_TextAnchoringType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextAnchoringType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextVertOverflowType byte

const (
	ST_TextVertOverflowTypeUnset    ST_TextVertOverflowType = 0
	ST_TextVertOverflowTypeOverflow ST_TextVertOverflowType = 1
	ST_TextVertOverflowTypeEllipsis ST_TextVertOverflowType = 2
	ST_TextVertOverflowTypeClip     ST_TextVertOverflowType = 3
)

func (e ST_TextVertOverflowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextVertOverflowTypeUnset:
		attr.Value = ""
	case ST_TextVertOverflowTypeOverflow:
		attr.Value = "overflow"
	case ST_TextVertOverflowTypeEllipsis:
		attr.Value = "ellipsis"
	case ST_TextVertOverflowTypeClip:
		attr.Value = "clip"
	}
	return attr, nil
}
func (e *ST_TextVertOverflowType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "overflow":
		*e = 1
	case "ellipsis":
		*e = 2
	case "clip":
		*e = 3
	}
	return nil
}
func (m ST_TextVertOverflowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextVertOverflowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "overflow":
			*m = 1
		case "ellipsis":
			*m = 2
		case "clip":
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
func (m ST_TextVertOverflowType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "overflow"
	case 2:
		return "ellipsis"
	case 3:
		return "clip"
	}
	return ""
}
func (m ST_TextVertOverflowType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextVertOverflowType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextHorzOverflowType byte

const (
	ST_TextHorzOverflowTypeUnset    ST_TextHorzOverflowType = 0
	ST_TextHorzOverflowTypeOverflow ST_TextHorzOverflowType = 1
	ST_TextHorzOverflowTypeClip     ST_TextHorzOverflowType = 2
)

func (e ST_TextHorzOverflowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextHorzOverflowTypeUnset:
		attr.Value = ""
	case ST_TextHorzOverflowTypeOverflow:
		attr.Value = "overflow"
	case ST_TextHorzOverflowTypeClip:
		attr.Value = "clip"
	}
	return attr, nil
}
func (e *ST_TextHorzOverflowType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "overflow":
		*e = 1
	case "clip":
		*e = 2
	}
	return nil
}
func (m ST_TextHorzOverflowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextHorzOverflowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "overflow":
			*m = 1
		case "clip":
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
func (m ST_TextHorzOverflowType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "overflow"
	case 2:
		return "clip"
	}
	return ""
}
func (m ST_TextHorzOverflowType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextHorzOverflowType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextVerticalType byte

const (
	ST_TextVerticalTypeUnset          ST_TextVerticalType = 0
	ST_TextVerticalTypeHorz           ST_TextVerticalType = 1
	ST_TextVerticalTypeVert           ST_TextVerticalType = 2
	ST_TextVerticalTypeVert270        ST_TextVerticalType = 3
	ST_TextVerticalTypeWordArtVert    ST_TextVerticalType = 4
	ST_TextVerticalTypeEaVert         ST_TextVerticalType = 5
	ST_TextVerticalTypeMongolianVert  ST_TextVerticalType = 6
	ST_TextVerticalTypeWordArtVertRtl ST_TextVerticalType = 7
)

func (e ST_TextVerticalType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextVerticalTypeUnset:
		attr.Value = ""
	case ST_TextVerticalTypeHorz:
		attr.Value = "horz"
	case ST_TextVerticalTypeVert:
		attr.Value = "vert"
	case ST_TextVerticalTypeVert270:
		attr.Value = "vert270"
	case ST_TextVerticalTypeWordArtVert:
		attr.Value = "wordArtVert"
	case ST_TextVerticalTypeEaVert:
		attr.Value = "eaVert"
	case ST_TextVerticalTypeMongolianVert:
		attr.Value = "mongolianVert"
	case ST_TextVerticalTypeWordArtVertRtl:
		attr.Value = "wordArtVertRtl"
	}
	return attr, nil
}
func (e *ST_TextVerticalType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "horz":
		*e = 1
	case "vert":
		*e = 2
	case "vert270":
		*e = 3
	case "wordArtVert":
		*e = 4
	case "eaVert":
		*e = 5
	case "mongolianVert":
		*e = 6
	case "wordArtVertRtl":
		*e = 7
	}
	return nil
}
func (m ST_TextVerticalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextVerticalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "horz":
			*m = 1
		case "vert":
			*m = 2
		case "vert270":
			*m = 3
		case "wordArtVert":
			*m = 4
		case "eaVert":
			*m = 5
		case "mongolianVert":
			*m = 6
		case "wordArtVertRtl":
			*m = 7
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
func (m ST_TextVerticalType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "horz"
	case 2:
		return "vert"
	case 3:
		return "vert270"
	case 4:
		return "wordArtVert"
	case 5:
		return "eaVert"
	case 6:
		return "mongolianVert"
	case 7:
		return "wordArtVertRtl"
	}
	return ""
}
func (m ST_TextVerticalType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextVerticalType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextWrappingType byte

const (
	ST_TextWrappingTypeUnset  ST_TextWrappingType = 0
	ST_TextWrappingTypeNone   ST_TextWrappingType = 1
	ST_TextWrappingTypeSquare ST_TextWrappingType = 2
)

func (e ST_TextWrappingType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextWrappingTypeUnset:
		attr.Value = ""
	case ST_TextWrappingTypeNone:
		attr.Value = "none"
	case ST_TextWrappingTypeSquare:
		attr.Value = "square"
	}
	return attr, nil
}
func (e *ST_TextWrappingType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "square":
		*e = 2
	}
	return nil
}
func (m ST_TextWrappingType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextWrappingType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "square":
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
func (m ST_TextWrappingType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "square"
	}
	return ""
}
func (m ST_TextWrappingType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextWrappingType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAutonumberScheme byte

const (
	ST_TextAutonumberSchemeUnset                 ST_TextAutonumberScheme = 0
	ST_TextAutonumberSchemeAlphaLcParenBoth      ST_TextAutonumberScheme = 1
	ST_TextAutonumberSchemeAlphaUcParenBoth      ST_TextAutonumberScheme = 2
	ST_TextAutonumberSchemeAlphaLcParenR         ST_TextAutonumberScheme = 3
	ST_TextAutonumberSchemeAlphaUcParenR         ST_TextAutonumberScheme = 4
	ST_TextAutonumberSchemeAlphaLcPeriod         ST_TextAutonumberScheme = 5
	ST_TextAutonumberSchemeAlphaUcPeriod         ST_TextAutonumberScheme = 6
	ST_TextAutonumberSchemeArabicParenBoth       ST_TextAutonumberScheme = 7
	ST_TextAutonumberSchemeArabicParenR          ST_TextAutonumberScheme = 8
	ST_TextAutonumberSchemeArabicPeriod          ST_TextAutonumberScheme = 9
	ST_TextAutonumberSchemeArabicPlain           ST_TextAutonumberScheme = 10
	ST_TextAutonumberSchemeRomanLcParenBoth      ST_TextAutonumberScheme = 11
	ST_TextAutonumberSchemeRomanUcParenBoth      ST_TextAutonumberScheme = 12
	ST_TextAutonumberSchemeRomanLcParenR         ST_TextAutonumberScheme = 13
	ST_TextAutonumberSchemeRomanUcParenR         ST_TextAutonumberScheme = 14
	ST_TextAutonumberSchemeRomanLcPeriod         ST_TextAutonumberScheme = 15
	ST_TextAutonumberSchemeRomanUcPeriod         ST_TextAutonumberScheme = 16
	ST_TextAutonumberSchemeCircleNumDbPlain      ST_TextAutonumberScheme = 17
	ST_TextAutonumberSchemeCircleNumWdBlackPlain ST_TextAutonumberScheme = 18
	ST_TextAutonumberSchemeCircleNumWdWhitePlain ST_TextAutonumberScheme = 19
	ST_TextAutonumberSchemeArabicDbPeriod        ST_TextAutonumberScheme = 20
	ST_TextAutonumberSchemeArabicDbPlain         ST_TextAutonumberScheme = 21
	ST_TextAutonumberSchemeEa1ChsPeriod          ST_TextAutonumberScheme = 22
	ST_TextAutonumberSchemeEa1ChsPlain           ST_TextAutonumberScheme = 23
	ST_TextAutonumberSchemeEa1ChtPeriod          ST_TextAutonumberScheme = 24
	ST_TextAutonumberSchemeEa1ChtPlain           ST_TextAutonumberScheme = 25
	ST_TextAutonumberSchemeEa1JpnChsDbPeriod     ST_TextAutonumberScheme = 26
	ST_TextAutonumberSchemeEa1JpnKorPlain        ST_TextAutonumberScheme = 27
	ST_TextAutonumberSchemeEa1JpnKorPeriod       ST_TextAutonumberScheme = 28
	ST_TextAutonumberSchemeArabic1Minus          ST_TextAutonumberScheme = 29
	ST_TextAutonumberSchemeArabic2Minus          ST_TextAutonumberScheme = 30
	ST_TextAutonumberSchemeHebrew2Minus          ST_TextAutonumberScheme = 31
	ST_TextAutonumberSchemeThaiAlphaPeriod       ST_TextAutonumberScheme = 32
	ST_TextAutonumberSchemeThaiAlphaParenR       ST_TextAutonumberScheme = 33
	ST_TextAutonumberSchemeThaiAlphaParenBoth    ST_TextAutonumberScheme = 34
	ST_TextAutonumberSchemeThaiNumPeriod         ST_TextAutonumberScheme = 35
	ST_TextAutonumberSchemeThaiNumParenR         ST_TextAutonumberScheme = 36
	ST_TextAutonumberSchemeThaiNumParenBoth      ST_TextAutonumberScheme = 37
	ST_TextAutonumberSchemeHindiAlphaPeriod      ST_TextAutonumberScheme = 38
	ST_TextAutonumberSchemeHindiNumPeriod        ST_TextAutonumberScheme = 39
	ST_TextAutonumberSchemeHindiNumParenR        ST_TextAutonumberScheme = 40
	ST_TextAutonumberSchemeHindiAlpha1Period     ST_TextAutonumberScheme = 41
)

func (e ST_TextAutonumberScheme) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAutonumberSchemeUnset:
		attr.Value = ""
	case ST_TextAutonumberSchemeAlphaLcParenBoth:
		attr.Value = "alphaLcParenBoth"
	case ST_TextAutonumberSchemeAlphaUcParenBoth:
		attr.Value = "alphaUcParenBoth"
	case ST_TextAutonumberSchemeAlphaLcParenR:
		attr.Value = "alphaLcParenR"
	case ST_TextAutonumberSchemeAlphaUcParenR:
		attr.Value = "alphaUcParenR"
	case ST_TextAutonumberSchemeAlphaLcPeriod:
		attr.Value = "alphaLcPeriod"
	case ST_TextAutonumberSchemeAlphaUcPeriod:
		attr.Value = "alphaUcPeriod"
	case ST_TextAutonumberSchemeArabicParenBoth:
		attr.Value = "arabicParenBoth"
	case ST_TextAutonumberSchemeArabicParenR:
		attr.Value = "arabicParenR"
	case ST_TextAutonumberSchemeArabicPeriod:
		attr.Value = "arabicPeriod"
	case ST_TextAutonumberSchemeArabicPlain:
		attr.Value = "arabicPlain"
	case ST_TextAutonumberSchemeRomanLcParenBoth:
		attr.Value = "romanLcParenBoth"
	case ST_TextAutonumberSchemeRomanUcParenBoth:
		attr.Value = "romanUcParenBoth"
	case ST_TextAutonumberSchemeRomanLcParenR:
		attr.Value = "romanLcParenR"
	case ST_TextAutonumberSchemeRomanUcParenR:
		attr.Value = "romanUcParenR"
	case ST_TextAutonumberSchemeRomanLcPeriod:
		attr.Value = "romanLcPeriod"
	case ST_TextAutonumberSchemeRomanUcPeriod:
		attr.Value = "romanUcPeriod"
	case ST_TextAutonumberSchemeCircleNumDbPlain:
		attr.Value = "circleNumDbPlain"
	case ST_TextAutonumberSchemeCircleNumWdBlackPlain:
		attr.Value = "circleNumWdBlackPlain"
	case ST_TextAutonumberSchemeCircleNumWdWhitePlain:
		attr.Value = "circleNumWdWhitePlain"
	case ST_TextAutonumberSchemeArabicDbPeriod:
		attr.Value = "arabicDbPeriod"
	case ST_TextAutonumberSchemeArabicDbPlain:
		attr.Value = "arabicDbPlain"
	case ST_TextAutonumberSchemeEa1ChsPeriod:
		attr.Value = "ea1ChsPeriod"
	case ST_TextAutonumberSchemeEa1ChsPlain:
		attr.Value = "ea1ChsPlain"
	case ST_TextAutonumberSchemeEa1ChtPeriod:
		attr.Value = "ea1ChtPeriod"
	case ST_TextAutonumberSchemeEa1ChtPlain:
		attr.Value = "ea1ChtPlain"
	case ST_TextAutonumberSchemeEa1JpnChsDbPeriod:
		attr.Value = "ea1JpnChsDbPeriod"
	case ST_TextAutonumberSchemeEa1JpnKorPlain:
		attr.Value = "ea1JpnKorPlain"
	case ST_TextAutonumberSchemeEa1JpnKorPeriod:
		attr.Value = "ea1JpnKorPeriod"
	case ST_TextAutonumberSchemeArabic1Minus:
		attr.Value = "arabic1Minus"
	case ST_TextAutonumberSchemeArabic2Minus:
		attr.Value = "arabic2Minus"
	case ST_TextAutonumberSchemeHebrew2Minus:
		attr.Value = "hebrew2Minus"
	case ST_TextAutonumberSchemeThaiAlphaPeriod:
		attr.Value = "thaiAlphaPeriod"
	case ST_TextAutonumberSchemeThaiAlphaParenR:
		attr.Value = "thaiAlphaParenR"
	case ST_TextAutonumberSchemeThaiAlphaParenBoth:
		attr.Value = "thaiAlphaParenBoth"
	case ST_TextAutonumberSchemeThaiNumPeriod:
		attr.Value = "thaiNumPeriod"
	case ST_TextAutonumberSchemeThaiNumParenR:
		attr.Value = "thaiNumParenR"
	case ST_TextAutonumberSchemeThaiNumParenBoth:
		attr.Value = "thaiNumParenBoth"
	case ST_TextAutonumberSchemeHindiAlphaPeriod:
		attr.Value = "hindiAlphaPeriod"
	case ST_TextAutonumberSchemeHindiNumPeriod:
		attr.Value = "hindiNumPeriod"
	case ST_TextAutonumberSchemeHindiNumParenR:
		attr.Value = "hindiNumParenR"
	case ST_TextAutonumberSchemeHindiAlpha1Period:
		attr.Value = "hindiAlpha1Period"
	}
	return attr, nil
}
func (e *ST_TextAutonumberScheme) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "alphaLcParenBoth":
		*e = 1
	case "alphaUcParenBoth":
		*e = 2
	case "alphaLcParenR":
		*e = 3
	case "alphaUcParenR":
		*e = 4
	case "alphaLcPeriod":
		*e = 5
	case "alphaUcPeriod":
		*e = 6
	case "arabicParenBoth":
		*e = 7
	case "arabicParenR":
		*e = 8
	case "arabicPeriod":
		*e = 9
	case "arabicPlain":
		*e = 10
	case "romanLcParenBoth":
		*e = 11
	case "romanUcParenBoth":
		*e = 12
	case "romanLcParenR":
		*e = 13
	case "romanUcParenR":
		*e = 14
	case "romanLcPeriod":
		*e = 15
	case "romanUcPeriod":
		*e = 16
	case "circleNumDbPlain":
		*e = 17
	case "circleNumWdBlackPlain":
		*e = 18
	case "circleNumWdWhitePlain":
		*e = 19
	case "arabicDbPeriod":
		*e = 20
	case "arabicDbPlain":
		*e = 21
	case "ea1ChsPeriod":
		*e = 22
	case "ea1ChsPlain":
		*e = 23
	case "ea1ChtPeriod":
		*e = 24
	case "ea1ChtPlain":
		*e = 25
	case "ea1JpnChsDbPeriod":
		*e = 26
	case "ea1JpnKorPlain":
		*e = 27
	case "ea1JpnKorPeriod":
		*e = 28
	case "arabic1Minus":
		*e = 29
	case "arabic2Minus":
		*e = 30
	case "hebrew2Minus":
		*e = 31
	case "thaiAlphaPeriod":
		*e = 32
	case "thaiAlphaParenR":
		*e = 33
	case "thaiAlphaParenBoth":
		*e = 34
	case "thaiNumPeriod":
		*e = 35
	case "thaiNumParenR":
		*e = 36
	case "thaiNumParenBoth":
		*e = 37
	case "hindiAlphaPeriod":
		*e = 38
	case "hindiNumPeriod":
		*e = 39
	case "hindiNumParenR":
		*e = 40
	case "hindiAlpha1Period":
		*e = 41
	}
	return nil
}
func (m ST_TextAutonumberScheme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextAutonumberScheme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "alphaLcParenBoth":
			*m = 1
		case "alphaUcParenBoth":
			*m = 2
		case "alphaLcParenR":
			*m = 3
		case "alphaUcParenR":
			*m = 4
		case "alphaLcPeriod":
			*m = 5
		case "alphaUcPeriod":
			*m = 6
		case "arabicParenBoth":
			*m = 7
		case "arabicParenR":
			*m = 8
		case "arabicPeriod":
			*m = 9
		case "arabicPlain":
			*m = 10
		case "romanLcParenBoth":
			*m = 11
		case "romanUcParenBoth":
			*m = 12
		case "romanLcParenR":
			*m = 13
		case "romanUcParenR":
			*m = 14
		case "romanLcPeriod":
			*m = 15
		case "romanUcPeriod":
			*m = 16
		case "circleNumDbPlain":
			*m = 17
		case "circleNumWdBlackPlain":
			*m = 18
		case "circleNumWdWhitePlain":
			*m = 19
		case "arabicDbPeriod":
			*m = 20
		case "arabicDbPlain":
			*m = 21
		case "ea1ChsPeriod":
			*m = 22
		case "ea1ChsPlain":
			*m = 23
		case "ea1ChtPeriod":
			*m = 24
		case "ea1ChtPlain":
			*m = 25
		case "ea1JpnChsDbPeriod":
			*m = 26
		case "ea1JpnKorPlain":
			*m = 27
		case "ea1JpnKorPeriod":
			*m = 28
		case "arabic1Minus":
			*m = 29
		case "arabic2Minus":
			*m = 30
		case "hebrew2Minus":
			*m = 31
		case "thaiAlphaPeriod":
			*m = 32
		case "thaiAlphaParenR":
			*m = 33
		case "thaiAlphaParenBoth":
			*m = 34
		case "thaiNumPeriod":
			*m = 35
		case "thaiNumParenR":
			*m = 36
		case "thaiNumParenBoth":
			*m = 37
		case "hindiAlphaPeriod":
			*m = 38
		case "hindiNumPeriod":
			*m = 39
		case "hindiNumParenR":
			*m = 40
		case "hindiAlpha1Period":
			*m = 41
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
func (m ST_TextAutonumberScheme) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "alphaLcParenBoth"
	case 2:
		return "alphaUcParenBoth"
	case 3:
		return "alphaLcParenR"
	case 4:
		return "alphaUcParenR"
	case 5:
		return "alphaLcPeriod"
	case 6:
		return "alphaUcPeriod"
	case 7:
		return "arabicParenBoth"
	case 8:
		return "arabicParenR"
	case 9:
		return "arabicPeriod"
	case 10:
		return "arabicPlain"
	case 11:
		return "romanLcParenBoth"
	case 12:
		return "romanUcParenBoth"
	case 13:
		return "romanLcParenR"
	case 14:
		return "romanUcParenR"
	case 15:
		return "romanLcPeriod"
	case 16:
		return "romanUcPeriod"
	case 17:
		return "circleNumDbPlain"
	case 18:
		return "circleNumWdBlackPlain"
	case 19:
		return "circleNumWdWhitePlain"
	case 20:
		return "arabicDbPeriod"
	case 21:
		return "arabicDbPlain"
	case 22:
		return "ea1ChsPeriod"
	case 23:
		return "ea1ChsPlain"
	case 24:
		return "ea1ChtPeriod"
	case 25:
		return "ea1ChtPlain"
	case 26:
		return "ea1JpnChsDbPeriod"
	case 27:
		return "ea1JpnKorPlain"
	case 28:
		return "ea1JpnKorPeriod"
	case 29:
		return "arabic1Minus"
	case 30:
		return "arabic2Minus"
	case 31:
		return "hebrew2Minus"
	case 32:
		return "thaiAlphaPeriod"
	case 33:
		return "thaiAlphaParenR"
	case 34:
		return "thaiAlphaParenBoth"
	case 35:
		return "thaiNumPeriod"
	case 36:
		return "thaiNumParenR"
	case 37:
		return "thaiNumParenBoth"
	case 38:
		return "hindiAlphaPeriod"
	case 39:
		return "hindiNumPeriod"
	case 40:
		return "hindiNumParenR"
	case 41:
		return "hindiAlpha1Period"
	}
	return ""
}
func (m ST_TextAutonumberScheme) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextAutonumberScheme) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PitchFamily byte

const (
	ST_PitchFamilyUnset ST_PitchFamily = 0
	ST_PitchFamily00    ST_PitchFamily = 1
	ST_PitchFamily01    ST_PitchFamily = 2
	ST_PitchFamily02    ST_PitchFamily = 3
	ST_PitchFamily16    ST_PitchFamily = 4
	ST_PitchFamily17    ST_PitchFamily = 5
	ST_PitchFamily18    ST_PitchFamily = 6
	ST_PitchFamily32    ST_PitchFamily = 7
	ST_PitchFamily33    ST_PitchFamily = 8
	ST_PitchFamily34    ST_PitchFamily = 9
	ST_PitchFamily48    ST_PitchFamily = 10
	ST_PitchFamily49    ST_PitchFamily = 11
	ST_PitchFamily50    ST_PitchFamily = 12
	ST_PitchFamily64    ST_PitchFamily = 13
	ST_PitchFamily65    ST_PitchFamily = 14
	ST_PitchFamily66    ST_PitchFamily = 15
	ST_PitchFamily80    ST_PitchFamily = 16
	ST_PitchFamily81    ST_PitchFamily = 17
	ST_PitchFamily82    ST_PitchFamily = 18
)

func (e ST_PitchFamily) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PitchFamilyUnset:
		attr.Value = ""
	case ST_PitchFamily00:
		attr.Value = "00"
	case ST_PitchFamily01:
		attr.Value = "01"
	case ST_PitchFamily02:
		attr.Value = "02"
	case ST_PitchFamily16:
		attr.Value = "16"
	case ST_PitchFamily17:
		attr.Value = "17"
	case ST_PitchFamily18:
		attr.Value = "18"
	case ST_PitchFamily32:
		attr.Value = "32"
	case ST_PitchFamily33:
		attr.Value = "33"
	case ST_PitchFamily34:
		attr.Value = "34"
	case ST_PitchFamily48:
		attr.Value = "48"
	case ST_PitchFamily49:
		attr.Value = "49"
	case ST_PitchFamily50:
		attr.Value = "50"
	case ST_PitchFamily64:
		attr.Value = "64"
	case ST_PitchFamily65:
		attr.Value = "65"
	case ST_PitchFamily66:
		attr.Value = "66"
	case ST_PitchFamily80:
		attr.Value = "80"
	case ST_PitchFamily81:
		attr.Value = "81"
	case ST_PitchFamily82:
		attr.Value = "82"
	}
	return attr, nil
}
func (e *ST_PitchFamily) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "00":
		*e = 1
	case "01":
		*e = 2
	case "02":
		*e = 3
	case "16":
		*e = 4
	case "17":
		*e = 5
	case "18":
		*e = 6
	case "32":
		*e = 7
	case "33":
		*e = 8
	case "34":
		*e = 9
	case "48":
		*e = 10
	case "49":
		*e = 11
	case "50":
		*e = 12
	case "64":
		*e = 13
	case "65":
		*e = 14
	case "66":
		*e = 15
	case "80":
		*e = 16
	case "81":
		*e = 17
	case "82":
		*e = 18
	}
	return nil
}
func (m ST_PitchFamily) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PitchFamily) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "00":
			*m = 1
		case "01":
			*m = 2
		case "02":
			*m = 3
		case "16":
			*m = 4
		case "17":
			*m = 5
		case "18":
			*m = 6
		case "32":
			*m = 7
		case "33":
			*m = 8
		case "34":
			*m = 9
		case "48":
			*m = 10
		case "49":
			*m = 11
		case "50":
			*m = 12
		case "64":
			*m = 13
		case "65":
			*m = 14
		case "66":
			*m = 15
		case "80":
			*m = 16
		case "81":
			*m = 17
		case "82":
			*m = 18
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
func (m ST_PitchFamily) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "00"
	case 2:
		return "01"
	case 3:
		return "02"
	case 4:
		return "16"
	case 5:
		return "17"
	case 6:
		return "18"
	case 7:
		return "32"
	case 8:
		return "33"
	case 9:
		return "34"
	case 10:
		return "48"
	case 11:
		return "49"
	case 12:
		return "50"
	case 13:
		return "64"
	case 14:
		return "65"
	case 15:
		return "66"
	case 16:
		return "80"
	case 17:
		return "81"
	case 18:
		return "82"
	}
	return ""
}
func (m ST_PitchFamily) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PitchFamily) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextUnderlineType byte

const (
	ST_TextUnderlineTypeUnset           ST_TextUnderlineType = 0
	ST_TextUnderlineTypeNone            ST_TextUnderlineType = 1
	ST_TextUnderlineTypeWords           ST_TextUnderlineType = 2
	ST_TextUnderlineTypeSng             ST_TextUnderlineType = 3
	ST_TextUnderlineTypeDbl             ST_TextUnderlineType = 4
	ST_TextUnderlineTypeHeavy           ST_TextUnderlineType = 5
	ST_TextUnderlineTypeDotted          ST_TextUnderlineType = 6
	ST_TextUnderlineTypeDottedHeavy     ST_TextUnderlineType = 7
	ST_TextUnderlineTypeDash            ST_TextUnderlineType = 8
	ST_TextUnderlineTypeDashHeavy       ST_TextUnderlineType = 9
	ST_TextUnderlineTypeDashLong        ST_TextUnderlineType = 10
	ST_TextUnderlineTypeDashLongHeavy   ST_TextUnderlineType = 11
	ST_TextUnderlineTypeDotDash         ST_TextUnderlineType = 12
	ST_TextUnderlineTypeDotDashHeavy    ST_TextUnderlineType = 13
	ST_TextUnderlineTypeDotDotDash      ST_TextUnderlineType = 14
	ST_TextUnderlineTypeDotDotDashHeavy ST_TextUnderlineType = 15
	ST_TextUnderlineTypeWavy            ST_TextUnderlineType = 16
	ST_TextUnderlineTypeWavyHeavy       ST_TextUnderlineType = 17
	ST_TextUnderlineTypeWavyDbl         ST_TextUnderlineType = 18
)

func (e ST_TextUnderlineType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextUnderlineTypeUnset:
		attr.Value = ""
	case ST_TextUnderlineTypeNone:
		attr.Value = "none"
	case ST_TextUnderlineTypeWords:
		attr.Value = "words"
	case ST_TextUnderlineTypeSng:
		attr.Value = "sng"
	case ST_TextUnderlineTypeDbl:
		attr.Value = "dbl"
	case ST_TextUnderlineTypeHeavy:
		attr.Value = "heavy"
	case ST_TextUnderlineTypeDotted:
		attr.Value = "dotted"
	case ST_TextUnderlineTypeDottedHeavy:
		attr.Value = "dottedHeavy"
	case ST_TextUnderlineTypeDash:
		attr.Value = "dash"
	case ST_TextUnderlineTypeDashHeavy:
		attr.Value = "dashHeavy"
	case ST_TextUnderlineTypeDashLong:
		attr.Value = "dashLong"
	case ST_TextUnderlineTypeDashLongHeavy:
		attr.Value = "dashLongHeavy"
	case ST_TextUnderlineTypeDotDash:
		attr.Value = "dotDash"
	case ST_TextUnderlineTypeDotDashHeavy:
		attr.Value = "dotDashHeavy"
	case ST_TextUnderlineTypeDotDotDash:
		attr.Value = "dotDotDash"
	case ST_TextUnderlineTypeDotDotDashHeavy:
		attr.Value = "dotDotDashHeavy"
	case ST_TextUnderlineTypeWavy:
		attr.Value = "wavy"
	case ST_TextUnderlineTypeWavyHeavy:
		attr.Value = "wavyHeavy"
	case ST_TextUnderlineTypeWavyDbl:
		attr.Value = "wavyDbl"
	}
	return attr, nil
}
func (e *ST_TextUnderlineType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "words":
		*e = 2
	case "sng":
		*e = 3
	case "dbl":
		*e = 4
	case "heavy":
		*e = 5
	case "dotted":
		*e = 6
	case "dottedHeavy":
		*e = 7
	case "dash":
		*e = 8
	case "dashHeavy":
		*e = 9
	case "dashLong":
		*e = 10
	case "dashLongHeavy":
		*e = 11
	case "dotDash":
		*e = 12
	case "dotDashHeavy":
		*e = 13
	case "dotDotDash":
		*e = 14
	case "dotDotDashHeavy":
		*e = 15
	case "wavy":
		*e = 16
	case "wavyHeavy":
		*e = 17
	case "wavyDbl":
		*e = 18
	}
	return nil
}
func (m ST_TextUnderlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextUnderlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "words":
			*m = 2
		case "sng":
			*m = 3
		case "dbl":
			*m = 4
		case "heavy":
			*m = 5
		case "dotted":
			*m = 6
		case "dottedHeavy":
			*m = 7
		case "dash":
			*m = 8
		case "dashHeavy":
			*m = 9
		case "dashLong":
			*m = 10
		case "dashLongHeavy":
			*m = 11
		case "dotDash":
			*m = 12
		case "dotDashHeavy":
			*m = 13
		case "dotDotDash":
			*m = 14
		case "dotDotDashHeavy":
			*m = 15
		case "wavy":
			*m = 16
		case "wavyHeavy":
			*m = 17
		case "wavyDbl":
			*m = 18
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
func (m ST_TextUnderlineType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "words"
	case 3:
		return "sng"
	case 4:
		return "dbl"
	case 5:
		return "heavy"
	case 6:
		return "dotted"
	case 7:
		return "dottedHeavy"
	case 8:
		return "dash"
	case 9:
		return "dashHeavy"
	case 10:
		return "dashLong"
	case 11:
		return "dashLongHeavy"
	case 12:
		return "dotDash"
	case 13:
		return "dotDashHeavy"
	case 14:
		return "dotDotDash"
	case 15:
		return "dotDotDashHeavy"
	case 16:
		return "wavy"
	case 17:
		return "wavyHeavy"
	case 18:
		return "wavyDbl"
	}
	return ""
}
func (m ST_TextUnderlineType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextUnderlineType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextStrikeType byte

const (
	ST_TextStrikeTypeUnset     ST_TextStrikeType = 0
	ST_TextStrikeTypeNoStrike  ST_TextStrikeType = 1
	ST_TextStrikeTypeSngStrike ST_TextStrikeType = 2
	ST_TextStrikeTypeDblStrike ST_TextStrikeType = 3
)

func (e ST_TextStrikeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextStrikeTypeUnset:
		attr.Value = ""
	case ST_TextStrikeTypeNoStrike:
		attr.Value = "noStrike"
	case ST_TextStrikeTypeSngStrike:
		attr.Value = "sngStrike"
	case ST_TextStrikeTypeDblStrike:
		attr.Value = "dblStrike"
	}
	return attr, nil
}
func (e *ST_TextStrikeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "noStrike":
		*e = 1
	case "sngStrike":
		*e = 2
	case "dblStrike":
		*e = 3
	}
	return nil
}
func (m ST_TextStrikeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextStrikeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "noStrike":
			*m = 1
		case "sngStrike":
			*m = 2
		case "dblStrike":
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
func (m ST_TextStrikeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "noStrike"
	case 2:
		return "sngStrike"
	case 3:
		return "dblStrike"
	}
	return ""
}
func (m ST_TextStrikeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextStrikeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextCapsType byte

const (
	ST_TextCapsTypeUnset ST_TextCapsType = 0
	ST_TextCapsTypeNone  ST_TextCapsType = 1
	ST_TextCapsTypeSmall ST_TextCapsType = 2
	ST_TextCapsTypeAll   ST_TextCapsType = 3
)

func (e ST_TextCapsType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextCapsTypeUnset:
		attr.Value = ""
	case ST_TextCapsTypeNone:
		attr.Value = "none"
	case ST_TextCapsTypeSmall:
		attr.Value = "small"
	case ST_TextCapsTypeAll:
		attr.Value = "all"
	}
	return attr, nil
}
func (e *ST_TextCapsType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "small":
		*e = 2
	case "all":
		*e = 3
	}
	return nil
}
func (m ST_TextCapsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextCapsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "small":
			*m = 2
		case "all":
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
func (m ST_TextCapsType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "small"
	case 3:
		return "all"
	}
	return ""
}
func (m ST_TextCapsType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextCapsType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextTabAlignType byte

const (
	ST_TextTabAlignTypeUnset ST_TextTabAlignType = 0
	ST_TextTabAlignTypeL     ST_TextTabAlignType = 1
	ST_TextTabAlignTypeCtr   ST_TextTabAlignType = 2
	ST_TextTabAlignTypeR     ST_TextTabAlignType = 3
	ST_TextTabAlignTypeDec   ST_TextTabAlignType = 4
)

func (e ST_TextTabAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextTabAlignTypeUnset:
		attr.Value = ""
	case ST_TextTabAlignTypeL:
		attr.Value = "l"
	case ST_TextTabAlignTypeCtr:
		attr.Value = "ctr"
	case ST_TextTabAlignTypeR:
		attr.Value = "r"
	case ST_TextTabAlignTypeDec:
		attr.Value = "dec"
	}
	return attr, nil
}
func (e *ST_TextTabAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "ctr":
		*e = 2
	case "r":
		*e = 3
	case "dec":
		*e = 4
	}
	return nil
}
func (m ST_TextTabAlignType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextTabAlignType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "l":
			*m = 1
		case "ctr":
			*m = 2
		case "r":
			*m = 3
		case "dec":
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
func (m ST_TextTabAlignType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "ctr"
	case 3:
		return "r"
	case 4:
		return "dec"
	}
	return ""
}
func (m ST_TextTabAlignType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextTabAlignType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAlignType byte

const (
	ST_TextAlignTypeUnset    ST_TextAlignType = 0
	ST_TextAlignTypeL        ST_TextAlignType = 1
	ST_TextAlignTypeCtr      ST_TextAlignType = 2
	ST_TextAlignTypeR        ST_TextAlignType = 3
	ST_TextAlignTypeJust     ST_TextAlignType = 4
	ST_TextAlignTypeJustLow  ST_TextAlignType = 5
	ST_TextAlignTypeDist     ST_TextAlignType = 6
	ST_TextAlignTypeThaiDist ST_TextAlignType = 7
)

func (e ST_TextAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAlignTypeUnset:
		attr.Value = ""
	case ST_TextAlignTypeL:
		attr.Value = "l"
	case ST_TextAlignTypeCtr:
		attr.Value = "ctr"
	case ST_TextAlignTypeR:
		attr.Value = "r"
	case ST_TextAlignTypeJust:
		attr.Value = "just"
	case ST_TextAlignTypeJustLow:
		attr.Value = "justLow"
	case ST_TextAlignTypeDist:
		attr.Value = "dist"
	case ST_TextAlignTypeThaiDist:
		attr.Value = "thaiDist"
	}
	return attr, nil
}
func (e *ST_TextAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "ctr":
		*e = 2
	case "r":
		*e = 3
	case "just":
		*e = 4
	case "justLow":
		*e = 5
	case "dist":
		*e = 6
	case "thaiDist":
		*e = 7
	}
	return nil
}
func (m ST_TextAlignType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextAlignType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "l":
			*m = 1
		case "ctr":
			*m = 2
		case "r":
			*m = 3
		case "just":
			*m = 4
		case "justLow":
			*m = 5
		case "dist":
			*m = 6
		case "thaiDist":
			*m = 7
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
func (m ST_TextAlignType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "ctr"
	case 3:
		return "r"
	case 4:
		return "just"
	case 5:
		return "justLow"
	case 6:
		return "dist"
	case 7:
		return "thaiDist"
	}
	return ""
}
func (m ST_TextAlignType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextAlignType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextFontAlignType byte

const (
	ST_TextFontAlignTypeUnset ST_TextFontAlignType = 0
	ST_TextFontAlignTypeAuto  ST_TextFontAlignType = 1
	ST_TextFontAlignTypeT     ST_TextFontAlignType = 2
	ST_TextFontAlignTypeCtr   ST_TextFontAlignType = 3
	ST_TextFontAlignTypeBase  ST_TextFontAlignType = 4
	ST_TextFontAlignTypeB     ST_TextFontAlignType = 5
)

func (e ST_TextFontAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextFontAlignTypeUnset:
		attr.Value = ""
	case ST_TextFontAlignTypeAuto:
		attr.Value = "auto"
	case ST_TextFontAlignTypeT:
		attr.Value = "t"
	case ST_TextFontAlignTypeCtr:
		attr.Value = "ctr"
	case ST_TextFontAlignTypeBase:
		attr.Value = "base"
	case ST_TextFontAlignTypeB:
		attr.Value = "b"
	}
	return attr, nil
}
func (e *ST_TextFontAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "t":
		*e = 2
	case "ctr":
		*e = 3
	case "base":
		*e = 4
	case "b":
		*e = 5
	}
	return nil
}
func (m ST_TextFontAlignType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TextFontAlignType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "t":
			*m = 2
		case "ctr":
			*m = 3
		case "base":
			*m = 4
		case "b":
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
func (m ST_TextFontAlignType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "t"
	case 3:
		return "ctr"
	case 4:
		return "base"
	case 5:
		return "b"
	}
	return ""
}
func (m ST_TextFontAlignType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TextFontAlignType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AudioFile", NewCT_AudioFile)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_VideoFile", NewCT_VideoFile)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_QuickTimeFile", NewCT_QuickTimeFile)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AudioCDTime", NewCT_AudioCDTime)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AudioCD", NewCT_AudioCD)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorScheme", NewCT_ColorScheme)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_CustomColor", NewCT_CustomColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SupplementalFont", NewCT_SupplementalFont)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_CustomColorList", NewCT_CustomColorList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FontCollection", NewCT_FontCollection)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectStyleItem", NewCT_EffectStyleItem)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FontScheme", NewCT_FontScheme)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FillStyleList", NewCT_FillStyleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineStyleList", NewCT_LineStyleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectStyleList", NewCT_EffectStyleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BackgroundFillStyleList", NewCT_BackgroundFillStyleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_StyleMatrix", NewCT_StyleMatrix)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BaseStyles", NewCT_BaseStyles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_OfficeArtExtension", NewCT_OfficeArtExtension)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Angle", NewCT_Angle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PositiveFixedAngle", NewCT_PositiveFixedAngle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Percentage", NewCT_Percentage)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PositivePercentage", NewCT_PositivePercentage)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FixedPercentage", NewCT_FixedPercentage)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PositiveFixedPercentage", NewCT_PositiveFixedPercentage)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Ratio", NewCT_Ratio)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Point2D", NewCT_Point2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PositiveSize2D", NewCT_PositiveSize2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ComplementTransform", NewCT_ComplementTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_InverseTransform", NewCT_InverseTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GrayscaleTransform", NewCT_GrayscaleTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GammaTransform", NewCT_GammaTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_InverseGammaTransform", NewCT_InverseGammaTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ScRgbColor", NewCT_ScRgbColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SRgbColor", NewCT_SRgbColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_HslColor", NewCT_HslColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SystemColor", NewCT_SystemColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SchemeColor", NewCT_SchemeColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PresetColor", NewCT_PresetColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_OfficeArtExtensionList", NewCT_OfficeArtExtensionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Scale2D", NewCT_Scale2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Transform2D", NewCT_Transform2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GroupTransform2D", NewCT_GroupTransform2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Point3D", NewCT_Point3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Vector3D", NewCT_Vector3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SphereCoords", NewCT_SphereCoords)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_RelativeRect", NewCT_RelativeRect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Color", NewCT_Color)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorMRU", NewCT_ColorMRU)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EmbeddedWAVAudioFile", NewCT_EmbeddedWAVAudioFile)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Hyperlink", NewCT_Hyperlink)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ConnectorLocking", NewCT_ConnectorLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ShapeLocking", NewCT_ShapeLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PictureLocking", NewCT_PictureLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GroupLocking", NewCT_GroupLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GraphicalObjectFrameLocking", NewCT_GraphicalObjectFrameLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ContentPartLocking", NewCT_ContentPartLocking)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualDrawingProps", NewCT_NonVisualDrawingProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualDrawingShapeProps", NewCT_NonVisualDrawingShapeProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualConnectorProperties", NewCT_NonVisualConnectorProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualPictureProperties", NewCT_NonVisualPictureProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualGroupDrawingShapeProps", NewCT_NonVisualGroupDrawingShapeProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualGraphicFrameProperties", NewCT_NonVisualGraphicFrameProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NonVisualContentPartProperties", NewCT_NonVisualContentPartProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GraphicalObjectData", NewCT_GraphicalObjectData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GraphicalObject", NewCT_GraphicalObject)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationDgmElement", NewCT_AnimationDgmElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationChartElement", NewCT_AnimationChartElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationElementChoice", NewCT_AnimationElementChoice)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationDgmBuildProperties", NewCT_AnimationDgmBuildProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationChartBuildProperties", NewCT_AnimationChartBuildProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AnimationGraphicalObjectBuildProperties", NewCT_AnimationGraphicalObjectBuildProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BackgroundFormatting", NewCT_BackgroundFormatting)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_WholeE2oFormatting", NewCT_WholeE2oFormatting)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlUseShapeRectangle", NewCT_GvmlUseShapeRectangle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlTextShape", NewCT_GvmlTextShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlShapeNonVisual", NewCT_GvmlShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlShape", NewCT_GvmlShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlConnectorNonVisual", NewCT_GvmlConnectorNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlConnector", NewCT_GvmlConnector)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlPictureNonVisual", NewCT_GvmlPictureNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlPicture", NewCT_GvmlPicture)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlGraphicFrameNonVisual", NewCT_GvmlGraphicFrameNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlGraphicalObjectFrame", NewCT_GvmlGraphicalObjectFrame)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlGroupShapeNonVisual", NewCT_GvmlGroupShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GvmlGroupShape", NewCT_GvmlGroupShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Camera", NewCT_Camera)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LightRig", NewCT_LightRig)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Scene3D", NewCT_Scene3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Backdrop", NewCT_Backdrop)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Bevel", NewCT_Bevel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Shape3D", NewCT_Shape3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FlatText", NewCT_FlatText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaBiLevelEffect", NewCT_AlphaBiLevelEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaCeilingEffect", NewCT_AlphaCeilingEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaFloorEffect", NewCT_AlphaFloorEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaInverseEffect", NewCT_AlphaInverseEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaModulateFixedEffect", NewCT_AlphaModulateFixedEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaOutsetEffect", NewCT_AlphaOutsetEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaReplaceEffect", NewCT_AlphaReplaceEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BiLevelEffect", NewCT_BiLevelEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BlurEffect", NewCT_BlurEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorChangeEffect", NewCT_ColorChangeEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorReplaceEffect", NewCT_ColorReplaceEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_DuotoneEffect", NewCT_DuotoneEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GlowEffect", NewCT_GlowEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GrayscaleEffect", NewCT_GrayscaleEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_HSLEffect", NewCT_HSLEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_InnerShadowEffect", NewCT_InnerShadowEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LuminanceEffect", NewCT_LuminanceEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_OuterShadowEffect", NewCT_OuterShadowEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PresetShadowEffect", NewCT_PresetShadowEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ReflectionEffect", NewCT_ReflectionEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_RelativeOffsetEffect", NewCT_RelativeOffsetEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SoftEdgesEffect", NewCT_SoftEdgesEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TintEffect", NewCT_TintEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TransformEffect", NewCT_TransformEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_NoFillProperties", NewCT_NoFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_SolidColorFillProperties", NewCT_SolidColorFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LinearShadeProperties", NewCT_LinearShadeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PathShadeProperties", NewCT_PathShadeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GradientStop", NewCT_GradientStop)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GradientStopList", NewCT_GradientStopList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GradientFillProperties", NewCT_GradientFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TileInfoProperties", NewCT_TileInfoProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_StretchInfoProperties", NewCT_StretchInfoProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Blip", NewCT_Blip)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BlipFillProperties", NewCT_BlipFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PatternFillProperties", NewCT_PatternFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GroupFillProperties", NewCT_GroupFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FillProperties", NewCT_FillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FillEffect", NewCT_FillEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FillOverlayEffect", NewCT_FillOverlayEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectReference", NewCT_EffectReference)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectContainer", NewCT_EffectContainer)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AlphaModulateEffect", NewCT_AlphaModulateEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BlendEffect", NewCT_BlendEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectList", NewCT_EffectList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EffectProperties", NewCT_EffectProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GeomGuide", NewCT_GeomGuide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GeomGuideList", NewCT_GeomGuideList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AdjPoint2D", NewCT_AdjPoint2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GeomRect", NewCT_GeomRect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_XYAdjustHandle", NewCT_XYAdjustHandle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PolarAdjustHandle", NewCT_PolarAdjustHandle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ConnectionSite", NewCT_ConnectionSite)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_AdjustHandleList", NewCT_AdjustHandleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ConnectionSiteList", NewCT_ConnectionSiteList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Connection", NewCT_Connection)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DMoveTo", NewCT_Path2DMoveTo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DLineTo", NewCT_Path2DLineTo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DArcTo", NewCT_Path2DArcTo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DQuadBezierTo", NewCT_Path2DQuadBezierTo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DCubicBezierTo", NewCT_Path2DCubicBezierTo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DClose", NewCT_Path2DClose)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2D", NewCT_Path2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Path2DList", NewCT_Path2DList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PresetGeometry2D", NewCT_PresetGeometry2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PresetTextShape", NewCT_PresetTextShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_CustomGeometry2D", NewCT_CustomGeometry2D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineEndProperties", NewCT_LineEndProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineJoinBevel", NewCT_LineJoinBevel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineJoinRound", NewCT_LineJoinRound)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineJoinMiterProperties", NewCT_LineJoinMiterProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_PresetLineDashProperties", NewCT_PresetLineDashProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_DashStop", NewCT_DashStop)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_DashStopList", NewCT_DashStopList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_LineProperties", NewCT_LineProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ShapeProperties", NewCT_ShapeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_GroupShapeProperties", NewCT_GroupShapeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_StyleMatrixReference", NewCT_StyleMatrixReference)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_FontReference", NewCT_FontReference)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ShapeStyle", NewCT_ShapeStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_DefaultShapeDefinition", NewCT_DefaultShapeDefinition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ObjectStyleDefaults", NewCT_ObjectStyleDefaults)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_EmptyElement", NewCT_EmptyElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorMapping", NewCT_ColorMapping)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorMappingOverride", NewCT_ColorMappingOverride)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorSchemeAndMapping", NewCT_ColorSchemeAndMapping)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ColorSchemeList", NewCT_ColorSchemeList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_OfficeStyleSheet", NewCT_OfficeStyleSheet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_BaseStylesOverride", NewCT_BaseStylesOverride)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ClipboardStyleSheet", NewCT_ClipboardStyleSheet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableCellProperties", NewCT_TableCellProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Headers", NewCT_Headers)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableCol", NewCT_TableCol)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableGrid", NewCT_TableGrid)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableCell", NewCT_TableCell)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableRow", NewCT_TableRow)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableProperties", NewCT_TableProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Table", NewCT_Table)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Cell3D", NewCT_Cell3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_ThemeableLineStyle", NewCT_ThemeableLineStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableStyleTextStyle", NewCT_TableStyleTextStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableCellBorderStyle", NewCT_TableCellBorderStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableBackgroundStyle", NewCT_TableBackgroundStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableStyleCellStyle", NewCT_TableStyleCellStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TablePartStyle", NewCT_TablePartStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableStyle", NewCT_TableStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TableStyleList", NewCT_TableStyleList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextParagraph", NewCT_TextParagraph)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextListStyle", NewCT_TextListStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextNormalAutofit", NewCT_TextNormalAutofit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextShapeAutofit", NewCT_TextShapeAutofit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextNoAutofit", NewCT_TextNoAutofit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBodyProperties", NewCT_TextBodyProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBody", NewCT_TextBody)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBulletColorFollowText", NewCT_TextBulletColorFollowText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBulletSizeFollowText", NewCT_TextBulletSizeFollowText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBulletSizePercent", NewCT_TextBulletSizePercent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBulletSizePoint", NewCT_TextBulletSizePoint)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBulletTypefaceFollowText", NewCT_TextBulletTypefaceFollowText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextAutonumberBullet", NewCT_TextAutonumberBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextCharBullet", NewCT_TextCharBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextBlipBullet", NewCT_TextBlipBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextNoBullet", NewCT_TextNoBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextFont", NewCT_TextFont)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextUnderlineLineFollowText", NewCT_TextUnderlineLineFollowText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextUnderlineFillFollowText", NewCT_TextUnderlineFillFollowText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextUnderlineFillGroupWrapper", NewCT_TextUnderlineFillGroupWrapper)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextCharacterProperties", NewCT_TextCharacterProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_Boolean", NewCT_Boolean)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextSpacingPercent", NewCT_TextSpacingPercent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextSpacingPoint", NewCT_TextSpacingPoint)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextTabStop", NewCT_TextTabStop)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextTabStopList", NewCT_TextTabStopList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextLineBreak", NewCT_TextLineBreak)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextSpacing", NewCT_TextSpacing)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextParagraphProperties", NewCT_TextParagraphProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_TextField", NewCT_TextField)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "CT_RegularTextRun", NewCT_RegularTextRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "videoFile", NewVideoFile)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "graphic", NewGraphic)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "blip", NewBlip)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "theme", NewTheme)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "themeOverride", NewThemeOverride)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "themeManager", NewThemeManager)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "tbl", NewTbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "tblStyleLst", NewTblStyleLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_Media", NewEG_Media)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ColorTransform", NewEG_ColorTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_OfficeArtExtensionList", NewEG_OfficeArtExtensionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ColorChoice", NewEG_ColorChoice)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_Text3D", NewEG_Text3D)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ShadeProperties", NewEG_ShadeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_FillModeProperties", NewEG_FillModeProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_FillProperties", NewEG_FillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_Effect", NewEG_Effect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_EffectProperties", NewEG_EffectProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_Geometry", NewEG_Geometry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextGeometry", NewEG_TextGeometry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_LineFillProperties", NewEG_LineFillProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_LineJoinProperties", NewEG_LineJoinProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_LineDashProperties", NewEG_LineDashProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ThemeableFillStyle", NewEG_ThemeableFillStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ThemeableEffectStyle", NewEG_ThemeableEffectStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_ThemeableFontStyles", NewEG_ThemeableFontStyles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextAutofit", NewEG_TextAutofit)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextBulletColor", NewEG_TextBulletColor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextBulletSize", NewEG_TextBulletSize)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextBulletTypeface", NewEG_TextBulletTypeface)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextBullet", NewEG_TextBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextUnderlineLine", NewEG_TextUnderlineLine)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextUnderlineFill", NewEG_TextUnderlineFill)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "EG_TextRun", NewEG_TextRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "AG_Blob", NewAG_Blob)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/main", "AG_Locking", NewAG_Locking)
}
