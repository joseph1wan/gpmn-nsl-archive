package inmemory

import (
"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

type InMemory struct {
maintenance map[int]datastore.MaintenanceRequest
}

func (m *InMemory) Init() error {
m.maintenance = map[int]datastore.MaintenanceRequest{}
return nil
}
