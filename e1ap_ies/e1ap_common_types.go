package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// AdditionalRRMPriorityIndex From: 9_4_5_Information_Element_Definitions.txt:109
// ASN.1 Data Type: BIT STRING
type AdditionalRRMPriorityIndex aper.BitString

// AlternativeQoSParaSetItemAlternativeQoSParameterIndex From: 9_4_5_Information_Element_Definitions.txt:115
// ASN.1 Data Type: INTEGER
type AlternativeQoSParaSetItemAlternativeQoSParameterIndex aper.Integer

// AveragingWindow From: 9_4_5_Information_Element_Definitions.txt:111
// ASN.1 Data Type: INTEGER
type AveragingWindow aper.Integer

// BITSTRING From: builtin:-1
// ASN.1 Data Type: BIT STRING
type BITSTRING aper.BitString

// BitRate From: 9_4_5_Information_Element_Definitions.txt:137
// ASN.1 Data Type: INTEGER
type BitRate aper.Integer

// BurstArrivalTime From: 9_4_5_Information_Element_Definitions.txt:2268
// ASN.1 Data Type: OCTET STRING
type BurstArrivalTime aper.OctetString

// CellGroupID From: 9_4_5_Information_Element_Definitions.txt:230
// ASN.1 Data Type: INTEGER
type CellGroupID aper.Integer

// CommonNetworkInstance From: 9_4_5_Information_Element_Definitions.txt:249
// ASN.1 Data Type: OCTET STRING
type CommonNetworkInstance aper.OctetString

// DRBBStatusTransferReceiveStatusofPDCPSDU From: 9_4_5_Information_Element_Definitions.txt:1634
// ASN.1 Data Type: BIT STRING
type DRBBStatusTransferReceiveStatusofPDCPSDU aper.BitString

// DRBID From: 9_4_5_Information_Element_Definitions.txt:542
// ASN.1 Data Type: INTEGER
type DRBID aper.Integer

// DRBRemovedItemDRBAccumulatedSessionTime From: 9_4_5_Information_Element_Definitions.txt:577
// ASN.1 Data Type: OCTET STRING
type DRBRemovedItemDRBAccumulatedSessionTime aper.OctetString

// Validate checks if the value is within the specified range.
func (v DRBRemovedItemDRBAccumulatedSessionTime) Validate() error {
	if len(v) < int(5) || len(v) > int(5) {
		return fmt.Errorf("value for DRBRemovedItemDRBAccumulatedSessionTime is outside the valid length (5..5)")
	}
	return nil
}

// DRBUsageReportItemEndTimeStamp From: 9_4_5_Information_Element_Definitions.txt:920
// ASN.1 Data Type: OCTET STRING
type DRBUsageReportItemEndTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v DRBUsageReportItemEndTimeStamp) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for DRBUsageReportItemEndTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// DRBUsageReportItemStartTimeStamp From: 9_4_5_Information_Element_Definitions.txt:920
// ASN.1 Data Type: OCTET STRING
type DRBUsageReportItemStartTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v DRBUsageReportItemStartTimeStamp) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for DRBUsageReportItemStartTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// DRBUsageReportItemUsageCountDL From: 9_4_5_Information_Element_Definitions.txt:920
// ASN.1 Data Type: INTEGER
type DRBUsageReportItemUsageCountDL aper.Integer

// DRBUsageReportItemUsageCountUL From: 9_4_5_Information_Element_Definitions.txt:920
// ASN.1 Data Type: INTEGER
type DRBUsageReportItemUsageCountUL aper.Integer

// Dynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:940
// ASN.1 Data Type: INTEGER
type Dynamic5QIDescriptorFiveQI aper.Integer

// EncryptionKey From: 9_4_5_Information_Element_Definitions.txt:1015
// ASN.1 Data Type: OCTET STRING
type EncryptionKey aper.OctetString

// ExtendedPacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1035
// ASN.1 Data Type: INTEGER
type ExtendedPacketDelayBudget aper.Integer

// GNBCUCPMeasurementID From: manual_patch:-1
// ASN.1 Data Type: INTEGER
type GNBCUCPMeasurementID aper.Integer

const (
	MinGNBCUCPMeasurementID GNBCUCPMeasurementID = 1
	MaxGNBCUCPMeasurementID GNBCUCPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v GNBCUCPMeasurementID) Validate() error {
	if v < MinGNBCUCPMeasurementID || v > MaxGNBCUCPMeasurementID {
		return fmt.Errorf("value for GNBCUCPMeasurementID is outside the valid range (MinGNBCUCPMeasurementID..MaxGNBCUCPMeasurementID)")
	}
	return nil
}

// GNBCUCPName From: builtin:-1
// ASN.1 Data Type: PrintableString
type GNBCUCPName aper.OctetString

// GNBCUCPUEE1APID From: 9_4_5_Information_Element_Definitions.txt:1082
// ASN.1 Data Type: INTEGER
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

// GNBCUUPCapacity From: 9_4_5_Information_Element_Definitions.txt:1084
// ASN.1 Data Type: INTEGER
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

// GNBCUUPID From: 9_4_5_Information_Element_Definitions.txt:1099
// ASN.1 Data Type: INTEGER
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

// GNBCUUPMeasurementID From: manual_patch:-1
// ASN.1 Data Type: INTEGER
type GNBCUUPMeasurementID aper.Integer

const (
	MinGNBCUUPMeasurementID GNBCUUPMeasurementID = 1
	MaxGNBCUUPMeasurementID GNBCUUPMeasurementID = 4095
)

// Validate checks if the value is within the specified range.
func (v GNBCUUPMeasurementID) Validate() error {
	if v < MinGNBCUUPMeasurementID || v > MaxGNBCUUPMeasurementID {
		return fmt.Errorf("value for GNBCUUPMeasurementID is outside the valid range (MinGNBCUUPMeasurementID..MaxGNBCUUPMeasurementID)")
	}
	return nil
}

// GNBCUUPName From: builtin:-1
// ASN.1 Data Type: PrintableString
type GNBCUUPName aper.OctetString

// GNBCUUPUEE1APID From: 9_4_5_Information_Element_Definitions.txt:1103
// ASN.1 Data Type: INTEGER
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

// GNBDUID From: 9_4_5_Information_Element_Definitions.txt:1221
// ASN.1 Data Type: INTEGER
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

// GTPTEID From: 9_4_5_Information_Element_Definitions.txt:1194
// ASN.1 Data Type: OCTET STRING
type GTPTEID aper.OctetString

// Validate checks if the value is within the specified range.
func (v GTPTEID) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for GTPTEID is outside the valid length (4..4)")
	}
	return nil
}

// HFN From: 9_4_5_Information_Element_Definitions.txt:1226
// ASN.1 Data Type: INTEGER
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

// HWCapacityIndicatorAvailableThroughput From: 9_4_5_Information_Element_Definitions.txt:1228
// ASN.1 Data Type: INTEGER
type HWCapacityIndicatorAvailableThroughput aper.Integer

// HWCapacityIndicatorOfferedThroughput From: 9_4_5_Information_Element_Definitions.txt:1228
// ASN.1 Data Type: INTEGER
type HWCapacityIndicatorOfferedThroughput aper.Integer

// INTEGER From: builtin:-1
// ASN.1 Data Type: INTEGER
type INTEGER aper.Integer

// InactivityTimer From: 9_4_5_Information_Element_Definitions.txt:1264
// ASN.1 Data Type: INTEGER
type InactivityTimer aper.Integer

// IntegrityProtectionKey From: 9_4_5_Information_Element_Definitions.txt:1256
// ASN.1 Data Type: OCTET STRING
type IntegrityProtectionKey aper.OctetString

// InterfacesToTrace From: 9_4_5_Information_Element_Definitions.txt:1266
// ASN.1 Data Type: BIT STRING
type InterfacesToTrace aper.BitString

// M7period From: 9_4_5_Information_Element_Definitions.txt:1376
// ASN.1 Data Type: INTEGER
type M7period aper.Integer

// MRDCDataUsageReportItemEndTimeStamp From: 9_4_5_Information_Element_Definitions.txt:1315
// ASN.1 Data Type: OCTET STRING
type MRDCDataUsageReportItemEndTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v MRDCDataUsageReportItemEndTimeStamp) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for MRDCDataUsageReportItemEndTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// MRDCDataUsageReportItemStartTimeStamp From: 9_4_5_Information_Element_Definitions.txt:1315
// ASN.1 Data Type: OCTET STRING
type MRDCDataUsageReportItemStartTimeStamp aper.OctetString

// Validate checks if the value is within the specified range.
func (v MRDCDataUsageReportItemStartTimeStamp) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for MRDCDataUsageReportItemStartTimeStamp is outside the valid length (4..4)")
	}
	return nil
}

// MRDCDataUsageReportItemUsageCountDL From: 9_4_5_Information_Element_Definitions.txt:1315
// ASN.1 Data Type: INTEGER
type MRDCDataUsageReportItemUsageCountDL aper.Integer

// MRDCDataUsageReportItemUsageCountUL From: 9_4_5_Information_Element_Definitions.txt:1315
// ASN.1 Data Type: INTEGER
type MRDCDataUsageReportItemUsageCountUL aper.Integer

// MaxDataBurstVolume From: 9_4_5_Information_Element_Definitions.txt:1295
// ASN.1 Data Type: INTEGER
type MaxDataBurstVolume aper.Integer

// MaxPacketLossRate From: 9_4_5_Information_Element_Definitions.txt:1313
// ASN.1 Data Type: INTEGER
type MaxPacketLossRate aper.Integer

// MeasurementsToActivate From: 9_4_5_Information_Element_Definitions.txt:1403
// ASN.1 Data Type: BIT STRING
type MeasurementsToActivate aper.BitString

// NID From: 9_4_5_Information_Element_Definitions.txt:1437
// ASN.1 Data Type: BIT STRING
type NID aper.BitString

// NRCellIdentity From: 9_4_5_Information_Element_Definitions.txt:1490
// ASN.1 Data Type: BIT STRING
type NRCellIdentity aper.BitString

// NULL From: builtin:-1
// ASN.1 Data Type: NULL
type NULL bool

// NetworkInstance From: 9_4_5_Information_Element_Definitions.txt:1409
// ASN.1 Data Type: INTEGER
type NetworkInstance aper.Integer

// NonDynamic5QIDescriptorFiveQI From: 9_4_5_Information_Element_Definitions.txt:1439
// ASN.1 Data Type: INTEGER
type NonDynamic5QIDescriptorFiveQI aper.Integer

// OCTET From: builtin:-1
// ASN.1 Data Type: OCTET STRING
type OCTET aper.OctetString

// OCTETSTRING From: builtin:-1
// ASN.1 Data Type: OCTET STRING
type OCTETSTRING aper.OctetString

// PDCPSN From: 9_4_5_Information_Element_Definitions.txt:1608
// ASN.1 Data Type: INTEGER
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

// PDUSessionID From: 9_4_5_Information_Element_Definitions.txt:1645
// ASN.1 Data Type: INTEGER
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

// PERExponent From: 9_4_5_Information_Element_Definitions.txt:1537
// ASN.1 Data Type: INTEGER
type PERExponent aper.Integer

// PERScalar From: 9_4_5_Information_Element_Definitions.txt:1536
// ASN.1 Data Type: INTEGER
type PERScalar aper.Integer

// PLMNIdentity From: 9_4_5_Information_Element_Definitions.txt:1908
// ASN.1 Data Type: OCTET STRING
type PLMNIdentity aper.OctetString

// Validate checks if the value is within the specified range.
func (v PLMNIdentity) Validate() error {
	if len(v) < int(3) || len(v) > int(3) {
		return fmt.Errorf("value for PLMNIdentity is outside the valid length (3..3)")
	}
	return nil
}

// PPI From: 9_4_5_Information_Element_Definitions.txt:1912
// ASN.1 Data Type: INTEGER
type PPI aper.Integer

// PacketDelayBudget From: 9_4_5_Information_Element_Definitions.txt:1523
// ASN.1 Data Type: INTEGER
type PacketDelayBudget aper.Integer

// Periodicity From: 9_4_5_Information_Element_Definitions.txt:2266
// ASN.1 Data Type: INTEGER
type Periodicity aper.Integer

// PortNumber From: 9_4_5_Information_Element_Definitions.txt:1910
// ASN.1 Data Type: BIT STRING
type PortNumber aper.BitString

// PrintableString From: builtin:-1
// ASN.1 Data Type: PrintableString
type PrintableString aper.OctetString

// PriorityLevel From: 9_4_5_Information_Element_Definitions.txt:1914
// ASN.1 Data Type: INTEGER
type PriorityLevel aper.Integer

// PrivateIEIDLocal From: 9_4_6_Common_Definitions.txt:38
// ASN.1 Data Type: INTEGER
type PrivateIEIDLocal aper.Integer

// ProcedureCode From: 9_4_6_Common_Definitions.txt:43
// ASN.1 Data Type: BASE_TYPE
type ProcedureCode aper.Integer

// ProtocolExtensionID From: 9_4_6_Common_Definitions.txt:45
// ASN.1 Data Type: BASE_TYPE
type ProtocolExtensionID aper.Integer

// ProtocolIEID From: 9_4_6_Common_Definitions.txt:47
// ASN.1 Data Type: BASE_TYPE
type ProtocolIEID aper.Integer

// QCI From: 9_4_5_Information_Element_Definitions.txt:1934
// ASN.1 Data Type: INTEGER
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

// QOSFlowIdentifier From: 9_4_5_Information_Element_Definitions.txt:1946
// ASN.1 Data Type: INTEGER
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

// QOSFlowRemovedItemQOSFlowAccumulatedSessionTime From: 9_4_5_Information_Element_Definitions.txt:2037
// ASN.1 Data Type: OCTET STRING
type QOSFlowRemovedItemQOSFlowAccumulatedSessionTime aper.OctetString

// Validate checks if the value is within the specified range.
func (v QOSFlowRemovedItemQOSFlowAccumulatedSessionTime) Validate() error {
	if len(v) < int(5) || len(v) > int(5) {
		return fmt.Errorf("value for QOSFlowRemovedItemQOSFlowAccumulatedSessionTime is outside the valid length (5..5)")
	}
	return nil
}

// QOSMappingInformationDscp From: 9_4_5_Information_Element_Definitions.txt:2049
// ASN.1 Data Type: BIT STRING
type QOSMappingInformationDscp aper.BitString

// QOSMappingInformationFlowLabel From: 9_4_5_Information_Element_Definitions.txt:2049
// ASN.1 Data Type: BIT STRING
type QOSMappingInformationFlowLabel aper.BitString

// QoSFlowLevelQoSParametersPagingPolicyIndicator From: 9_4_5_Information_Element_Definitions.txt:2019
// ASN.1 Data Type: INTEGER
type QoSFlowLevelQoSParametersPagingPolicyIndicator aper.Integer

// QoSPriorityLevel From: 9_4_5_Information_Element_Definitions.txt:2000
// ASN.1 Data Type: INTEGER
type QoSPriorityLevel aper.Integer

// RANUEID From: 9_4_5_Information_Element_Definitions.txt:2057
// ASN.1 Data Type: OCTET STRING
type RANUEID aper.OctetString

// Validate checks if the value is within the specified range.
func (v RANUEID) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for RANUEID is outside the valid length (8..8)")
	}
	return nil
}

// ROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2115
// ASN.1 Data Type: INTEGER
type ROHCMaxCID aper.Integer

// ROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2115
// ASN.1 Data Type: INTEGER
type ROHCROHCProfiles aper.Integer

// ReportCharacteristics From: 9_4_5_Information_Element_Definitions.txt:2088
// ASN.1 Data Type: BIT STRING
type ReportCharacteristics aper.BitString

// SNSSAISD From: 9_4_5_Information_Element_Definitions.txt:2185
// ASN.1 Data Type: OCTET STRING
type SNSSAISD aper.OctetString

// Validate checks if the value is within the specified range.
func (v SNSSAISD) Validate() error {
	if len(v) < int(3) || len(v) > int(3) {
		return fmt.Errorf("value for SNSSAISD is outside the valid length (3..3)")
	}
	return nil
}

// SNSSAISST From: 9_4_5_Information_Element_Definitions.txt:2185
// ASN.1 Data Type: OCTET STRING
type SNSSAISST aper.OctetString

// Validate checks if the value is within the specified range.
func (v SNSSAISST) Validate() error {
	if len(v) < int(1) || len(v) > int(1) {
		return fmt.Errorf("value for SNSSAISST is outside the valid length (1..1)")
	}
	return nil
}

// STRING From: builtin:-1
// ASN.1 Data Type: OCTET STRING
type STRING aper.OctetString

// SubscriberProfileIDforRFP From: 9_4_5_Information_Element_Definitions.txt:2220
// ASN.1 Data Type: INTEGER
type SubscriberProfileIDforRFP aper.Integer

// TNLAvailableCapacityIndicatorDLTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2233
// ASN.1 Data Type: INTEGER
type TNLAvailableCapacityIndicatorDLTNLAvailableCapacity aper.Integer

// TNLAvailableCapacityIndicatorDLTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2233
// ASN.1 Data Type: INTEGER
type TNLAvailableCapacityIndicatorDLTNLOfferedCapacity aper.Integer

// TNLAvailableCapacityIndicatorULTNLAvailableCapacity From: 9_4_5_Information_Element_Definitions.txt:2233
// ASN.1 Data Type: INTEGER
type TNLAvailableCapacityIndicatorULTNLAvailableCapacity aper.Integer

// TNLAvailableCapacityIndicatorULTNLOfferedCapacity From: 9_4_5_Information_Element_Definitions.txt:2233
// ASN.1 Data Type: INTEGER
type TNLAvailableCapacityIndicatorULTNLOfferedCapacity aper.Integer

// TraceID From: 9_4_5_Information_Element_Definitions.txt:2295
// ASN.1 Data Type: OCTET STRING
type TraceID aper.OctetString

// Validate checks if the value is within the specified range.
func (v TraceID) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for TraceID is outside the valid length (8..8)")
	}
	return nil
}

// TransactionID From: 9_4_5_Information_Element_Definitions.txt:2299
// ASN.1 Data Type: INTEGER
type TransactionID aper.Integer

// TransportLayerAddress From: 9_4_5_Information_Element_Definitions.txt:2297
// ASN.1 Data Type: BIT STRING
type TransportLayerAddress aper.BitString

// URIaddress From: builtin:-1
// ASN.1 Data Type: VisibleString
type URIaddress aper.OctetString

// UTF8String From: builtin:-1
// ASN.1 Data Type: UTF8String
type UTF8String aper.OctetString

// UplinkOnlyROHCMaxCID From: 9_4_5_Information_Element_Definitions.txt:2430
// ASN.1 Data Type: INTEGER
type UplinkOnlyROHCMaxCID aper.Integer

// UplinkOnlyROHCROHCProfiles From: 9_4_5_Information_Element_Definitions.txt:2430
// ASN.1 Data Type: INTEGER
type UplinkOnlyROHCROHCProfiles aper.Integer

// VisibleString From: builtin:-1
// ASN.1 Data Type: VisibleString
type VisibleString aper.OctetString
