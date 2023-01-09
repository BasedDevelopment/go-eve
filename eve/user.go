package eve

import (
	"time"

	"github.com/google/uuid"
)

// GetSelf returns the data for the current authenticated user
func (c *Client) GetSelf() (User, error) {
	_, body, err := c.makeRequest("/me", "GET", nil)

	if err != nil {
		return User{}, err
	}

	m, err := time.Parse(time.RFC3339, body["created"].(string))

	if err != nil {
		return User{}, err
	}

	l, err := time.Parse(time.RFC3339, body["last_login"].(string))

	if err != nil {
		return User{}, err
	}

	uid, err := uuid.Parse(body["id"].(string))

	if err != nil {
		return User{}, nil
	}

	return User{
		Name:      body["name"].(string),
		ID:        uid,
		Email:     body["email"].(string),
		Created:   m,
		LastLogin: l,
	}, nil
}

// UpdateSelf updates the current authenticated user with the data in the u
// argument and returns the User back with the changes applied.
func (c *Client) UpdateSelf(u User) (User, error) {
	return User{}, nil
}

// GetVirtualMachines fetches all of the users accessible VMs on the network
func (c *Client) GetVirtualMachines() ([]VirtualMachine, error) {
	return []VirtualMachine{}, nil
}

// GetVirtualMachine fetches the virtual machine with id `id` and returns it
func (c *Client) GetVirtualMachine(id uuid.UUID) (VirtualMachine, error) {
	return VirtualMachine{}, nil
}

// UpdateVirtualMachines updates the VM with id `id` with the data provided in
// `v`
func (c *Client) UpdateVirtualMachine(id uuid.UUID, v VirtualMachine) (VirtualMachine, error) {
	return VirtualMachine{}, nil
}
