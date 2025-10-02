package e1ap_ies

// SystemGNBCUUPCounterCheckRequest From: 9_4_4_PDU_Definitions.txt:1190
const (
	SystemGNBCUUPCounterCheckRequestPresentNothing uint64 = iota
	SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN
)

type SystemGNBCUUPCounterCheckRequest struct {
	Choice                              uint64
	DRBsSubjectToCounterCheckListEUTRAN []DRBsSubjectToCounterCheckItemEUTRAN
}
