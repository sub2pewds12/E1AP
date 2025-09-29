package e1ap_ies

// BearerContextModificationRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:793
type BearerContextModificationRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
