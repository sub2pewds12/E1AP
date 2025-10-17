package e1ap_ies

// PDUSessionResourceSetupItem is a generated SEQUENCE type.
type PDUSessionResourceSetupItem struct {
	PDUSessionID                                PDUSessionID                                `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult                             `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation                            `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation                  `aper:"optional,ext"`
	NGDLUPUnchanged                             *PDUSessionResourceSetupItemNGDLUPUnchanged `aper:"optional,ext"`
	DRBSetupListNGRAN                           []DRBSetupItemNGRAN                         `aper:"mandatory,ext"`
	DRBFailedListNGRAN                          []DRBFailedItemNGRAN                        `aper:"optional,ext"`
	IEExtensions                                *ProtocolExtensionContainer                 `aper:"optional,ext"`
}
