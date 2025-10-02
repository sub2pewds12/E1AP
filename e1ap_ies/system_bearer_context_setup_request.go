package e1ap_ies

// SystemBearerContextSetupRequest From: 9_4_4_PDU_Definitions.txt:696
const (
	SystemBearerContextSetupRequestPresentNothing uint64 = iota
	SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN
	SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex
)

type SystemBearerContextSetupRequest struct {
	Choice                     uint64
	DRBToSetupListEUTRAN       []DRBToSetupItemEUTRAN
	SubscriberProfileIDforRFP  *int64
	AdditionalRRMPriorityIndex *[]byte
}
