package e1ap_ies

// CriticalityDiagnosticsIEItem is a generated SEQUENCE type.
type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `aper:"mandatory"`
	IEID          ProtocolIEID `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	TypeOfError   TypeOfError  `aper:"mandatory"`
}
