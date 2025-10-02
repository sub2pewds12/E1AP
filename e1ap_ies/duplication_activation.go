package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DuplicationActivation From: 9_4_5_Information_Element_Definitions.txt:933
const (
	DuplicationActivationActive   aper.Enumerated = 0
	DuplicationActivationInactive aper.Enumerated = 1
)

type DuplicationActivation struct {
	Value aper.Enumerated
}
