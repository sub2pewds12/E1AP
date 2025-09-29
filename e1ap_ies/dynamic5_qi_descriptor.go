package e1ap_ies

// Dynamic5QIDescriptor represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:940
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   int64                             `asn1:"lb:0,ub:127,mandatory,ext"`
	PacketDelayBudget  int64                             `asn1:"lb:0,ub:1023,mandatory,ext"`
	PacketErrorRate    PacketErrorRate                   `asn1:"mandatory,ext"`
	FiveQI             *int64                            `asn1:"lb:0,ub:255,optional,ext"`
	DelayCritical      Dynamic5QIDescriptorDelayCritical `asn1:"mandatory,ext"`
	AveragingWindow    *int64                            `asn1:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *int64                            `asn1:"lb:0,ub:2000000,optional,ext"`
}
