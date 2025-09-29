package e1ap_ies

type ProtocolIEField struct {
	Id          ProtocolIEID
	Criticality Criticality
	Value       any
}

type ProtocolIESingleContainer = ProtocolIEField

type ProtocolIEContainer []ProtocolIEField

type ProtocolExtensionField struct {
	Id          ProtocolExtensionID
	Criticality Criticality
	Value       any
}

type ProtocolExtensionContainer []ProtocolExtensionField

type PrivateIEField struct {
	Id          PrivateIEID
	Criticality Criticality
	Value       any
}

type PrivateIEContainer []PrivateIEField
