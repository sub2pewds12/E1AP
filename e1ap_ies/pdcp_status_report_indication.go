package e1ap_ies

// PDCPStatusReportIndication represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1623
type PDCPStatusReportIndication int32

const (
	PDCPStatusReportIndication_Downlink PDCPStatusReportIndication = 0
	PDCPStatusReportIndication_Uplink   PDCPStatusReportIndication = 1
	PDCPStatusReportIndication_Both     PDCPStatusReportIndication = 2
)
