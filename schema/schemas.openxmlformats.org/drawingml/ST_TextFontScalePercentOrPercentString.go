// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"fmt"
)

// ST_TextFontScalePercentOrPercentString is a union type
type ST_TextFontScalePercentOrPercentString struct {
	ST_TextFontScalePercent *int32
	ST_Percentage           *string
}

func (m *ST_TextFontScalePercentOrPercentString) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_TextFontScalePercentOrPercentString) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextFontScalePercent != nil {
		mems = append(mems, "ST_TextFontScalePercent")
	}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_TextFontScalePercentOrPercentString) String() string {
	if m.ST_TextFontScalePercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextFontScalePercent)
	}
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	return ""
}
