package e1ap_ies

import "github.com/lvdund/ngap/aper"

// RSN From: 9_4_5_Information_Element_Definitions.txt:2077
const (
	RSNV1 aper.Enumerated = 0
	RSNV2 aper.Enumerated = 1
)

type RSN struct {
	Value aper.Enumerated
}
