package e1ap_ies

// MDTMode is a generated CHOICE type.
type MDTMode struct {
	Choice       uint64
	ImmediateMDT *ImmediateMDT
}

const (
	MDTModePresentNothing uint64 = iota
	MDTModePresentImmediateMDT
)
