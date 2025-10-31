package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper" // NOTE: Ensure this is the correct import path for YOUR aper library
)

// =================================================================================
// == IE Container (Manually Written)
// =================================================================================

// E1APMessageIE is the generic container for all Information Elements in a PDU.
// This is a foundational type for the PDU encoding/decoding pattern.
type E1APMessageIE struct {
	Id          ProtocolIEID
	Criticality Criticality
	Value       aper.IE // This interface requires all values to have Encode/Decode methods
}

// =================================================================================
// == Primitive Type Wrappers (Manually Written)
// =================================================================================
// These structs are used as temporary wrappers inside the toIes and decodeIE
// methods to bundle a primitive value with its encoding constraints.

// --- INTEGER WRAPPER ---
type INTEGER struct {
	Value aper.Integer
	c     aper.Constraint
	ext   bool
}

func NewINTEGER(v int64, c aper.Constraint, ext bool) INTEGER {
	return INTEGER{
		Value: aper.Integer(v),
		c:     c,
		ext:   ext,
	}
}
func (t *INTEGER) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(t.Value), &t.c, t.ext)
	return
}
func (t *INTEGER) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadInteger(&t.c, t.ext)
	t.Value = aper.Integer(v)
	return
}

// --- OCTETSTRING WRAPPER ---
type OCTETSTRING struct {
	Value aper.OctetString
	c     aper.Constraint
	ext   bool
}

func NewOCTETSTRING(v []byte, c aper.Constraint, ext bool) OCTETSTRING {
	return OCTETSTRING{
		Value: v,
		c:     c,
		ext:   ext,
	}
}
func (t *OCTETSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		err = w.WriteOctetString(t.Value, nil, t.ext)
	} else {
		err = w.WriteOctetString(t.Value, &t.c, t.ext)
	}
	return
}
func (t *OCTETSTRING) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		if v, err = r.ReadOctetString(nil, t.ext); err != nil {
			return
		}
	} else {
		if v, err = r.ReadOctetString(&t.c, t.ext); err != nil {
			return
		}
	}
	t.Value = v
	return
}

// --- BITSTRING WRAPPER ---
type BITSTRING struct {
	Value aper.BitString
	c     aper.Constraint
	ext   bool
}

func NewBITSTRING(v aper.BitString, c aper.Constraint, ext bool) BITSTRING {
	return BITSTRING{
		Value: aper.BitString{
			Bytes:   v.Bytes,
			NumBits: v.NumBits,
		},
		c:   c,
		ext: ext,
	}
}
func (t *BITSTRING) Encode(w *aper.AperWriter) (err error) {
	// If NumBits is not set, assume the full length of the byte slice.
	if t.Value.NumBits == 0 && len(t.Value.Bytes) > 0 {
		t.Value.NumBits = uint64(len(t.Value.Bytes) * 8)
	}
	err = w.WriteBitString(t.Value.Bytes, uint(t.Value.NumBits), &t.c, t.ext)
	return
}
func (t *BITSTRING) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&t.c, t.ext); err != nil {
		return
	}
	t.Value.Bytes = v
	t.Value.NumBits = uint64(n)
	return
}

// --- ENUMERATED WRAPPER ---
type ENUMERATED struct {
	Value aper.Enumerated
	c     aper.Constraint
	ext   bool
}

func NewENUMERATED(v int64, c aper.Constraint, ext bool) ENUMERATED {
	return ENUMERATED{
		Value: aper.Enumerated(v),
		c:     c,
		ext:   ext,
	}
}
func (t *ENUMERATED) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(t.Value), t.c, t.ext)
	return
}
func (t *ENUMERATED) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(t.c, t.ext)
	t.Value = aper.Enumerated(v)
	return
}

// =================================================================================
// == Generic SEQUENCE OF / List Wrapper (Manually Written)
// =================================================================================

type Sequence[T aper.IE] struct {
	Value []T
	c     aper.Constraint
	ext   bool
}

func NewSequence[T aper.IE](items []T, c aper.Constraint, ext bool) Sequence[T] {
	return Sequence[T]{
		Value: items,
		c:     c,
		ext:   ext,
	}
}

func (s *Sequence[T]) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[T](s.Value, w, &s.c, s.ext); err != nil {
		return
	}
	return
}

func (s *Sequence[T]) Decode(r *AperReader) (err error) {
	// This function now has the correct signature but is a placeholder.
	// Decoding of lists will be handled specially inside the decodeIE switch case.
	return fmt.Errorf("Decode for generic Sequence[T] should not be called directly. Use the specific list decoder logic.")
}

type ProtocolExtensionContainer struct {
	Value []ProtocolExtensionField
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolExtensionContainer) Encode(w *aper.AperWriter) (err error) {
	// This is a complex encoding process involving a bitmap and open types.
	// For now, a placeholder is acceptable to get the code compiling.
	// TODO: Implement the real APER extension container encoding logic.
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolExtensionContainer) Decode(r *aper.AperReader) (err error) {
	// TODO: Implement the real APER extension container decoding logic.
	return nil
}

// ProtocolExtensionField is the item within the container.
type ProtocolExtensionField struct {
	Id             ProtocolIEID
	Criticality    Criticality
	ExtensionValue aper.IE
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolExtensionField) Encode(w *aper.AperWriter) (err error) {
	// TODO: Implement the real APER encoding logic for an extension field.
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolExtensionField) Decode(r *aper.AperReader) (err error) {
	// TODO: Implement the real APER decoding logic for an extension field.
	return nil
}
