package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GBRQoSFlowInformation From: 9_4_5_Information_Element_Definitions.txt:1178
// ASN.1 Data Type: SEQUENCE
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        aper.Integer  `aper:"mandatory,reject,ext"`
	MaxFlowBitRateUplink          aper.Integer  `aper:"mandatory,reject,ext"`
	GuaranteedFlowBitRateDownlink aper.Integer  `aper:"mandatory,reject,ext"`
	GuaranteedFlowBitRateUplink   aper.Integer  `aper:"mandatory,reject,ext"`
	MaxPacketLossRateDownlink     *aper.Integer `aper:"optional,ext"`
	MaxPacketLossRateUplink       *aper.Integer `aper:"optional,ext"`
}
