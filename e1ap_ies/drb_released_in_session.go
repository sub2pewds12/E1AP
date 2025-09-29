package e1ap_ies

// DRBReleasedInSession represents the ASN.1 definition from manual_patch:-1
type DRBReleasedInSession int32

const (
	DRBReleasedInSession_ReleasedInSession    DRBReleasedInSession = 0
	DRBReleasedInSession_NotReleasedInSession DRBReleasedInSession = 1
)
