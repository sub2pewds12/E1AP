package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPCount From: 9_4_5_Information_Element_Definitions.txt:1564
// ASN.1 Data Type: SEQUENCE
type PDCPCount struct {
	PDCPSN aper.Integer `aper:"lb:0,ub:262143,mandatory,ext"`
	HFN    aper.Integer `aper:"lb:0,ub:4294967295,mandatory,ext"`
}
