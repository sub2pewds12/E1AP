package e1ap_ies

// QOSFlowItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1950
type QOSFlowItem struct {
	QOSFlowIdentifier int64 `asn1:"lb:0,ub:63,mandatory,ext"`
}
