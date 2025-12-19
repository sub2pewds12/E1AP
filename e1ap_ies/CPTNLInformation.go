package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CPTNLInformation is a generated CHOICE type.
type CPTNLInformation struct {
	Choice            uint64 `json:"-"`
	EndpointIPAddress *TransportLayerAddress
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	CPTNLInformationPresentNothing uint64 = iota
	CPTNLInformationPresentEndpointIPAddress
	CPTNLInformationPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *CPTNLInformation) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE CPTNLInformation | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for CPTNLInformation: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case CPTNLInformationPresentEndpointIPAddress:
		if err = s.EndpointIPAddress.Encode(w); err != nil {
			return fmt.Errorf("Encode EndpointIPAddress failed: %w", err)
		}
	case CPTNLInformationPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of CPTNLInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CPTNLInformation) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.EndpointIPAddress = new(TransportLayerAddress)
		if err = s.EndpointIPAddress.Decode(r); err != nil {
			return fmt.Errorf("Decode EndpointIPAddress failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of CPTNLInformation with unknown choice index %d", choice)
	}
	return nil
}
