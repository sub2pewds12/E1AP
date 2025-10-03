package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBModifiedItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:546
// ASN.1 Data Type: SEQUENCE
type DRBModifiedItemEUTRAN struct {
	DRBID                   aper.Integer             `aper:"mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation        `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `aper:"optional,ext"`
	ULUPTransportParameters []UPParametersItem       `aper:"optional,ext"`
}
