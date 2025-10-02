package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EHCUplinkParametersDRBContinueEHCUL From: 9_4_5_Information_Element_Definitions.txt:995
const (
	EHCUplinkParametersDRBContinueEHCULTrue aper.Enumerated = 0
)

type EHCUplinkParametersDRBContinueEHCUL struct {
	Value aper.Enumerated
}
