package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// EndpointIPAddressAndPort is a generated SEQUENCE type.
type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress
	PortNumber        PortNumber
	IEExtensions      *EndpointIPAddressAndPortExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *EndpointIPAddressAndPort) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "endpoint-IP-Address", Optional: false},
			per.ComponentInfo{Name: "portNumber", Optional: false},
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

	if err = w.EncodeBitString(s.EndpointIPAddress.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)}); err != nil {
		return fmt.Errorf("encode EndpointIPAddress failed: %w", err)
	}
	if err = w.EncodeBitString(s.PortNumber.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(16), Max: int64Ptr(16)}); err != nil {
		return fmt.Errorf("encode PortNumber failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EndpointIPAddressAndPort) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "endpoint-IP-Address", Optional: false},
			per.ComponentInfo{Name: "portNumber", Optional: false},
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
			return fmt.Errorf("decode EndpointIPAddress failed: %w", err)
		}
		s.EndpointIPAddress.Value = val
	}

	{
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(16), Max: int64Ptr(16)})
		if err != nil {
			return fmt.Errorf("decode PortNumber failed: %w", err)
		}
		s.PortNumber.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(EndpointIPAddressAndPortExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
