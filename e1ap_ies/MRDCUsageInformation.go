package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// MRDCUsageInformation is a generated SEQUENCE type.
type MRDCUsageInformation struct {
	DataUsagePerPDUSessionReport *DataUsagePerPDUSessionReport
	DataUsagePerQOSFlowList      *DataUsagePerQOSFlowList
	IEExtensions                 *MRDCUsageInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCUsageInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-Usage-per-PDU-Session-Report", Optional: true},
			per.ComponentInfo{Name: "data-Usage-per-QoS-Flow-List", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.DataUsagePerPDUSessionReport != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DataUsagePerQOSFlowList != nil {
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

	if s.DataUsagePerPDUSessionReport != nil {
		if err = s.DataUsagePerPDUSessionReport.Encode(w); err != nil {
			return fmt.Errorf("encode DataUsagePerPDUSessionReport failed: %w", err)
		}
	}

	if s.DataUsagePerQOSFlowList != nil {
		if err = s.DataUsagePerQOSFlowList.Encode(w); err != nil {
			return fmt.Errorf("encode DataUsagePerQOSFlowList failed: %w", err)
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
func (s *MRDCUsageInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "data-Usage-per-PDU-Session-Report", Optional: true},
			per.ComponentInfo{Name: "data-Usage-per-QoS-Flow-List", Optional: true},
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
		s.DataUsagePerPDUSessionReport = new(DataUsagePerPDUSessionReport)
		if err = s.DataUsagePerPDUSessionReport.Decode(r); err != nil {
			return fmt.Errorf("Decode DataUsagePerPDUSessionReport failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.DataUsagePerQOSFlowList = new(DataUsagePerQOSFlowList)
		if err = s.DataUsagePerQOSFlowList.Decode(r); err != nil {
			return fmt.Errorf("Decode DataUsagePerQOSFlowList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(MRDCUsageInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
