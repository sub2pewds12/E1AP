package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NonDynamic5QIDescriptor is a generated SEQUENCE type.
type NonDynamic5QIDescriptor struct {
	FiveQI             aper.Integer  `aper:"lb:0,ub:255,mandatory,ext"`
	QoSPriorityLevel   *aper.Integer `aper:"lb:0,ub:127,optional,ext"`
	AveragingWindow    *aper.Integer `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *aper.Integer `aper:"lb:0,ub:4095,optional,ext"`
}
