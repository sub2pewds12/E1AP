package e1ap_ies

// EUTRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel               `aper:"lb:0,ub:0,mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}
