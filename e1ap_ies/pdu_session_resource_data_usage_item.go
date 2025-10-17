package e1ap_ies

// PDUSessionResourceDataUsageItem is a generated SEQUENCE type.
type PDUSessionResourceDataUsageItem struct {
	PDUSessionID         PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	MRDCUsageInformation MRDCUsageInformation        `aper:"mandatory,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}
