package e1ap_ies

// Dynamic5QIDescriptor From: 9_4_5_Information_Element_Definitions.txt:940
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   int64                              `asn1:"mandatory,ext"`
	PacketDelayBudget  int64                              `asn1:"mandatory,ext"`
	PacketErrorRate    PacketErrorRate                    `asn1:"mandatory,ext"`
	FiveQI             *int64                             `asn1:"optional,ext"`
	DelayCritical      *Dynamic5QIDescriptorDelayCritical `asn1:"optional,ext"`
	AveragingWindow    *int64                             `asn1:"optional,ext"`
	MaxDataBurstVolume *int64                             `asn1:"optional,ext"`
}
