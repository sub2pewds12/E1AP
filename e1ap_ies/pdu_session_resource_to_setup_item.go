package e1ap_ies

// PDUSessionResourceToSetupItem is a generated SEQUENCE type.
type PDUSessionResourceToSetupItem struct {
	PDUSessionID                               PDUSessionID                      `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                    `aper:"mandatory,ext"`
	SNSSAI                                     SNSSAI                            `aper:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                `aper:"mandatory,ext"`
	PDUSessionResourceDLAMBR                   *BitRate                          `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	NGULUPTNLInformation                       UPTNLInformation                  `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *InactivityTimer                  `aper:"lb:1,ub:7200,optional,reject,ext"`
	ExistingAllocatedNGDLUPTNLInfo             *UPTNLInformation                 `aper:"optional,ext"`
	NetworkInstance                            *NetworkInstance                  `aper:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `aper:"mandatory,ext"`
	IEExtensions                               *ProtocolExtensionContainer       `aper:"optional,ext"`
	// Possible extensions:
	// - CommonNetworkInstance (ID: id-CommonNetworkInstance)
	// - UPTNLInformation (ID: id-redundant-nG-UL-UP-TNL-Information)
	// - CommonNetworkInstance (ID: id-RedundantCommonNetworkInstance)
	// - RedundantPDUSessionInformation (ID: id-RedundantPDUSessionInformation)
}
