package e1ap_ies

// SystemBearerContextModificationResponse From: 9_4_4_PDU_Definitions.txt:861
const (
	SystemBearerContextModificationResponsePresentNothing uint64 = iota
	SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN
	SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo
)

type SystemBearerContextModificationResponse struct {
	Choice                        uint64
	DRBSetupModListEUTRAN         []DRBSetupModItemEUTRAN
	DRBFailedModListEUTRAN        []DRBFailedModItemEUTRAN
	DRBModifiedListEUTRAN         []DRBModifiedItemEUTRAN
	DRBFailedToModifyListEUTRAN   []DRBFailedToModifyItemEUTRAN
	RetainabilityMeasurementsInfo []DRBRemovedItem
}
