package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPE1SetupRequest is a generated SEQUENCE type.
type GNBCUUPE1SetupRequest struct {
	TransactionID             aper.Integer               `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUUPID                 aper.Integer               `aper:"lb:0,ub:68719476735,mandatory,reject,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	CNSupport                 CNSupport                  `aper:"mandatory,reject,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,mandatory,reject,ext"`
	GNBCUUPCapacity           *aper.Integer              `aper:"lb:0,ub:255,optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPE1SetupRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUUPE1SetupRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPE1SetupRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUUPE1SetupRequest")
}
