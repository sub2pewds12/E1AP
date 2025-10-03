package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ResourceStatusUpdate From: 9_4_4_PDU_Definitions.txt:1401
// ASN.1 Data Type: SEQUENCE
type ResourceStatusUpdate struct {
	TransactionID                 aper.Integer                   `aper:"mandatory,reject,ext"`
	TNLAvailableCapacityIndicator *TNLAvailableCapacityIndicator `aper:"optional,ignore,ext"`
	HWCapacityIndicator           HWCapacityIndicator            `aper:"mandatory,ignore,ext"`
}
