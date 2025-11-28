package e1ap_ies

import (
	"math"

	"github.com/lvdund/ngap/aper"
)

// AdditionalRRMPriorityIndex From: 9_4_5_Information_Element_Definitions.txt:126
type AdditionalRRMPriorityIndex struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *AdditionalRRMPriorityIndex) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 32, Ub: 32}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *AdditionalRRMPriorityIndex) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 32, Ub: 32}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// AlternativeQoSParaSetItemAlternativeQoSParameterIndex From: unknown:-1
type AlternativeQoSParaSetItemAlternativeQoSParameterIndex struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *AlternativeQoSParaSetItemAlternativeQoSParameterIndex) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 8}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *AlternativeQoSParaSetItemAlternativeQoSParameterIndex) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// AveragingWindow From: unknown:-1
type AveragingWindow struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *AveragingWindow) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *AveragingWindow) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// BitRate From: 9_4_5_Information_Element_Definitions.txt:154
type BitRate struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *BitRate) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BitRate) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// BurstArrivalTime From: unknown:-1
type BurstArrivalTime struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *BurstArrivalTime) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 0, Ub: 0}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BurstArrivalTime) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false)
	return err
}

// CellGroupID From: 9_4_5_Information_Element_Definitions.txt:252
type CellGroupID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellGroupID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 3}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CellGroupID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// CommonNetworkInstance From: 9_4_5_Information_Element_Definitions.txt:273
type CommonNetworkInstance struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *CommonNetworkInstance) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 0, Ub: 0}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CommonNetworkInstance) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false)
	return err
}

// DRBBStatusTransferReceiveStatusofPDCPSDU From: 9_4_5_Information_Element_Definitions.txt:1747
type DRBBStatusTransferReceiveStatusofPDCPSDU struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBBStatusTransferReceiveStatusofPDCPSDU) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 131072}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBBStatusTransferReceiveStatusofPDCPSDU) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 131072}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// DRBID From: 9_4_5_Information_Element_Definitions.txt:586
type DRBID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 32}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// DRBMeasurementResultsInformationItemULD1Result From: unknown:-1
type DRBMeasurementResultsInformationItemULD1Result struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBMeasurementResultsInformationItemULD1Result) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 10000}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBMeasurementResultsInformationItemULD1Result) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10000}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// DRBUsageReportItemEndTimeStamp From: unknown:-1
type DRBUsageReportItemEndTimeStamp struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBUsageReportItemEndTimeStamp) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 4, Ub: 4}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBUsageReportItemEndTimeStamp) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	return err
}

// DRBUsageReportItemStartTimeStamp From: unknown:-1
type DRBUsageReportItemStartTimeStamp struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBUsageReportItemStartTimeStamp) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 4, Ub: 4}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBUsageReportItemStartTimeStamp) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	return err
}

// DRBUsageReportItemUsageCountDL From: unknown:-1
type DRBUsageReportItemUsageCountDL struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBUsageReportItemUsageCountDL) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBUsageReportItemUsageCountDL) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// DRBUsageReportItemUsageCountUL From: unknown:-1
type DRBUsageReportItemUsageCountUL struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBUsageReportItemUsageCountUL) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBUsageReportItemUsageCountUL) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// Dynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:1001
type Dynamic5QIDescriptorFiveQI struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *Dynamic5QIDescriptorFiveQI) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Dynamic5QIDescriptorFiveQI) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// EncryptionKey From: 9_4_5_Information_Element_Definitions.txt:1079
type EncryptionKey struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *EncryptionKey) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 0, Ub: 0}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EncryptionKey) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false)
	return err
}

// ExtendedPacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1099
type ExtendedPacketDelayBudget struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedPacketDelayBudget) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 65535}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ExtendedPacketDelayBudget) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 65535}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GNBCUCPName From: 9_4_5_Information_Element_Definitions.txt:1144
type GNBCUCPName struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPName) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPName) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUCPNameUTF8String From: unknown:-1
type GNBCUCPNameUTF8String struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPNameUTF8String) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPNameUTF8String) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUCPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1157
type GNBCUCPNameVisibleString struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPNameVisibleString) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPNameVisibleString) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUCPUEE1APID From: unknown:-1
type GNBCUCPUEE1APID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPUEE1APID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPUEE1APID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GNBCUUPCapacity From: unknown:-1
type GNBCUUPCapacity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPCapacity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPCapacity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GNBCUUPID From: 9_4_5_Information_Element_Definitions.txt:1179
type GNBCUUPID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 68719476735}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GNBCUUPName From: unknown:-1
type GNBCUUPName struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPName) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPName) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUUPNameUTF8String From: unknown:-1
type GNBCUUPNameUTF8String struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPNameUTF8String) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPNameUTF8String) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUUPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1194
type GNBCUUPNameVisibleString struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPNameVisibleString) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 150}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPNameVisibleString) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, false)
	return err
}

// GNBCUUPUEE1APID From: unknown:-1
type GNBCUUPUEE1APID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPUEE1APID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPUEE1APID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GNBDUID From: unknown:-1
type GNBDUID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBDUID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 68719476735}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBDUID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// GTPTEID From: 9_4_5_Information_Element_Definitions.txt:1289
type GTPTEID struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *GTPTEID) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 4, Ub: 4}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GTPTEID) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	return err
}

// HFN From: unknown:-1
type HFN struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *HFN) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *HFN) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// HWCapacityIndicatorAvailableThroughput From: unknown:-1
type HWCapacityIndicatorAvailableThroughput struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *HWCapacityIndicatorAvailableThroughput) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 100}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *HWCapacityIndicatorAvailableThroughput) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// HWCapacityIndicatorOfferedThroughput From: unknown:-1
type HWCapacityIndicatorOfferedThroughput struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *HWCapacityIndicatorOfferedThroughput) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 16777216}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *HWCapacityIndicatorOfferedThroughput) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16777216}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// InactivityTimer From: 9_4_5_Information_Element_Definitions.txt:1364
type InactivityTimer struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *InactivityTimer) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *InactivityTimer) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// IntegrityProtectionKey From: 9_4_5_Information_Element_Definitions.txt:1356
type IntegrityProtectionKey struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *IntegrityProtectionKey) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 0, Ub: 0}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *IntegrityProtectionKey) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false)
	return err
}

// InterfacesToTrace From: unknown:-1
type InterfacesToTrace struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *InterfacesToTrace) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *InterfacesToTrace) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// M7period From: 9_4_5_Information_Element_Definitions.txt:1478
type M7period struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *M7period) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 60}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *M7period) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 60}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MRDCDataUsageReportItemEndTimeStamp From: unknown:-1
type MRDCDataUsageReportItemEndTimeStamp struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItemEndTimeStamp) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 4, Ub: 4}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItemEndTimeStamp) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	return err
}

// MRDCDataUsageReportItemStartTimeStamp From: unknown:-1
type MRDCDataUsageReportItemStartTimeStamp struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItemStartTimeStamp) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 4, Ub: 4}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItemStartTimeStamp) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	return err
}

// MRDCDataUsageReportItemUsageCountDL From: unknown:-1
type MRDCDataUsageReportItemUsageCountDL struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItemUsageCountDL) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItemUsageCountDL) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MRDCDataUsageReportItemUsageCountUL From: unknown:-1
type MRDCDataUsageReportItemUsageCountUL struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItemUsageCountUL) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItemUsageCountUL) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MaxCIDEHCDL From: unknown:-1
type MaxCIDEHCDL struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *MaxCIDEHCDL) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 32767}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MaxCIDEHCDL) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32767}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MaxDataBurstVolume From: 9_4_5_Information_Element_Definitions.txt:1395
type MaxDataBurstVolume struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *MaxDataBurstVolume) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MaxDataBurstVolume) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MaxPacketLossRate From: 9_4_5_Information_Element_Definitions.txt:1413
type MaxPacketLossRate struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *MaxPacketLossRate) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 1000}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MaxPacketLossRate) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1000}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// MeasurementsToActivate From: 9_4_5_Information_Element_Definitions.txt:1505
type MeasurementsToActivate struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *MeasurementsToActivate) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MeasurementsToActivate) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// NID From: 9_4_5_Information_Element_Definitions.txt:1539
type NID struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *NID) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 44, Ub: 44}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NID) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 44, Ub: 44}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// NRCellIdentity From: 9_4_5_Information_Element_Definitions.txt:1592
type NRCellIdentity struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *NRCellIdentity) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 36, Ub: 36}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NRCellIdentity) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 36, Ub: 36}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// NetworkInstance From: unknown:-1
type NetworkInstance struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *NetworkInstance) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 256}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NetworkInstance) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// NonDynamic5QIDescriptorFiveQI From: unknown:-1
type NonDynamic5QIDescriptorFiveQI struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *NonDynamic5QIDescriptorFiveQI) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NonDynamic5QIDescriptorFiveQI) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// NumberOfTunnels From: unknown:-1
type NumberOfTunnels struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *NumberOfTunnels) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NumberOfTunnels) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PDCPSN From: 9_4_5_Information_Element_Definitions.txt:1721
type PDCPSN struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPSN) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 262143}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPSN) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 262143}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PDUSessionID From: 9_4_5_Information_Element_Definitions.txt:1758
type PDUSessionID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PERExponent From: unknown:-1
type PERExponent struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PERExponent) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 9}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PERExponent) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PERScalar From: 9_4_5_Information_Element_Definitions.txt:1649
type PERScalar struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PERScalar) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 9}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PERScalar) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PLMNIdentity From: 9_4_5_Information_Element_Definitions.txt:2018
type PLMNIdentity struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *PLMNIdentity) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 3, Ub: 3}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PLMNIdentity) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	return err
}

// PPI From: unknown:-1
type PPI struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PPI) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 7}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PPI) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1636
type PacketDelayBudget struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PacketDelayBudget) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 1023}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PacketDelayBudget) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// Periodicity From: 9_4_5_Information_Element_Definitions.txt:2409
type Periodicity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *Periodicity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 640000}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Periodicity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 640000}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PortNumber From: unknown:-1
type PortNumber struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *PortNumber) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PortNumber) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// PriorityLevel From: unknown:-1
type PriorityLevel struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PriorityLevel) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 15}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PriorityLevel) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// PrivateIEIDLocal From: unknown:-1
type PrivateIEIDLocal struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *PrivateIEIDLocal) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: MaxPrivateIEs}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PrivateIEIDLocal) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: MaxPrivateIEs}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ProcedureCode From: 9_4_6_Common_Definitions.txt:43
type ProcedureCode struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProcedureCode) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProcedureCode) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ProtocolExtensionID From: unknown:-1
type ProtocolExtensionID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolExtensionID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: MaxProtocolExtensions}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolExtensionID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: MaxProtocolExtensions}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ProtocolIEID From: unknown:-1
type ProtocolIEID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolIEID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: MaxProtocolIEs}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIEID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: MaxProtocolIEs}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// QCI From: 9_4_5_Information_Element_Definitions.txt:2044
type QCI struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *QCI) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QCI) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// QOSFlowIdentifier From: 9_4_5_Information_Element_Definitions.txt:2056
type QOSFlowIdentifier struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowIdentifier) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 63}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowIdentifier) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// QOSFlowRemovedItemQOSFlowAccumulatedSessionTime From: unknown:-1
type QOSFlowRemovedItemQOSFlowAccumulatedSessionTime struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 5, Ub: 5}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 5, Ub: 5}, false)
	return err
}

// QOSMappingInformationDscp From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationDscp struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSMappingInformationDscp) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSMappingInformationDscp) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// QOSMappingInformationFlowLabel From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationFlowLabel struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSMappingInformationFlowLabel) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 20, Ub: 20}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSMappingInformationFlowLabel) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 20, Ub: 20}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// QoSFlowLevelQoSParametersPagingPolicyIndicator From: 9_4_5_Information_Element_Definitions.txt:2131
type QoSFlowLevelQoSParametersPagingPolicyIndicator struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParametersPagingPolicyIndicator) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 8}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSFlowLevelQoSParametersPagingPolicyIndicator) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// QoSPriorityLevel From: 9_4_5_Information_Element_Definitions.txt:2112
type QoSPriorityLevel struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSPriorityLevel) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 127}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSPriorityLevel) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// QosMonitoringReportingFrequency From: unknown:-1
type QosMonitoringReportingFrequency struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *QosMonitoringReportingFrequency) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 1800}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QosMonitoringReportingFrequency) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 1800}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// RANUEID From: 9_4_5_Information_Element_Definitions.txt:2200
type RANUEID struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *RANUEID) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 8, Ub: 8}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *RANUEID) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 8, Ub: 8}, false)
	return err
}

// ROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCMaxCID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ROHCMaxCID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 16383}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHCMaxCID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCROHCProfiles struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ROHCROHCProfiles) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 511}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHCROHCProfiles) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 511}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ReportCharacteristics From: 9_4_5_Information_Element_Definitions.txt:2231
type ReportCharacteristics struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *ReportCharacteristics) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 36, Ub: 36}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ReportCharacteristics) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 36, Ub: 36}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// ResourceStatusFailureIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUCPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusFailureIEsIDGNBCUCPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusFailureIEsIDGNBCUCPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusFailureIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUUPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusFailureIEsIDGNBCUUPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusFailureIEsIDGNBCUUPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusRequestIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUCPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusRequestIEsIDGNBCUCPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusRequestIEsIDGNBCUCPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusRequestIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUUPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusRequestIEsIDGNBCUUPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusRequestIEsIDGNBCUUPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusResponseIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUCPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusResponseIEsIDGNBCUCPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusResponseIEsIDGNBCUCPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusResponseIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUUPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusResponseIEsIDGNBCUUPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusResponseIEsIDGNBCUUPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUCPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusUpdateIEsIDGNBCUCPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusUpdateIEsIDGNBCUCPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUUPMeasurementID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 4095}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// SNSSAISD From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISD struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *SNSSAISD) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 3, Ub: 3}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SNSSAISD) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	return err
}

// SNSSAISST From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISST struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *SNSSAISST) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 1, Ub: 1}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SNSSAISST) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	return err
}

// SubscriberProfileIDforRFP From: 9_4_5_Information_Element_Definitions.txt:2363
type SubscriberProfileIDforRFP struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *SubscriberProfileIDforRFP) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 1, Ub: 256}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SubscriberProfileIDforRFP) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLAvailableCapacity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicatorDLTNLAvailableCapacity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 100}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicatorDLTNLAvailableCapacity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLOfferedCapacity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicatorDLTNLOfferedCapacity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicatorDLTNLOfferedCapacity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16777216}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TNLAvailableCapacityIndicatorULTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLAvailableCapacity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicatorULTNLAvailableCapacity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 100}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicatorULTNLAvailableCapacity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TNLAvailableCapacityIndicatorULTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLOfferedCapacity struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicatorULTNLOfferedCapacity) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicatorULTNLOfferedCapacity) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16777216}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TraceID From: 9_4_5_Information_Element_Definitions.txt:2438
type TraceID struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceID) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 8, Ub: 8}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceID) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 8, Ub: 8}, false)
	return err
}

// TransactionID From: unknown:-1
type TransactionID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransactionID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 255}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransactionID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// TransportLayerAddress From: unknown:-1
type TransportLayerAddress struct {
	Value aper.BitString
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportLayerAddress) Encode(w *aper.AperWriter) error {
	return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransportLayerAddress) Decode(r *aper.AperReader) error {

	var err error
	var numBits uint
	s.Value.Bytes, numBits, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 160}, false)
	if err == nil {
		s.Value.NumBits = uint64(numBits)
	}
	return err
}

// URIaddress From: 9_4_5_Information_Element_Definitions.txt:2582
type URIaddress struct {
	Value aper.OctetString
}

// Encode implements the aper.AperMarshaller interface.
func (s *URIaddress) Encode(w *aper.AperWriter) error {
	return w.WriteOctetString([]byte(s.Value), &aper.Constraint{Lb: 0, Ub: 0}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *URIaddress) Decode(r *aper.AperReader) error {

	var err error
	s.Value, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false)
	return err
}

// UplinkOnlyROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCMaxCID struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *UplinkOnlyROHCMaxCID) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 16383}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UplinkOnlyROHCMaxCID) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}

// UplinkOnlyROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCROHCProfiles struct {
	Value aper.Integer
}

// Encode implements the aper.AperMarshaller interface.
func (s *UplinkOnlyROHCROHCProfiles) Encode(w *aper.AperWriter) error {
	return w.WriteInteger(int64(s.Value), &aper.Constraint{Lb: 0, Ub: 511}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UplinkOnlyROHCROHCProfiles) Decode(r *aper.AperReader) error {

	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 511}, true)
	if err != nil {
		return err
	}
	s.Value = aper.Integer(val)
	return nil
}
