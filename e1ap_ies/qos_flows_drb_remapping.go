package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QOSFlowsDRBRemapping is a generated ENUMERATED type.
type QOSFlowsDRBRemapping struct {
	Value aper.Enumerated
}

const (
	QOSFlowsDRBRemappingUpdate              aper.Enumerated = 0
	QOSFlowsDRBRemappingSourceConfiguration aper.Enumerated = 1
)
