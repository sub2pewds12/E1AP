package e1ap_ies

// AlternativeQoSParaSetItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:115
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex int64            `asn1:"lb:1,ub:8,mandatory,ext"`
	GuaranteedFlowBitRateDL      *int64           `asn1:"lb:0,ub:4000000000000,optional,ext"`
	GuaranteedFlowBitRateUL      *int64           `asn1:"lb:0,ub:4000000000000,optional,ext"`
	PacketDelayBudget            *int64           `asn1:"lb:0,ub:1023,optional,ext"`
	PacketErrorRate              *PacketErrorRate `asn1:"optional,ext"`
}
