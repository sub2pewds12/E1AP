package e1ap_ies

// BearerContextSetupResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:726
type BearerContextSetupResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
