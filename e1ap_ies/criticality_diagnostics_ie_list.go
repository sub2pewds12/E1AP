package e1ap_ies

// CriticalityDiagnosticsIEList From: manual_patch:-1
type CriticalityDiagnosticsIEList []CriticalityDiagnosticsIEItem

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `asn1:"choice:ignore"`
	IEID          ProtocolIEID `asn1:"choice:ignore"`
	TypeOfError   TypeOfError  `asn1:"choice:ignore"`
}
