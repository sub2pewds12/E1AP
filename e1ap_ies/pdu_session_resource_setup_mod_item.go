package e1ap_ies

// PDUSessionResourceSetupModItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1777
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                int64                      `asn1:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult            `asn1:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation           `asn1:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN     `asn1:"lb:1,ub:Item,mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN    `asn1:"lb:1,ub:Item,optional,ext"`
}
