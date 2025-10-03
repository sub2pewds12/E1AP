package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPCounterCheckRequest From: 9_4_4_PDU_Definitions.txt:1178
// ASN.1 Data Type: SEQUENCE
type GNBCUUPCounterCheckRequest struct {
	GNBCUCPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemGNBCUUPCounterCheckRequest SystemGNBCUUPCounterCheckRequest `aper:"mandatory,reject,ext"`
}
