package e1ap_ies

// SystemBearerContextModificationRequest is a generated CHOICE type.
type SystemBearerContextModificationRequest struct {
	Choice                     uint64
	DRBToSetupModListEUTRAN    []DRBToSetupModItemEUTRAN
	DRBToModifyListEUTRAN      []DRBToModifyItemEUTRAN
	DRBToRemoveListEUTRAN      []DRBToRemoveItemEUTRAN
	SubscriberProfileIDforRFP  *SubscriberProfileIDforRFP
	AdditionalRRMPriorityIndex *AdditionalRRMPriorityIndex
}

const (
	SystemBearerContextModificationRequestPresentNothing uint64 = iota
	SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN
	SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex
)
