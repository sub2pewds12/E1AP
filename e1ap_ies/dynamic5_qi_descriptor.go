package e1ap_ies

// Dynamic5QIDescriptor is a generated SEQUENCE type.
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   QoSPriorityLevel                   `aper:"lb:0,ub:127,mandatory,ext"`
	PacketDelayBudget  PacketDelayBudget                  `aper:"lb:0,ub:1023,mandatory,ext"`
	PacketErrorRate    PacketErrorRate                    `aper:"mandatory,ext"`
	FiveQI             *Dynamic5QIDescriptorFiveQI        `aper:"lb:0,ub:255,optional,ext"`
	DelayCritical      *Dynamic5QIDescriptorDelayCritical `aper:"optional,ext"`
	AveragingWindow    *AveragingWindow                   `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *MaxDataBurstVolume                `aper:"lb:0,ub:4095,optional,ext"`
	IEExtensions       *ProtocolExtensionContainer        `aper:"optional,ext"`
	// Possible extensions:
	// - ExtendedPacketDelayBudget (ID: id-ExtendedPacketDelayBudget)
	// - ExtendedPacketDelayBudget (ID: id-CNPacketDelayBudgetDownlink)
	// - ExtendedPacketDelayBudget (ID: id-CNPacketDelayBudgetUplink)
}
