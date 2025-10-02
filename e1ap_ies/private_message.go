package e1ap_ies

// PrivateMessage From: 9_4_4_PDU_Definitions.txt:1325
type PrivateMessage struct {
	PrivateIEs PrivateIEContainer `asn1:"mandatory,ext"`
}
