package e1ap_ies

// SystemBearerContextModificationResponse is a generated CHOICE type.
type SystemBearerContextModificationResponse struct {
	Choice                        uint64
	DRBSetupModListEUTRAN         []DRBSetupModItemEUTRAN
	DRBFailedModListEUTRAN        []DRBFailedModItemEUTRAN
	DRBModifiedListEUTRAN         []DRBModifiedItemEUTRAN
	DRBFailedToModifyListEUTRAN   []DRBFailedToModifyItemEUTRAN
	RetainabilityMeasurementsInfo []DRBRemovedItem
}

const (
	SystemBearerContextModificationResponsePresentNothing uint64 = iota
	SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN
	SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo
)
