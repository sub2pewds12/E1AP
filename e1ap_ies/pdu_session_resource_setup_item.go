package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceSetupItem From: 9_4_5_Information_Element_Definitions.txt:1757
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceSetupItem struct {
	PDUSessionID                                aper.Integer                                `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult                             `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation                            `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation                  `aper:"optional,ext"`
	NGDLUPUnchanged                             *PDUSessionResourceSetupItemNGDLUPUnchanged `aper:"optional,ext"`
	DRBSetupListNGRAN                           []DRBSetupItemNGRAN                         `aper:"mandatory,ext"`
	DRBFailedListNGRAN                          []DRBFailedItemNGRAN                        `aper:"optional,ext"`
}
