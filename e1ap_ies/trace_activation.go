package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TraceActivation From: 9_4_5_Information_Element_Definitions.txt:2270
// ASN.1 Data Type: SEQUENCE
type TraceActivation struct {
	TraceID                        aper.OctetString `aper:"lb:8,ub:8,mandatory,ignore,ext"`
	InterfacesToTrace              aper.BitString   `aper:"lb:8,ub:8,mandatory,ext"`
	TraceDepth                     TraceDepth       `aper:"mandatory,ext"`
	TraceCollectionEntityIPAddress aper.BitString   `aper:"mandatory,ignore,ext"`
}
