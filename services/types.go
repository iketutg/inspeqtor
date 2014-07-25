package services

type ProcessId int32
type Status uint8

const (
	Unknown Status = iota
	Down
	Starting
	Up
	Stopping
	// if we try to start the service and it does not start, we mark it as broken so we
	// don't continually try to start a broken service.
	Broken
)

//
// Your init system(s) manages services.  We use
// the init system to:
// 1. find the associated PID
// 2. start/stop/restart the service
//
type InitSystem interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Look up PID for the given service name, returns
	// positive integer if successful, -1 if the service
	// name was not found or error if there was an
	// unexpected failure.
	LookupService(name string) (ProcessId, Status, error)

	Start(name string) error
	Stop(name string) error
	Status(name string) error
}