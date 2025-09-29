package e1ap_ies

// GBRQoSFlowInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1178
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        int64  `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	MaxFlowBitRateUplink          int64  `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	GuaranteedFlowBitRateDownlink int64  `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	GuaranteedFlowBitRateUplink   int64  `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	MaxPacketLossRateDownlink     *int64 `asn1:"lb:0,ub:1000,optional,ext"`
	MaxPacketLossRateUplink       *int64 `asn1:"lb:0,ub:1000,optional,ext"`
}
