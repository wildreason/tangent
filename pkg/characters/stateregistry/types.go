package stateregistry

// StateFrame represents a single frame in a state animation
type StateFrame struct {
	Lines []string `json:"lines"`
}

// StateDefinition represents a complete state with all its frames
type StateDefinition struct {
	Name   string       `json:"name"`
	Frames []StateFrame `json:"frames"`
	FPS    int          `json:"fps,omitempty"` // Optional FPS override (default: 5)
}

// Registry holds all available states
type Registry struct {
	states map[string]StateDefinition
}

// NewRegistry creates a new empty registry
func NewRegistry() *Registry {
	return &Registry{
		states: make(map[string]StateDefinition),
	}
}

// Register adds a state to the registry
func (r *Registry) Register(state StateDefinition) {
	r.states[state.Name] = state
}

// Get retrieves a state by name
func (r *Registry) Get(name string) (StateDefinition, bool) {
	state, ok := r.states[name]
	return state, ok
}

// List returns all state names
func (r *Registry) List() []string {
	names := make([]string, 0, len(r.states))
	for name := range r.states {
		names = append(names, name)
	}
	return names
}

// All returns all states
func (r *Registry) All() map[string]StateDefinition {
	return r.states
}
