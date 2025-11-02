package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdate is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdate struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ext"`
	GNBCUCPTNLAToAddList      []GNBCUCPTNLAToAddItem     `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAToRemoveList   []GNBCUCPTNLAToRemoveItem  `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAToUpdateList   []GNBCUCPTNLAToUpdateItem  `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUCPName       *ExtendedGNBCUCPName       `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPConfigurationUpdate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUCPName != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.GNBCUCPTNLAToAddList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.GNBCUCPTNLAToRemoveList != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.GNBCUCPTNLAToUpdateList != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.TransportLayerAddressInfo != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.ExtendedGNBCUCPName != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(6), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.TransactionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if s.GNBCUCPName != nil {
		if err = w.WriteOctetString([]byte(s.GNBCUCPName.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode GNBCUCPName failed: %w", err)
		}
	}
	if s.GNBCUCPTNLAToAddList != nil {
		if err = s.GNBCUCPTNLAToAddList.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUCPTNLAToAddList failed: %w", err)
		}
	}
	if s.GNBCUCPTNLAToRemoveList != nil {
		if err = s.GNBCUCPTNLAToRemoveList.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUCPTNLAToRemoveList failed: %w", err)
		}
	}
	if s.GNBCUCPTNLAToUpdateList != nil {
		if err = s.GNBCUCPTNLAToUpdateList.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUCPTNLAToUpdateList failed: %w", err)
		}
	}
	if s.TransportLayerAddressInfo != nil {
		if err = s.TransportLayerAddressInfo.Encode(w); err != nil {
			return fmt.Errorf("Encode TransportLayerAddressInfo failed: %w", err)
		}
	}
	if s.ExtendedGNBCUCPName != nil {
		if err = s.ExtendedGNBCUCPName.Encode(w); err != nil {
			return fmt.Errorf("Encode ExtendedGNBCUCPName failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPConfigurationUpdate) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Decode TransactionID failed: %w", err)
		}
		s.TransactionID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode GNBCUCPName failed: %w", err)
			}
			s.GNBCUCPName = new(aper.OctetString)
			s.GNBCUCPName.Value = aper.OctetString(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.GNBCUCPTNLAToAddList = new(GNBCUCPTNLAToAddList)
		if err = s.GNBCUCPTNLAToAddList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPTNLAToAddList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.GNBCUCPTNLAToRemoveList = new(GNBCUCPTNLAToRemoveList)
		if err = s.GNBCUCPTNLAToRemoveList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPTNLAToRemoveList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.GNBCUCPTNLAToUpdateList = new(GNBCUCPTNLAToUpdateList)
		if err = s.GNBCUCPTNLAToUpdateList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPTNLAToUpdateList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = s.TransportLayerAddressInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.ExtendedGNBCUCPName = new(ExtendedGNBCUCPName)
		if err = s.ExtendedGNBCUCPName.Decode(r); err != nil {
			return fmt.Errorf("Decode ExtendedGNBCUCPName failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GNBCUCPConfigurationUpdate")
	}
	return nil
}
