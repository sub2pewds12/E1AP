package e1ap_ies

// SystemBearerContextModificationConfirm represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:979
// SystemBearerContextModificationConfirm represents a CHOICE type.
type SystemBearerContextModificationConfirm struct {
	EUTRANBearerContextModificationConfirm ProtocolIEContainer `asn1:"mandatory"`
	NGRANBearerContextModificationConfirm  ProtocolIEContainer `asn1:"mandatory"`
}
