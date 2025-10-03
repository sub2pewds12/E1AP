package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PreEmptionCapability From: 9_4_5_Information_Element_Definitions.txt:1916
// ASN.1 Data Type: ENUMERATED
const (
	PreEmptionCapabilityShallNotTriggerPreEmption aper.Enumerated = 0
	PreEmptionCapabilityMayTriggerPreEmption      aper.Enumerated = 1
)

type PreEmptionCapability struct {
	Value aper.Enumerated
}
