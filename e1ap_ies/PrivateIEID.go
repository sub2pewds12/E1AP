package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PrivateIEID is a generated CHOICE type.
type PrivateIEID struct {
	Choice uint64 `json:"-"`
	Local  *INTEGER
	Global *string
}

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

// Encode implements the aper.AperMarshaller interface.
func (s *PrivateIEID) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE PrivateIEID | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for PrivateIEID: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case PrivateIEIDPresentLocal:
		if err = s.Local.Encode(w); err != nil {
			return fmt.Errorf("encode Local failed: %w", err)
		}
	case PrivateIEIDPresentGlobal:
		if err = w.WriteOctetString([]byte(*s.Global), nil, false); err != nil {
			return fmt.Errorf("encode Global failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of PrivateIEID with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PrivateIEID) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(1, false)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.Local = new(INTEGER)
		if err = s.Local.Decode(r); err != nil {
			return fmt.Errorf("decode Local failed: %w", err)
		}
	case 2:
		val, err := r.ReadOctetString(nil, false)
		if err != nil {
			return fmt.Errorf("decode Global failed: %w", err)
		}
		tmpStr := string(val)
		s.Global = &tmpStr
	default:
		return fmt.Errorf("decode choice of PrivateIEID with unknown choice index %d", choice)
	}
	return nil
}
