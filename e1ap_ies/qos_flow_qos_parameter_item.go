package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowQOSParameterItem From: 9_4_5_Information_Element_Definitions.txt:2005
// ASN.1 Data Type: SEQUENCE
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         aper.Integer              `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters `aper:"mandatory,ext"`
	QoSFlowMappingIndication  *QOSFlowMappingIndication `aper:"optional,ext"`
}
