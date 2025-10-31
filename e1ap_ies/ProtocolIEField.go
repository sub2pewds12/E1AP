package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ProtocolIEField is a generated SEQUENCE type.
type ProtocolIEField struct {
	ID          ProtocolIEID     `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	Criticality Criticality      `aper:"mandatory"`
	Value       aper.OctetString `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolIEField) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ID), &aper.Constraint{Lb: 0, Ub: maxProtocolIEs}, false); err != nil {
		return fmt.Errorf("Encode ID failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.Criticality.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode Criticality failed: %w", err)
	}
	if err = s.Value.Encode(w); err != nil {
		return fmt.Errorf("Encode Value failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIEField) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxProtocolIEs}, false); err != nil {
			return fmt.Errorf("Decode ID failed: %w", err)
		}
		s.ID = ProtocolIEID(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("Decode Criticality failed: %w", err)
		}
		s.Criticality.Value = aper.Enumerated(val)
	}

	{
		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode Value failed: %w", err)
		}
		s.Value = aper.OctetString(val)
	}

	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ProtocolIEField")
	}
	return nil
}
