package e1ap_ies

// PDCPCount is a generated SEQUENCE type.
type PDCPCount struct {
	PDCPSN       PDCPSN                      `aper:"lb:0,ub:262143,mandatory,ext"`
	HFN          HFN                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
