package e1ap_ies

// NonDynamic5QIDescriptor is a generated SEQUENCE type.
type NonDynamic5QIDescriptor struct {
	FiveQI             NonDynamic5QIDescriptorFiveQI `aper:"lb:0,ub:255,mandatory,ext"`
	QoSPriorityLevel   *QoSPriorityLevel             `aper:"lb:0,ub:127,optional,ext"`
	AveragingWindow    *AveragingWindow              `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *MaxDataBurstVolume           `aper:"lb:0,ub:4095,optional,ext"`
	IEExtensions       *ProtocolExtensionContainer   `aper:"optional,ext"`
	// Possible extensions:
	// - ExtendedPacketDelayBudget (ID: id-CNPacketDelayBudgetDownlink)
	// - ExtendedPacketDelayBudget (ID: id-CNPacketDelayBudgetUplink)
}
