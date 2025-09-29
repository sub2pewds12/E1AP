package e1ap_ies

// RLCMode represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2095
type RLCMode int32

const (
	RLCMode_RlcTm                 RLCMode = 0
	RLCMode_RlcAm                 RLCMode = 1
	RLCMode_RlcUmBidirectional    RLCMode = 2
	RLCMode_RlcUmUnidirectionalUl RLCMode = 3
	RLCMode_RlcUmUnidirectionalDl RLCMode = 4
)
