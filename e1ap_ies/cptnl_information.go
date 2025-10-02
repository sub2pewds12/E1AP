package e1ap_ies

// CPTNLInformation From: 9_4_5_Information_Element_Definitions.txt:266
const (
	CPTNLInformationPresentNothing uint64 = iota
	CPTNLInformationPresentEndpointIPAddress
)

type CPTNLInformation struct {
	Choice            uint64
	EndpointIPAddress *[]byte
}
