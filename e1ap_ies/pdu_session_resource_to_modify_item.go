package e1ap_ies

// PDUSessionResourceToModifyItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1796
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               int64                             `asn1:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication               `asn1:"optional,ext"`
	PDUSessionResourceDLAMBR                   *int64                            `asn1:"lb:0,ub:4000000000000,optional,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                 `asn1:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation        `asn1:"optional,ext"`
	PDUSessionInactivityTimer                  *int64                            `asn1:"lb:1,ub:7200,optional,ext"`
	NetworkInstance                            *int64                            `asn1:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `asn1:"lb:1,ub:Item,optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN            `asn1:"lb:1,ub:Item,optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN            `asn1:"lb:1,ub:Item,optional,ext"`
}
