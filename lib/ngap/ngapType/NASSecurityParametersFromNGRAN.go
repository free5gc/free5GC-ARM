//go:binary-only-package

package ngapType

import "free5gc/lib/aper"

// Need to import "free5gc/lib/aper" if it uses "aper"

type NASSecurityParametersFromNGRAN struct {
	Value aper.OctetString
}
