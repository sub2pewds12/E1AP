package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataUsageReport is a generated SEQUENCE type.
type DataUsageReport struct {
	GNBCUCPUEE1APID     GNBCUCPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID     GNBCUUPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DataUsageReportList []DataUsageReportItem `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataUsageReport) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUCPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUCPUEE1APID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUUPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUUPUEE1APID failed: %w", err)
	}
	if err = s.DataUsageReportList.Encode(w); err != nil {
		return fmt.Errorf("Encode DataUsageReportList failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataUsageReport) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
		}
		s.GNBCUCPUEE1APID.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
		}
		s.GNBCUUPUEE1APID.Value = aper.Integer(val)
	}
	if err = s.DataUsageReportList.Decode(r); err != nil {
		return fmt.Errorf("Decode DataUsageReportList failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DataUsageReport")
	}
	return nil
}
