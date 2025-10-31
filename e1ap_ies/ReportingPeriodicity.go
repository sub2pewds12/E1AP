package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ReportingPeriodicity is a generated ENUMERATED type.
type ReportingPeriodicity struct {
	Value aper.Enumerated
}

const (
	ReportingPeriodicityMs500    aper.Enumerated = 0
	ReportingPeriodicityMs1000   aper.Enumerated = 1
	ReportingPeriodicityMs2000   aper.Enumerated = 2
	ReportingPeriodicityMs5000   aper.Enumerated = 3
	ReportingPeriodicityMs10000  aper.Enumerated = 4
	ReportingPeriodicityMs20000  aper.Enumerated = 5
	ReportingPeriodicityMs30000  aper.Enumerated = 6
	ReportingPeriodicityMs40000  aper.Enumerated = 7
	ReportingPeriodicityMs50000  aper.Enumerated = 8
	ReportingPeriodicityMs60000  aper.Enumerated = 9
	ReportingPeriodicityMs70000  aper.Enumerated = 10
	ReportingPeriodicityMs80000  aper.Enumerated = 11
	ReportingPeriodicityMs90000  aper.Enumerated = 12
	ReportingPeriodicityMs100000 aper.Enumerated = 13
	ReportingPeriodicityMs110000 aper.Enumerated = 14
	ReportingPeriodicityMs120000 aper.Enumerated = 15
)

func (e *ReportingPeriodicity) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ReportingPeriodicity to be generated here.
	return fmt.Errorf("Encode not implemented for enum ReportingPeriodicity")
}

func (e *ReportingPeriodicity) Decode(r *aper.AperReader) error {
	// Decode logic for enum ReportingPeriodicity to be generated here.
	return fmt.Errorf("Decode not implemented for enum ReportingPeriodicity")
}
