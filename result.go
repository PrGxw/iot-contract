package contract

// Outcome is how an intention ended.
type Outcome string

const (
	OutcomeDone   Outcome = "done"
	OutcomeFailed Outcome = "failed"
)

// ResultCode is a contracted failure reason.
type ResultCode string

const (
	CodeDeviceTypeUnrecognised ResultCode = "device_type_unrecognised"
	CodeTimeout                ResultCode = "timeout"
	CodeDeviceRejected         ResultCode = "device_rejected"
	CodeInvalidIntention       ResultCode = "invalid_intention"
)

// Result settles an intention. It never carries a raw firmware reply.
type Result struct {
	Outcome   Outcome    `json:"outcome"`
	Code      ResultCode `json:"code,omitempty"`
	Detail    string     `json:"detail,omitempty"`     // adapter text, shown verbatim
	Type      DeviceType `json:"type,omitempty"`       // query_type only
	IdleNodes []string   `json:"idle_nodes,omitempty"` // mesh_idle_nodes only
}
