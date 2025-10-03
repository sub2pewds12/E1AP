package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataUsagePerQOSFlowItemSecondaryRATType From: 9_4_5_Information_Element_Definitions.txt:359
// ASN.1 Data Type: ENUMERATED
const (
	DataUsagePerQOSFlowItemSecondaryRATTypeNR    aper.Enumerated = 0
	DataUsagePerQOSFlowItemSecondaryRATTypeEUTRA aper.Enumerated = 1
)

type DataUsagePerQOSFlowItemSecondaryRATType struct {
	Value aper.Enumerated
}
