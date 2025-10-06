package e1ap_ies

import (
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
