package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EarlyForwardingSNTransfer From: 9_4_4_PDU_Definitions.txt:1495
// ASN.1 Data Type: SEQUENCE
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  aper.Integer                       `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  aper.Integer                       `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem `aper:"mandatory,reject,ext"`
}
