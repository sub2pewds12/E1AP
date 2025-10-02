package e1ap_ies

// PDCPCount From: 9_4_5_Information_Element_Definitions.txt:1564
type PDCPCount struct {
	PDCPSN int64 `asn1:"lb:0,ub:262143,mandatory,ext"`
	HFN    int64 `asn1:"lb:0,ub:4294967295,mandatory,ext"`
}
