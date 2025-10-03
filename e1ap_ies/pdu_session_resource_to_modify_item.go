package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceToModifyItem From: 9_4_5_Information_Element_Definitions.txt:1796
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               aper.Integer                      `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication               `aper:"optional,ext"`
	PDUSessionResourceDLAMBR                   *aper.Integer                     `aper:"optional,reject,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                 `aper:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation        `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *aper.Integer                     `aper:"optional,reject,ext"`
	NetworkInstance                            *aper.Integer                     `aper:"optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `aper:"optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN            `aper:"optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN            `aper:"optional,ext"`
}
