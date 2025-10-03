package e1ap_ies

// SystemBearerContextSetupResponse From: 9_4_4_PDU_Definitions.txt:739
// ASN.1 Data Type: CHOICE
const (
	SystemBearerContextSetupResponsePresentNothing uint64 = iota
	SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN
	SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN
)

type SystemBearerContextSetupResponse struct {
	Choice              uint64
	DRBSetupListEUTRAN  []DRBSetupItemEUTRAN
	DRBFailedListEUTRAN []DRBFailedItemEUTRAN
}
