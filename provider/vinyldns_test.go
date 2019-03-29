/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

var mockVinylDNSProvider vinyldnsProvider
var vinyldnsListZonesResponse []vinyldns.Zone

type mockVinylDNSServiceInterface struct {
	mock.Mock
}

func TestVinylDNSServices(t *testing.T) {
	// Setup example responses
	firstZone := vinyldns.Zone{
		ID:   "1",
		Name: "example.com",
	}
	secondZone := vinyldns.Zone{
		ID:   "2",
		Name: "example-beta.com",
	}
	vinyldnsListZonesResponse = []vinyldns.Zone{firstZone, secondZone}

	// Setup mock services
	mockDNS := &mockVinylDNSServiceInterface{}
	mockDNS.On("Zones").Return(&vinyldnsListZonesResponse, nil)
	mockDNS.On("RecordSet", "1").Return(&vinyldnsListZonesResponse, nil)
	mockDNS.On("RecordSets", "1").Return(&vinyldnsListZonesResponse, nil)
	mockDNS.On("RecordSetCreate", "1").Return(&vinyldnsListZonesResponse, nil)
	mockDNS.On("RecordSetUpdate", "1").Return(&vinyldnsListZonesResponse, nil)
	mockDNS.On("RecordSetDelete", "1").Return(&vinyldnsListZonesResponse, nil)

	mockVinylDNSProvider = vinyldnsProvider{client: mockDNS}
}

func (_m *mockVinylDNSServiceInterface) Zones() ([]vinyldns.Zone, error) {
	return nil, nil
}

func (_m *mockVinylDNSServiceInterface) RecordSets(id string) ([]vinyldns.RecordSet, error) {
	return nil, nil
}

func (_m *mockVinylDNSServiceInterface) RecordSet(zoneID, recordSet string) (vinyldns.RecordSet, error) {
	return vinyldns.RecordSet{}, nil
}

func (_m *mockVinylDNSServiceInterface) RecordSetCreate(zoneID string, rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	return nil, nil
}

func (_m *mockVinylDNSServiceInterface) RecordSetUpdate(zoneID, recordSetID string, rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	return nil, nil
}

func (_m *mockVinylDNSServiceInterface) RecordSetDelete(zoneID, recordSetID string) (*vinyldns.RecordSetUpdateResponse, error) {
	return nil, nil
}
