package e1ap_ies

// QOSFlowFailedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1963
type QOSFlowFailedItem struct {
	QOSFlowIdentifier int64 `asn1:"lb:0,ub:63,mandatory,ext"`
	Cause             Cause `asn1:"mandatory,ext"`
}
