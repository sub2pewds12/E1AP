package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NonDynamic5QIDescriptor From: 9_4_5_Information_Element_Definitions.txt:1439
// ASN.1 Data Type: SEQUENCE
type NonDynamic5QIDescriptor struct {
	FiveQI             aper.Integer  `aper:"mandatory,ext"`
	QoSPriorityLevel   *aper.Integer `aper:"optional,ext"`
	AveragingWindow    *aper.Integer `aper:"optional,ext"`
	MaxDataBurstVolume *aper.Integer `aper:"optional,ext"`
}
