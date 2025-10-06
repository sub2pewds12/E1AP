package e1ap_ies

import "github.com/lvdund/ngap/aper"

// Dynamic5QIDescriptor is a generated SEQUENCE type.
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   aper.Integer                       `aper:"lb:0,ub:127,mandatory,ext"`
	PacketDelayBudget  aper.Integer                       `aper:"lb:0,ub:1023,mandatory,ext"`
	PacketErrorRate    PacketErrorRate                    `aper:"mandatory,ext"`
	FiveQI             *aper.Integer                      `aper:"lb:0,ub:255,optional,ext"`
	DelayCritical      *Dynamic5QIDescriptorDelayCritical `aper:"optional,ext"`
	AveragingWindow    *aper.Integer                      `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *aper.Integer                      `aper:"lb:0,ub:4095,optional,ext"`
}
