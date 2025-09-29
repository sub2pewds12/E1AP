package e1ap_ies

// SystemBearerContextModificationRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:815
// SystemBearerContextModificationRequest represents a CHOICE type.
type SystemBearerContextModificationRequest struct {
	EUTRANBearerContextModificationRequest ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextModificationRequest  ProtocolIEContainer `asn1:"mandatory"`
}
