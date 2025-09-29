package e1ap_ies

// QosMonitoringRequest represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2035
type QosMonitoringRequest int32

const (
	QosMonitoringRequest_Ul   QosMonitoringRequest = 0
	QosMonitoringRequest_Dl   QosMonitoringRequest = 1
	QosMonitoringRequest_Both QosMonitoringRequest = 2
)
