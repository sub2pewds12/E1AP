package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDCPStatusReportIndication From: 9_4_5_Information_Element_Definitions.txt:1623
// ASN.1 Data Type: ENUMERATED
const (
	PDCPStatusReportIndicationDownlink aper.Enumerated = 0
	PDCPStatusReportIndicationUplink   aper.Enumerated = 1
	PDCPStatusReportIndicationBoth     aper.Enumerated = 2
)

type PDCPStatusReportIndication struct {
	Value aper.Enumerated
}
