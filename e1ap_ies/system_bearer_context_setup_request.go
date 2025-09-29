package e1ap_ies

// SystemBearerContextSetupRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:696
// SystemBearerContextSetupRequest represents a CHOICE type.
type SystemBearerContextSetupRequest struct {
	EUTRANBearerContextSetupRequest ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextSetupRequest  ProtocolIEContainer `asn1:"mandatory"`
}
