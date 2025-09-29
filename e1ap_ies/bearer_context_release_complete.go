package e1ap_ies

// BearerContextReleaseComplete represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1030
type BearerContextReleaseComplete struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
