package e1ap_ies

// DLDataNotification represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1109
type DLDataNotification struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
