package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DRBSetupItemEUTRANS1DLUPUnchanged From: 9_4_5_Information_Element_Definitions.txt:623
// ASN.1 Data Type: ENUMERATED
const (
	DRBSetupItemEUTRANS1DLUPUnchangedTrue aper.Enumerated = 0
)

type DRBSetupItemEUTRANS1DLUPUnchanged struct {
	Value aper.Enumerated
}
