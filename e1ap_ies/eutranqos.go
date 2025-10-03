package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EUTRANQOS From: 9_4_5_Information_Element_Definitions.txt:1053
// ASN.1 Data Type: SEQUENCE
type EUTRANQOS struct {
	QCI                                  aper.Integer                         `aper:"lb:0,ub:255,mandatory,ext"`
	EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority `aper:"mandatory,ext"`
	GbrQosInformation                    *GBRQosInformation                   `aper:"optional,ext"`
}
