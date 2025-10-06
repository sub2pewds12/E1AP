package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CriticalityDiagnosticsIEItem is a generated SEQUENCE type.
type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `aper:"mandatory,ignore"`
	IEID          aper.Integer `aper:"lb:0,ub:MaxProtocolIEs,mandatory,ignore"`
	TypeOfError   TypeOfError  `aper:"mandatory,ignore"`
}
