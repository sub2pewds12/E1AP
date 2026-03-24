package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DataForwardingInformationRequest is a generated SEQUENCE type.
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest
	QOSFlowsForwardedOnFwdTunnels *QOSFlowMappingList
	IEExtensions                  *DataForwardingInformationRequestExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformationRequest) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-Forwarding-Request", Optional: false},
			per.ComponentInfo{Name: "qoS-Flows-Forwarded-On-Fwd-Tunnels", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.QOSFlowsForwardedOnFwdTunnels != nil {
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.DataForwardingRequest.Value), enumC); err != nil {
			return fmt.Errorf("encode DataForwardingRequest failed: %w", err)
		}
	}

	if s.QOSFlowsForwardedOnFwdTunnels != nil {
		if err = s.QOSFlowsForwardedOnFwdTunnels.Encode(w); err != nil {
			return fmt.Errorf("encode QOSFlowsForwardedOnFwdTunnels failed: %w", err)
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
func (s *DataForwardingInformationRequest) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-Forwarding-Request", Optional: false},
			per.ComponentInfo{Name: "qoS-Flows-Forwarded-On-Fwd-Tunnels", Optional: true},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode DataForwardingRequest failed: %w", err)
		}
		s.DataForwardingRequest.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.QOSFlowsForwardedOnFwdTunnels = new(QOSFlowMappingList)
		if err = s.QOSFlowsForwardedOnFwdTunnels.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSFlowsForwardedOnFwdTunnels failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(DataForwardingInformationRequestExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
