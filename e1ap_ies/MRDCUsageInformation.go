package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// MRDCUsageInformation is a generated SEQUENCE type.
type MRDCUsageInformation struct {
	DataUsagePerPDUSessionReport *DataUsagePerPDUSessionReport `aper:"lb:1,ub:Maxnooftimeperiods,optional,ext"`
	DataUsagePerQOSFlowList      *DataUsagePerQOSFlowList      `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer   `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCUsageInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DataUsagePerPDUSessionReport != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DataUsagePerQOSFlowList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
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
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCUsageInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DataUsagePerPDUSessionReport = new(DataUsagePerPDUSessionReport)
		if err = s.DataUsagePerPDUSessionReport.Decode(r); err != nil {
			return fmt.Errorf("decode DataUsagePerPDUSessionReport failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DataUsagePerQOSFlowList = new(DataUsagePerQOSFlowList)
		if err = s.DataUsagePerQOSFlowList.Decode(r); err != nil {
			return fmt.Errorf("decode DataUsagePerQOSFlowList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for MRDCUsageInformation */
	}
	return nil
}
