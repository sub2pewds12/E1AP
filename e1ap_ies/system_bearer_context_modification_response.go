package e1ap_ies

// SystemBearerContextModificationResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:861
// SystemBearerContextModificationResponse represents a CHOICE type.
type SystemBearerContextModificationResponse struct {
	EUTRANBearerContextModificationResponse ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextModificationResponse  ProtocolIEContainer `asn1:"mandatory"`
}
