package e1ap_ies

// GNBCUUPOverloadInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1219
type GNBCUUPOverloadInformation int32

const (
	GNBCUUPOverloadInformation_Overloaded    GNBCUUPOverloadInformation = 0
	GNBCUUPOverloadInformation_NotOverloaded GNBCUUPOverloadInformation = 1
)
