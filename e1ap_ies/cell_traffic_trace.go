package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CellTrafficTrace From: 9_4_4_PDU_Definitions.txt:1303
// ASN.1 Data Type: SEQUENCE
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                aper.Integer      `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                aper.Integer      `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceID                        aper.OctetString  `aper:"lb:8,ub:8,mandatory,ignore,ext"`
	TraceCollectionEntityIPAddress aper.BitString    `aper:"mandatory,ignore,ext"`
	PrivacyIndicator               *PrivacyIndicator `aper:"optional,ignore,ext"`
	URIaddress                     *aper.OctetString `aper:"optional,ignore,ext"`
}
