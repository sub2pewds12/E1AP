package e1ap_ies

// MRDCDataUsageReport represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1242
type MRDCDataUsageReport struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
