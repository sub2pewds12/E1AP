package e1ap_ies

import "github.com/lvdund/ngap/aper"

// RLCMode From: 9_4_5_Information_Element_Definitions.txt:2095
const (
	RLCModeRlcTm                 aper.Enumerated = 0
	RLCModeRlcAm                 aper.Enumerated = 1
	RLCModeRlcUmBidirectional    aper.Enumerated = 2
	RLCModeRlcUmUnidirectionalUl aper.Enumerated = 3
	RLCModeRlcUmUnidirectionalDl aper.Enumerated = 4
)

type RLCMode struct {
	Value aper.Enumerated
}
