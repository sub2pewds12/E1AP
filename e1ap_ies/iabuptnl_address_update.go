package e1ap_ies

import (
	"fmt"
	"io"
)

// IABUPTNLAddressUpdate is a generated SEQUENCE type.
type IABUPTNLAddressUpdate struct {
	TransactionID              TransactionID                `aper:"lb:0,ub:255,mandatory,reject,ext"`
	DLUPTNLAddressToUpdateList []DLUPTNLAddressToUpdateItem `aper:"ub:MaxnoofTNLAddresses,optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *IABUPTNLAddressUpdate) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for IABUPTNLAddressUpdate")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *IABUPTNLAddressUpdate) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for IABUPTNLAddressUpdate")
}
