package e1ap_ies

// NonDynamic5QIDescriptor represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1439
type NonDynamic5QIDescriptor struct {
	FiveQI             int64  `asn1:"lb:0,ub:255,mandatory,ext"`
	QoSPriorityLevel   *int64 `asn1:"lb:0,ub:127,optional,ext"`
	AveragingWindow    *int64 `asn1:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *int64 `asn1:"lb:0,ub:2000000,optional,ext"`
}
