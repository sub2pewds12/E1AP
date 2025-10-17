package e1ap_ies

// TSCTrafficInformation is a generated SEQUENCE type.
type TSCTrafficInformation struct {
	Periodicity      Periodicity                 `aper:"lb:1,ub:640000,mandatory"`
	BurstArrivalTime *BurstArrivalTime           `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}
