package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CriticalityDiagnosticsIEItem is a generated SEQUENCE type.
type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `aper:"mandatory"`
	IEID          ProtocolIEID `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	TypeOfError   TypeOfError  `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CriticalityDiagnosticsIEItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = s.IECriticality.Encode(w); err != nil {
		return fmt.Errorf("Encode IECriticality failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.IEID.Value), &aper.Constraint{Lb: 0, Ub: MaxProtocolIEs}, false); err != nil {
		return fmt.Errorf("Encode IEID failed: %w", err)
	}
	if err = s.TypeOfError.Encode(w); err != nil {
		return fmt.Errorf("Encode TypeOfError failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CriticalityDiagnosticsIEItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	if err = s.IECriticality.Decode(r); err != nil {
		return fmt.Errorf("Decode IECriticality failed: %w", err)
	}
	if err = s.IEID.Decode(r); err != nil {
		return fmt.Errorf("Decode IEID failed: %w", err)
	}
	if err = s.TypeOfError.Decode(r); err != nil {
		return fmt.Errorf("Decode TypeOfError failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for CriticalityDiagnosticsIEItem")
	}
	return nil
}
