package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TReordering is a generated ENUMERATED type.
type TReordering struct {
	Value aper.Enumerated
}

const (
	TReorderingMs0    aper.Enumerated = 0
	TReorderingMs1    aper.Enumerated = 1
	TReorderingMs2    aper.Enumerated = 2
	TReorderingMs4    aper.Enumerated = 3
	TReorderingMs5    aper.Enumerated = 4
	TReorderingMs8    aper.Enumerated = 5
	TReorderingMs10   aper.Enumerated = 6
	TReorderingMs15   aper.Enumerated = 7
	TReorderingMs20   aper.Enumerated = 8
	TReorderingMs30   aper.Enumerated = 9
	TReorderingMs40   aper.Enumerated = 10
	TReorderingMs50   aper.Enumerated = 11
	TReorderingMs60   aper.Enumerated = 12
	TReorderingMs80   aper.Enumerated = 13
	TReorderingMs100  aper.Enumerated = 14
	TReorderingMs120  aper.Enumerated = 15
	TReorderingMs140  aper.Enumerated = 16
	TReorderingMs160  aper.Enumerated = 17
	TReorderingMs180  aper.Enumerated = 18
	TReorderingMs200  aper.Enumerated = 19
	TReorderingMs220  aper.Enumerated = 20
	TReorderingMs240  aper.Enumerated = 21
	TReorderingMs260  aper.Enumerated = 22
	TReorderingMs280  aper.Enumerated = 23
	TReorderingMs300  aper.Enumerated = 24
	TReorderingMs500  aper.Enumerated = 25
	TReorderingMs750  aper.Enumerated = 26
	TReorderingMs1000 aper.Enumerated = 27
	TReorderingMs1250 aper.Enumerated = 28
	TReorderingMs1500 aper.Enumerated = 29
	TReorderingMs1750 aper.Enumerated = 30
	TReorderingMs2000 aper.Enumerated = 31
	TReorderingMs2250 aper.Enumerated = 32
	TReorderingMs2500 aper.Enumerated = 33
	TReorderingMs2750 aper.Enumerated = 34
	TReorderingMs3000 aper.Enumerated = 35
)

func (e *TReordering) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TReordering to be generated here.
	return fmt.Errorf("Encode not implemented for enum TReordering")
}

func (e *TReordering) Decode(r *aper.AperReader) error {
	// Decode logic for enum TReordering to be generated here.
	return fmt.Errorf("Decode not implemented for enum TReordering")
}
