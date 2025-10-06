package e1ap_ies

// DataForwardingInformation is a generated SEQUENCE type.
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation `aper:"optional,ext"`
	DLDataForwarding *UPTNLInformation `aper:"optional,ext"`
}
