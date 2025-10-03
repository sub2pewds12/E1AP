package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// M4period From: 9_4_5_Information_Element_Definitions.txt:1350
// ASN.1 Data Type: ENUMERATED
const (
	M4periodMs1024  aper.Enumerated = 0
	M4periodMs2048  aper.Enumerated = 1
	M4periodMs5120  aper.Enumerated = 2
	M4periodMs10240 aper.Enumerated = 3
	M4periodMin1    aper.Enumerated = 4
)

type M4period struct {
	Value aper.Enumerated
}
