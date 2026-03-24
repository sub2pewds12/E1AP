package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// IE is the interface that all encodable/decodable types must implement.
type IE interface {
	Encode(w *per.Encoder) error
	Decode(r *per.Decoder) error
}

type E1APMessageIE struct {
	ID          ProtocolIEID
	Criticality Criticality
	Value       IE
}

type ProtocolIEField = E1APMessageIE
type ProtocolIESingleContainer = E1APMessageIE
type ProtocolIEContainer = []E1APMessageIE
type ProtocolIEContainerList = []ProtocolIEContainer
type PrivateIEContainer = []ProtocolExtensionField

func (ie *E1APMessageIE) Encode(w *per.Encoder) error {
	if err := w.EncodeInteger(ie.ID.Value, per.Constrained(0, 65535)); err != nil {
		return fmt.Errorf("encode IE ID failed: %w", err)
	}
	if err := ie.Criticality.Encode(w); err != nil {
		return fmt.Errorf("encode IE Criticality failed: %w", err)
	}
	valEncoder := per.NewEncoder(per.APER)
	if err := ie.Value.Encode(valEncoder); err != nil {
		return fmt.Errorf("encode IE Value failed: %w", err)
	}
	if err := w.EncodeOctetString(valEncoder.Bytes(), per.SizeConstraints{Extensible: false, Min: nil, Max: nil}); err != nil {
		return fmt.Errorf("encode IE Value as OpenType failed: %w", err)
	}
	return nil
}

func (ie *E1APMessageIE) Decode(r *per.Decoder) error {
	return fmt.Errorf("E1APMessageIE.Decode should not be called directly")
}

type INTEGER struct {
	Value int64
	c     per.IntegerConstraints
}

func NewINTEGER(v int64, c per.IntegerConstraints) INTEGER {
	return INTEGER{Value: v, c: c}
}
func (t *INTEGER) Encode(w *per.Encoder) (err error) { return w.EncodeInteger(t.Value, t.c) }
func (t *INTEGER) Decode(r *per.Decoder) (err error) {
	v, err := r.DecodeInteger(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type OCTETSTRING struct {
	Value []byte
	c     per.SizeConstraints
}

func NewOCTETSTRING(v []byte, c per.SizeConstraints) OCTETSTRING {
	return OCTETSTRING{Value: v, c: c}
}
func (t *OCTETSTRING) Encode(w *per.Encoder) (err error) { return w.EncodeOctetString(t.Value, t.c) }
func (t *OCTETSTRING) Decode(r *per.Decoder) (err error) {
	v, err := r.DecodeOctetString(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type BITSTRING struct {
	Value per.BitString
	c     per.SizeConstraints
}

func NewBITSTRING(v per.BitString, c per.SizeConstraints) BITSTRING {
	return BITSTRING{Value: v, c: c}
}
func (t *BITSTRING) Encode(w *per.Encoder) (err error) { return w.EncodeBitString(t.Value, t.c) }
func (t *BITSTRING) Decode(r *per.Decoder) (err error) {
	v, err := r.DecodeBitString(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type ENUMERATED struct {
	Value int64
	c     per.EnumeratedConstraints
}

func NewENUMERATED(v int64, c per.EnumeratedConstraints) ENUMERATED {
	return ENUMERATED{Value: v, c: c}
}
func (t *ENUMERATED) Encode(w *per.Encoder) (err error) { return w.EncodeEnumerated(t.Value, t.c) }
func (t *ENUMERATED) Decode(r *per.Decoder) (err error) {
	v, err := r.DecodeEnumerated(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type Sequence[T IE] struct {
	Value []T
	c     per.SizeConstraints
}

func NewSequence[T IE](items []T, c per.SizeConstraints) Sequence[T] {
	return Sequence[T]{Value: items, c: c}
}

func (s *Sequence[T]) Encode(w *per.Encoder) (err error) {
	if err = w.EncodeLengthDeterminant(int64(len(s.Value)), s.c); err != nil {
		return fmt.Errorf("encode sequence length failed: %w", err)
	}
	for i := range s.Value {
		if err = s.Value[i].Encode(w); err != nil {
			return fmt.Errorf("encode sequence item %d failed: %w", i, err)
		}
	}
	return
}

func (s *Sequence[T]) Decode(r *per.Decoder) (err error) {
	return fmt.Errorf("decode for generic Sequence[T] should not be called directly")
}

type ProtocolExtensionContainer struct {
	Value []ProtocolExtensionField
}

func (s *ProtocolExtensionContainer) Encode(w *per.Encoder) (err error) { return nil }
func (s *ProtocolExtensionContainer) Decode(r *per.Decoder) (err error) { return nil }

type ProtocolExtensionField struct {
	ID             ProtocolIEID
	Criticality    Criticality
	ExtensionValue IE
	ValueBytes     []byte
}

func (s *ProtocolExtensionField) Encode(w *per.Encoder) (err error) {
	if err = w.EncodeInteger(s.ID.Value, per.Constrained(0, 65535)); err != nil {
		return fmt.Errorf("encode ProtocolExtensionField ID failed: %w", err)
	}
	if err = s.Criticality.Encode(w); err != nil {
		return fmt.Errorf("encode ProtocolExtensionField Criticality failed: %w", err)
	}

	valEncoder := per.NewEncoder(per.APER)
	if err = s.ExtensionValue.Encode(valEncoder); err != nil {
		return fmt.Errorf("encode ProtocolExtensionField ExtensionValue failed: %w", err)
	}

	if err = w.EncodeOctetString(valEncoder.Bytes(), per.SizeConstraints{Extensible: false, Min: nil, Max: nil}); err != nil {
		return fmt.Errorf("encode ProtocolExtensionField ExtensionValue as OpenType failed: %w", err)
	}
	return nil
}

func (s *ProtocolExtensionField) Decode(r *per.Decoder) (err error) {
	val, err := r.DecodeInteger(per.Constrained(0, 65535))
	if err != nil {
		return fmt.Errorf("decode ProtocolExtensionField ID failed: %w", err)
	}
	s.ID = ProtocolIEID{Value: val}
	if err = s.Criticality.Decode(r); err != nil {
		return fmt.Errorf("decode ProtocolExtensionField Criticality failed: %w", err)
	}
	if s.ValueBytes, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: nil, Max: nil}); err != nil {
		return fmt.Errorf("decode ProtocolExtensionField ExtensionValue as OpenType failed: %w", err)
	}
	return nil
}
