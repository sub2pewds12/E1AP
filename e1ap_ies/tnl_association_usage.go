package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TNLAssociationUsage From: 9_4_5_Information_Element_Definitions.txt:2226
const (
	TNLAssociationUsageUe    aper.Enumerated = 0
	TNLAssociationUsageNonUe aper.Enumerated = 1
	TNLAssociationUsageBoth  aper.Enumerated = 2
)

type TNLAssociationUsage struct {
	Value aper.Enumerated
}
