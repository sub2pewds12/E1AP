package e1ap_ies

// SystemBearerContextModificationRequired is a generated CHOICE type.
type SystemBearerContextModificationRequired struct {
	Choice                        uint64
	DRBRequiredToModifyListEUTRAN []DRBRequiredToModifyItemEUTRAN
	DRBRequiredToRemoveListEUTRAN []DRBRequiredToRemoveItemEUTRAN
}

const (
	SystemBearerContextModificationRequiredPresentNothing uint64 = iota
	SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN
	SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN
)
