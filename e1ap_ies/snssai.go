package e1ap_ies

// SNSSAI is a generated SEQUENCE type.
type SNSSAI struct {
	SST          SNSSAISST                   `aper:"lb:1,ub:1,mandatory,ext"`
	SD           *SNSSAISD                   `aper:"lb:3,ub:3,optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
