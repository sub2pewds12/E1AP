package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                aper.Integer               `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult            `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation           `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN     `aper:"mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN    `aper:"optional,ext"`
}
