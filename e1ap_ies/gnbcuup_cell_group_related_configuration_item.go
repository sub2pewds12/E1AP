package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPCellGroupRelatedConfigurationItem From: 9_4_5_Information_Element_Definitions.txt:1088
// ASN.1 Data Type: SEQUENCE
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      aper.Integer     `aper:"mandatory"`
	UPTNLInformation UPTNLInformation `aper:"mandatory"`
	ULConfiguration  *ULConfiguration `aper:"optional"`
}
