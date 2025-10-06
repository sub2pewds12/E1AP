package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowQOSParameterItem is a generated SEQUENCE type.
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         aper.Integer              `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters `aper:"mandatory,ext"`
	QoSFlowMappingIndication  *QOSFlowMappingIndication `aper:"optional,ext"`
}
