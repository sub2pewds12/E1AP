package e1ap_ies

import (
	"fmt"

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

func (e *Dynamic5QIDescriptorDelayCritical) Encode(w *aper.AperWriter) error {
	// Encode logic for enum Dynamic5QIDescriptorDelayCritical to be generated here.
	return fmt.Errorf("Encode not implemented for enum Dynamic5QIDescriptorDelayCritical")
}

func (e *Dynamic5QIDescriptorDelayCritical) Decode(r *aper.AperReader) error {
	// Decode logic for enum Dynamic5QIDescriptorDelayCritical to be generated here.
	return fmt.Errorf("Decode not implemented for enum Dynamic5QIDescriptorDelayCritical")
}
