package e1ap_ies

// BearerContextSetupFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:768
type BearerContextSetupFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
