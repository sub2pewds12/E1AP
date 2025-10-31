package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TraceStart is a generated SEQUENCE type.
type TraceStart struct {
	GNBCUCPUEE1APID GNBCUCPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID GNBCUUPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,ext"`
	TraceActivation TraceActivation `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceStart) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUCPUEE1APID), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUCPUEE1APID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUUPUEE1APID), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUUPUEE1APID failed: %w", err)
	}
	if err = s.TraceActivation.Encode(w); err != nil {
		return fmt.Errorf("Encode TraceActivation failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceStart) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
		}
		s.GNBCUCPUEE1APID = GNBCUCPUEE1APID(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
		}
		s.GNBCUUPUEE1APID = GNBCUUPUEE1APID(val)
	}

	if err = s.TraceActivation.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceActivation failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for TraceStart")
	}
	return nil
}
