package new

import "testing"

func TestNewService(t *testing.T) {
	service := NewService()
	service.Start()
	service.Stop()
}
