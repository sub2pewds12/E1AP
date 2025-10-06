package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupRequest is a generated CHOICE type.
type SystemBearerContextSetupRequest struct {
	Choice                     uint64
	DRBToSetupListEUTRAN       []DRBToSetupItemEUTRAN
	SubscriberProfileIDforRFP  *aper.Integer
	AdditionalRRMPriorityIndex *aper.BitString
}

const (
	SystemBearerContextSetupRequestPresentNothing uint64 = iota
	SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN
	SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex
)
