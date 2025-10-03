package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequest From: 9_4_4_PDU_Definitions.txt:815
// ASN.1 Data Type: CHOICE
const (
	SystemBearerContextModificationRequestPresentNothing uint64 = iota
	SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN
	SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex
)

type SystemBearerContextModificationRequest struct {
	Choice                     uint64
	DRBToSetupModListEUTRAN    []DRBToSetupModItemEUTRAN
	DRBToModifyListEUTRAN      []DRBToModifyItemEUTRAN
	DRBToRemoveListEUTRAN      []DRBToRemoveItemEUTRAN
	SubscriberProfileIDforRFP  *aper.Integer
	AdditionalRRMPriorityIndex *aper.BitString
}
