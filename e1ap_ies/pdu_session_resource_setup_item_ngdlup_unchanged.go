package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceSetupItemNGDLUPUnchanged From: 9_4_5_Information_Element_Definitions.txt:1757
const (
	PDUSessionResourceSetupItemNGDLUPUnchangedTrue aper.Enumerated = 0
)

type PDUSessionResourceSetupItemNGDLUPUnchanged struct {
	Value aper.Enumerated
}
