package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// TransportLayerAddressInfo is a generated SEQUENCE type.
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    *TransportUPLayerAddressesInfoToAddList
	TransportUPLayerAddressesInfoToRemoveList *TransportUPLayerAddressesInfoToRemoveList
	IEExtensions                              *TransportLayerAddressInfoExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportLayerAddressInfo) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "transport-UP-Layer-Addresses-Info-To-Add-List", Optional: true},
			per.ComponentInfo{Name: "transport-UP-Layer-Addresses-Info-To-Remove-List", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.TransportUPLayerAddressesInfoToAddList != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.TransportUPLayerAddressesInfoToRemoveList != nil {
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

	if s.TransportUPLayerAddressesInfoToAddList != nil {
		if err = s.TransportUPLayerAddressesInfoToAddList.Encode(w); err != nil {
			return fmt.Errorf("encode TransportUPLayerAddressesInfoToAddList failed: %w", err)
		}
	}

	if s.TransportUPLayerAddressesInfoToRemoveList != nil {
		if err = s.TransportUPLayerAddressesInfoToRemoveList.Encode(w); err != nil {
			return fmt.Errorf("encode TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
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
func (s *TransportLayerAddressInfo) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "transport-UP-Layer-Addresses-Info-To-Add-List", Optional: true},
			per.ComponentInfo{Name: "transport-UP-Layer-Addresses-Info-To-Remove-List", Optional: true},
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

	if seqDecoder.IsComponentPresent(0) {
		s.TransportUPLayerAddressesInfoToAddList = new(TransportUPLayerAddressesInfoToAddList)
		if err = s.TransportUPLayerAddressesInfoToAddList.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportUPLayerAddressesInfoToAddList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.TransportUPLayerAddressesInfoToRemoveList = new(TransportUPLayerAddressesInfoToRemoveList)
		if err = s.TransportUPLayerAddressesInfoToRemoveList.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(TransportLayerAddressInfoExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
