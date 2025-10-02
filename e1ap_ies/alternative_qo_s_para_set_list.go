package e1ap_ies

// AlternativeQoSParaSetList From: 9_4_5_Information_Element_Definitions.txt:113
type AlternativeQoSParaSetList []AlternativeQoSParaSetItem

// AlternativeQoSParaSetItem From: 9_4_5_Information_Element_Definitions.txt:115
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex int64            `asn1:"mandatory,ext"`
	GuaranteedFlowBitRateDL      *int64           `asn1:"optional,ext"`
	GuaranteedFlowBitRateUL      *int64           `asn1:"optional,ext"`
	PacketDelayBudget            *int64           `asn1:"optional,ext"`
	PacketErrorRate              *PacketErrorRate `asn1:"optional,ext"`
}
