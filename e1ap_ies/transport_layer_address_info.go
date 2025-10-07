package e1ap_ies

// TransportLayerAddressInfo is a generated SEQUENCE type.
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    []TransportUPLayerAddressesInfoToAddItem    `aper:"optional,ext"`
	TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem `aper:"optional,ext"`
	IEExtensions                              *ProtocolExtensionContainer                 `aper:"optional,ext"`
}
