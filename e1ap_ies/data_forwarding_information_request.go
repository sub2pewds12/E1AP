package e1ap_ies

// DataForwardingInformationRequest From: 9_4_5_Information_Element_Definitions.txt:317
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest `asn1:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels []QOSFlowMappingItem  `asn1:"optional,ext"`
}
