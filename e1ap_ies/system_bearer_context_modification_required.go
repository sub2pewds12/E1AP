package e1ap_ies

// SystemBearerContextModificationRequired represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:936
// SystemBearerContextModificationRequired represents a CHOICE type.
type SystemBearerContextModificationRequired struct {
	EUTRANBearerContextModificationRequired ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextModificationRequired  ProtocolIEContainer `asn1:"mandatory"`
}
