package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Dynamic5QIDescriptorDelayCritical is a generated ENUMERATED type.
type Dynamic5QIDescriptorDelayCritical struct {
	Value aper.Enumerated
}

const (
	Dynamic5QIDescriptorDelayCriticalDelayCritical    aper.Enumerated = 0
	Dynamic5QIDescriptorDelayCriticalNonDelayCritical aper.Enumerated = 1
)
