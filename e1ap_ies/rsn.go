package e1ap_ies

// RSN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2077
type RSN int32

const (
	RSN_V1 RSN = 0
	RSN_V2 RSN = 1
)
