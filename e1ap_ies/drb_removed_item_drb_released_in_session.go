package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRemovedItemDRBReleasedInSession From: 9_4_5_Information_Element_Definitions.txt:577
const (
	DRBRemovedItemDRBReleasedInSessionReleasedInSession    aper.Enumerated = 0
	DRBRemovedItemDRBReleasedInSessionNotReleasedInSession aper.Enumerated = 1
)

type DRBRemovedItemDRBReleasedInSession struct {
	Value aper.Enumerated
}
