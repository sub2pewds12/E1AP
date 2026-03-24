package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// UPSecuritykey is a generated SEQUENCE type.
type UPSecuritykey struct {
	EncryptionKey          EncryptionKey
	IntegrityProtectionKey *IntegrityProtectionKey
	IEExtensions           *UPSecuritykeyExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *UPSecuritykey) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "encryptionKey", Optional: false},
			per.ComponentInfo{Name: "integrityProtectionKey", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IntegrityProtectionKey != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeOctetString([]byte(s.EncryptionKey.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)}); err != nil {
		return fmt.Errorf("encode EncryptionKey failed: %w", err)
	}

	if s.IntegrityProtectionKey != nil {
		if err = w.EncodeOctetString([]byte((*s.IntegrityProtectionKey).Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)}); err != nil {
			return fmt.Errorf("encode IntegrityProtectionKey failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPSecuritykey) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "encryptionKey", Optional: false},
			per.ComponentInfo{Name: "integrityProtectionKey", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	{
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
		if err != nil {
			return fmt.Errorf("decode EncryptionKey failed: %w", err)
		}
		s.EncryptionKey.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IntegrityProtectionKey = new(IntegrityProtectionKey)

		{
			val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
			if err != nil {
				return fmt.Errorf("decode IntegrityProtectionKey failed: %w", err)
			}
			s.IntegrityProtectionKey.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(UPSecuritykeyExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
