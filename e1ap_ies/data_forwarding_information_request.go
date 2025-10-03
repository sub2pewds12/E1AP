package e1ap_ies

// DataForwardingInformationRequest From: 9_4_5_Information_Element_Definitions.txt:317
// ASN.1 Data Type: SEQUENCE
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest `aper:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels []QOSFlowMappingItem  `aper:"optional,ext"`
}
