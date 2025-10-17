package e1ap_ies

// PDUSessionResourceSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult             `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation            `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN      `aper:"mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN     `aper:"optional,ext"`
	IEExtensions                                *ProtocolExtensionContainer `aper:"optional,ext"`
}
