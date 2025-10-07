package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               aper.Integer                      `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication               `aper:"optional,ext"`
	PDUSessionResourceDLAMBR                   *aper.Integer                     `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                 `aper:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation        `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *aper.Integer                     `aper:"lb:1,ub:7200,optional,reject,ext"`
	NetworkInstance                            *aper.Integer                     `aper:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `aper:"optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN            `aper:"optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN            `aper:"optional,ext"`
	IEExtensions                               *ProtocolExtensionContainer       `aper:"optional,ext"`
}
