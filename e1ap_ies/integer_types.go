package e1ap_ies

type AveragingWindow int64
const (
	AveragingWindowMinValue int64 = 0
	AveragingWindowMaxValue int64 = 4095
)

type BitRate int64
const (
	BitRateMinValue int64 = 0
	BitRateMaxValue int64 = 4000000000000
)

type CellGroupId int64
const (
	CellGroupIdMinValue int64 = 0
	CellGroupIdMaxValue int64 = 3
)

type DrbId int64
const (
	DrbIdMinValue int64 = 1
	DrbIdMaxValue int64 = 32
)

type ExtendedPacketDelayBudget int64
const (
	ExtendedPacketDelayBudgetMinValue int64 = 1
	ExtendedPacketDelayBudgetMaxValue int64 = 65535
)

type GnbCuCpUeE1apId int64
const (
	GnbCuCpUeE1apIdMinValue int64 = 0
	GnbCuCpUeE1apIdMaxValue int64 = 4294967295
)

type GnbCuUpCapacity int64
const (
	GnbCuUpCapacityMinValue int64 = 0
	GnbCuUpCapacityMaxValue int64 = 255
)

type GnbCuUpId int64
const (
	GnbCuUpIdMinValue int64 = 0
	GnbCuUpIdMaxValue int64 = 68719476735
)

type GnbCuUpUeE1apId int64
const (
	GnbCuUpUeE1apIdMinValue int64 = 0
	GnbCuUpUeE1apIdMaxValue int64 = 4294967295
)

type GnbDuId int64
const (
	GnbDuIdMinValue int64 = 0
	GnbDuIdMaxValue int64 = 68719476735
)

type Hfn int64
const (
	HfnMinValue int64 = 0
	HfnMaxValue int64 = 4294967295
)

type InactivityTimer int64
const (
	InactivityTimerMinValue int64 = 1
	InactivityTimerMaxValue int64 = 7200
)

type M7period int64
const (
	M7periodMinValue int64 = 1
	M7periodMaxValue int64 = 60
)

type MaxDataBurstVolume int64
const (
	MaxDataBurstVolumeMinValue int64 = 0
	MaxDataBurstVolumeMaxValue int64 = 4095
)

type MaxPacketLossRate int64
const (
	MaxPacketLossRateMinValue int64 = 0
	MaxPacketLossRateMaxValue int64 = 1000
)

type NetworkInstance int64
const (
	NetworkInstanceMinValue int64 = 1
	NetworkInstanceMaxValue int64 = 256
)

type PdcpSn int64
const (
	PdcpSnMinValue int64 = 0
	PdcpSnMaxValue int64 = 262143
)

type PduSessionId int64
const (
	PduSessionIdMinValue int64 = 0
	PduSessionIdMaxValue int64 = 255
)

type PerExponent int64
const (
	PerExponentMinValue int64 = 0
	PerExponentMaxValue int64 = 9
)

type PerScalar int64
const (
	PerScalarMinValue int64 = 0
	PerScalarMaxValue int64 = 9
)

type Ppi int64
const (
	PpiMinValue int64 = 0
	PpiMaxValue int64 = 7
)

type PacketDelayBudget int64
const (
	PacketDelayBudgetMinValue int64 = 0
	PacketDelayBudgetMaxValue int64 = 1023
)

type Periodicity int64
const (
	PeriodicityMinValue int64 = 1
	PeriodicityMaxValue int64 = 640000
)

type PriorityLevel int64
const (
	PriorityLevelMinValue int64 = 0
	PriorityLevelMaxValue int64 = 0
)

type ProcedureCode int64
const (
	ProcedureCodeMinValue int64 = 0
	ProcedureCodeMaxValue int64 = 255
)

type ProtocolExtensionId int64
const (
	ProtocolExtensionIdMinValue int64 = 0
	ProtocolExtensionIdMaxValue int64 = MaxProtocolExtensions
)

type ProtocolIeId int64
const (
	ProtocolIeIdMinValue int64 = 0
	ProtocolIeIdMaxValue int64 = MaxProtocolIEs
)

type Qci int64
const (
	QciMinValue int64 = 0
	QciMaxValue int64 = 255
)

type QoSFlowIdentifier int64
const (
	QoSFlowIdentifierMinValue int64 = 0
	QoSFlowIdentifierMaxValue int64 = 63
)

type QoSPriorityLevel int64
const (
	QoSPriorityLevelMinValue int64 = 0
	QoSPriorityLevelMaxValue int64 = 127
)

type SubscriberProfileIDforRFP int64
const (
	SubscriberProfileIDforRFPMinValue int64 = 1
	SubscriberProfileIDforRFPMaxValue int64 = 256
)

type TransactionID int64
const (
	TransactionIDMinValue int64 = 0
	TransactionIDMaxValue int64 = 255
)

type AlternativeQoSParameterIndex int64
const (
	AlternativeQoSParameterIndexMinValue int64 = 1
	AlternativeQoSParameterIndexMaxValue int64 = 8
)

type AvailableThroughput int64
const (
	AvailableThroughputMinValue int64 = 0
	AvailableThroughputMaxValue int64 = 100
)

type DLTnlAvailableCapacity int64
const (
	DLTnlAvailableCapacityMinValue int64 = 0
	DLTnlAvailableCapacityMaxValue int64 = 100
)

type DLTnlOfferedCapacity int64
const (
	DLTnlOfferedCapacityMinValue int64 = 0
	DLTnlOfferedCapacityMaxValue int64 = 16777216
)

type FiveQI int64
const (
	FiveQIMinValue int64 = 0
	FiveQIMaxValue int64 = 255
)

type GNBCuCpMeasurementId int64
const (
	GNBCuCpMeasurementIdMinValue int64 = 1
	GNBCuCpMeasurementIdMaxValue int64 = 4095
)

type GNBCuUpMeasurementId int64
const (
	GNBCuUpMeasurementIdMinValue int64 = 1
	GNBCuUpMeasurementIdMaxValue int64 = 4095
)

type Local int64
const (
	LocalMinValue int64 = 0
	LocalMaxValue int64 = MaxPrivateIEs
)

type MaxCID int64
const (
	MaxCIDMinValue int64 = 0
	MaxCIDMaxValue int64 = 16383
)

type OfferedThroughput int64
const (
	OfferedThroughputMinValue int64 = 1
	OfferedThroughputMaxValue int64 = 16777216
)

type PagingPolicyIndicator int64
const (
	PagingPolicyIndicatorMinValue int64 = 1
	PagingPolicyIndicatorMaxValue int64 = 8
)

type ROHCProfiles int64
const (
	ROHCProfilesMinValue int64 = 0
	ROHCProfilesMaxValue int64 = 511
)

type ULTnlAvailableCapacity int64
const (
	ULTnlAvailableCapacityMinValue int64 = 0
	ULTnlAvailableCapacityMaxValue int64 = 100
)

type ULTnlOfferedCapacity int64
const (
	ULTnlOfferedCapacityMinValue int64 = 0
	ULTnlOfferedCapacityMaxValue int64 = 16777216
)

type UsageCountDL uint64
const (
	UsageCountDLMinValue uint64 = 0
	UsageCountDLMaxValue uint64 = 18446744073709551615
)

type UsageCountUL uint64
const (
	UsageCountULMinValue uint64 = 0
	UsageCountULMaxValue uint64 = 18446744073709551615
)

