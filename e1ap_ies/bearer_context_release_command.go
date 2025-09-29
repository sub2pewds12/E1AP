package e1ap_ies

// BearerContextReleaseCommand represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1012
type BearerContextReleaseCommand struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
