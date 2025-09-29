package e1ap_ies

// BearerContextInactivityNotification represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1084
type BearerContextInactivityNotification struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
