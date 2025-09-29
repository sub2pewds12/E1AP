package e1ap_ies

// PDUSessionResourceActivityItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1655
type PDUSessionResourceActivityItem struct {
	PDUSessionID               int64                      `asn1:"lb:0,ub:255,mandatory,ext"`
	PDUSessionResourceActivity PDUSessionResourceActivity `asn1:"mandatory,ext"`
}
