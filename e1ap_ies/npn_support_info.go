package e1ap_ies

// NPNSupportInfo is a generated CHOICE type.
type NPNSupportInfo struct {
	Choice uint64
	SNPN   *NPNSupportInfoSNPN
}

const (
	NPNSupportInfoPresentNothing uint64 = iota
	NPNSupportInfoPresentSNPN
)
