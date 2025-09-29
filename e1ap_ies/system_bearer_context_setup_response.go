package e1ap_ies

// SystemBearerContextSetupResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:739
// SystemBearerContextSetupResponse represents a CHOICE type.
type SystemBearerContextSetupResponse struct {
	EUTRANBearerContextSetupResponse ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextSetupResponse  ProtocolIEContainer `asn1:"mandatory"`
}
