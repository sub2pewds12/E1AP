package e1ap_ies

// UEAssociatedLogicalE1ConnectionItem From: 9_4_5_Information_Element_Definitions.txt:2364
type UEAssociatedLogicalE1ConnectionItem struct {
	GNBCUCPUEE1APID *int64 `asn1:"lb:0,ub:4294967295,optional,ext"`
	GNBCUUPUEE1APID *int64 `asn1:"lb:0,ub:4294967295,optional,ext"`
}
