package e1ap_ies

// ResetType From: 9_4_4_PDU_Definitions.txt:261
const (
	ResetTypePresentNothing uint64 = iota
	ResetTypePresentE1Interface
	ResetTypePresentPartOfE1Interface
)

type ResetType struct {
	Choice            uint64
	E1Interface       *ResetAll
	PartOfE1Interface []ProtocolIESingleContainer
}
