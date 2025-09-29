package e1ap_ies

// UplinkOnlyROHC represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2430
type UplinkOnlyROHC struct {
	MaxCID       int64                      `asn1:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles int64                      `asn1:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC UplinkOnlyROHCContinueROHC `asn1:"mandatory,ext"`
}
