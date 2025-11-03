package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CellTrafficTrace is a generated SEQUENCE type.
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                GNBCUCPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                GNBCUUPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	TraceID                        TraceID               `aper:"lb:8,ub:8,mandatory,ext"`
	TraceCollectionEntityIPAddress TransportLayerAddress `aper:"lb:1,ub:160,mandatory,ext"`
	PrivacyIndicator               *PrivacyIndicator     `aper:"optional,ext"`
	URIaddress                     *URIaddress           `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellTrafficTrace) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.PrivacyIndicator != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.URIaddress != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUCPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUCPUEE1APID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUUPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUUPUEE1APID failed: %w", err)
	}
	if err = s.TraceID.Encode(w); err != nil {
		return fmt.Errorf("Encode TraceID failed: %w", err)
	}
	if err = s.TraceCollectionEntityIPAddress.Encode(w); err != nil {
		return fmt.Errorf("Encode TraceCollectionEntityIPAddress failed: %w", err)
	}
	if s.PrivacyIndicator != nil {
		if err = w.WriteEnumerate(uint64(s.PrivacyIndicator.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Encode PrivacyIndicator failed: %w", err)
		}
	}
	if s.URIaddress != nil {
		if err = w.WriteOctetString([]byte(s.URIaddress.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode URIaddress failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CellTrafficTrace) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
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
	if err = s.TraceID.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceID failed: %w", err)
	}
	if err = s.TraceCollectionEntityIPAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceCollectionEntityIPAddress failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.PrivacyIndicator = new(PrivacyIndicator)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode PrivacyIndicator failed: %w", err)
			}
			s.PrivacyIndicator.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode URIaddress failed: %w", err)
			}
			s.URIaddress = new(URIaddress)
			s.URIaddress.Value = aper.OctetString(val)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for CellTrafficTrace")
	}
	return nil
}
