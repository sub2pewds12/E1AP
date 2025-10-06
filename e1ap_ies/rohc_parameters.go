package e1ap_ies

// ROHCParameters is a generated CHOICE type.
type ROHCParameters struct {
	Choice          uint64
	ROHC            *ROHC
	UPlinkOnlyROHC  *UplinkOnlyROHC
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	ROHCParametersPresentNothing uint64 = iota
	ROHCParametersPresentROHC
	ROHCParametersPresentUPlinkOnlyROHC
	ROHCParametersPresentChoiceExtension
)
