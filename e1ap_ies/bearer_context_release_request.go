package e1ap_ies

// BearerContextReleaseRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1056
type BearerContextReleaseRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
