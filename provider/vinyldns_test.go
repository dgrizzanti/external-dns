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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

type mockVinyldnsZoneInterface struct {
	mock.Mock
}

var mockVinylDNSProvider vinyldnsProvider

var vinylDNSZones []vinyldns.Zone
var vinylDNSRecords []vinyldns.RecordSet
var vinylDNSRecordSetUpdateResponse *vinyldns.RecordSetUpdateResponse

func TestVinylDNSServices(t *testing.T) {
	firstZone := vinyldns.Zone{
		ID:   "0",
		Name: "example.com",
	}
	secondZone := vinyldns.Zone{
		ID:   "1",
		Name: "example-beta.com",
	}
	vinylDNSZones = []vinyldns.Zone{firstZone, secondZone}

	firstRecord := vinyldns.RecordSet{
		ID:     "1",
		ZoneID: "example.com",
		Name:   "example",
		TTL:    3600,
		Type:   "CNAME",
	}
	vinylDNSRecords = []vinyldns.RecordSet{firstRecord}

	mockVinylDNS := &mockVinyldnsZoneInterface{}
	mockVinylDNS.On("Zones").Return(vinylDNSZones, nil)
	mockVinylDNS.On("RecordSets").Return(vinylDNSRecords, nil)

	mockVinylDNSProvider = vinyldnsProvider{client: mockVinylDNS}

	// Run tests on mock services
	t.Run("Zones", testVinylDNSProviderZones)
	t.Run("Records", testVinylDNSProviderRecords)
	// t.Run("ApplyChanges", testDnsimpleProviderApplyChanges)
	// t.Run("ApplyChanges/SkipUnknownZone", testDnsimpleProviderApplyChangesSkipsUnknown)
	// t.Run("SuitableZone", testDnsimpleSuitableZone)
	// t.Run("GetRecordID", testDnsimpleGetRecordID
}

func testVinylDNSProviderZones(t *testing.T) {
	result, err := mockVinylDNSProvider.client.Zones()
	assert.Nil(t, err)
	validateVinylDNSZones(t, result, vinylDNSZones)
}

func testVinylDNSProviderRecords(t *testing.T) {
	mockVinylDNSProvider.zoneFilter = NewZoneIDFilter([]string{"1"})
	result, err := mockVinylDNSProvider.Records()
	assert.Nil(t, err)
	assert.Equal(t, len(vinylDNSRecords), len(result))
}

func validateVinylDNSZones(t *testing.T, zones []vinyldns.Zone, expected []vinyldns.Zone) {
	require.Len(t, zones, len(expected))

	for i, v := range expected {
		assert.Equal(t, zones[i].Name, v.Name)
	}
}

func (m *mockVinyldnsZoneInterface) Zones() ([]vinyldns.Zone, error) {
	args := m.Called()
	var r0 []vinyldns.Zone

	if args.Get(0) != nil {
		r0 = args.Get(0).([]vinyldns.Zone)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSet(zoneID, recordSet string) (vinyldns.RecordSet, error) {
	args := m.Called(zoneID, recordSet)
	var r0 vinyldns.RecordSet

	if args.Get(0) != nil {
		r0 = args.Get(0).(vinyldns.RecordSet)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSets(id string) ([]vinyldns.RecordSet, error) {
	args := m.Called(id)
	var r0 []vinyldns.RecordSet

	if args.Get(0) != nil {
		r0 = args.Get(0).([]vinyldns.RecordSet)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetCreate(rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(rs)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetUpdate(rs *vinyldns.RecordSet) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(rs)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}

func (m *mockVinyldnsZoneInterface) RecordSetDelete(zoneID, recordSetID string) (*vinyldns.RecordSetUpdateResponse, error) {
	args := m.Called(zoneID, recordSetID)
	var r0 *vinyldns.RecordSetUpdateResponse

	if args.Get(0) != nil {
		r0 = args.Get(0).(*vinyldns.RecordSetUpdateResponse)
	}

	return r0, args.Error(1)
}
