package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EHCDownlinkParametersDRBContinueEHCDL From: 9_4_5_Information_Element_Definitions.txt:986
// ASN.1 Data Type: ENUMERATED
const (
	EHCDownlinkParametersDRBContinueEHCDLTrue aper.Enumerated = 0
)

type EHCDownlinkParametersDRBContinueEHCDL struct {
	Value aper.Enumerated
}
