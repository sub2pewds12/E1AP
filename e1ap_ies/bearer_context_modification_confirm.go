package e1ap_ies

// BearerContextModificationConfirm represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:966
type BearerContextModificationConfirm struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
