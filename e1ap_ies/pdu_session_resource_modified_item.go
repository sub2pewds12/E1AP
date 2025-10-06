package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceModifiedItem is a generated SEQUENCE type.
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
