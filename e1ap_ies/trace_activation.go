package e1ap_ies

// TraceActivation From: 9_4_5_Information_Element_Definitions.txt:2270
type TraceActivation struct {
	TraceID                        []byte     `asn1:"lb:8,ub:8,mandatory,ext"`
	InterfacesToTrace              []byte     `asn1:"lb:8,ub:8,mandatory,ext"`
	TraceDepth                     TraceDepth `asn1:"mandatory,ext"`
	TraceCollectionEntityIPAddress []byte     `asn1:"mandatory,ext"`
}
