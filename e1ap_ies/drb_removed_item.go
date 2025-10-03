package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRemovedItem From: 9_4_5_Information_Element_Definitions.txt:577
// ASN.1 Data Type: SEQUENCE
type DRBRemovedItem struct {
	DRBID                     aper.Integer                        `aper:"mandatory,ext"`
	DRBReleasedInSession      *DRBRemovedItemDRBReleasedInSession `aper:"optional,ext"`
	DRBAccumulatedSessionTime *aper.OctetString                   `aper:"lb:5,ub:5,optional,ext"`
	QOSFlowRemovedList        *SEQUENCE                           `aper:"optional,ext"`
}
