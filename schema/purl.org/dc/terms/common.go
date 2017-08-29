// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import "baliance.com/gooxml"

func init() {
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "LCSH", NewLCSH)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "MESH", NewMESH)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "DDC", NewDDC)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "LCC", NewLCC)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "UDC", NewUDC)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "Period", NewPeriod)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "W3CDTF", NewW3CDTF)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "DCMIType", NewDCMIType)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "IMT", NewIMT)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "URI", NewURI)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "ISO639-2", NewISO639_2)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "RFC1766", NewRFC1766)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "RFC3066", NewRFC3066)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "Point", NewPoint)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "ISO3166", NewISO3166)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "Box", NewBox)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "TGN", NewTGN)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "elementOrRefinementContainer", NewElementOrRefinementContainer)
	gooxml.RegisterConstructor("http://purl.org/dc/terms/", "elementsAndRefinementsGroup", NewElementsAndRefinementsGroup)
}
