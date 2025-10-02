package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPSNStatusRequest From: 9_4_5_Information_Element_Definitions.txt:1575
const (
	PDCPSNStatusRequestRequested aper.Enumerated = 0
)

type PDCPSNStatusRequest struct {
	Value aper.Enumerated
}
