package e1ap_ies

import "github.com/lvdund/ngap/aper"

// RATType From: 9_4_5_Information_Element_Definitions.txt:2059
const (
	RATTypeEUTRA aper.Enumerated = 0
	RATTypeNR    aper.Enumerated = 1
)

type RATType struct {
	Value aper.Enumerated
}
