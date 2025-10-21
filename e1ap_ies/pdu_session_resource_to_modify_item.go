package e1ap_ies

// PDUSessionResourceToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               PDUSessionID                      `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication               `aper:"optional,ext"`
	PDUSessionResourceDLAMBR                   *BitRate                          `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                 `aper:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation        `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *InactivityTimer                  `aper:"lb:1,ub:7200,optional,reject,ext"`
	NetworkInstance                            *NetworkInstance                  `aper:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `aper:"optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN            `aper:"optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN            `aper:"optional,ext"`
	IEExtensions                               *ProtocolExtensionContainer       `aper:"optional,ext"`
	// Possible extensions:
	// - SNSSAI (ID: id-SNSSAI)
	// - CommonNetworkInstance (ID: id-CommonNetworkInstance)
	// - UPTNLInformation (ID: id-redundant-nG-UL-UP-TNL-Information)
	// - CommonNetworkInstance (ID: id-RedundantCommonNetworkInstance)
	// - DataForwardingtoEUTRANInformationList (ID: id-DataForwardingtoE-UTRANInformationList)
}
