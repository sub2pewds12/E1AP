package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ULConfiguration From: 9_4_5_Information_Element_Definitions.txt:2375
const (
	ULConfigurationNoData aper.Enumerated = 0
	ULConfigurationShared aper.Enumerated = 1
	ULConfigurationOnly   aper.Enumerated = 2
)

type ULConfiguration struct {
	Value aper.Enumerated
}
