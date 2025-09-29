package e1ap_ies

// PDUSessionResourceModifiedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1721
type PDUSessionResourceModifiedItem struct {
	PDUSessionID                                int64                        `asn1:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation                        *UPTNLInformation            `asn1:"optional,ext"`
	SecurityResult                              *SecurityResult              `asn1:"optional,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation   `asn1:"optional,ext"`
	DRBSetupListNGRAN                           []DRBSetupItemNGRAN          `asn1:"lb:1,ub:Item,optional,ext"`
	DRBFailedListNGRAN                          []DRBFailedItemNGRAN         `asn1:"lb:1,ub:Item,optional,ext"`
	DRBModifiedListNGRAN                        []DRBModifiedItemNGRAN       `asn1:"lb:1,ub:Item,optional,ext"`
	DRBFailedToModifyListNGRAN                  []DRBFailedToModifyItemNGRAN `asn1:"lb:1,ub:Item,optional,ext"`
}
