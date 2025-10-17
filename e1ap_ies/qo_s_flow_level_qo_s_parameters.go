package e1ap_ies

// QoSFlowLevelQoSParameters is a generated SEQUENCE type.
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                 `aper:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority                `aper:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                             `aper:"optional,ext"`
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute   `aper:"optional,ext"`
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation `aper:"optional,ext"`
	PagingPolicyIndicator            *QoSFlowLevelQoSParametersPagingPolicyIndicator    `aper:"lb:1,ub:8,optional,ext"`
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator   `aper:"optional,ext"`
	IEExtensions                     *ProtocolExtensionContainer                        `aper:"optional,ext"`
}
