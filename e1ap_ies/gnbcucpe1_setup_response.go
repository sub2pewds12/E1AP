package e1ap_ies

// GNBCUCPE1SetupResponse From: 9_4_4_PDU_Definitions.txt:445
type GNBCUCPE1SetupResponse struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	GNBCUUPID                 int64                      `asn1:"lb:0,ub:68719476735,mandatory,reject,ext"`
	GNBCUUPName               *[]byte                    `asn1:"optional,ignore,ext"`
	CNSupport                 CNSupport                  `asn1:"mandatory,reject,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `asn1:"lb:1,ub:MaxnoofSPLMNs,mandatory,reject,ext"`
	GNBCUUPCapacity           *int64                     `asn1:"lb:0,ub:255,optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
