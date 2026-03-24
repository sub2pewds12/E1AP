package e1ap_ies

import (
	"asn1go/per"
)

// ReportingPeriodicity is a generated ENUMERATED type.
type ReportingPeriodicity struct {
	Value int64
}

const (
	ReportingPeriodicityMs500    int64 = 0
	ReportingPeriodicityMs1000   int64 = 1
	ReportingPeriodicityMs2000   int64 = 2
	ReportingPeriodicityMs5000   int64 = 3
	ReportingPeriodicityMs10000  int64 = 4
	ReportingPeriodicityMs20000  int64 = 5
	ReportingPeriodicityMs30000  int64 = 6
	ReportingPeriodicityMs40000  int64 = 7
	ReportingPeriodicityMs50000  int64 = 8
	ReportingPeriodicityMs60000  int64 = 9
	ReportingPeriodicityMs70000  int64 = 10
	ReportingPeriodicityMs80000  int64 = 11
	ReportingPeriodicityMs90000  int64 = 12
	ReportingPeriodicityMs100000 int64 = 13
	ReportingPeriodicityMs110000 int64 = 14
	ReportingPeriodicityMs120000 int64 = 15
)

// Encode implements the MessageEncoder interface for ReportingPeriodicity.
func (e *ReportingPeriodicity) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 16), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for ReportingPeriodicity.
func (e *ReportingPeriodicity) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 16), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
