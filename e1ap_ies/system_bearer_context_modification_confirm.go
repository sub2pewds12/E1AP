package e1ap_ies

// SystemBearerContextModificationConfirm From: 9_4_4_PDU_Definitions.txt:979
// ASN.1 Data Type: CHOICE
const (
	SystemBearerContextModificationConfirmPresentNothing uint64 = iota
	SystemBearerContextModificationConfirmPresentDRBConfirmModifiedListEUTRAN
)

type SystemBearerContextModificationConfirm struct {
	Choice                       uint64
	DRBConfirmModifiedListEUTRAN []DRBConfirmModifiedItemEUTRAN
}
