package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupRequest From: 9_4_4_PDU_Definitions.txt:696
// ASN.1 Data Type: CHOICE
const (
	SystemBearerContextSetupRequestPresentNothing uint64 = iota
	SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN
	SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex
)

type SystemBearerContextSetupRequest struct {
	Choice                     uint64
	DRBToSetupListEUTRAN       []DRBToSetupItemEUTRAN
	SubscriberProfileIDforRFP  *aper.Integer
	AdditionalRRMPriorityIndex *aper.BitString
}
