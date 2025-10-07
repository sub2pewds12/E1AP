package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceDataUsageItem is a generated SEQUENCE type.
type PDUSessionResourceDataUsageItem struct {
	PDUSessionID         aper.Integer                `aper:"lb:0,ub:255,mandatory,ext"`
	MRDCUsageInformation MRDCUsageInformation        `aper:"mandatory,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}
