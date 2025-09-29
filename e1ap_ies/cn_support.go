package e1ap_ies

// CNSupport represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:242
type CNSupport int32

const (
	CNSupport_CEpc CNSupport = 0
	CNSupport_C5gc CNSupport = 1
	CNSupport_Both CNSupport = 2
)
