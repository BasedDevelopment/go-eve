package eve

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	LastLogin time.Time `json:"last_login"`
	Created   time.Time
}

type VirtualMachine struct {
	ID          uuid.UUID
	HV          uuid.UUID
	HostName    string
	User        uuid.UUID
	CPU         int
	Memory      int64
	Nics        map[string]Nic
	Storages    map[string]Storage
	Created     time.Time
	Updated     time.Time
	Remarks     string
	State       Status
	StateStr    string
	StateReason string
}

type Nic struct {
	ID      uuid.UUID
	Name    string
	MAC     string
	IP      []net.IP
	Created time.Time
	Updated time.Time
	Remarks string
	State   string
}

type Storage struct {
	ID      uuid.UUID
	Size    int
	Created time.Time
	Updated time.Time
	Remarks string
}
