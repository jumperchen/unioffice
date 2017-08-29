// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_ModifyVerifier struct {
	// Cryptographic Algorithm Name
	AlgorithmNameAttr *string
	// Password Hash Value
	HashValueAttr *string
	// Salt Value for Password Verifier
	SaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	SpinValueAttr *uint32
	// Cryptographic Provider Type
	CryptProviderTypeAttr sharedTypes.ST_CryptProv
	// Cryptographic Algorithm Class
	CryptAlgorithmClassAttr sharedTypes.ST_AlgClass
	// Cryptographic Algorithm Type
	CryptAlgorithmTypeAttr sharedTypes.ST_AlgType
	// Cryptographic Hashing Algorithm
	CryptAlgorithmSidAttr *uint32
	// Iterations to Run Hashing Algorithm
	SpinCountAttr *uint32
	// Salt for Password Verifier
	SaltDataAttr *string
	// Password Hash
	HashDataAttr *string
	// Cryptographic Provider
	CryptProviderAttr *string
	// Cryptographic Algorithm Extensibility
	AlgIdExtAttr *uint32
	// Algorithm Extensibility Source
	AlgIdExtSourceAttr *string
	// Cryptographic Provider Type Extensibility
	CryptProviderTypeExtAttr *uint32
	// Provider Type Extensibility Source
	CryptProviderTypeExtSourceAttr *string
}

func NewCT_ModifyVerifier() *CT_ModifyVerifier {
	ret := &CT_ModifyVerifier{}
	return ret
}
func (m *CT_ModifyVerifier) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spinValue"},
			Value: fmt.Sprintf("%v", *m.SpinValueAttr)})
	}
	if m.CryptProviderTypeAttr != sharedTypes.ST_CryptProvUnset {
		attr, err := m.CryptProviderTypeAttr.MarshalXMLAttr(xml.Name{Local: "cryptProviderType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmClassAttr != sharedTypes.ST_AlgClassUnset {
		attr, err := m.CryptAlgorithmClassAttr.MarshalXMLAttr(xml.Name{Local: "cryptAlgorithmClass"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmTypeAttr != sharedTypes.ST_AlgTypeUnset {
		attr, err := m.CryptAlgorithmTypeAttr.MarshalXMLAttr(xml.Name{Local: "cryptAlgorithmType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmSidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cryptAlgorithmSid"},
			Value: fmt.Sprintf("%v", *m.CryptAlgorithmSidAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	if m.SaltDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saltData"},
			Value: fmt.Sprintf("%v", *m.SaltDataAttr)})
	}
	if m.HashDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hashData"},
			Value: fmt.Sprintf("%v", *m.HashDataAttr)})
	}
	if m.CryptProviderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cryptProvider"},
			Value: fmt.Sprintf("%v", *m.CryptProviderAttr)})
	}
	if m.AlgIdExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algIdExt"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtAttr)})
	}
	if m.AlgIdExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algIdExtSource"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtSourceAttr)})
	}
	if m.CryptProviderTypeExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cryptProviderTypeExt"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtAttr)})
	}
	if m.CryptProviderTypeExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cryptProviderTypeExtSource"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtSourceAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ModifyVerifier) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
		}
		if attr.Name.Local == "spinValue" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.SpinValueAttr = &pt
		}
		if attr.Name.Local == "cryptProviderType" {
			m.CryptProviderTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmClass" {
			m.CryptAlgorithmClassAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmType" {
			m.CryptAlgorithmTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cryptAlgorithmSid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CryptAlgorithmSidAttr = &pt
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.SpinCountAttr = &pt
		}
		if attr.Name.Local == "saltData" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltDataAttr = &parsed
		}
		if attr.Name.Local == "hashData" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashDataAttr = &parsed
		}
		if attr.Name.Local == "cryptProvider" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderAttr = &parsed
		}
		if attr.Name.Local == "algIdExt" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.AlgIdExtAttr = &pt
		}
		if attr.Name.Local == "algIdExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgIdExtSourceAttr = &parsed
		}
		if attr.Name.Local == "cryptProviderTypeExt" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CryptProviderTypeExtAttr = &pt
		}
		if attr.Name.Local == "cryptProviderTypeExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderTypeExtSourceAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ModifyVerifier: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ModifyVerifier) Validate() error {
	return m.ValidateWithPath("CT_ModifyVerifier")
}
func (m *CT_ModifyVerifier) ValidateWithPath(path string) error {
	if err := m.CryptProviderTypeAttr.ValidateWithPath(path + "/CryptProviderTypeAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmClassAttr.ValidateWithPath(path + "/CryptAlgorithmClassAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmTypeAttr.ValidateWithPath(path + "/CryptAlgorithmTypeAttr"); err != nil {
		return err
	}
	return nil
}
