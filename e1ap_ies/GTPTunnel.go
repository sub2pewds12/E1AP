package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// GTPTunnel is a generated SEQUENCE type.
type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress
	GTPTEID               GTPTEID
	IEExtensions          *GTPTunnelExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GTPTunnel) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "transportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "gTP-TEID", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeBitString(s.TransportLayerAddress.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)}); err != nil {
		return fmt.Errorf("encode TransportLayerAddress failed: %w", err)
	}
	if err = w.EncodeOctetString([]byte(s.GTPTEID.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)}); err != nil {
		return fmt.Errorf("encode GTPTEID failed: %w", err)
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
func (s *GTPTunnel) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "transportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "gTP-TEID", Optional: false},
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
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
		if err != nil {
			return fmt.Errorf("decode TransportLayerAddress failed: %w", err)
		}
		s.TransportLayerAddress.Value = val
	}

	{
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
		if err != nil {
			return fmt.Errorf("decode GTPTEID failed: %w", err)
		}
		s.GTPTEID.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(GTPTunnelExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
