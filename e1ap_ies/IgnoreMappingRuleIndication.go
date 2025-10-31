package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// IgnoreMappingRuleIndication is a generated ENUMERATED type.
type IgnoreMappingRuleIndication struct {
	Value aper.Enumerated
}

const (
	IgnoreMappingRuleIndicationTrue aper.Enumerated = 0
)

func (e *IgnoreMappingRuleIndication) Encode(w *aper.AperWriter) error {
	// Encode logic for enum IgnoreMappingRuleIndication to be generated here.
	return fmt.Errorf("Encode not implemented for enum IgnoreMappingRuleIndication")
}

func (e *IgnoreMappingRuleIndication) Decode(r *aper.AperReader) error {
	// Decode logic for enum IgnoreMappingRuleIndication to be generated here.
	return fmt.Errorf("Decode not implemented for enum IgnoreMappingRuleIndication")
}
