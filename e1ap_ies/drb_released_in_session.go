package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBReleasedInSession From: manual_patch:-1
const (
	DRBReleasedInSessionReleasedInSession    aper.Enumerated = 0
	DRBReleasedInSessionNotReleasedInSession aper.Enumerated = 1
)

type DRBReleasedInSession struct {
	Value aper.Enumerated
}
