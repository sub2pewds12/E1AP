package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TransportUPLayerAddressesInfoToRemoveList is a generated LIST type.
type TransportUPLayerAddressesInfoToRemoveList struct {
	Value []TransportUPLayerAddressesInfoToRemoveItem
}

func (s *TransportUPLayerAddressesInfoToRemoveList) Encode(w *aper.AperWriter) (err error) {

	// 1. Create a temporary slice of the INTERFACE type.
	itemPointers := make([]aper.AperMarshaller, len(s.Value)) // <-- FIX: Changed from []*aper.AperMarshaller
	for i := 0; i < len(s.Value); i++ {
		// Assigning a pointer-to-struct to an interface value is correct.
		itemPointers[i] = &s.Value[i]
	}

	// 2. Call the generic WriteSequenceOf helper with the slice of interfaces.
	if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return fmt.Errorf("WriteSequenceOf for TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
	}
	return nil
}

func (s *TransportUPLayerAddressesInfoToRemoveList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*TransportUPLayerAddressesInfoToRemoveItem, error) {
		item := new(TransportUPLayerAddressesInfoToRemoveItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []TransportUPLayerAddressesInfoToRemoveItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return fmt.Errorf("ReadSequenceOf for TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
