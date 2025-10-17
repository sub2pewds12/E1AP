package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// AdditionalRRMPriorityIndex From: 9_4_5_Information_Element_Definitions.txt:126
type AdditionalRRMPriorityIndex aper.OctetString

// Validate checks if the value is within the specified range.
func (v AdditionalRRMPriorityIndex) Validate() error {
	if len(v) < 32 || len(v) > 32 {
		return fmt.Errorf("value for AdditionalRRMPriorityIndex is outside the valid length (32..32)")
	}
	return nil
}

// AlternativeQoSParaSetItemAlternativeQoSParameterIndex From: 9_4_5_Information_Element_Definitions.txt:132
type AlternativeQoSParaSetItemAlternativeQoSParameterIndex aper.Integer

const (
	MinAlternativeQoSParaSetItemAlternativeQoSParameterIndex AlternativeQoSParaSetItemAlternativeQoSParameterIndex = 1
	MaxAlternativeQoSParaSetItemAlternativeQoSParameterIndex AlternativeQoSParaSetItemAlternativeQoSParameterIndex = 8
)

// Validate checks if the value is within the specified range.
func (v AlternativeQoSParaSetItemAlternativeQoSParameterIndex) Validate() error {
	if v < MinAlternativeQoSParaSetItemAlternativeQoSParameterIndex || v > MaxAlternativeQoSParaSetItemAlternativeQoSParameterIndex {
		return fmt.Errorf("value for AlternativeQoSParaSetItemAlternativeQoSParameterIndex is outside the valid range (MinAlternativeQoSParaSetItemAlternativeQoSParameterIndex..MaxAlternativeQoSParaSetItemAlternativeQoSParameterIndex)")
	}
	return nil
}

// AveragingWindow From: 9_4_5_Information_Element_Definitions.txt:128
type AveragingWindow aper.Integer

const (
	MinAveragingWindow AveragingWindow = 0
	MaxAveragingWindow AveragingWindow = 4095
)

// Validate checks if the value is within the specified range.
func (v AveragingWindow) Validate() error {
	if v < MinAveragingWindow || v > MaxAveragingWindow {
		return fmt.Errorf("value for AveragingWindow is outside the valid range (MinAveragingWindow..MaxAveragingWindow)")
	}
	return nil
}

// BitRate From: 9_4_5_Information_Element_Definitions.txt:154
type BitRate aper.Integer

const (
	MinBitRate BitRate = 0
	MaxBitRate BitRate = 4000000000000
)

// Validate checks if the value is within the specified range.
func (v BitRate) Validate() error {
	if v < MinBitRate || v > MaxBitRate {
		return fmt.Errorf("value for BitRate is outside the valid range (MinBitRate..MaxBitRate)")
	}
	return nil
}

// BurstArrivalTime From: 9_4_5_Information_Element_Definitions.txt:2411
type BurstArrivalTime aper.OctetString

// CellGroupID From: 9_4_5_Information_Element_Definitions.txt:252
type CellGroupID aper.Integer

const (
	MinCellGroupID CellGroupID = 0
	MaxCellGroupID CellGroupID = 3
)

// Validate checks if the value is within the specified range.
func (v CellGroupID) Validate() error {
	if v < MinCellGroupID || v > MaxCellGroupID {
		return fmt.Errorf("value for CellGroupID is outside the valid range (MinCellGroupID..MaxCellGroupID)")
	}
	return nil
}

// CommonNetworkInstance From: 9_4_5_Information_Element_Definitions.txt:273
type CommonNetworkInstance aper.OctetString

// DRBBStatusTransferReceiveStatusofPDCPSDU From: 9_4_5_Information_Element_Definitions.txt:1747
type DRBBStatusTransferReceiveStatusofPDCPSDU aper.OctetString

const (
	MinDRBBStatusTransferReceiveStatusofPDCPSDU int = 1
	MaxDRBBStatusTransferReceiveStatusofPDCPSDU int = 131072
)

// Validate checks if the value is within the specified range.
func (v DRBBStatusTransferReceiveStatusofPDCPSDU) Validate() error {
	if len(v) < MinDRBBStatusTransferReceiveStatusofPDCPSDU || len(v) > MaxDRBBStatusTransferReceiveStatusofPDCPSDU {
		return fmt.Errorf("value for DRBBStatusTransferReceiveStatusofPDCPSDU is outside the valid length (MinDRBBStatusTransferReceiveStatusofPDCPSDU..MaxDRBBStatusTransferReceiveStatusofPDCPSDU)")
	}
	return nil
}

// DRBID From: 9_4_5_Information_Element_Definitions.txt:586
type DRBID aper.Integer

const (
	MinDRBID DRBID = 1
	MaxDRBID DRBID = 32
)

// Validate checks if the value is within the specified range.
func (v DRBID) Validate() error {
	if v < MinDRBID || v > MaxDRBID {
		return fmt.Errorf("value for DRBID is outside the valid range (MinDRBID..MaxDRBID)")
	}
	return nil
}

// DRBMeasurementResultsInformationItemULD1Result From: 9_4_5_Information_Element_Definitions.txt:588
type DRBMeasurementResultsInformationItemULD1Result aper.Integer

const (
	MinDRBMeasurementResultsInformationItemULD1Result DRBMeasurementResultsInformationItemULD1Result = 0
	MaxDRBMeasurementResultsInformationItemULD1Result DRBMeasurementResultsInformationItemULD1Result = 10000
)

// Validate checks if the value is within the specified range.
func (v DRBMeasurementResultsInformationItemULD1Result) Validate() error {
	if v < MinDRBMeasurementResultsInformationItemULD1Result || v > MaxDRBMeasurementResultsInformationItemULD1Result {
		return fmt.Errorf("value for DRBMeasurementResultsInformationItemULD1Result is outside the valid range (MinDRBMeasurementResultsInformationItemULD1Result..MaxDRBMeasurementResultsInformationItemULD1Result)")
	}
	return nil
}

// DRBUsageReportItemEndTimeStamp From: 9_4_5_Information_Element_Definitions.txt:981
type DRBUsageReportItemEndTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v DRBUsageReportItemEndTimeStamp) Validate() error {
	if len(v) < 4 || len(v) > 4 {
		return fmt.Errorf("value for DRBUsageReportItemEndTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// DRBUsageReportItemStartTimeStamp From: 9_4_5_Information_Element_Definitions.txt:981
type DRBUsageReportItemStartTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v DRBUsageReportItemStartTimeStamp) Validate() error {
	if len(v) < 4 || len(v) > 4 {
		return fmt.Errorf("value for DRBUsageReportItemStartTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// DRBUsageReportItemUsageCountDL From: 9_4_5_Information_Element_Definitions.txt:981
type DRBUsageReportItemUsageCountDL aper.Integer

// DRBUsageReportItemUsageCountUL From: 9_4_5_Information_Element_Definitions.txt:981
type DRBUsageReportItemUsageCountUL aper.Integer

// Dynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:1001
type Dynamic5QIDescriptorFiveQI aper.Integer

const (
	MinDynamic5QIDescriptorFiveQI Dynamic5QIDescriptorFiveQI = 0
	MaxDynamic5QIDescriptorFiveQI Dynamic5QIDescriptorFiveQI = 255
)

// Validate checks if the value is within the specified range.
func (v Dynamic5QIDescriptorFiveQI) Validate() error {
	if v < MinDynamic5QIDescriptorFiveQI || v > MaxDynamic5QIDescriptorFiveQI {
		return fmt.Errorf("value for Dynamic5QIDescriptorFiveQI is outside the valid range (MinDynamic5QIDescriptorFiveQI..MaxDynamic5QIDescriptorFiveQI)")
	}
	return nil
}

// EncryptionKey From: 9_4_5_Information_Element_Definitions.txt:1079
type EncryptionKey aper.OctetString

// ExtendedPacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1099
type ExtendedPacketDelayBudget aper.Integer

const (
	MinExtendedPacketDelayBudget ExtendedPacketDelayBudget = 1
	MaxExtendedPacketDelayBudget ExtendedPacketDelayBudget = 65535
)

// Validate checks if the value is within the specified range.
func (v ExtendedPacketDelayBudget) Validate() error {
	if v < MinExtendedPacketDelayBudget || v > MaxExtendedPacketDelayBudget {
		return fmt.Errorf("value for ExtendedPacketDelayBudget is outside the valid range (MinExtendedPacketDelayBudget..MaxExtendedPacketDelayBudget)")
	}
	return nil
}

// GNBCUCPName From: 9_4_5_Information_Element_Definitions.txt:1144
type GNBCUCPName aper.OctetString

// GNBCUCPNameUTF8String From: 9_4_5_Information_Element_Definitions.txt:1159
type GNBCUCPNameUTF8String aper.OctetString

// GNBCUCPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1157
type GNBCUCPNameVisibleString aper.OctetString

// GNBCUCPUEE1APID From: 9_4_5_Information_Element_Definitions.txt:1162
type GNBCUCPUEE1APID aper.Integer

const (
	MinGNBCUCPUEE1APID GNBCUCPUEE1APID = 0
	MaxGNBCUCPUEE1APID GNBCUCPUEE1APID = 4294967295
)

// Validate checks if the value is within the specified range.
func (v GNBCUCPUEE1APID) Validate() error {
	if v < MinGNBCUCPUEE1APID || v > MaxGNBCUCPUEE1APID {
		return fmt.Errorf("value for GNBCUCPUEE1APID is outside the valid range (MinGNBCUCPUEE1APID..MaxGNBCUCPUEE1APID)")
	}
	return nil
}

// GNBCUUPCapacity From: 9_4_5_Information_Element_Definitions.txt:1164
type GNBCUUPCapacity aper.Integer

const (
	MinGNBCUUPCapacity GNBCUUPCapacity = 0
	MaxGNBCUUPCapacity GNBCUUPCapacity = 255
)

// Validate checks if the value is within the specified range.
func (v GNBCUUPCapacity) Validate() error {
	if v < MinGNBCUUPCapacity || v > MaxGNBCUUPCapacity {
		return fmt.Errorf("value for GNBCUUPCapacity is outside the valid range (MinGNBCUUPCapacity..MaxGNBCUUPCapacity)")
	}
	return nil
}

// GNBCUUPID From: 9_4_5_Information_Element_Definitions.txt:1179
type GNBCUUPID aper.Integer

const (
	MinGNBCUUPID GNBCUUPID = 0
	MaxGNBCUUPID GNBCUUPID = 68719476735
)

// Validate checks if the value is within the specified range.
func (v GNBCUUPID) Validate() error {
	if v < MinGNBCUUPID || v > MaxGNBCUUPID {
		return fmt.Errorf("value for GNBCUUPID is outside the valid range (MinGNBCUUPID..MaxGNBCUUPID)")
	}
	return nil
}

// GNBCUUPName From: 9_4_5_Information_Element_Definitions.txt:1181
type GNBCUUPName aper.OctetString

// GNBCUUPNameUTF8String From: 9_4_5_Information_Element_Definitions.txt:1196
type GNBCUUPNameUTF8String aper.OctetString

// GNBCUUPNameVisibleString From: 9_4_5_Information_Element_Definitions.txt:1194
type GNBCUUPNameVisibleString aper.OctetString

// GNBCUUPUEE1APID From: 9_4_5_Information_Element_Definitions.txt:1198
type GNBCUUPUEE1APID aper.Integer

const (
	MinGNBCUUPUEE1APID GNBCUUPUEE1APID = 0
	MaxGNBCUUPUEE1APID GNBCUUPUEE1APID = 4294967295
)

// Validate checks if the value is within the specified range.
func (v GNBCUUPUEE1APID) Validate() error {
	if v < MinGNBCUUPUEE1APID || v > MaxGNBCUUPUEE1APID {
		return fmt.Errorf("value for GNBCUUPUEE1APID is outside the valid range (MinGNBCUUPUEE1APID..MaxGNBCUUPUEE1APID)")
	}
	return nil
}

// GNBDUID From: 9_4_5_Information_Element_Definitions.txt:1316
type GNBDUID aper.Integer

const (
	MinGNBDUID GNBDUID = 0
	MaxGNBDUID GNBDUID = 68719476735
)

// Validate checks if the value is within the specified range.
func (v GNBDUID) Validate() error {
	if v < MinGNBDUID || v > MaxGNBDUID {
		return fmt.Errorf("value for GNBDUID is outside the valid range (MinGNBDUID..MaxGNBDUID)")
	}
	return nil
}

// GTPTEID From: 9_4_5_Information_Element_Definitions.txt:1289
type GTPTEID aper.OctetString

// Validate checks if the value is within the specified range.
func (v GTPTEID) Validate() error {
	if len(v) < 4 || len(v) > 4 {
		return fmt.Errorf("value for GTPTEID is outside the valid length (4..4)")
	}
	return nil
}

// HFN From: 9_4_5_Information_Element_Definitions.txt:1321
type HFN aper.Integer

const (
	MinHFN HFN = 0
	MaxHFN HFN = 4294967295
)

// Validate checks if the value is within the specified range.
func (v HFN) Validate() error {
	if v < MinHFN || v > MaxHFN {
		return fmt.Errorf("value for HFN is outside the valid range (MinHFN..MaxHFN)")
	}
	return nil
}

// HWCapacityIndicatorAvailableThroughput From: 9_4_5_Information_Element_Definitions.txt:1323
type HWCapacityIndicatorAvailableThroughput aper.Integer

const (
	MinHWCapacityIndicatorAvailableThroughput HWCapacityIndicatorAvailableThroughput = 0
	MaxHWCapacityIndicatorAvailableThroughput HWCapacityIndicatorAvailableThroughput = 100
)

// Validate checks if the value is within the specified range.
func (v HWCapacityIndicatorAvailableThroughput) Validate() error {
	if v < MinHWCapacityIndicatorAvailableThroughput || v > MaxHWCapacityIndicatorAvailableThroughput {
		return fmt.Errorf("value for HWCapacityIndicatorAvailableThroughput is outside the valid range (MinHWCapacityIndicatorAvailableThroughput..MaxHWCapacityIndicatorAvailableThroughput)")
	}
	return nil
}

// HWCapacityIndicatorOfferedThroughput From: 9_4_5_Information_Element_Definitions.txt:1323
type HWCapacityIndicatorOfferedThroughput aper.Integer

const (
	MinHWCapacityIndicatorOfferedThroughput HWCapacityIndicatorOfferedThroughput = 1
	MaxHWCapacityIndicatorOfferedThroughput HWCapacityIndicatorOfferedThroughput = 16777216
)

// Validate checks if the value is within the specified range.
func (v HWCapacityIndicatorOfferedThroughput) Validate() error {
	if v < MinHWCapacityIndicatorOfferedThroughput || v > MaxHWCapacityIndicatorOfferedThroughput {
		return fmt.Errorf("value for HWCapacityIndicatorOfferedThroughput is outside the valid range (MinHWCapacityIndicatorOfferedThroughput..MaxHWCapacityIndicatorOfferedThroughput)")
	}
	return nil
}

// InactivityTimer From: 9_4_5_Information_Element_Definitions.txt:1364
type InactivityTimer aper.Integer

const (
	MinInactivityTimer InactivityTimer = 1
	MaxInactivityTimer InactivityTimer = 7200
)

// Validate checks if the value is within the specified range.
func (v InactivityTimer) Validate() error {
	if v < MinInactivityTimer || v > MaxInactivityTimer {
		return fmt.Errorf("value for InactivityTimer is outside the valid range (MinInactivityTimer..MaxInactivityTimer)")
	}
	return nil
}

// IntegrityProtectionKey From: 9_4_5_Information_Element_Definitions.txt:1356
type IntegrityProtectionKey aper.OctetString

// InterfacesToTrace From: 9_4_5_Information_Element_Definitions.txt:1366
type InterfacesToTrace aper.OctetString

// Validate checks if the value is within the specified range.
func (v InterfacesToTrace) Validate() error {
	if len(v) < 8 || len(v) > 8 {
		return fmt.Errorf("value for InterfacesToTrace is outside the valid length (8..8)")
	}
	return nil
}

// M7period From: 9_4_5_Information_Element_Definitions.txt:1478
type M7period aper.Integer

const (
	MinM7period M7period = 1
	MaxM7period M7period = 60
)

// Validate checks if the value is within the specified range.
func (v M7period) Validate() error {
	if v < MinM7period || v > MaxM7period {
		return fmt.Errorf("value for M7period is outside the valid range (MinM7period..MaxM7period)")
	}
	return nil
}

// MRDCDataUsageReportItemEndTimeStamp From: 9_4_5_Information_Element_Definitions.txt:1417
type MRDCDataUsageReportItemEndTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v MRDCDataUsageReportItemEndTimeStamp) Validate() error {
	if len(v) < 4 || len(v) > 4 {
		return fmt.Errorf("value for MRDCDataUsageReportItemEndTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// MRDCDataUsageReportItemStartTimeStamp From: 9_4_5_Information_Element_Definitions.txt:1417
type MRDCDataUsageReportItemStartTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v MRDCDataUsageReportItemStartTimeStamp) Validate() error {
	if len(v) < 4 || len(v) > 4 {
		return fmt.Errorf("value for MRDCDataUsageReportItemStartTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// MRDCDataUsageReportItemUsageCountDL From: 9_4_5_Information_Element_Definitions.txt:1417
type MRDCDataUsageReportItemUsageCountDL aper.Integer

// MRDCDataUsageReportItemUsageCountUL From: 9_4_5_Information_Element_Definitions.txt:1417
type MRDCDataUsageReportItemUsageCountUL aper.Integer

// MaxCIDEHCDL From: 9_4_5_Information_Element_Definitions.txt:1415
type MaxCIDEHCDL aper.Integer

const (
	MinMaxCIDEHCDL MaxCIDEHCDL = 1
	MaxMaxCIDEHCDL MaxCIDEHCDL = 32767
)

// Validate checks if the value is within the specified range.
func (v MaxCIDEHCDL) Validate() error {
	if v < MinMaxCIDEHCDL || v > MaxMaxCIDEHCDL {
		return fmt.Errorf("value for MaxCIDEHCDL is outside the valid range (MinMaxCIDEHCDL..MaxMaxCIDEHCDL)")
	}
	return nil
}

// MaxDataBurstVolume From: 9_4_5_Information_Element_Definitions.txt:1395
type MaxDataBurstVolume aper.Integer

const (
	MinMaxDataBurstVolume MaxDataBurstVolume = 0
	MaxMaxDataBurstVolume MaxDataBurstVolume = 4095
)

// Validate checks if the value is within the specified range.
func (v MaxDataBurstVolume) Validate() error {
	if v < MinMaxDataBurstVolume || v > MaxMaxDataBurstVolume {
		return fmt.Errorf("value for MaxDataBurstVolume is outside the valid range (MinMaxDataBurstVolume..MaxMaxDataBurstVolume)")
	}
	return nil
}

// MaxPacketLossRate From: 9_4_5_Information_Element_Definitions.txt:1413
type MaxPacketLossRate aper.Integer

const (
	MinMaxPacketLossRate MaxPacketLossRate = 0
	MaxMaxPacketLossRate MaxPacketLossRate = 1000
)

// Validate checks if the value is within the specified range.
func (v MaxPacketLossRate) Validate() error {
	if v < MinMaxPacketLossRate || v > MaxMaxPacketLossRate {
		return fmt.Errorf("value for MaxPacketLossRate is outside the valid range (MinMaxPacketLossRate..MaxMaxPacketLossRate)")
	}
	return nil
}

// MeasurementsToActivate From: 9_4_5_Information_Element_Definitions.txt:1505
type MeasurementsToActivate aper.OctetString

// Validate checks if the value is within the specified range.
func (v MeasurementsToActivate) Validate() error {
	if len(v) < 8 || len(v) > 8 {
		return fmt.Errorf("value for MeasurementsToActivate is outside the valid length (8..8)")
	}
	return nil
}

// NID From: 9_4_5_Information_Element_Definitions.txt:1539
type NID aper.OctetString

// Validate checks if the value is within the specified range.
func (v NID) Validate() error {
	if len(v) < 44 || len(v) > 44 {
		return fmt.Errorf("value for NID is outside the valid length (44..44)")
	}
	return nil
}

// NRCellIdentity From: 9_4_5_Information_Element_Definitions.txt:1592
type NRCellIdentity aper.OctetString

// Validate checks if the value is within the specified range.
func (v NRCellIdentity) Validate() error {
	if len(v) < 36 || len(v) > 36 {
		return fmt.Errorf("value for NRCellIdentity is outside the valid length (36..36)")
	}
	return nil
}

// NetworkInstance From: 9_4_5_Information_Element_Definitions.txt:1511
type NetworkInstance aper.Integer

const (
	MinNetworkInstance NetworkInstance = 1
	MaxNetworkInstance NetworkInstance = 256
)

// Validate checks if the value is within the specified range.
func (v NetworkInstance) Validate() error {
	if v < MinNetworkInstance || v > MaxNetworkInstance {
		return fmt.Errorf("value for NetworkInstance is outside the valid range (MinNetworkInstance..MaxNetworkInstance)")
	}
	return nil
}

// NonDynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:1541
type NonDynamic5QIDescriptorFiveQI aper.Integer

const (
	MinNonDynamic5QIDescriptorFiveQI NonDynamic5QIDescriptorFiveQI = 0
	MaxNonDynamic5QIDescriptorFiveQI NonDynamic5QIDescriptorFiveQI = 255
)

// Validate checks if the value is within the specified range.
func (v NonDynamic5QIDescriptorFiveQI) Validate() error {
	if v < MinNonDynamic5QIDescriptorFiveQI || v > MaxNonDynamic5QIDescriptorFiveQI {
		return fmt.Errorf("value for NonDynamic5QIDescriptorFiveQI is outside the valid range (MinNonDynamic5QIDescriptorFiveQI..MaxNonDynamic5QIDescriptorFiveQI)")
	}
	return nil
}

// NumberOfTunnels From: 9_4_5_Information_Element_Definitions.txt:256
type NumberOfTunnels aper.Integer

const (
	MinNumberOfTunnels NumberOfTunnels = 1
	MaxNumberOfTunnels NumberOfTunnels = 4
)

// Validate checks if the value is within the specified range.
func (v NumberOfTunnels) Validate() error {
	if v < MinNumberOfTunnels || v > MaxNumberOfTunnels {
		return fmt.Errorf("value for NumberOfTunnels is outside the valid range (MinNumberOfTunnels..MaxNumberOfTunnels)")
	}
	return nil
}

// PDCPSN From: 9_4_5_Information_Element_Definitions.txt:1721
type PDCPSN aper.Integer

const (
	MinPDCPSN PDCPSN = 0
	MaxPDCPSN PDCPSN = 262143
)

// Validate checks if the value is within the specified range.
func (v PDCPSN) Validate() error {
	if v < MinPDCPSN || v > MaxPDCPSN {
		return fmt.Errorf("value for PDCPSN is outside the valid range (MinPDCPSN..MaxPDCPSN)")
	}
	return nil
}

// PDUSessionID From: 9_4_5_Information_Element_Definitions.txt:1758
type PDUSessionID aper.Integer

const (
	MinPDUSessionID PDUSessionID = 0
	MaxPDUSessionID PDUSessionID = 255
)

// Validate checks if the value is within the specified range.
func (v PDUSessionID) Validate() error {
	if v < MinPDUSessionID || v > MaxPDUSessionID {
		return fmt.Errorf("value for PDUSessionID is outside the valid range (MinPDUSessionID..MaxPDUSessionID)")
	}
	return nil
}

// PERExponent From: 9_4_5_Information_Element_Definitions.txt:1650
type PERExponent aper.Integer

const (
	MinPERExponent PERExponent = 0
	MaxPERExponent PERExponent = 9
)

// Validate checks if the value is within the specified range.
func (v PERExponent) Validate() error {
	if v < MinPERExponent || v > MaxPERExponent {
		return fmt.Errorf("value for PERExponent is outside the valid range (MinPERExponent..MaxPERExponent)")
	}
	return nil
}

// PERScalar From: 9_4_5_Information_Element_Definitions.txt:1649
type PERScalar aper.Integer

const (
	MinPERScalar PERScalar = 0
	MaxPERScalar PERScalar = 9
)

// Validate checks if the value is within the specified range.
func (v PERScalar) Validate() error {
	if v < MinPERScalar || v > MaxPERScalar {
		return fmt.Errorf("value for PERScalar is outside the valid range (MinPERScalar..MaxPERScalar)")
	}
	return nil
}

// PLMNIdentity From: 9_4_5_Information_Element_Definitions.txt:2018
type PLMNIdentity aper.OctetString

// Validate checks if the value is within the specified range.
func (v PLMNIdentity) Validate() error {
	if len(v) < 3 || len(v) > 3 {
		return fmt.Errorf("value for PLMNIdentity is outside the valid length (3..3)")
	}
	return nil
}

// PPI From: 9_4_5_Information_Element_Definitions.txt:2022
type PPI aper.Integer

const (
	MinPPI PPI = 0
	MaxPPI PPI = 7
)

// Validate checks if the value is within the specified range.
func (v PPI) Validate() error {
	if v < MinPPI || v > MaxPPI {
		return fmt.Errorf("value for PPI is outside the valid range (MinPPI..MaxPPI)")
	}
	return nil
}

// PacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1636
type PacketDelayBudget aper.Integer

const (
	MinPacketDelayBudget PacketDelayBudget = 0
	MaxPacketDelayBudget PacketDelayBudget = 1023
)

// Validate checks if the value is within the specified range.
func (v PacketDelayBudget) Validate() error {
	if v < MinPacketDelayBudget || v > MaxPacketDelayBudget {
		return fmt.Errorf("value for PacketDelayBudget is outside the valid range (MinPacketDelayBudget..MaxPacketDelayBudget)")
	}
	return nil
}

// Periodicity From: 9_4_5_Information_Element_Definitions.txt:2409
type Periodicity aper.Integer

const (
	MinPeriodicity Periodicity = 1
	MaxPeriodicity Periodicity = 640000
)

// Validate checks if the value is within the specified range.
func (v Periodicity) Validate() error {
	if v < MinPeriodicity || v > MaxPeriodicity {
		return fmt.Errorf("value for Periodicity is outside the valid range (MinPeriodicity..MaxPeriodicity)")
	}
	return nil
}

// PortNumber From: 9_4_5_Information_Element_Definitions.txt:2020
type PortNumber aper.OctetString

// Validate checks if the value is within the specified range.
func (v PortNumber) Validate() error {
	if len(v) < 16 || len(v) > 16 {
		return fmt.Errorf("value for PortNumber is outside the valid length (16..16)")
	}
	return nil
}

// PriorityLevel From: 9_4_5_Information_Element_Definitions.txt:2024
type PriorityLevel aper.Integer

// Validate checks if the value is within the specified range.
func (v PriorityLevel) Validate() error {
	if v < 0 || v > 0 {
		return fmt.Errorf("value for PriorityLevel is outside the valid range (0..0)")
	}
	return nil
}

// PrivateIEIDLocal From: 9_4_6_Common_Definitions.txt:38
type PrivateIEIDLocal aper.Integer

// ProcedureCode From: 9_4_6_Common_Definitions.txt:43
type ProcedureCode aper.Integer

const (
	MinProcedureCode ProcedureCode = 0
	MaxProcedureCode ProcedureCode = 255
)

// Validate checks if the value is within the specified range.
func (v ProcedureCode) Validate() error {
	if v < MinProcedureCode || v > MaxProcedureCode {
		return fmt.Errorf("value for ProcedureCode is outside the valid range (MinProcedureCode..MaxProcedureCode)")
	}
	return nil
}

// ProtocolExtensionID From: 9_4_6_Common_Definitions.txt:45
type ProtocolExtensionID aper.Integer

const (
	MinProtocolExtensionID ProtocolExtensionID = 0
	MaxProtocolExtensionID ProtocolExtensionID = ProtocolExtensionID(MaxProtocolExtensions)
)

// Validate checks if the value is within the specified range.
func (v ProtocolExtensionID) Validate() error {
	if v < MinProtocolExtensionID || v > MaxProtocolExtensionID {
		return fmt.Errorf("value for ProtocolExtensionID is outside the valid range (MinProtocolExtensionID..MaxProtocolExtensionID)")
	}
	return nil
}

// ProtocolIEID From: 9_4_6_Common_Definitions.txt:47
type ProtocolIEID aper.Integer

const (
	MinProtocolIEID ProtocolIEID = 0
	MaxProtocolIEID ProtocolIEID = ProtocolIEID(MaxProtocolIEs)
)

// Validate checks if the value is within the specified range.
func (v ProtocolIEID) Validate() error {
	if v < MinProtocolIEID || v > MaxProtocolIEID {
		return fmt.Errorf("value for ProtocolIEID is outside the valid range (MinProtocolIEID..MaxProtocolIEID)")
	}
	return nil
}

// QCI From: 9_4_5_Information_Element_Definitions.txt:2044
type QCI aper.Integer

const (
	MinQCI QCI = 0
	MaxQCI QCI = 255
)

// Validate checks if the value is within the specified range.
func (v QCI) Validate() error {
	if v < MinQCI || v > MaxQCI {
		return fmt.Errorf("value for QCI is outside the valid range (MinQCI..MaxQCI)")
	}
	return nil
}

// QOSFlowIdentifier From: 9_4_5_Information_Element_Definitions.txt:2056
type QOSFlowIdentifier aper.Integer

const (
	MinQOSFlowIdentifier QOSFlowIdentifier = 0
	MaxQOSFlowIdentifier QOSFlowIdentifier = 63
)

// Validate checks if the value is within the specified range.
func (v QOSFlowIdentifier) Validate() error {
	if v < MinQOSFlowIdentifier || v > MaxQOSFlowIdentifier {
		return fmt.Errorf("value for QOSFlowIdentifier is outside the valid range (MinQOSFlowIdentifier..MaxQOSFlowIdentifier)")
	}
	return nil
}

// QOSFlowRemovedItemQOSFlowAccumulatedSessionTime From: 9_4_5_Information_Element_Definitions.txt:2156
type QOSFlowRemovedItemQOSFlowAccumulatedSessionTime aper.OctetString

// Validate checks if the value is within the specified range.
func (v QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Validate() error {
	if len(v) < 5 || len(v) > 5 {
		return fmt.Errorf("value for QOSFlowRemovedItemQOSFlowAccumulatedSessionTime is outside the valid length (5..5)")
	}
	return nil
}

// QOSMappingInformationDscp From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationDscp aper.OctetString

// Validate checks if the value is within the specified range.
func (v QOSMappingInformationDscp) Validate() error {
	if len(v) < 6 || len(v) > 6 {
		return fmt.Errorf("value for QOSMappingInformationDscp is outside the valid length (6..6)")
	}
	return nil
}

// QOSMappingInformationFlowLabel From: 9_4_5_Information_Element_Definitions.txt:2180
type QOSMappingInformationFlowLabel aper.OctetString

// Validate checks if the value is within the specified range.
func (v QOSMappingInformationFlowLabel) Validate() error {
	if len(v) < 20 || len(v) > 20 {
		return fmt.Errorf("value for QOSMappingInformationFlowLabel is outside the valid length (20..20)")
	}
	return nil
}

// QoSFlowLevelQoSParametersPagingPolicyIndicator From: 9_4_5_Information_Element_Definitions.txt:2131
type QoSFlowLevelQoSParametersPagingPolicyIndicator aper.Integer

const (
	MinQoSFlowLevelQoSParametersPagingPolicyIndicator QoSFlowLevelQoSParametersPagingPolicyIndicator = 1
	MaxQoSFlowLevelQoSParametersPagingPolicyIndicator QoSFlowLevelQoSParametersPagingPolicyIndicator = 8
)

// Validate checks if the value is within the specified range.
func (v QoSFlowLevelQoSParametersPagingPolicyIndicator) Validate() error {
	if v < MinQoSFlowLevelQoSParametersPagingPolicyIndicator || v > MaxQoSFlowLevelQoSParametersPagingPolicyIndicator {
		return fmt.Errorf("value for QoSFlowLevelQoSParametersPagingPolicyIndicator is outside the valid range (MinQoSFlowLevelQoSParametersPagingPolicyIndicator..MaxQoSFlowLevelQoSParametersPagingPolicyIndicator)")
	}
	return nil
}

// QoSPriorityLevel From: 9_4_5_Information_Element_Definitions.txt:2112
type QoSPriorityLevel aper.Integer

const (
	MinQoSPriorityLevel QoSPriorityLevel = 0
	MaxQoSPriorityLevel QoSPriorityLevel = 127
)

// Validate checks if the value is within the specified range.
func (v QoSPriorityLevel) Validate() error {
	if v < MinQoSPriorityLevel || v > MaxQoSPriorityLevel {
		return fmt.Errorf("value for QoSPriorityLevel is outside the valid range (MinQoSPriorityLevel..MaxQoSPriorityLevel)")
	}
	return nil
}

// QosMonitoringReportingFrequency From: 9_4_5_Information_Element_Definitions.txt:2152
type QosMonitoringReportingFrequency aper.Integer

const (
	MinQosMonitoringReportingFrequency QosMonitoringReportingFrequency = 1
	MaxQosMonitoringReportingFrequency QosMonitoringReportingFrequency = 1800
)

// Validate checks if the value is within the specified range.
func (v QosMonitoringReportingFrequency) Validate() error {
	if v < MinQosMonitoringReportingFrequency || v > MaxQosMonitoringReportingFrequency {
		return fmt.Errorf("value for QosMonitoringReportingFrequency is outside the valid range (MinQosMonitoringReportingFrequency..MaxQosMonitoringReportingFrequency)")
	}
	return nil
}

// RANUEID From: 9_4_5_Information_Element_Definitions.txt:2200
type RANUEID aper.OctetString

// Validate checks if the value is within the specified range.
func (v RANUEID) Validate() error {
	if len(v) < 8 || len(v) > 8 {
		return fmt.Errorf("value for RANUEID is outside the valid length (8..8)")
	}
	return nil
}

// ROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCMaxCID aper.Integer

const (
	MinROHCMaxCID ROHCMaxCID = 0
	MaxROHCMaxCID ROHCMaxCID = 16383
)

// Validate checks if the value is within the specified range.
func (v ROHCMaxCID) Validate() error {
	if v < MinROHCMaxCID || v > MaxROHCMaxCID {
		return fmt.Errorf("value for ROHCMaxCID is outside the valid range (MinROHCMaxCID..MaxROHCMaxCID)")
	}
	return nil
}

// ROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2258
type ROHCROHCProfiles aper.Integer

const (
	MinROHCROHCProfiles ROHCROHCProfiles = 0
	MaxROHCROHCProfiles ROHCROHCProfiles = 511
)

// Validate checks if the value is within the specified range.
func (v ROHCROHCProfiles) Validate() error {
	if v < MinROHCROHCProfiles || v > MaxROHCROHCProfiles {
		return fmt.Errorf("value for ROHCROHCProfiles is outside the valid range (MinROHCROHCProfiles..MaxROHCROHCProfiles)")
	}
	return nil
}

// ReportCharacteristics From: 9_4_5_Information_Element_Definitions.txt:2231
type ReportCharacteristics aper.OctetString

// Validate checks if the value is within the specified range.
func (v ReportCharacteristics) Validate() error {
	if len(v) < 36 || len(v) > 36 {
		return fmt.Errorf("value for ReportCharacteristics is outside the valid length (36..36)")
	}
	return nil
}

// ResourceStatusFailureIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUCPMeasurementID aper.Integer

const (
	MinResourceStatusFailureIEsIDGNBCUCPMeasurementID ResourceStatusFailureIEsIDGNBCUCPMeasurementID = 1
	MaxResourceStatusFailureIEsIDGNBCUCPMeasurementID ResourceStatusFailureIEsIDGNBCUCPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusFailureIEsIDGNBCUCPMeasurementID) Validate() error {
	if v < MinResourceStatusFailureIEsIDGNBCUCPMeasurementID || v > MaxResourceStatusFailureIEsIDGNBCUCPMeasurementID {
		return fmt.Errorf("value for ResourceStatusFailureIEsIDGNBCUCPMeasurementID is outside the valid range (MinResourceStatusFailureIEsIDGNBCUCPMeasurementID..MaxResourceStatusFailureIEsIDGNBCUCPMeasurementID)")
	}
	return nil
}

// ResourceStatusFailureIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusFailureIEsIDGNBCUUPMeasurementID aper.Integer

const (
	MinResourceStatusFailureIEsIDGNBCUUPMeasurementID ResourceStatusFailureIEsIDGNBCUUPMeasurementID = 1
	MaxResourceStatusFailureIEsIDGNBCUUPMeasurementID ResourceStatusFailureIEsIDGNBCUUPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusFailureIEsIDGNBCUUPMeasurementID) Validate() error {
	if v < MinResourceStatusFailureIEsIDGNBCUUPMeasurementID || v > MaxResourceStatusFailureIEsIDGNBCUUPMeasurementID {
		return fmt.Errorf("value for ResourceStatusFailureIEsIDGNBCUUPMeasurementID is outside the valid range (MinResourceStatusFailureIEsIDGNBCUUPMeasurementID..MaxResourceStatusFailureIEsIDGNBCUUPMeasurementID)")
	}
	return nil
}

// ResourceStatusRequestIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUCPMeasurementID aper.Integer

const (
	MinResourceStatusRequestIEsIDGNBCUCPMeasurementID ResourceStatusRequestIEsIDGNBCUCPMeasurementID = 1
	MaxResourceStatusRequestIEsIDGNBCUCPMeasurementID ResourceStatusRequestIEsIDGNBCUCPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusRequestIEsIDGNBCUCPMeasurementID) Validate() error {
	if v < MinResourceStatusRequestIEsIDGNBCUCPMeasurementID || v > MaxResourceStatusRequestIEsIDGNBCUCPMeasurementID {
		return fmt.Errorf("value for ResourceStatusRequestIEsIDGNBCUCPMeasurementID is outside the valid range (MinResourceStatusRequestIEsIDGNBCUCPMeasurementID..MaxResourceStatusRequestIEsIDGNBCUCPMeasurementID)")
	}
	return nil
}

// ResourceStatusRequestIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusRequestIEsIDGNBCUUPMeasurementID aper.Integer

const (
	MinResourceStatusRequestIEsIDGNBCUUPMeasurementID ResourceStatusRequestIEsIDGNBCUUPMeasurementID = 1
	MaxResourceStatusRequestIEsIDGNBCUUPMeasurementID ResourceStatusRequestIEsIDGNBCUUPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusRequestIEsIDGNBCUUPMeasurementID) Validate() error {
	if v < MinResourceStatusRequestIEsIDGNBCUUPMeasurementID || v > MaxResourceStatusRequestIEsIDGNBCUUPMeasurementID {
		return fmt.Errorf("value for ResourceStatusRequestIEsIDGNBCUUPMeasurementID is outside the valid range (MinResourceStatusRequestIEsIDGNBCUUPMeasurementID..MaxResourceStatusRequestIEsIDGNBCUUPMeasurementID)")
	}
	return nil
}

// ResourceStatusResponseIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUCPMeasurementID aper.Integer

const (
	MinResourceStatusResponseIEsIDGNBCUCPMeasurementID ResourceStatusResponseIEsIDGNBCUCPMeasurementID = 1
	MaxResourceStatusResponseIEsIDGNBCUCPMeasurementID ResourceStatusResponseIEsIDGNBCUCPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusResponseIEsIDGNBCUCPMeasurementID) Validate() error {
	if v < MinResourceStatusResponseIEsIDGNBCUCPMeasurementID || v > MaxResourceStatusResponseIEsIDGNBCUCPMeasurementID {
		return fmt.Errorf("value for ResourceStatusResponseIEsIDGNBCUCPMeasurementID is outside the valid range (MinResourceStatusResponseIEsIDGNBCUCPMeasurementID..MaxResourceStatusResponseIEsIDGNBCUCPMeasurementID)")
	}
	return nil
}

// ResourceStatusResponseIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusResponseIEsIDGNBCUUPMeasurementID aper.Integer

const (
	MinResourceStatusResponseIEsIDGNBCUUPMeasurementID ResourceStatusResponseIEsIDGNBCUUPMeasurementID = 1
	MaxResourceStatusResponseIEsIDGNBCUUPMeasurementID ResourceStatusResponseIEsIDGNBCUUPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusResponseIEsIDGNBCUUPMeasurementID) Validate() error {
	if v < MinResourceStatusResponseIEsIDGNBCUUPMeasurementID || v > MaxResourceStatusResponseIEsIDGNBCUUPMeasurementID {
		return fmt.Errorf("value for ResourceStatusResponseIEsIDGNBCUUPMeasurementID is outside the valid range (MinResourceStatusResponseIEsIDGNBCUUPMeasurementID..MaxResourceStatusResponseIEsIDGNBCUUPMeasurementID)")
	}
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUCPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUCPMeasurementID aper.Integer

const (
	MinResourceStatusUpdateIEsIDGNBCUCPMeasurementID ResourceStatusUpdateIEsIDGNBCUCPMeasurementID = 1
	MaxResourceStatusUpdateIEsIDGNBCUCPMeasurementID ResourceStatusUpdateIEsIDGNBCUCPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusUpdateIEsIDGNBCUCPMeasurementID) Validate() error {
	if v < MinResourceStatusUpdateIEsIDGNBCUCPMeasurementID || v > MaxResourceStatusUpdateIEsIDGNBCUCPMeasurementID {
		return fmt.Errorf("value for ResourceStatusUpdateIEsIDGNBCUCPMeasurementID is outside the valid range (MinResourceStatusUpdateIEsIDGNBCUCPMeasurementID..MaxResourceStatusUpdateIEsIDGNBCUCPMeasurementID)")
	}
	return nil
}

// ResourceStatusUpdateIEsIDGNBCUUPMeasurementID From: inline_in_ie_set:-1
type ResourceStatusUpdateIEsIDGNBCUUPMeasurementID aper.Integer

const (
	MinResourceStatusUpdateIEsIDGNBCUUPMeasurementID ResourceStatusUpdateIEsIDGNBCUUPMeasurementID = 1
	MaxResourceStatusUpdateIEsIDGNBCUUPMeasurementID ResourceStatusUpdateIEsIDGNBCUUPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v ResourceStatusUpdateIEsIDGNBCUUPMeasurementID) Validate() error {
	if v < MinResourceStatusUpdateIEsIDGNBCUUPMeasurementID || v > MaxResourceStatusUpdateIEsIDGNBCUUPMeasurementID {
		return fmt.Errorf("value for ResourceStatusUpdateIEsIDGNBCUUPMeasurementID is outside the valid range (MinResourceStatusUpdateIEsIDGNBCUUPMeasurementID..MaxResourceStatusUpdateIEsIDGNBCUUPMeasurementID)")
	}
	return nil
}

// SNSSAISD From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISD aper.OctetString

// Validate checks if the value is within the specified range.
func (v SNSSAISD) Validate() error {
	if len(v) < 3 || len(v) > 3 {
		return fmt.Errorf("value for SNSSAISD is outside the valid length (3..3)")
	}
	return nil
}

// SNSSAISST From: 9_4_5_Information_Element_Definitions.txt:2328
type SNSSAISST aper.OctetString

// Validate checks if the value is within the specified range.
func (v SNSSAISST) Validate() error {
	if len(v) < 1 || len(v) > 1 {
		return fmt.Errorf("value for SNSSAISST is outside the valid length (1..1)")
	}
	return nil
}

// SubscriberProfileIDforRFP From: 9_4_5_Information_Element_Definitions.txt:2363
type SubscriberProfileIDforRFP aper.Integer

const (
	MinSubscriberProfileIDforRFP SubscriberProfileIDforRFP = 1
	MaxSubscriberProfileIDforRFP SubscriberProfileIDforRFP = 256
)

// Validate checks if the value is within the specified range.
func (v SubscriberProfileIDforRFP) Validate() error {
	if v < MinSubscriberProfileIDforRFP || v > MaxSubscriberProfileIDforRFP {
		return fmt.Errorf("value for SubscriberProfileIDforRFP is outside the valid range (MinSubscriberProfileIDforRFP..MaxSubscriberProfileIDforRFP)")
	}
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLAvailableCapacity aper.Integer

const (
	MinTNLAvailableCapacityIndicatorDLTNLAvailableCapacity TNLAvailableCapacityIndicatorDLTNLAvailableCapacity = 0
	MaxTNLAvailableCapacityIndicatorDLTNLAvailableCapacity TNLAvailableCapacityIndicatorDLTNLAvailableCapacity = 100
)

// Validate checks if the value is within the specified range.
func (v TNLAvailableCapacityIndicatorDLTNLAvailableCapacity) Validate() error {
	if v < MinTNLAvailableCapacityIndicatorDLTNLAvailableCapacity || v > MaxTNLAvailableCapacityIndicatorDLTNLAvailableCapacity {
		return fmt.Errorf("value for TNLAvailableCapacityIndicatorDLTNLAvailableCapacity is outside the valid range (MinTNLAvailableCapacityIndicatorDLTNLAvailableCapacity..MaxTNLAvailableCapacityIndicatorDLTNLAvailableCapacity)")
	}
	return nil
}

// TNLAvailableCapacityIndicatorDLTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorDLTNLOfferedCapacity aper.Integer

const (
	MinTNLAvailableCapacityIndicatorDLTNLOfferedCapacity TNLAvailableCapacityIndicatorDLTNLOfferedCapacity = 0
	MaxTNLAvailableCapacityIndicatorDLTNLOfferedCapacity TNLAvailableCapacityIndicatorDLTNLOfferedCapacity = 16777216
)

// Validate checks if the value is within the specified range.
func (v TNLAvailableCapacityIndicatorDLTNLOfferedCapacity) Validate() error {
	if v < MinTNLAvailableCapacityIndicatorDLTNLOfferedCapacity || v > MaxTNLAvailableCapacityIndicatorDLTNLOfferedCapacity {
		return fmt.Errorf("value for TNLAvailableCapacityIndicatorDLTNLOfferedCapacity is outside the valid range (MinTNLAvailableCapacityIndicatorDLTNLOfferedCapacity..MaxTNLAvailableCapacityIndicatorDLTNLOfferedCapacity)")
	}
	return nil
}

// TNLAvailableCapacityIndicatorULTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLAvailableCapacity aper.Integer

const (
	MinTNLAvailableCapacityIndicatorULTNLAvailableCapacity TNLAvailableCapacityIndicatorULTNLAvailableCapacity = 0
	MaxTNLAvailableCapacityIndicatorULTNLAvailableCapacity TNLAvailableCapacityIndicatorULTNLAvailableCapacity = 100
)

// Validate checks if the value is within the specified range.
func (v TNLAvailableCapacityIndicatorULTNLAvailableCapacity) Validate() error {
	if v < MinTNLAvailableCapacityIndicatorULTNLAvailableCapacity || v > MaxTNLAvailableCapacityIndicatorULTNLAvailableCapacity {
		return fmt.Errorf("value for TNLAvailableCapacityIndicatorULTNLAvailableCapacity is outside the valid range (MinTNLAvailableCapacityIndicatorULTNLAvailableCapacity..MaxTNLAvailableCapacityIndicatorULTNLAvailableCapacity)")
	}
	return nil
}

// TNLAvailableCapacityIndicatorULTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2376
type TNLAvailableCapacityIndicatorULTNLOfferedCapacity aper.Integer

const (
	MinTNLAvailableCapacityIndicatorULTNLOfferedCapacity TNLAvailableCapacityIndicatorULTNLOfferedCapacity = 0
	MaxTNLAvailableCapacityIndicatorULTNLOfferedCapacity TNLAvailableCapacityIndicatorULTNLOfferedCapacity = 16777216
)

// Validate checks if the value is within the specified range.
func (v TNLAvailableCapacityIndicatorULTNLOfferedCapacity) Validate() error {
	if v < MinTNLAvailableCapacityIndicatorULTNLOfferedCapacity || v > MaxTNLAvailableCapacityIndicatorULTNLOfferedCapacity {
		return fmt.Errorf("value for TNLAvailableCapacityIndicatorULTNLOfferedCapacity is outside the valid range (MinTNLAvailableCapacityIndicatorULTNLOfferedCapacity..MaxTNLAvailableCapacityIndicatorULTNLOfferedCapacity)")
	}
	return nil
}

// TraceID From: 9_4_5_Information_Element_Definitions.txt:2438
type TraceID aper.OctetString

// Validate checks if the value is within the specified range.
func (v TraceID) Validate() error {
	if len(v) < 8 || len(v) > 8 {
		return fmt.Errorf("value for TraceID is outside the valid length (8..8)")
	}
	return nil
}

// TransactionID From: 9_4_5_Information_Element_Definitions.txt:2442
type TransactionID aper.Integer

const (
	MinTransactionID TransactionID = 0
	MaxTransactionID TransactionID = 255
)

// Validate checks if the value is within the specified range.
func (v TransactionID) Validate() error {
	if v < MinTransactionID || v > MaxTransactionID {
		return fmt.Errorf("value for TransactionID is outside the valid range (MinTransactionID..MaxTransactionID)")
	}
	return nil
}

// TransportLayerAddress From: 9_4_5_Information_Element_Definitions.txt:2440
type TransportLayerAddress aper.OctetString

const (
	MinTransportLayerAddress int = 1
	MaxTransportLayerAddress int = 160
)

// Validate checks if the value is within the specified range.
func (v TransportLayerAddress) Validate() error {
	if len(v) < MinTransportLayerAddress || len(v) > MaxTransportLayerAddress {
		return fmt.Errorf("value for TransportLayerAddress is outside the valid length (MinTransportLayerAddress..MaxTransportLayerAddress)")
	}
	return nil
}

// URIaddress From: 9_4_5_Information_Element_Definitions.txt:2582
type URIaddress aper.OctetString

// UplinkOnlyROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCMaxCID aper.Integer

const (
	MinUplinkOnlyROHCMaxCID UplinkOnlyROHCMaxCID = 0
	MaxUplinkOnlyROHCMaxCID UplinkOnlyROHCMaxCID = 16383
)

// Validate checks if the value is within the specified range.
func (v UplinkOnlyROHCMaxCID) Validate() error {
	if v < MinUplinkOnlyROHCMaxCID || v > MaxUplinkOnlyROHCMaxCID {
		return fmt.Errorf("value for UplinkOnlyROHCMaxCID is outside the valid range (MinUplinkOnlyROHCMaxCID..MaxUplinkOnlyROHCMaxCID)")
	}
	return nil
}

// UplinkOnlyROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2571
type UplinkOnlyROHCROHCProfiles aper.Integer

const (
	MinUplinkOnlyROHCROHCProfiles UplinkOnlyROHCROHCProfiles = 0
	MaxUplinkOnlyROHCROHCProfiles UplinkOnlyROHCROHCProfiles = 511
)

// Validate checks if the value is within the specified range.
func (v UplinkOnlyROHCROHCProfiles) Validate() error {
	if v < MinUplinkOnlyROHCROHCProfiles || v > MaxUplinkOnlyROHCROHCProfiles {
		return fmt.Errorf("value for UplinkOnlyROHCROHCProfiles is outside the valid range (MinUplinkOnlyROHCROHCProfiles..MaxUplinkOnlyROHCROHCProfiles)")
	}
	return nil
}
