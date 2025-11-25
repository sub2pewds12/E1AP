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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.DataUsagePerPDUSessionReport != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DataUsagePerPDUSessionReport.Value))
			for i := 0; i < len(s.DataUsagePerPDUSessionReport.Value); i++ {
				itemPointers[i] = &(s.DataUsagePerPDUSessionReport.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: Maxnooftimeperiods}, true); err != nil {
				return fmt.Errorf("Encode DataUsagePerPDUSessionReport failed: %w", err)
			}
		}
	}
	if s.DataUsagePerQOSFlowList != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DataUsagePerQOSFlowList.Value))
			for i := 0; i < len(s.DataUsagePerQOSFlowList.Value); i++ {
				itemPointers[i] = &(s.DataUsagePerQOSFlowList.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows}, false); err != nil {
				return fmt.Errorf("Encode DataUsagePerQOSFlowList failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCUsageInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DataUsagePerPDUSessionReport = new(DataUsagePerPDUSessionReport)
		if err = s.DataUsagePerPDUSessionReport.Decode(r); err != nil {
			return fmt.Errorf("Decode DataUsagePerPDUSessionReport failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DataUsagePerQOSFlowList = new(DataUsagePerQOSFlowList)
		if err = s.DataUsagePerQOSFlowList.Decode(r); err != nil {
			return fmt.Errorf("Decode DataUsagePerQOSFlowList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for MRDCUsageInformation */
	}
	return nil
}
