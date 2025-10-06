package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// CellTrafficTrace is a generated SEQUENCE type.
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                aper.Integer      `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                aper.Integer      `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceID                        aper.OctetString  `aper:"lb:8,ub:8,mandatory,ignore,ext"`
	TraceCollectionEntityIPAddress aper.BitString    `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	PrivacyIndicator               *PrivacyIndicator `aper:"optional,ignore,ext"`
	URIaddress                     *aper.OctetString `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellTrafficTrace) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for CellTrafficTrace")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CellTrafficTrace) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for CellTrafficTrace")
}
