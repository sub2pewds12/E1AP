package e1ap_ies

import "github.com/lvdund/ngap/aper"

// MDTActivation From: 9_4_5_Information_Element_Definitions.txt:1378
const (
	MDTActivationImmediateMDTOnly     aper.Enumerated = 0
	MDTActivationImmediateMDTAndTrace aper.Enumerated = 1
)

type MDTActivation struct {
	Value aper.Enumerated
}
