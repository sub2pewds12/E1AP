package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPCellGroupRelatedConfigurationItem is a generated SEQUENCE type.
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      aper.Integer                `aper:"lb:0,ub:3,mandatory"`
	UPTNLInformation UPTNLInformation            `aper:"mandatory"`
	ULConfiguration  *ULConfiguration            `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}
