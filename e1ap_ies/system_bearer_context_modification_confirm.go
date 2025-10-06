package e1ap_ies

// SystemBearerContextModificationConfirm is a generated CHOICE type.
type SystemBearerContextModificationConfirm struct {
	Choice                       uint64
	DRBConfirmModifiedListEUTRAN []DRBConfirmModifiedItemEUTRAN
}

const (
	SystemBearerContextModificationConfirmPresentNothing uint64 = iota
	SystemBearerContextModificationConfirmPresentDRBConfirmModifiedListEUTRAN
)
