package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersReflectiveQOSAttribute From: 9_4_5_Information_Element_Definitions.txt:2019
// ASN.1 Data Type: ENUMERATED
const (
	QoSFlowLevelQoSParametersReflectiveQOSAttributeSubjectTo aper.Enumerated = 0
)

type QoSFlowLevelQoSParametersReflectiveQOSAttribute struct {
	Value aper.Enumerated
}
