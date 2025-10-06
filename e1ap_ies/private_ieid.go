package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PrivateIEID is a generated CHOICE type.
type PrivateIEID struct {
	Choice uint64
	Local  *aper.Integer
	Global *string
}

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)
