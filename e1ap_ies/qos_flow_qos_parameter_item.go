package e1ap_ies

// QOSFlowQOSParameterItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2005
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         int64                     `asn1:"lb:0,ub:63,mandatory,ext"`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters `asn1:"mandatory,ext"`
	QoSFlowMappingIndication  *QOSFlowMappingIndication `asn1:"optional,ext"`
}
