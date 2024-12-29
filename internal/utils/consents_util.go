package utils

type ConsentStatus int

const (
	ContractUnknown  ConsentStatus = iota
	ContractPending  ConsentStatus = 1
	ContractAccepted ConsentStatus = 2
	ContractDenied   ConsentStatus = 3
	ContractInvalid  ConsentStatus = 4
)
