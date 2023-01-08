package eve

// Status "enum" type for VMs, HVs, etc...
type Status int8

// Status constants for HVs, VMs, etc..
// https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainState
const (
	StatusUnknown Status = iota
	StatusRunning
	StatusBlocked
	StatusPaused
	StatusShutdown
	StatusShutoff
	StatusCrashed
	StatusPMSuspended
)

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "unknown"
	case StatusRunning:
		return "running"
	case StatusBlocked:
		return "blocked"
	case StatusPaused:
		return "paused"
	case StatusShutdown:
		return "shutdown"
	case StatusShutoff:
		return "shutoff"
	case StatusCrashed:
		return "crashed"
	case StatusPMSuspended:
		return "pmsuspended"
	default:
		return "unknown"
	}
}
