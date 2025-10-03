package e1ap_ies

// ROHCParameters From: 9_4_5_Information_Element_Definitions.txt:2105
// ASN.1 Data Type: CHOICE
const (
	ROHCParametersPresentNothing uint64 = iota
	ROHCParametersPresentROHC
	ROHCParametersPresentUPlinkOnlyROHC
	ROHCParametersPresentChoiceExtension
)

type ROHCParameters struct {
	Choice          uint64
	ROHC            *ROHC
	UPlinkOnlyROHC  *UplinkOnlyROHC
	ChoiceExtension *ProtocolIESingleContainer
}
