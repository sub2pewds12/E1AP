package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceActivity From: 9_4_5_Information_Element_Definitions.txt:1647
// ASN.1 Data Type: ENUMERATED
const (
	PDUSessionResourceActivityActive    aper.Enumerated = 0
	PDUSessionResourceActivityNotActive aper.Enumerated = 1
)

type PDUSessionResourceActivity struct {
	Value aper.Enumerated
}
