package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSMappingInformation From: 9_4_5_Information_Element_Definitions.txt:2049
// ASN.1 Data Type: SEQUENCE
type QOSMappingInformation struct {
	Dscp      *aper.BitString `aper:"lb:6,ub:6,optional,ext"`
	FlowLabel *aper.BitString `aper:"lb:20,ub:20,optional,ext"`
}
