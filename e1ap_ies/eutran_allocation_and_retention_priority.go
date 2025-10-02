package e1ap_ies

// EUTRANAllocationAndRetentionPriority From: 9_4_5_Information_Element_Definitions.txt:1027
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           int64                   `asn1:"mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability    `asn1:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability `asn1:"mandatory,ext"`
}
