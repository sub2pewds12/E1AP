package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QOSFlowMappingIndication is a generated ENUMERATED type.
type QOSFlowMappingIndication struct {
	Value aper.Enumerated
}

const (
	QOSFlowMappingIndicationUl aper.Enumerated = 0
	QOSFlowMappingIndicationDl aper.Enumerated = 1
)
