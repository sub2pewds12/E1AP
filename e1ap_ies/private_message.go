package e1ap_ies

// PrivateMessage From: 9_4_4_PDU_Definitions.txt:1325
// ASN.1 Data Type: SEQUENCE
type PrivateMessage struct {
	PrivateIEs PrivateIEContainer `aper:"mandatory,ext"`
}
