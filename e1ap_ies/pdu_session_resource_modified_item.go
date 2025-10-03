package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceModifiedItem From: 9_4_5_Information_Element_Definitions.txt:1721
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceModifiedItem struct {
	PDUSessionID                                aper.Integer                 `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation                        *UPTNLInformation            `aper:"optional,ext"`
	SecurityResult                              *SecurityResult              `aper:"optional,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation   `aper:"optional,ext"`
	DRBSetupListNGRAN                           []DRBSetupItemNGRAN          `aper:"optional,ext"`
	DRBFailedListNGRAN                          []DRBFailedItemNGRAN         `aper:"optional,ext"`
	DRBModifiedListNGRAN                        []DRBModifiedItemNGRAN       `aper:"optional,ext"`
	DRBFailedToModifyListNGRAN                  []DRBFailedToModifyItemNGRAN `aper:"optional,ext"`
}
