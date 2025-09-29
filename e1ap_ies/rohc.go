package e1ap_ies

// ROHC represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2115
type ROHC struct {
	MaxCID       int64            `asn1:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles int64            `asn1:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC ROHCContinueROHC `asn1:"mandatory,ext"`
}
