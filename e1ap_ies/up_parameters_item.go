package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UPParametersItem From: 9_4_5_Information_Element_Definitions.txt:2398
// ASN.1 Data Type: SEQUENCE
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation `aper:"mandatory,ext"`
	CellGroupID      aper.Integer     `aper:"mandatory,ext"`
}
