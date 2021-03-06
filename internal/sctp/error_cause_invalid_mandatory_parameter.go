package sctp

// errorCauseInvalidMandatoryParameter represents an SCTP error cause
type errorCauseInvalidMandatoryParameter struct {
	errorCauseHeader
}

func (e *errorCauseInvalidMandatoryParameter) marshal() ([]byte, error) {
	return e.errorCauseHeader.marshal()
}

func (e *errorCauseInvalidMandatoryParameter) unmarshal(raw []byte) error {
	return e.errorCauseHeader.unmarshal(raw)
}
