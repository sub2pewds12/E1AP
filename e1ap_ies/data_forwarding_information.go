package e1ap_ies

// DataForwardingInformation is a generated SEQUENCE type.
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation           `aper:"optional,ext"`
	DLDataForwarding *UPTNLInformation           `aper:"optional,ext"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional,ext"`
	// Possible extensions:
	// - DataForwardingtoNGRANQoSFlowInformationList (ID: id-DataForwardingtoNG-RANQoSFlowInformationList)
}
