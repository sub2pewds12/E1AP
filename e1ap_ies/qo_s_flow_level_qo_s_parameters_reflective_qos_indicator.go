package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QoSFlowLevelQoSParametersReflectiveQOSIndicator From: 9_4_5_Information_Element_Definitions.txt:2019
const (
	QoSFlowLevelQoSParametersReflectiveQOSIndicatorEnabled aper.Enumerated = 0
)

type QoSFlowLevelQoSParametersReflectiveQOSIndicator struct {
	Value aper.Enumerated
}
