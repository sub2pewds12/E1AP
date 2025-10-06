package e1ap_ies

// SystemBearerContextSetupResponse is a generated CHOICE type.
type SystemBearerContextSetupResponse struct {
	Choice              uint64
	DRBSetupListEUTRAN  []DRBSetupItemEUTRAN
	DRBFailedListEUTRAN []DRBFailedItemEUTRAN
}

const (
	SystemBearerContextSetupResponsePresentNothing uint64 = iota
	SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN
	SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN
)
