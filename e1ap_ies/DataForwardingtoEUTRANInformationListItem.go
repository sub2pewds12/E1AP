package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DataForwardingtoEUTRANInformationListItem is a generated SEQUENCE type.
type DataForwardingtoEUTRANInformationListItem struct {
	DataForwardingTunnelInformation UPTNLInformation
	QOSFlowsToBeForwardedList       QOSFlowsToBeForwardedList
	IEExtensions                    *DataForwardingtoEUTRANInformationListItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingtoEUTRANInformationListItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-forwarding-tunnel-information", Optional: false},
			per.ComponentInfo{Name: "qoS-Flows-to-be-forwarded-List", Optional: false},
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

	if err = s.DataForwardingTunnelInformation.Encode(w); err != nil {
		return fmt.Errorf("encode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Encode(w); err != nil {
		return fmt.Errorf("encode QOSFlowsToBeForwardedList failed: %w", err)
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
func (s *DataForwardingtoEUTRANInformationListItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-forwarding-tunnel-information", Optional: false},
			per.ComponentInfo{Name: "qoS-Flows-to-be-forwarded-List", Optional: false},
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

	if err = s.DataForwardingTunnelInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSFlowsToBeForwardedList failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(DataForwardingtoEUTRANInformationListItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
