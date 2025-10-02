package e1ap_ies

// SystemBearerContextModificationRequired From: 9_4_4_PDU_Definitions.txt:936
const (
	SystemBearerContextModificationRequiredPresentNothing uint64 = iota
	SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN
	SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN
)

type SystemBearerContextModificationRequired struct {
	Choice                        uint64
	DRBRequiredToModifyListEUTRAN []DRBRequiredToModifyItemEUTRAN
	DRBRequiredToRemoveListEUTRAN []DRBRequiredToRemoveItemEUTRAN
}
