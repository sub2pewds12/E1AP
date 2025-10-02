package e1ap_ies

// NonDynamic5QIDescriptor From: 9_4_5_Information_Element_Definitions.txt:1439
type NonDynamic5QIDescriptor struct {
	FiveQI             int64  `asn1:"mandatory,ext"`
	QoSPriorityLevel   *int64 `asn1:"optional,ext"`
	AveragingWindow    *int64 `asn1:"optional,ext"`
	MaxDataBurstVolume *int64 `asn1:"optional,ext"`
}
