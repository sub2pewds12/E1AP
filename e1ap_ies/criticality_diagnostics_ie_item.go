package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CriticalityDiagnosticsIEItem From: manual_patch:-1
// ASN.1 Data Type: SEQUENCE
type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `aper:"mandatory,ignore"`
	IEID          aper.Integer `aper:"lb:0,ub:MaxProtocolIEs,mandatory,ignore"`
	TypeOfError   TypeOfError  `aper:"mandatory,ignore"`
}
