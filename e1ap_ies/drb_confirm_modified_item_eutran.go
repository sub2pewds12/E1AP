package e1ap_ies

// DRBConfirmModifiedItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:440
type DRBConfirmModifiedItemEUTRAN struct {
	DRBID                int64                      `asn1:"lb:1,ub:32,mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `asn1:"lb:1,ub:Item,optional,ext"`
}
