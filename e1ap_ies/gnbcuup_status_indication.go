package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPStatusIndication From: 9_4_4_PDU_Definitions.txt:1224
// ASN.1 Data Type: SEQUENCE
type GNBCUUPStatusIndication struct {
	TransactionID              aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUUPOverloadInformation GNBCUUPOverloadInformation `aper:"mandatory,reject,ext"`
}
