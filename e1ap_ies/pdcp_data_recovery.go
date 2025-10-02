package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPDataRecovery From: 9_4_5_Information_Element_Definitions.txt:1580
const (
	PDCPDataRecoveryTrue aper.Enumerated = 0
)

type PDCPDataRecovery struct {
	Value aper.Enumerated
}
