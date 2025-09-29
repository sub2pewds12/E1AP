package e1ap_ies

// BearerContextModificationFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:896
type BearerContextModificationFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
