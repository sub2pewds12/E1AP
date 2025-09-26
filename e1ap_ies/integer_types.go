package e1ap_ies

// AveragingWindow ::= INTEGER
type Averagingwindow int64

// BitRate ::= INTEGER
type Bitrate int64

// Cell-Group-ID ::= INTEGER
type CellGroupId int64

// DRB-ID ::= INTEGER
type DrbId int64

// DRB-Usage-Report-Item-usageCountDL ::= INTEGER
type DrbUsageReportItemUsagecountdl int64
const (
	DrbUsageReportItemUsagecountdlMinValue = 0
	DrbUsageReportItemUsagecountdlMaxValue = 18446744073709551615
)

// DRB-Usage-Report-Item-usageCountUL ::= INTEGER
type DrbUsageReportItemUsagecountul int64
const (
	DrbUsageReportItemUsagecountulMinValue = 0
	DrbUsageReportItemUsagecountulMaxValue = 18446744073709551615
)

// ExtendedPacketDelayBudget ::= INTEGER
type Extendedpacketdelaybudget int64

// GNB-CU-CP-UE-E1AP-ID ::= INTEGER
type GnbCuCpUeE1apId int64
const (
	GnbCuCpUeE1apIdMinValue = 0
	GnbCuCpUeE1apIdMaxValue = 4294967295
)

// GNB-CU-UP-Capacity ::= INTEGER
type GnbCuUpCapacity int64
const (
	GnbCuUpCapacityMinValue = 0
	GnbCuUpCapacityMaxValue = 255
)

// GNB-CU-UP-ID ::= INTEGER
type GnbCuUpId int64
const (
	GnbCuUpIdMinValue = 0
	GnbCuUpIdMaxValue = 68719476735
)

// GNB-CU-UP-UE-E1AP-ID ::= INTEGER
type GnbCuUpUeE1apId int64
const (
	GnbCuUpUeE1apIdMinValue = 0
	GnbCuUpUeE1apIdMaxValue = 4294967295
)

// GNB-DU-ID ::= INTEGER
type GnbDuId int64
const (
	GnbDuIdMinValue = 0
	GnbDuIdMaxValue = 68719476735
)

// HFN ::= INTEGER
type Hfn int64
const (
	HfnMinValue = 0
	HfnMaxValue = 4294967295
)

// Inactivity-Timer ::= INTEGER
type InactivityTimer int64

// M7period ::= INTEGER
type M7period int64

// MRDC-Data-Usage-Report-Item-usageCountDL ::= INTEGER
type MrdcDataUsageReportItemUsagecountdl int64
const (
	MrdcDataUsageReportItemUsagecountdlMinValue = 0
	MrdcDataUsageReportItemUsagecountdlMaxValue = 18446744073709551615
)

// MRDC-Data-Usage-Report-Item-usageCountUL ::= INTEGER
type MrdcDataUsageReportItemUsagecountul int64
const (
	MrdcDataUsageReportItemUsagecountulMinValue = 0
	MrdcDataUsageReportItemUsagecountulMaxValue = 18446744073709551615
)

// MaxDataBurstVolume ::= INTEGER
type Maxdataburstvolume int64

// MaxPacketLossRate ::= INTEGER
type Maxpacketlossrate int64

// NetworkInstance ::= INTEGER
type Networkinstance int64

// PDCP-SN ::= INTEGER
type PdcpSn int64
const (
	PdcpSnMinValue = 0
	PdcpSnMaxValue = 262143
)

// PDU-Session-ID ::= INTEGER
type PduSessionId int64
const (
	PduSessionIdMinValue = 0
	PduSessionIdMaxValue = 255
)

// PER-Exponent ::= INTEGER
type PerExponent int64

// PER-Scalar ::= INTEGER
type PerScalar int64

// PPI ::= INTEGER
type Ppi int64

// PacketDelayBudget ::= INTEGER
type Packetdelaybudget int64

// Periodicity ::= INTEGER
type Periodicity int64

// PrivateIE-ID-local ::= INTEGER
type PrivateieIdLocal int64
const (
	PrivateieIdLocalMinValue = 0
	PrivateieIdLocalMaxValue = Maxprivateies
)

// ProcedureCode ::= INTEGER
type Procedurecode int64
const (
	ProcedurecodeMinValue = 0
	ProcedurecodeMaxValue = 255
)

// ProtocolExtensionID ::= INTEGER
type Protocolextensionid int64
const (
	ProtocolextensionidMinValue = 0
	ProtocolextensionidMaxValue = Maxprotocolextensions
)

// ProtocolIE-ContainerList { ::= INTEGER
type ProtocolieContainerlist int64

// ProtocolIE-ID ::= INTEGER
type ProtocolieId int64
const (
	ProtocolieIdMinValue = 0
	ProtocolieIdMaxValue = Maxprotocolies
)

// QCI ::= INTEGER
type Qci int64
const (
	QciMinValue = 0
	QciMaxValue = 255
)

// QoS-Flow-Identifier ::= INTEGER
type QosFlowIdentifier int64
const (
	QosFlowIdentifierMinValue = 0
	QosFlowIdentifierMaxValue = 63
)

// QoSPriorityLevel ::= INTEGER
type Qosprioritylevel int64

// SubscriberProfileIDforRFP ::= INTEGER
type Subscriberprofileidforrfp int64

// TransactionID ::= INTEGER
type Transactionid int64

