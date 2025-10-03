package e1ap_ies

import "github.com/lvdund/ngap/aper"

// Dynamic5QIDescriptor From: 9_4_5_Information_Element_Definitions.txt:940
// ASN.1 Data Type: SEQUENCE
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   aper.Integer                       `aper:"mandatory,ext"`
	PacketDelayBudget  aper.Integer                       `aper:"mandatory,ext"`
	PacketErrorRate    PacketErrorRate                    `aper:"mandatory,ext"`
	FiveQI             *aper.Integer                      `aper:"optional,ext"`
	DelayCritical      *Dynamic5QIDescriptorDelayCritical `aper:"optional,ext"`
	AveragingWindow    *aper.Integer                      `aper:"optional,ext"`
	MaxDataBurstVolume *aper.Integer                      `aper:"optional,ext"`
}
