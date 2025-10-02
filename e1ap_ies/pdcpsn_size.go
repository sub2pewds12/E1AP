package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPSNSize From: 9_4_5_Information_Element_Definitions.txt:1610
const (
	PDCPSNSizeS12 aper.Enumerated = 0
	PDCPSNSizeS18 aper.Enumerated = 1
)

type PDCPSNSize struct {
	Value aper.Enumerated
}
