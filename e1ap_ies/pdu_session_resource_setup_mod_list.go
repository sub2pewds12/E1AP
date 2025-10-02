package e1ap_ies

// PDUSessionResourceSetupModList From: 9_4_5_Information_Element_Definitions.txt:1775
type PDUSessionResourceSetupModList []PDUSessionResourceSetupModItem

// PDUSessionResourceSetupModItem From: 9_4_5_Information_Element_Definitions.txt:1777
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                int64                      `asn1:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult            `asn1:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation           `asn1:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN     `asn1:"mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN    `asn1:"optional,ext"`
}
