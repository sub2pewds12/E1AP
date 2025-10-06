package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequest is a generated CHOICE type.
type SystemBearerContextModificationRequest struct {
	Choice                     uint64
	DRBToSetupModListEUTRAN    []DRBToSetupModItemEUTRAN
	DRBToModifyListEUTRAN      []DRBToModifyItemEUTRAN
	DRBToRemoveListEUTRAN      []DRBToRemoveItemEUTRAN
	SubscriberProfileIDforRFP  *aper.Integer
	AdditionalRRMPriorityIndex *aper.BitString
}

const (
	SystemBearerContextModificationRequestPresentNothing uint64 = iota
	SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN
	SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex
)
