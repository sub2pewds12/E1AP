package e1ap_ies

// ResetType is a generated CHOICE type.
type ResetType struct {
	Choice            uint64
	E1Interface       *ResetAll
	PartOfE1Interface *UEAssociatedLogicalE1ConnectionListRes
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	ResetTypePresentNothing uint64 = iota
	ResetTypePresentE1Interface
	ResetTypePresentPartOfE1Interface
	ResetTypePresentChoiceExtension
)
