package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PrivateIEID is a generated CHOICE type.
type PrivateIEID struct {
	Choice uint64 `json:"-"`
	Local  *INTEGER
	Global *ObjectIdentifier
}

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

// Encode implements the aper.AperMarshaller interface.
func (s *PrivateIEID) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case PrivateIEIDPresentLocal:
		if err = s.Local.Encode(w); err != nil {
			return fmt.Errorf("Encode Local failed: %w", err)
		}
	case PrivateIEIDPresentGlobal:
		if err = s.Global.Encode(w); err != nil {
			return fmt.Errorf("Encode Global failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of PrivateIEID with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PrivateIEID) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		s.Local = new(INTEGER)
		if err = s.Local.Decode(r); err != nil {
			return fmt.Errorf("Decode Local failed: %w", err)
		}
	case 1:
		s.Global = new(string)
		if err = s.Global.Decode(r); err != nil {
			return fmt.Errorf("Decode Global failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of PrivateIEID with unknown choice index %d", choice)
	}
	return nil
}
