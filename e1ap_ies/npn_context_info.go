package e1ap_ies

// NPNContextInfo From: 9_4_5_Information_Element_Definitions.txt:1472
const (
	NPNContextInfoPresentNothing uint64 = iota
	NPNContextInfoPresentSNPN
)

type NPNContextInfo struct {
	Choice uint64
	SNPN   *NPNContextInfoSNPN
}
