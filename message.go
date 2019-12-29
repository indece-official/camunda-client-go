package camunda_client_go

// Message a client for Message API
type Message struct {
	client *Client
}

// CorrelationKey defines a key-value-pair used for correlating running processes
type CorrelationKey struct {
	// The variable's value
	Value interface{} `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
}

// ReqMessage a JSON object corresponding to the Message interface in the engine
type ReqMessage struct {
	// MessageName is the name of the message to deliver.
	MessageName string `json:"messageName"`

	// BusinessKey is used for correlation of process instances that wait for incoming
	// messages. Will only correlate to executions that belong to a process instance with
	// the provided business key.
	BusinessKey *string `json:"businessKey,omitempty"`

	// TenantID is used to correlate the message for a tenant with the given id. Will only
	// correlate to executions and process definitions which belong to the tenant.
	TenantID *string `json:"tenantId,omitempty"`

	// WithoutTenantID is a boolean value that indicates whether the message should only
	// be correlated to executions and process definitions which belong to no tenant or
	// not. Value may only be true, as false is the default behavior.
	WithoutTenantID *bool `json:"withoutTenantId,omitempty"`

	// ProcessInstanceID is used to correlate the message to the process instance with the
	// given id. Must not be supplied in conjunction with a tenantId.
	ProcessInstanceID *string `json:"processInstanceId,omitempty"`

	// CorrelationKeys is used for correlation of process instances that wait for incoming
	// messages. Has to be a JSON object containing key-value pairs that are matched
	// against process instance variables during correlation. Each key is a variable name
	// and each value a JSON variable value object with the following properties.
	//
	// Note: Process instance variables are the global variables of a process instance.
	// Local variables of child executions (such as in subprocesses) are not considered!
	CorrelationKeys *map[string]CorrelationKey `json:"correlationKeys,omitempty"`

	// ProcessVariables	is a map of variables that is injected into the triggered execution
	// or process instance after the message has been delivered. Each key is a variable
	// name and each value a JSON variable value object with the following properties.
	//
	// For variables of type File the value has to be submitted as Base64 encoded string.
	ProcessVariables *map[string]Variable `json:"processVariables,omitempty"`

	// All is a boolean value that indicates whether the message should be correlated to
	// exactly one entity or multiple entities. If the value is set to false, the message
	// will be correlated to exactly one entity (execution or process definition). If the
	// value is set to true, the message will be correlated to multiple executions and a
	// process definition that can be instantiated by this message in one go.
	All *bool `json:"all,omitempty"`
}

func (m *Message) SendMessage(reqMsg ReqMessage) (err error) {
	res, err := m.client.doPostJson("/message", map[string]string{}, reqMsg)
	if err != nil {
		return
	}

	err = m.client.checkResponse(res)
	return
}
