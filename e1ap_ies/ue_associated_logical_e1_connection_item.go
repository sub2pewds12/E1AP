package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// UEAssociatedLogicalE1ConnectionItem is a generated SEQUENCE type.
type UEAssociatedLogicalE1ConnectionItem struct {
	GNBCUCPUEE1APID *aper.Integer `aper:"lb:0,ub:4294967295,optional,reject,ext"`
	GNBCUUPUEE1APID *aper.Integer `aper:"lb:0,ub:4294967295,optional,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for UEAssociatedLogicalE1ConnectionItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for UEAssociatedLogicalE1ConnectionItem")
}
