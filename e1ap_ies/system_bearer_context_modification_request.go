package e1ap_ies

// SystemBearerContextModificationRequest From: 9_4_4_PDU_Definitions.txt:815
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
	SubscriberProfileIDforRFP  *int64
	AdditionalRRMPriorityIndex *[]byte
}
