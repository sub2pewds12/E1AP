package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// RSN is a generated ENUMERATED type.
type RSN struct {
	Value aper.Enumerated
}

const (
	RSNV1 aper.Enumerated = 0
	RSNV2 aper.Enumerated = 1
)
