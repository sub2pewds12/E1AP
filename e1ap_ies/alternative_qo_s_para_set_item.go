package e1ap_ies

import "github.com/lvdund/ngap/aper"

// AlternativeQoSParaSetItem From: 9_4_5_Information_Element_Definitions.txt:115
// ASN.1 Data Type: SEQUENCE
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex aper.Integer     `aper:"mandatory,ext"`
	GuaranteedFlowBitRateDL      *aper.Integer    `aper:"optional,reject,ext"`
	GuaranteedFlowBitRateUL      *aper.Integer    `aper:"optional,reject,ext"`
	PacketDelayBudget            *aper.Integer    `aper:"optional,ext"`
	PacketErrorRate              *PacketErrorRate `aper:"optional,ext"`
}
