package e1ap_ies

// ResetAcknowledge represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:291
type ResetAcknowledge struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
