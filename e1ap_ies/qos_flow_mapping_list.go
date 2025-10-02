package e1ap_ies

// QOSFlowMappingList From: 9_4_5_Information_Element_Definitions.txt:1974
type QOSFlowMappingList []QOSFlowMappingItem

// QOSFlowMappingItem From: 9_4_5_Information_Element_Definitions.txt:1976
type QOSFlowMappingItem struct {
	QOSFlowIdentifier        int64                     `asn1:"lb:0,ub:63,mandatory,ext"`
	QoSFlowMappingIndication *QOSFlowMappingIndication `asn1:"optional,ext"`
}
