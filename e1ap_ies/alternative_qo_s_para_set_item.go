package e1ap_ies

import "github.com/lvdund/ngap/aper"

// AlternativeQoSParaSetItem is a generated SEQUENCE type.
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex aper.Integer                `aper:"lb:1,ub:8,mandatory,ext"`
	GuaranteedFlowBitRateDL      *aper.Integer               `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	GuaranteedFlowBitRateUL      *aper.Integer               `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	PacketDelayBudget            *aper.Integer               `aper:"lb:0,ub:1023,optional,ext"`
	PacketErrorRate              *PacketErrorRate            `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer `aper:"optional,ext"`
}
