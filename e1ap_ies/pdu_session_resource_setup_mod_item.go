package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceSetupModItem From: 9_4_5_Information_Element_Definitions.txt:1777
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                aper.Integer               `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult            `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation           `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN     `aper:"mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN    `aper:"optional,ext"`
}
