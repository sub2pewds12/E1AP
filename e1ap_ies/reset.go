package e1ap_ies

// Reset represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:249
type Reset struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
