// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"fmt"
)

// ST_HpsMeasure is a union type
type ST_HpsMeasure struct {
	ST_UnsignedDecimalNumber    *uint64
	ST_PositiveUniversalMeasure *string
}

func (m *ST_HpsMeasure) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_HpsMeasure) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_UnsignedDecimalNumber != nil {
		mems = append(mems, "ST_UnsignedDecimalNumber")
	}
	if m.ST_PositiveUniversalMeasure != nil {
		mems = append(mems, "ST_PositiveUniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_HpsMeasure) String() string {
	if m.ST_UnsignedDecimalNumber != nil {
		return fmt.Sprintf("%v", *m.ST_UnsignedDecimalNumber)
	}
	if m.ST_PositiveUniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_PositiveUniversalMeasure)
	}
	return ""
}
