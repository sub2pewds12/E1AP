package e1ap_ies

import (
	"asn1go/per"
)

func int64Ptr(v int64) *int64 { return &v }

// AdditionalRRMPriorityIndex From: 9_4_5_Information_Element_Definitions.txt:126
type AdditionalRRMPriorityIndex struct {
	Value per.BitString
}

func (s *AdditionalRRMPriorityIndex) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(32), Max: int64Ptr(32)})
}

func (s *AdditionalRRMPriorityIndex) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(32), Max: int64Ptr(32)})
	return err
}

// AlternativeQoSParaSetItemAlternativeQoSParameterIndex From: unknown:-1
type AlternativeQoSParaSetItemAlternativeQoSParameterIndex struct {
	Value int64
}

func (s *AlternativeQoSParaSetItemAlternativeQoSParameterIndex) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 8))
}

func (s *AlternativeQoSParaSetItemAlternativeQoSParameterIndex) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 8))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// AveragingWindow From: unknown:-1
type AveragingWindow struct {
	Value int64
}

func (s *AveragingWindow) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 4095))
}

func (s *AveragingWindow) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// BitRate From: 9_4_5_Information_Element_Definitions.txt:154
type BitRate struct {
	Value int64
}

func (s *BitRate) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 4000000000000))
}

func (s *BitRate) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// BurstArrivalTime From: unknown:-1
type BurstArrivalTime struct {
	Value []byte
}

func (s *BurstArrivalTime) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
}

func (s *BurstArrivalTime) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
	return err
}

// CellGroupID From: 9_4_5_Information_Element_Definitions.txt:252
type CellGroupID struct {
	Value int64
}

func (s *CellGroupID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 3))
}

func (s *CellGroupID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 3))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// CommonNetworkInstance From: 9_4_5_Information_Element_Definitions.txt:273
type CommonNetworkInstance struct {
	Value []byte
}

func (s *CommonNetworkInstance) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
}

func (s *CommonNetworkInstance) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
	return err
}

// DRBBStatusTransferReceiveStatusofPDCPSDU From: 9_4_5_Information_Element_Definitions.txt:1747
type DRBBStatusTransferReceiveStatusofPDCPSDU struct {
	Value per.BitString
}

func (s *DRBBStatusTransferReceiveStatusofPDCPSDU) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(131072)})
}

func (s *DRBBStatusTransferReceiveStatusofPDCPSDU) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(131072)})
	return err
}

// DRBID From: 9_4_5_Information_Element_Definitions.txt:586
type DRBID struct {
	Value int64
}

func (s *DRBID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 32))
}

func (s *DRBID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 32))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// DRBMeasurementResultsInformationItemULD1Result From: unknown:-1
type DRBMeasurementResultsInformationItemULD1Result struct {
	Value int64
}

func (s *DRBMeasurementResultsInformationItemULD1Result) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 10000))
}

func (s *DRBMeasurementResultsInformationItemULD1Result) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 10000))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// DRBUsageReportItemEndTimeStamp From: unknown:-1
type DRBUsageReportItemEndTimeStamp struct {
	Value []byte
}

func (s *DRBUsageReportItemEndTimeStamp) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
}

func (s *DRBUsageReportItemEndTimeStamp) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
	return err
}

// DRBUsageReportItemStartTimeStamp From: unknown:-1
type DRBUsageReportItemStartTimeStamp struct {
	Value []byte
}

func (s *DRBUsageReportItemStartTimeStamp) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
}

func (s *DRBUsageReportItemStartTimeStamp) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
	return err
}

// DRBUsageReportItemUsageCountDL From: unknown:-1
type DRBUsageReportItemUsageCountDL struct {
	Value int64
}

func (s *DRBUsageReportItemUsageCountDL) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Unconstrained())
}

func (s *DRBUsageReportItemUsageCountDL) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Unconstrained())
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// DRBUsageReportItemUsageCountUL From: unknown:-1
type DRBUsageReportItemUsageCountUL struct {
	Value int64
}

func (s *DRBUsageReportItemUsageCountUL) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Unconstrained())
}

func (s *DRBUsageReportItemUsageCountUL) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Unconstrained())
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// Dynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:1001
type Dynamic5QIDescriptorFiveQI struct {
	Value int64
}

func (s *Dynamic5QIDescriptorFiveQI) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 255))
}

func (s *Dynamic5QIDescriptorFiveQI) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// EncryptionKey From: 9_4_5_Information_Element_Definitions.txt:1079
type EncryptionKey struct {
	Value []byte
}

func (s *EncryptionKey) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
}

func (s *EncryptionKey) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
	return err
}

// ExtendedPacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1099
type ExtendedPacketDelayBudget struct {
	Value int64
}

func (s *ExtendedPacketDelayBudget) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 65535))
}

func (s *ExtendedPacketDelayBudget) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 65535))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GNBCUCPName From: 9_4_5_Information_Element_Definitions.txt:1144
type GNBCUCPName struct {
	Value []byte
}

func (s *GNBCUCPName) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUCPName) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUCPNameUTF8String From: unknown:-1
type GNBCUCPNameUTF8String struct {
	Value []byte
}

func (s *GNBCUCPNameUTF8String) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUCPNameUTF8String) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUCPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1157
type GNBCUCPNameVisibleString struct {
	Value []byte
}

func (s *GNBCUCPNameVisibleString) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUCPNameVisibleString) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUCPUEE1APID From: unknown:-1
type GNBCUCPUEE1APID struct {
	Value int64
}

func (s *GNBCUCPUEE1APID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 4294967295))
}

func (s *GNBCUCPUEE1APID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 4294967295))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GNBCUUPCapacity From: unknown:-1
type GNBCUUPCapacity struct {
	Value int64
}

func (s *GNBCUUPCapacity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 255))
}

func (s *GNBCUUPCapacity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GNBCUUPID From: 9_4_5_Information_Element_Definitions.txt:1179
type GNBCUUPID struct {
	Value int64
}

func (s *GNBCUUPID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 68719476735))
}

func (s *GNBCUUPID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 68719476735))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GNBCUUPName From: unknown:-1
type GNBCUUPName struct {
	Value []byte
}

func (s *GNBCUUPName) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUUPName) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUUPNameUTF8String From: unknown:-1
type GNBCUUPNameUTF8String struct {
	Value []byte
}

func (s *GNBCUUPNameUTF8String) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUUPNameUTF8String) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUUPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1194
type GNBCUUPNameVisibleString struct {
	Value []byte
}

func (s *GNBCUUPNameVisibleString) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
}

func (s *GNBCUUPNameVisibleString) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(150)})
	return err
}

// GNBCUUPUEE1APID From: unknown:-1
type GNBCUUPUEE1APID struct {
	Value int64
}

func (s *GNBCUUPUEE1APID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 4294967295))
}

func (s *GNBCUUPUEE1APID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 4294967295))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GNBDUID From: unknown:-1
type GNBDUID struct {
	Value int64
}

func (s *GNBDUID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 68719476735))
}

func (s *GNBDUID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 68719476735))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// GTPTEID From: 9_4_5_Information_Element_Definitions.txt:1289
type GTPTEID struct {
	Value []byte
}

func (s *GTPTEID) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
}

func (s *GTPTEID) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
	return err
}

// HFN From: unknown:-1
type HFN struct {
	Value int64
}

func (s *HFN) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 4294967295))
}

func (s *HFN) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 4294967295))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// HWCapacityIndicatorAvailableThroughput From: unknown:-1
type HWCapacityIndicatorAvailableThroughput struct {
	Value int64
}

func (s *HWCapacityIndicatorAvailableThroughput) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 100))
}

func (s *HWCapacityIndicatorAvailableThroughput) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 100))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// HWCapacityIndicatorOfferedThroughput From: unknown:-1
type HWCapacityIndicatorOfferedThroughput struct {
	Value int64
}

func (s *HWCapacityIndicatorOfferedThroughput) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 16777216))
}

func (s *HWCapacityIndicatorOfferedThroughput) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 16777216))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// InactivityTimer From: 9_4_5_Information_Element_Definitions.txt:1364
type InactivityTimer struct {
	Value int64
}

func (s *InactivityTimer) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 7200))
}

func (s *InactivityTimer) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 7200))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// IntegrityProtectionKey From: 9_4_5_Information_Element_Definitions.txt:1356
type IntegrityProtectionKey struct {
	Value []byte
}

func (s *IntegrityProtectionKey) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
}

func (s *IntegrityProtectionKey) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
	return err
}

// InterfacesToTrace From: unknown:-1
type InterfacesToTrace struct {
	Value per.BitString
}

func (s *InterfacesToTrace) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
}

func (s *InterfacesToTrace) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
	return err
}

// M7period From: 9_4_5_Information_Element_Definitions.txt:1478
type M7period struct {
	Value int64
}

func (s *M7period) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 60))
}

func (s *M7period) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 60))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MRDCDataUsageReportItemEndTimeStamp From: unknown:-1
type MRDCDataUsageReportItemEndTimeStamp struct {
	Value []byte
}

func (s *MRDCDataUsageReportItemEndTimeStamp) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
}

func (s *MRDCDataUsageReportItemEndTimeStamp) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
	return err
}

// MRDCDataUsageReportItemStartTimeStamp From: unknown:-1
type MRDCDataUsageReportItemStartTimeStamp struct {
	Value []byte
}

func (s *MRDCDataUsageReportItemStartTimeStamp) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
}

func (s *MRDCDataUsageReportItemStartTimeStamp) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
	return err
}

// MRDCDataUsageReportItemUsageCountDL From: unknown:-1
type MRDCDataUsageReportItemUsageCountDL struct {
	Value int64
}

func (s *MRDCDataUsageReportItemUsageCountDL) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Unconstrained())
}

func (s *MRDCDataUsageReportItemUsageCountDL) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Unconstrained())
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MRDCDataUsageReportItemUsageCountUL From: unknown:-1
type MRDCDataUsageReportItemUsageCountUL struct {
	Value int64
}

func (s *MRDCDataUsageReportItemUsageCountUL) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Unconstrained())
}

func (s *MRDCDataUsageReportItemUsageCountUL) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Unconstrained())
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MaxCIDEHCDL From: unknown:-1
type MaxCIDEHCDL struct {
	Value int64
}

func (s *MaxCIDEHCDL) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 32767))
}

func (s *MaxCIDEHCDL) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 32767))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MaxDataBurstVolume From: 9_4_5_Information_Element_Definitions.txt:1395
type MaxDataBurstVolume struct {
	Value int64
}

func (s *MaxDataBurstVolume) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 4095))
}

func (s *MaxDataBurstVolume) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MaxPacketLossRate From: 9_4_5_Information_Element_Definitions.txt:1413
type MaxPacketLossRate struct {
	Value int64
}

func (s *MaxPacketLossRate) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 1000))
}

func (s *MaxPacketLossRate) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1000))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// MeasurementsToActivate From: 9_4_5_Information_Element_Definitions.txt:1505
type MeasurementsToActivate struct {
	Value per.BitString
}

func (s *MeasurementsToActivate) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
}

func (s *MeasurementsToActivate) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
	return err
}

// NID From: 9_4_5_Information_Element_Definitions.txt:1539
type NID struct {
	Value per.BitString
}

func (s *NID) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(44), Max: int64Ptr(44)})
}

func (s *NID) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(44), Max: int64Ptr(44)})
	return err
}

// NRCellIdentity From: 9_4_5_Information_Element_Definitions.txt:1592
type NRCellIdentity struct {
	Value per.BitString
}

func (s *NRCellIdentity) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)})
}

func (s *NRCellIdentity) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)})
	return err
}

// NetworkInstance From: unknown:-1
type NetworkInstance struct {
	Value int64
}

func (s *NetworkInstance) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 256))
}

func (s *NetworkInstance) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 256))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// NonDynamic5QIDescriptorFiveQI From: unknown:-1
type NonDynamic5QIDescriptorFiveQI struct {
	Value int64
}

func (s *NonDynamic5QIDescriptorFiveQI) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 255))
}

func (s *NonDynamic5QIDescriptorFiveQI) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// NumberOfTunnels From: unknown:-1
type NumberOfTunnels struct {
	Value int64
}

func (s *NumberOfTunnels) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4))
}

func (s *NumberOfTunnels) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PDCPSN From: 9_4_5_Information_Element_Definitions.txt:1721
type PDCPSN struct {
	Value int64
}

func (s *PDCPSN) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 262143))
}

func (s *PDCPSN) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 262143))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PDUSessionID From: 9_4_5_Information_Element_Definitions.txt:1758
type PDUSessionID struct {
	Value int64
}

func (s *PDUSessionID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 255))
}

func (s *PDUSessionID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PERExponent From: unknown:-1
type PERExponent struct {
	Value int64
}

func (s *PERExponent) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 9))
}

func (s *PERExponent) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 9))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PERScalar From: 9_4_5_Information_Element_Definitions.txt:1649
type PERScalar struct {
	Value int64
}

func (s *PERScalar) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 9))
}

func (s *PERScalar) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 9))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PLMNIdentity From: 9_4_5_Information_Element_Definitions.txt:2018
type PLMNIdentity struct {
	Value []byte
}

func (s *PLMNIdentity) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
}

func (s *PLMNIdentity) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
	return err
}

// PPI From: unknown:-1
type PPI struct {
	Value int64
}

func (s *PPI) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 7))
}

func (s *PPI) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 7))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1636
type PacketDelayBudget struct {
	Value int64
}

func (s *PacketDelayBudget) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 1023))
}

func (s *PacketDelayBudget) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1023))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// Periodicity From: 9_4_5_Information_Element_Definitions.txt:2409
type Periodicity struct {
	Value int64
}

func (s *Periodicity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 640000))
}

func (s *Periodicity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 640000))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PortNumber From: unknown:-1
type PortNumber struct {
	Value per.BitString
}

func (s *PortNumber) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(16), Max: int64Ptr(16)})
}

func (s *PortNumber) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(16), Max: int64Ptr(16)})
	return err
}

// PriorityLevel From: unknown:-1
type PriorityLevel struct {
	Value int64
}

func (s *PriorityLevel) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 15))
}

func (s *PriorityLevel) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 15))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// PrivateIEIDLocal From: unknown:-1
type PrivateIEIDLocal struct {
	Value int64
}

func (s *PrivateIEIDLocal) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, MaxPrivateIEs))
}

func (s *PrivateIEIDLocal) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, MaxPrivateIEs))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ProcedureCode From: 9_4_6_Common_Definitions.txt:43
type ProcedureCode struct {
	Value int64
}

func (s *ProcedureCode) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 255))
}

func (s *ProcedureCode) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ProtocolExtensionID From: unknown:-1
type ProtocolExtensionID struct {
	Value int64
}

func (s *ProtocolExtensionID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, MaxProtocolExtensions))
}

func (s *ProtocolExtensionID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, MaxProtocolExtensions))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ProtocolIEID From: unknown:-1
type ProtocolIEID struct {
	Value int64
}

func (s *ProtocolIEID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, MaxProtocolIEs))
}

func (s *ProtocolIEID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, MaxProtocolIEs))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// QCI From: 9_4_5_Information_Element_Definitions.txt:2044
type QCI struct {
	Value int64
}

func (s *QCI) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 255))
}

func (s *QCI) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// QOSFlowIdentifier From: 9_4_5_Information_Element_Definitions.txt:2056
type QOSFlowIdentifier struct {
	Value int64
}

func (s *QOSFlowIdentifier) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.Constrained(0, 63))
}

func (s *QOSFlowIdentifier) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.Constrained(0, 63))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// QOSFlowRemovedItemQOSFlowAccumulatedSessionTime From: unknown:-1
type QOSFlowRemovedItemQOSFlowAccumulatedSessionTime struct {
	Value []byte
}

func (s *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(5), Max: int64Ptr(5)})
}

func (s *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(5), Max: int64Ptr(5)})
	return err
}

// QOSMappingInformationDscp From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationDscp struct {
	Value per.BitString
}

func (s *QOSMappingInformationDscp) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(6), Max: int64Ptr(6)})
}

func (s *QOSMappingInformationDscp) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(6), Max: int64Ptr(6)})
	return err
}

// QOSMappingInformationFlowLabel From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationFlowLabel struct {
	Value per.BitString
}

func (s *QOSMappingInformationFlowLabel) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(20), Max: int64Ptr(20)})
}

func (s *QOSMappingInformationFlowLabel) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(20), Max: int64Ptr(20)})
	return err
}

// QoSFlowLevelQoSParametersPagingPolicyIndicator From: 9_4_5_Information_Element_Definitions.txt:2131
type QoSFlowLevelQoSParametersPagingPolicyIndicator struct {
	Value int64
}

func (s *QoSFlowLevelQoSParametersPagingPolicyIndicator) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 8))
}

func (s *QoSFlowLevelQoSParametersPagingPolicyIndicator) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 8))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// QoSPriorityLevel From: 9_4_5_Information_Element_Definitions.txt:2112
type QoSPriorityLevel struct {
	Value int64
}

func (s *QoSPriorityLevel) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 127))
}

func (s *QoSPriorityLevel) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 127))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// QosMonitoringReportingFrequency From: unknown:-1
type QosMonitoringReportingFrequency struct {
	Value int64
}

func (s *QosMonitoringReportingFrequency) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 1800))
}

func (s *QosMonitoringReportingFrequency) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 1800))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// RANUEID From: 9_4_5_Information_Element_Definitions.txt:2200
type RANUEID struct {
	Value []byte
}

func (s *RANUEID) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
}

func (s *RANUEID) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
	return err
}

// ROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCMaxCID struct {
	Value int64
}

func (s *ROHCMaxCID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 16383))
}

func (s *ROHCMaxCID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16383))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCROHCProfiles struct {
	Value int64
}

func (s *ROHCROHCProfiles) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 511))
}

func (s *ROHCROHCProfiles) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 511))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ReportCharacteristics From: 9_4_5_Information_Element_Definitions.txt:2231
type ReportCharacteristics struct {
	Value per.BitString
}

func (s *ReportCharacteristics) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)})
}

func (s *ReportCharacteristics) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)})
	return err
}

// ResourceStatusFailureIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUCPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusFailureIEsIDGNBCUCPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusFailureIEsIDGNBCUCPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusFailureIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUUPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusFailureIEsIDGNBCUUPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusFailureIEsIDGNBCUUPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusRequestIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUCPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusRequestIEsIDGNBCUCPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusRequestIEsIDGNBCUCPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusRequestIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUUPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusRequestIEsIDGNBCUUPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusRequestIEsIDGNBCUUPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusResponseIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUCPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusResponseIEsIDGNBCUCPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusResponseIEsIDGNBCUCPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusResponseIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUUPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusResponseIEsIDGNBCUUPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusResponseIEsIDGNBCUUPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUCPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusUpdateIEsIDGNBCUCPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusUpdateIEsIDGNBCUCPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUUPMeasurementID struct {
	Value int64
}

func (s *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 4095))
}

func (s *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 4095))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// SNSSAISD From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISD struct {
	Value []byte
}

func (s *SNSSAISD) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
}

func (s *SNSSAISD) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
	return err
}

// SNSSAISST From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISST struct {
	Value []byte
}

func (s *SNSSAISST) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(1)})
}

func (s *SNSSAISST) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(1)})
	return err
}

// SubscriberProfileIDforRFP From: 9_4_5_Information_Element_Definitions.txt:2363
type SubscriberProfileIDforRFP struct {
	Value int64
}

func (s *SubscriberProfileIDforRFP) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(1, 256))
}

func (s *SubscriberProfileIDforRFP) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 256))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLAvailableCapacity struct {
	Value int64
}

func (s *TNLAvailableCapacityIndicatorDLTNLAvailableCapacity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 100))
}

func (s *TNLAvailableCapacityIndicatorDLTNLAvailableCapacity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 100))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLOfferedCapacity struct {
	Value int64
}

func (s *TNLAvailableCapacityIndicatorDLTNLOfferedCapacity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 16777216))
}

func (s *TNLAvailableCapacityIndicatorDLTNLOfferedCapacity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16777216))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TNLAvailableCapacityIndicatorULTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLAvailableCapacity struct {
	Value int64
}

func (s *TNLAvailableCapacityIndicatorULTNLAvailableCapacity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 100))
}

func (s *TNLAvailableCapacityIndicatorULTNLAvailableCapacity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 100))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TNLAvailableCapacityIndicatorULTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLOfferedCapacity struct {
	Value int64
}

func (s *TNLAvailableCapacityIndicatorULTNLOfferedCapacity) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 16777216))
}

func (s *TNLAvailableCapacityIndicatorULTNLOfferedCapacity) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16777216))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TraceID From: 9_4_5_Information_Element_Definitions.txt:2438
type TraceID struct {
	Value []byte
}

func (s *TraceID) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
}

func (s *TraceID) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
	return err
}

// TransactionID From: unknown:-1
type TransactionID struct {
	Value int64
}

func (s *TransactionID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 255))
}

func (s *TransactionID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// TransportLayerAddress From: unknown:-1
type TransportLayerAddress struct {
	Value per.BitString
}

func (s *TransportLayerAddress) Encode(w *per.Encoder) error {
	return w.EncodeBitString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
}

func (s *TransportLayerAddress) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
	return err
}

// URIaddress From: 9_4_5_Information_Element_Definitions.txt:2582
type URIaddress struct {
	Value []byte
}

func (s *URIaddress) Encode(w *per.Encoder) error {
	return w.EncodeOctetString(s.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
}

func (s *URIaddress) Decode(r *per.Decoder) error {

	var err error
	s.Value, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
	return err
}

// UplinkOnlyROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCMaxCID struct {
	Value int64
}

func (s *UplinkOnlyROHCMaxCID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 16383))
}

func (s *UplinkOnlyROHCMaxCID) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16383))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}

// UplinkOnlyROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCROHCProfiles struct {
	Value int64
}

func (s *UplinkOnlyROHCROHCProfiles) Encode(w *per.Encoder) error {
	return w.EncodeInteger(s.Value, per.ConstrainedExtensible(0, 511))
}

func (s *UplinkOnlyROHCROHCProfiles) Decode(r *per.Decoder) error {

	val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 511))
	if err != nil {
		return err
	}
	s.Value = val
	return nil
}
