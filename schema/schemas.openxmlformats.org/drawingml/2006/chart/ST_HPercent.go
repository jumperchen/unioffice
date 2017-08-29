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

// ST_HPercent is a union type
type ST_HPercent struct {
	ST_HPercentWithSymbol *string
	ST_HPercentUShort     *uint16
}

func (m *ST_HPercent) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_HPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HPercentWithSymbol != nil {
		mems = append(mems, "ST_HPercentWithSymbol")
	}
	if m.ST_HPercentUShort != nil {
		mems = append(mems, "ST_HPercentUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_HPercent) String() string {
	if m.ST_HPercentWithSymbol != nil {
		return fmt.Sprintf("%v", *m.ST_HPercentWithSymbol)
	}
	if m.ST_HPercentUShort != nil {
		return fmt.Sprintf("%v", *m.ST_HPercentUShort)
	}
	return ""
}
