package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceToSetupModItem From: 9_4_5_Information_Element_Definitions.txt:1863
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceToSetupModItem struct {
	PDUSessionID                               aper.Integer                      `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                    `aper:"mandatory,ext"`
	SNSSAI                                     SNSSAI                            `aper:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                `aper:"mandatory,ext"`
	PDUSessionResourceAMBR                     *aper.Integer                     `aper:"optional,reject,ext"`
	NGULUPTNLInformation                       UPTNLInformation                  `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *aper.Integer                     `aper:"optional,reject,ext"`
	DRBToSetupModListNGRAN                     []DRBToSetupModItemNGRAN          `aper:"mandatory,ext"`
}
