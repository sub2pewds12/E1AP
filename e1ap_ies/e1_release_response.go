package e1ap_ies

// E1ReleaseResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:650
type E1ReleaseResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
