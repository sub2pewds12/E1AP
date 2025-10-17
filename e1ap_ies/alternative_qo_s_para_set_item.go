package e1ap_ies

// AlternativeQoSParaSetItem is a generated SEQUENCE type.
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex AlternativeQoSParaSetItemAlternativeQoSParameterIndex `aper:"lb:1,ub:8,mandatory,ext"`
	GuaranteedFlowBitRateDL      *BitRate                                              `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	GuaranteedFlowBitRateUL      *BitRate                                              `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	PacketDelayBudget            *PacketDelayBudget                                    `aper:"lb:0,ub:1023,optional,ext"`
	PacketErrorRate              *PacketErrorRate                                      `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer                           `aper:"optional,ext"`
}
