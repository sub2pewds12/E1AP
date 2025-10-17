package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// AdditionalHandoverInfo is a generated ENUMERATED type.
type AdditionalHandoverInfo struct {
	Value aper.Enumerated
}

const (
	AdditionalHandoverInfoDiscardPdpcSN aper.Enumerated = 0
)
