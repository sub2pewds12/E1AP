package e1ap_ies

// PDUSessionResourceConfirmModifiedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1669
type PDUSessionResourceConfirmModifiedItem struct {
	PDUSessionID                int64                         `asn1:"lb:0,ub:255,mandatory,ext"`
	DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN `asn1:"lb:1,ub:Item,optional,ext"`
}
