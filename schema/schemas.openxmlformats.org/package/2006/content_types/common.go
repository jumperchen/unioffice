// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"regexp"

	"baliance.com/gooxml"
)

const ST_ContentTypePattern = `(((([[[:ascii:]]-[\p{Cc}\(\)<>@,;:\\"/\[\]\?=\{\}\s\t]])+))/((([[[:ascii:]]-[\p{Cc}\(\)<>@,;:\\"/\[\]\?=\{\}\s\t]])+))((\s+)*;(\s+)*(((([[[:ascii:]]-[\p{Cc}\(\)<>@,;:\\"/\[\]\?=\{\}\s\t]])+))=((([[[:ascii:]]-[\p{Cc}\(\)<>@,;:\\"/\[\]\?=\{\}\s\t]])+)|("(([[^[:ascii:]][[:ascii:]]-[\p{Cc}"\n\r]]|(\s+))|(\\[[[:ascii:]]]))*"))))*)`

var ST_ContentTypePatternRe = regexp.MustCompile(ST_ContentTypePattern)

const ST_ExtensionPattern = `([!$&'\(\)\*\+,:=]|(%[0-9a-fA-F][0-9a-fA-F])|[:@]|[a-zA-Z0-9\-_~])+`

var ST_ExtensionPatternRe = regexp.MustCompile(ST_ExtensionPattern)

func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Types", NewCT_Types)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Default", NewCT_Default)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Override", NewCT_Override)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Types", NewTypes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Default", NewDefault)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Override", NewOverride)
}
