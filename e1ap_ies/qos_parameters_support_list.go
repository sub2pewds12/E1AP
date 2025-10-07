package e1ap_ies

// QOSParametersSupportList is a generated SEQUENCE type.
type QOSParametersSupportList struct {
	EUTRANQOSSupportList []EUTRANQOSSupportItem      `aper:"optional,ext"`
	NGRANQOSSupportList  []NGRANQOSSupportItem       `aper:"optional,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}
