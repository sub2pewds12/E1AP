package e1ap_ies

// QoSFlowLevelQoSParameters represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2019
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                `asn1:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority               `asn1:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                            `asn1:"optional,ext"`
	ReflectiveQOSAttribute           QoSFlowLevelQoSParametersReflectiveQOSAttribute   `asn1:"mandatory,ext"`
	AdditionalQOSInformation         QoSFlowLevelQoSParametersAdditionalQOSInformation `asn1:"mandatory,ext"`
	PagingPolicyIndicator            *int64                                            `asn1:"lb:1,ub:8,optional,ext"`
	ReflectiveQOSIndicator           QoSFlowLevelQoSParametersReflectiveQOSIndicator   `asn1:"mandatory,ext"`
}
