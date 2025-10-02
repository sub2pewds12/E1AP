package e1ap_ies

// NPNSupportInfo From: 9_4_5_Information_Element_Definitions.txt:1453
const (
	NPNSupportInfoPresentNothing uint64 = iota
	NPNSupportInfoPresentSNPN
)

type NPNSupportInfo struct {
	Choice uint64
	SNPN   *NPNSupportInfoSNPN
}
