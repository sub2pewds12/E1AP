package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ProtocolIEField is a generated SEQUENCE type.
type ProtocolIEField struct {
	ID          ProtocolIEID `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	Criticality Criticality  `aper:"mandatory"`
	Value       aper.IE      `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolIEField) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ID.Value), &aper.Constraint{Lb: 0, Ub: MaxProtocolIEs}, false); err != nil {
		return fmt.Errorf("encode ID failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.Criticality.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode Criticality failed: %w", err)
	}
	if err = s.Value.Encode(w); err != nil {
		return fmt.Errorf("encode Value failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIEField) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if err = s.ID.Decode(r); err != nil {
		return fmt.Errorf("decode ID failed: %w", err)
	}
	if err = s.Criticality.Decode(r); err != nil {
		return fmt.Errorf("decode Criticality failed: %w", err)
	}
	if err = s.Value.Decode(r); err != nil {
		return fmt.Errorf("decode Value failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for ProtocolIEField */
	}
	return nil
}
