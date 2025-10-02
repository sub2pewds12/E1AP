package e1ap_ies

// PDUSessionResourceToModifyList From: 9_4_5_Information_Element_Definitions.txt:1794
type PDUSessionResourceToModifyList []PDUSessionResourceToModifyItem

// PDUSessionResourceToModifyItem From: 9_4_5_Information_Element_Definitions.txt:1796
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               int64                             `asn1:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication               `asn1:"optional,ext"`
	PDUSessionResourceDLAMBR                   *int64                            `asn1:"optional,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                 `asn1:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation        `asn1:"optional,ext"`
	PDUSessionInactivityTimer                  *int64                            `asn1:"optional,ext"`
	NetworkInstance                            *int64                            `asn1:"optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `asn1:"optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN            `asn1:"optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN            `asn1:"optional,ext"`
}
