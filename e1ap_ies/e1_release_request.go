package e1ap_ies

// E1ReleaseRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:633
type E1ReleaseRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
