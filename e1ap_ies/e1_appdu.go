package e1ap_ies

// E1APPDU is a generated CHOICE type.
type E1APPDU struct {
	Choice              uint64
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}

const (
	E1APPDUPresentNothing uint64 = iota
	E1APPDUPresentInitiatingMessage
	E1APPDUPresentSuccessfulOutcome
	E1APPDUPresentUnsuccessfulOutcome
)
