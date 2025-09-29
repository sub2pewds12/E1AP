package e1ap_ies

// QOSFlowRemovedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2037
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             int64                                      `asn1:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      QOSFlowRemovedItemQOSFlowReleasedInSession `asn1:"mandatory,ext"`
	QOSFlowAccumulatedSessionTime *[]byte                                    `asn1:"optional,ext"`
}
