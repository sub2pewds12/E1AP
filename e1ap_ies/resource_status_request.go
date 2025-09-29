package e1ap_ies

// ResourceStatusRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1340
type ResourceStatusRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
