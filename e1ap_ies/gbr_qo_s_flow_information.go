package e1ap_ies

// GBRQoSFlowInformation From: 9_4_5_Information_Element_Definitions.txt:1178
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        int64  `asn1:"mandatory,ext"`
	MaxFlowBitRateUplink          int64  `asn1:"mandatory,ext"`
	GuaranteedFlowBitRateDownlink int64  `asn1:"mandatory,ext"`
	GuaranteedFlowBitRateUplink   int64  `asn1:"mandatory,ext"`
	MaxPacketLossRateDownlink     *int64 `asn1:"optional,ext"`
	MaxPacketLossRateUplink       *int64 `asn1:"optional,ext"`
}
