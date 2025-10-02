package e1ap_ies

// QOSFlowRemovedItem From: 9_4_5_Information_Element_Definitions.txt:2037
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             int64                                       `asn1:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession `asn1:"optional,ext"`
	QOSFlowAccumulatedSessionTime *[]byte                                     `asn1:"optional,ext"`
}
