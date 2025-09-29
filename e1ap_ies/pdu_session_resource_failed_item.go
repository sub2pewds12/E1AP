package e1ap_ies

// PDUSessionResourceFailedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1682
type PDUSessionResourceFailedItem struct {
	PDUSessionID int64 `asn1:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause `asn1:"mandatory,ext"`
}
