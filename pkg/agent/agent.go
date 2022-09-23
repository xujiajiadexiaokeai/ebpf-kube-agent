package agent

type Agent struct {
	Pod    string
	Status AgentStatus
	Jobs   []*EbpfJob
}

type AgentStatus string

type AgentOptions struct {
}

// NewAgent: create new agent
func NewAgent(ao *AgentOptions) *Agent {
	return &Agent{}
}
