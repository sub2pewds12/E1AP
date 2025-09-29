package e1ap_ies

// DataForwardingRequest represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:339
type DataForwardingRequest int32

const (
	DataForwardingRequest_UL   DataForwardingRequest = 0
	DataForwardingRequest_DL   DataForwardingRequest = 1
	DataForwardingRequest_Both DataForwardingRequest = 2
)
