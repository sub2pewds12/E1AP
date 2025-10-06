package e1ap_ies

// SystemGNBCUUPCounterCheckRequest is a generated CHOICE type.
type SystemGNBCUUPCounterCheckRequest struct {
	Choice                              uint64
	DRBsSubjectToCounterCheckListEUTRAN []DRBsSubjectToCounterCheckItemEUTRAN
}

const (
	SystemGNBCUUPCounterCheckRequestPresentNothing uint64 = iota
	SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN
)
