package e1ap_ies

// DataForwardingInformationRequest represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:317
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest `asn1:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels []QOSFlowMappingItem  `asn1:"lb:1,ub:Item,optional,ext"`
}
