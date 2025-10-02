package e1ap_ies

// UPTNLInformation From: 9_4_5_Information_Element_Definitions.txt:2421
const (
	UPTNLInformationPresentNothing uint64 = iota
	UPTNLInformationPresentGTPTunnel
)

type UPTNLInformation struct {
	Choice    uint64
	GTPTunnel *GTPTunnel
}
