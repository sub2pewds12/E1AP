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
	if err = w.WriteEnumerate(uint64(s.IECriticality.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode IECriticality failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.IEID), &aper.Constraint{Lb: 0, Ub: maxProtocolIEs}, false); err != nil {
		return fmt.Errorf("Encode IEID failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.TypeOfError.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
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

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("Decode IECriticality failed: %w", err)
		}
		s.IECriticality.Value = aper.Enumerated(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxProtocolIEs}, false); err != nil {
			return fmt.Errorf("Decode IEID failed: %w", err)
		}
		s.IEID = ProtocolIEID(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Decode TypeOfError failed: %w", err)
		}
		s.TypeOfError.Value = aper.Enumerated(val)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for CriticalityDiagnosticsIEItem")
	}
	return nil
}
