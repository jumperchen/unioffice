// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"
)

// ST_BubbleScale is a union type
type ST_BubbleScale struct {
	ST_BubbleScalePercent *string
	ST_BubbleScaleUInt    *uint32
}

func (m *ST_BubbleScale) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_BubbleScale) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_BubbleScalePercent != nil {
		mems = append(mems, "ST_BubbleScalePercent")
	}
	if m.ST_BubbleScaleUInt != nil {
		mems = append(mems, "ST_BubbleScaleUInt")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_BubbleScale) String() string {
	if m.ST_BubbleScalePercent != nil {
		return fmt.Sprintf("%v", *m.ST_BubbleScalePercent)
	}
	if m.ST_BubbleScaleUInt != nil {
		return fmt.Sprintf("%v", *m.ST_BubbleScaleUInt)
	}
	return ""
}
