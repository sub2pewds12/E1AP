package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersAdditionalQOSInformation From: 9_4_5_Information_Element_Definitions.txt:2019
// ASN.1 Data Type: ENUMERATED
const (
	QoSFlowLevelQoSParametersAdditionalQOSInformationMoreLikely aper.Enumerated = 0
)

type QoSFlowLevelQoSParametersAdditionalQOSInformation struct {
	Value aper.Enumerated
}
