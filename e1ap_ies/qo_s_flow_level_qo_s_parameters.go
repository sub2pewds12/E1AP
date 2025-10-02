package e1ap_ies

// QoSFlowLevelQoSParameters From: 9_4_5_Information_Element_Definitions.txt:2019
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                 `asn1:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority                `asn1:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                             `asn1:"optional,ext"`
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute   `asn1:"optional,ext"`
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation `asn1:"optional,ext"`
	PagingPolicyIndicator            *int64                                             `asn1:"optional,ext"`
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator   `asn1:"optional,ext"`
}
