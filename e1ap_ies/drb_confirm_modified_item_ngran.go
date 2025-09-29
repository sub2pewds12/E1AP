package e1ap_ies

// DRBConfirmModifiedItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:453
type DRBConfirmModifiedItemNGRAN struct {
	DRBID                int64                      `asn1:"lb:1,ub:32,mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `asn1:"lb:1,ub:Item,optional,ext"`
}
