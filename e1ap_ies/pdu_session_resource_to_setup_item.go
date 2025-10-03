package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceToSetupItem From: 9_4_5_Information_Element_Definitions.txt:1836
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceToSetupItem struct {
	PDUSessionID                               aper.Integer                      `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                    `aper:"mandatory,ext"`
	SNSSAI                                     SNSSAI                            `aper:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                `aper:"mandatory,ext"`
	PDUSessionResourceDLAMBR                   *aper.Integer                     `aper:"optional,reject,ext"`
	NGULUPTNLInformation                       UPTNLInformation                  `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *aper.Integer                     `aper:"optional,reject,ext"`
	ExistingAllocatedNGDLUPTNLInfo             *UPTNLInformation                 `aper:"optional,ext"`
	NetworkInstance                            *aper.Integer                     `aper:"optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `aper:"mandatory,ext"`
}
