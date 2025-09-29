package e1ap_ies

// BearerContextSetupRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:672
type BearerContextSetupRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
