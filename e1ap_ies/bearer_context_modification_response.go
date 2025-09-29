package e1ap_ies

// BearerContextModificationResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:848
type BearerContextModificationResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
