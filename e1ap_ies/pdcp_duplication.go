package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPDuplication From: 9_4_5_Information_Element_Definitions.txt:1585
const (
	PDCPDuplicationTrue aper.Enumerated = 0
)

type PDCPDuplication struct {
	Value aper.Enumerated
}
