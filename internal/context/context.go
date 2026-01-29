package context

// The Context interface provides access to shell package
// functionality while avoiding circular dependencies.
type Context interface {
	SetIsRunning(bool)
}
