package e1ap_ies

import "fmt"

type AdditionalRRMPriorityIndex []byte // From: 9_4_5_Information_Element_Definitions.txt:109

func (v AdditionalRRMPriorityIndex) Validate() error {
	if len(v) < int(32) || len(v) > int(32) {
		return fmt.Errorf("value for AdditionalRRMPriorityIndex is outside the valid length (32..32)")
	}
	return nil
}

type AveragingWindow int64 // From: 9_4_5_Information_Element_Definitions.txt:111

type BITSTRING []byte // From: builtin:-1

type BitRate int64 // From: 9_4_5_Information_Element_Definitions.txt:137

type BurstArrivalTime []byte // From: 9_4_5_Information_Element_Definitions.txt:2268

type CellGroupID int64 // From: 9_4_5_Information_Element_Definitions.txt:230

type CommonNetworkInstance []byte // From: 9_4_5_Information_Element_Definitions.txt:249

type DRBID int64 // From: 9_4_5_Information_Element_Definitions.txt:542

type EncryptionKey []byte // From: 9_4_5_Information_Element_Definitions.txt:1015

type ExtendedPacketDelayBudget int64 // From: 9_4_5_Information_Element_Definitions.txt:1035

type GNBCUCPMeasurementID int64 // From: manual_patch:-1

const (
	MinGNBCUCPMeasurementID GNBCUCPMeasurementID = 1
	MaxGNBCUCPMeasurementID GNBCUCPMeasurementID = 4095
)

func (v GNBCUCPMeasurementID) Validate() error {
	if v < MinGNBCUCPMeasurementID || v > MaxGNBCUCPMeasurementID {
		return fmt.Errorf("value for GNBCUCPMeasurementID is outside the valid range (MinGNBCUCPMeasurementID..MaxGNBCUCPMeasurementID)")
	}
	return nil
}

type GNBCUCPName []byte // From: 9_4_5_Information_Element_Definitions.txt:1080

type GNBCUCPUEE1APID int64 // From: 9_4_5_Information_Element_Definitions.txt:1082

const (
	MinGNBCUCPUEE1APID GNBCUCPUEE1APID = 0
	MaxGNBCUCPUEE1APID GNBCUCPUEE1APID = 4294967295
)

func (v GNBCUCPUEE1APID) Validate() error {
	if v < MinGNBCUCPUEE1APID || v > MaxGNBCUCPUEE1APID {
		return fmt.Errorf("value for GNBCUCPUEE1APID is outside the valid range (MinGNBCUCPUEE1APID..MaxGNBCUCPUEE1APID)")
	}
	return nil
}

type GNBCUUPCapacity int64 // From: 9_4_5_Information_Element_Definitions.txt:1084

const (
	MinGNBCUUPCapacity GNBCUUPCapacity = 0
	MaxGNBCUUPCapacity GNBCUUPCapacity = 255
)

func (v GNBCUUPCapacity) Validate() error {
	if v < MinGNBCUUPCapacity || v > MaxGNBCUUPCapacity {
		return fmt.Errorf("value for GNBCUUPCapacity is outside the valid range (MinGNBCUUPCapacity..MaxGNBCUUPCapacity)")
	}
	return nil
}

type GNBCUUPID int64 // From: 9_4_5_Information_Element_Definitions.txt:1099

const (
	MinGNBCUUPID GNBCUUPID = 0
	MaxGNBCUUPID GNBCUUPID = 68719476735
)

func (v GNBCUUPID) Validate() error {
	if v < MinGNBCUUPID || v > MaxGNBCUUPID {
		return fmt.Errorf("value for GNBCUUPID is outside the valid range (MinGNBCUUPID..MaxGNBCUUPID)")
	}
	return nil
}

type GNBCUUPMeasurementID int64 // From: manual_patch:-1

const (
	MinGNBCUUPMeasurementID GNBCUUPMeasurementID = 1
	MaxGNBCUUPMeasurementID GNBCUUPMeasurementID = 4095
)

func (v GNBCUUPMeasurementID) Validate() error {
	if v < MinGNBCUUPMeasurementID || v > MaxGNBCUUPMeasurementID {
		return fmt.Errorf("value for GNBCUUPMeasurementID is outside the valid range (MinGNBCUUPMeasurementID..MaxGNBCUUPMeasurementID)")
	}
	return nil
}

type GNBCUUPName []byte // From: 9_4_5_Information_Element_Definitions.txt:1101

type GNBCUUPUEE1APID int64 // From: 9_4_5_Information_Element_Definitions.txt:1103

const (
	MinGNBCUUPUEE1APID GNBCUUPUEE1APID = 0
	MaxGNBCUUPUEE1APID GNBCUUPUEE1APID = 4294967295
)

func (v GNBCUUPUEE1APID) Validate() error {
	if v < MinGNBCUUPUEE1APID || v > MaxGNBCUUPUEE1APID {
		return fmt.Errorf("value for GNBCUUPUEE1APID is outside the valid range (MinGNBCUUPUEE1APID..MaxGNBCUUPUEE1APID)")
	}
	return nil
}

type GNBDUID int64 // From: 9_4_5_Information_Element_Definitions.txt:1221

const (
	MinGNBDUID GNBDUID = 0
	MaxGNBDUID GNBDUID = 68719476735
)

func (v GNBDUID) Validate() error {
	if v < MinGNBDUID || v > MaxGNBDUID {
		return fmt.Errorf("value for GNBDUID is outside the valid range (MinGNBDUID..MaxGNBDUID)")
	}
	return nil
}

type GTPTEID []byte // From: 9_4_5_Information_Element_Definitions.txt:1194

func (v GTPTEID) Validate() error {
	if len(v) < int(4) || len(v) > int(4) {
		return fmt.Errorf("value for GTPTEID is outside the valid length (4..4)")
	}
	return nil
}

type HFN int64 // From: 9_4_5_Information_Element_Definitions.txt:1226

const (
	MinHFN HFN = 0
	MaxHFN HFN = 4294967295
)

func (v HFN) Validate() error {
	if v < MinHFN || v > MaxHFN {
		return fmt.Errorf("value for HFN is outside the valid range (MinHFN..MaxHFN)")
	}
	return nil
}

type INTEGER int64 // From: builtin:-1

type InactivityTimer int64 // From: 9_4_5_Information_Element_Definitions.txt:1264

type IntegrityProtectionKey []byte // From: 9_4_5_Information_Element_Definitions.txt:1256

type InterfacesToTrace []byte // From: 9_4_5_Information_Element_Definitions.txt:1266

func (v InterfacesToTrace) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for InterfacesToTrace is outside the valid length (8..8)")
	}
	return nil
}

type M7period int64 // From: 9_4_5_Information_Element_Definitions.txt:1376

type MaxDataBurstVolume int64 // From: 9_4_5_Information_Element_Definitions.txt:1295

type MaxPacketLossRate int64 // From: 9_4_5_Information_Element_Definitions.txt:1313

type MeasurementsToActivate []byte // From: 9_4_5_Information_Element_Definitions.txt:1403

func (v MeasurementsToActivate) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for MeasurementsToActivate is outside the valid length (8..8)")
	}
	return nil
}

type NID []byte // From: 9_4_5_Information_Element_Definitions.txt:1437

func (v NID) Validate() error {
	if len(v) < int(44) || len(v) > int(44) {
		return fmt.Errorf("value for NID is outside the valid length (44..44)")
	}
	return nil
}

type NRCellIdentity []byte // From: 9_4_5_Information_Element_Definitions.txt:1490

func (v NRCellIdentity) Validate() error {
	if len(v) < int(36) || len(v) > int(36) {
		return fmt.Errorf("value for NRCellIdentity is outside the valid length (36..36)")
	}
	return nil
}

type NULL bool // From: builtin:-1

type NetworkInstance int64 // From: 9_4_5_Information_Element_Definitions.txt:1409

type OCTET []byte // From: builtin:-1

type OCTETSTRING []byte // From: builtin:-1

type PDCPSN int64 // From: 9_4_5_Information_Element_Definitions.txt:1608

const (
	MinPDCPSN PDCPSN = 0
	MaxPDCPSN PDCPSN = 262143
)

func (v PDCPSN) Validate() error {
	if v < MinPDCPSN || v > MaxPDCPSN {
		return fmt.Errorf("value for PDCPSN is outside the valid range (MinPDCPSN..MaxPDCPSN)")
	}
	return nil
}

type PDUSessionID int64 // From: 9_4_5_Information_Element_Definitions.txt:1645

const (
	MinPDUSessionID PDUSessionID = 0
	MaxPDUSessionID PDUSessionID = 255
)

func (v PDUSessionID) Validate() error {
	if v < MinPDUSessionID || v > MaxPDUSessionID {
		return fmt.Errorf("value for PDUSessionID is outside the valid range (MinPDUSessionID..MaxPDUSessionID)")
	}
	return nil
}

type PERExponent int64 // From: 9_4_5_Information_Element_Definitions.txt:1537

type PERScalar int64 // From: 9_4_5_Information_Element_Definitions.txt:1536

type PLMNIdentity []byte // From: 9_4_5_Information_Element_Definitions.txt:1908

func (v PLMNIdentity) Validate() error {
	if len(v) < int(3) || len(v) > int(3) {
		return fmt.Errorf("value for PLMNIdentity is outside the valid length (3..3)")
	}
	return nil
}

type PPI int64 // From: 9_4_5_Information_Element_Definitions.txt:1912

type PacketDelayBudget int64 // From: 9_4_5_Information_Element_Definitions.txt:1523

type Periodicity int64 // From: 9_4_5_Information_Element_Definitions.txt:2266

type PortNumber []byte // From: 9_4_5_Information_Element_Definitions.txt:1910

func (v PortNumber) Validate() error {
	if len(v) < int(16) || len(v) > int(16) {
		return fmt.Errorf("value for PortNumber is outside the valid length (16..16)")
	}
	return nil
}

type PrintableString []byte // From: builtin:-1

type PriorityLevel int64 // From: 9_4_5_Information_Element_Definitions.txt:1914

type ProcedureCode int64 // From: 9_4_6_Common_Definitions.txt:43

const (
	MinProcedureCode ProcedureCode = 0
	MaxProcedureCode ProcedureCode = 255
)

func (v ProcedureCode) Validate() error {
	if v < MinProcedureCode || v > MaxProcedureCode {
		return fmt.Errorf("value for ProcedureCode is outside the valid range (MinProcedureCode..MaxProcedureCode)")
	}
	return nil
}

type ProtocolExtensionID int64 // From: 9_4_6_Common_Definitions.txt:45

const (
	MinProtocolExtensionID ProtocolExtensionID = 0
	MaxProtocolExtensionID ProtocolExtensionID = ProtocolExtensionID(MaxProtocolExtensions)
)

func (v ProtocolExtensionID) Validate() error {
	if v < MinProtocolExtensionID || v > MaxProtocolExtensionID {
		return fmt.Errorf("value for ProtocolExtensionID is outside the valid range (MinProtocolExtensionID..MaxProtocolExtensionID)")
	}
	return nil
}

type ProtocolIEID int64 // From: 9_4_6_Common_Definitions.txt:47

const (
	MinProtocolIEID ProtocolIEID = 0
	MaxProtocolIEID ProtocolIEID = ProtocolIEID(MaxProtocolIEs)
)

func (v ProtocolIEID) Validate() error {
	if v < MinProtocolIEID || v > MaxProtocolIEID {
		return fmt.Errorf("value for ProtocolIEID is outside the valid range (MinProtocolIEID..MaxProtocolIEID)")
	}
	return nil
}

type QCI int64 // From: 9_4_5_Information_Element_Definitions.txt:1934

const (
	MinQCI QCI = 0
	MaxQCI QCI = 255
)

func (v QCI) Validate() error {
	if v < MinQCI || v > MaxQCI {
		return fmt.Errorf("value for QCI is outside the valid range (MinQCI..MaxQCI)")
	}
	return nil
}

type QOSFlowIdentifier int64 // From: 9_4_5_Information_Element_Definitions.txt:1946

const (
	MinQOSFlowIdentifier QOSFlowIdentifier = 0
	MaxQOSFlowIdentifier QOSFlowIdentifier = 63
)

func (v QOSFlowIdentifier) Validate() error {
	if v < MinQOSFlowIdentifier || v > MaxQOSFlowIdentifier {
		return fmt.Errorf("value for QOSFlowIdentifier is outside the valid range (MinQOSFlowIdentifier..MaxQOSFlowIdentifier)")
	}
	return nil
}

type QoSPriorityLevel int64 // From: 9_4_5_Information_Element_Definitions.txt:2000

type RANUEID []byte // From: 9_4_5_Information_Element_Definitions.txt:2057

func (v RANUEID) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for RANUEID is outside the valid length (8..8)")
	}
	return nil
}

type ReportCharacteristics []byte // From: 9_4_5_Information_Element_Definitions.txt:2088

func (v ReportCharacteristics) Validate() error {
	if len(v) < int(36) || len(v) > int(36) {
		return fmt.Errorf("value for ReportCharacteristics is outside the valid length (36..36)")
	}
	return nil
}

type STRING []byte // From: builtin:-1

type SubscriberProfileIDforRFP int64 // From: 9_4_5_Information_Element_Definitions.txt:2220

type TraceID []byte // From: 9_4_5_Information_Element_Definitions.txt:2295

func (v TraceID) Validate() error {
	if len(v) < int(8) || len(v) > int(8) {
		return fmt.Errorf("value for TraceID is outside the valid length (8..8)")
	}
	return nil
}

type TransactionID int64 // From: 9_4_5_Information_Element_Definitions.txt:2299

type TransportLayerAddress []byte // From: 9_4_5_Information_Element_Definitions.txt:2297

type URIaddress []byte // From: 9_4_5_Information_Element_Definitions.txt:2441

type UTF8String []byte // From: builtin:-1

type VisibleString []byte // From: builtin:-1
