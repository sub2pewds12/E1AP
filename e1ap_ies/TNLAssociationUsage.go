package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TNLAssociationUsage is a generated ENUMERATED type.
type TNLAssociationUsage struct {
	Value aper.Enumerated
}

const (
	TNLAssociationUsageUe    aper.Enumerated = 0
	TNLAssociationUsageNonUe aper.Enumerated = 1
	TNLAssociationUsageBoth  aper.Enumerated = 2
)

func (e *TNLAssociationUsage) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TNLAssociationUsage to be generated here.
	return fmt.Errorf("Encode not implemented for enum TNLAssociationUsage")
}

func (e *TNLAssociationUsage) Decode(r *aper.AperReader) error {
	// Decode logic for enum TNLAssociationUsage to be generated here.
	return fmt.Errorf("Decode not implemented for enum TNLAssociationUsage")
}
