package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ResetAcknowledge From: 9_4_4_PDU_Definitions.txt:291
// ASN.1 Data Type: SEQUENCE
type ResetAcknowledge struct {
	TransactionID                             aper.Integer                `aper:"mandatory,reject,ext"`
	UEAssociatedLogicalE1ConnectionListResAck []ProtocolIESingleContainer `aper:"optional,ignore,ext"`
	CriticalityDiagnostics                    *CriticalityDiagnostics     `aper:"optional,ignore,ext"`
}
