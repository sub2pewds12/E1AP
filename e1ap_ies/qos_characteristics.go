package e1ap_ies

// QOSCharacteristics From: 9_4_5_Information_Element_Definitions.txt:1936
// ASN.1 Data Type: CHOICE
const (
	QOSCharacteristicsPresentNothing uint64 = iota
	QOSCharacteristicsPresentNonDynamic5QI
	QOSCharacteristicsPresentDynamic5QI
)

type QOSCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor
	Dynamic5QI    *Dynamic5QIDescriptor
}
