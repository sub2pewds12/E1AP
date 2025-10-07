package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EUTRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           aper.Integer                `aper:"lb:0,ub:0,mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}
