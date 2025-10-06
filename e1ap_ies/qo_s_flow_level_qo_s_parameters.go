package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QoSFlowLevelQoSParameters is a generated SEQUENCE type.
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                 `aper:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority                `aper:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                             `aper:"optional,ext"`
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute   `aper:"optional,ext"`
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation `aper:"optional,ext"`
	PagingPolicyIndicator            *aper.Integer                                      `aper:"lb:1,ub:8,optional,ext"`
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator   `aper:"optional,ext"`
}
