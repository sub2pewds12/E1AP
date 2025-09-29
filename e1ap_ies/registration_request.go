package e1ap_ies

// RegistrationRequest represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2081
type RegistrationRequest int32

const (
	RegistrationRequest_Start RegistrationRequest = 0
	RegistrationRequest_Stop  RegistrationRequest = 1
)
