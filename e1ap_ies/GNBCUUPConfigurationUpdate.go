package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdate is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdate struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUUPID                 GNBCUUPID                  `aper:"lb:0,ub:68719476735,mandatory,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,optional,ext"`
	GNBCUUPCapacity           *GNBCUUPCapacity           `aper:"lb:0,ub:255,optional,ext"`
	GNBCUUPTNLAToRemoveList   []GNBCUUPTNLAToRemoveItem  `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUUPName       *ExtendedGNBCUUPName       `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPConfigurationUpdate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUUPName != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.SupportedPLMNs != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.GNBCUUPCapacity != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.GNBCUUPTNLAToRemoveList != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.TransportLayerAddressInfo != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.ExtendedGNBCUUPName != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(6), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.TransactionID), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUUPID), &aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
		return fmt.Errorf("Encode GNBCUUPID failed: %w", err)
	}
	if s.GNBCUUPName != nil {
		if err = s.GNBCUUPName.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUUPName failed: %w", err)
		}
	}
	if s.SupportedPLMNs != nil {
		if err = s.SupportedPLMNs.Encode(w); err != nil {
			return fmt.Errorf("Encode SupportedPLMNs failed: %w", err)
		}
	}
	if s.GNBCUUPCapacity != nil {
		if err = w.WriteInteger(int64((*s.GNBCUUPCapacity)), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Encode GNBCUUPCapacity failed: %w", err)
		}
	}
	if s.GNBCUUPTNLAToRemoveList != nil {
		if err = s.GNBCUUPTNLAToRemoveList.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUUPTNLAToRemoveList failed: %w", err)
		}
	}
	if s.TransportLayerAddressInfo != nil {
		if err = s.TransportLayerAddressInfo.Encode(w); err != nil {
			return fmt.Errorf("Encode TransportLayerAddressInfo failed: %w", err)
		}
	}
	if s.ExtendedGNBCUUPName != nil {
		if err = s.ExtendedGNBCUUPName.Encode(w); err != nil {
			return fmt.Errorf("Encode ExtendedGNBCUUPName failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPConfigurationUpdate) Decode(r *aper.AperReader) (err error) {
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
		s.TransactionID = TransactionID(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
			return fmt.Errorf("Decode GNBCUUPID failed: %w", err)
		}
		s.GNBCUUPID = GNBCUUPID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPName failed: %w", err)
			}
			tmp := aper.OctetString(val)
			s.GNBCUUPName = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.SupportedPLMNs = new(SupportedPLMNsList)
		if err = s.SupportedPLMNs.Decode(r); err != nil {
			return fmt.Errorf("Decode SupportedPLMNs failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPCapacity failed: %w", err)
			}
			tmp := GNBCUUPCapacity(val)
			s.GNBCUUPCapacity = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.GNBCUUPTNLAToRemoveList = new(GNBCUUPTNLAToRemoveList)
		if err = s.GNBCUUPTNLAToRemoveList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPTNLAToRemoveList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = s.TransportLayerAddressInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.ExtendedGNBCUUPName = new(ExtendedGNBCUUPName)
		if err = s.ExtendedGNBCUUPName.Decode(r); err != nil {
			return fmt.Errorf("Decode ExtendedGNBCUUPName failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GNBCUUPConfigurationUpdate")
	}
	return nil
}
