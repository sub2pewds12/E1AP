package e1ap_ies

import (
	"fmt"
	"io"
)

// MRDCDataUsageReport is a generated SEQUENCE type.
type MRDCDataUsageReport struct {
	GNBCUCPUEE1APID                 GNBCUCPUEE1APID                   `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                 GNBCUUPUEE1APID                   `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PDUSessionResourceDataUsageList []PDUSessionResourceDataUsageItem `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReport) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for MRDCDataUsageReport")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReport) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for MRDCDataUsageReport")
}
