package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QoSFlowLevelQoSParameters From: 9_4_5_Information_Element_Definitions.txt:2019
// ASN.1 Data Type: SEQUENCE
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                 `aper:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority                `aper:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                             `aper:"optional,ext"`
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute   `aper:"optional,ext"`
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation `aper:"optional,ext"`
	PagingPolicyIndicator            *aper.Integer                                      `aper:"optional,ext"`
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator   `aper:"optional,ext"`
}
