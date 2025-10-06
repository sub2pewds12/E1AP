package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DiscardTimer is a generated ENUMERATED type.
type DiscardTimer struct {
	Value aper.Enumerated
}

const (
	DiscardTimerMs10     aper.Enumerated = 0
	DiscardTimerMs20     aper.Enumerated = 1
	DiscardTimerMs30     aper.Enumerated = 2
	DiscardTimerMs40     aper.Enumerated = 3
	DiscardTimerMs50     aper.Enumerated = 4
	DiscardTimerMs60     aper.Enumerated = 5
	DiscardTimerMs75     aper.Enumerated = 6
	DiscardTimerMs100    aper.Enumerated = 7
	DiscardTimerMs150    aper.Enumerated = 8
	DiscardTimerMs200    aper.Enumerated = 9
	DiscardTimerMs250    aper.Enumerated = 10
	DiscardTimerMs300    aper.Enumerated = 11
	DiscardTimerMs500    aper.Enumerated = 12
	DiscardTimerMs750    aper.Enumerated = 13
	DiscardTimerMs1500   aper.Enumerated = 14
	DiscardTimerInfinity aper.Enumerated = 15
)
