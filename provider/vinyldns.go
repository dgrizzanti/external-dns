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
	"fmt"
	"os"

	"github.com/vinyldns/go-vinyldns/vinyldns"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	log "github.com/sirupsen/logrus"
)

type vinyldnsZoneInterface interface {
	Zones() ([]vinyldns.Zone, error)
	RecordSets(id string) ([]vinyldns.RecordSet, error)
}

type vinyldnsProvider struct {
	client     vinyldnsZoneInterface
	dryRun     bool
	zoneFilter ZoneIDFilter
}

type vinyldnsChange struct {
	Action string
}

func NewVinylDNSProvider(zoneFilter ZoneIDFilter, dryRun bool) (Provider, error) {
	_, ok := os.LookupEnv("VINYLDNS_ACCESS_KEY")
	if !ok {
		return nil, fmt.Errorf("no vinyldns access key found")
	}

	client := vinyldns.NewClientFromEnv()

	provider := &vinyldnsProvider{
		client:     client,
		dryRun:     dryRun,
		zoneFilter: zoneFilter,
	}
	return provider, nil
}

func (p *vinyldnsProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.client.Zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		log.Infof(zone.Name + " " + zone.ID)
		if !p.zoneFilter.Match(zone.ID) {
			continue
		}

		records, err := p.client.RecordSets(zone.ID)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			if supportedRecordType(string(r.Type)) {
				name := fmt.Sprintf("%s.%s", r.Name, zone.Name)

				// root name is identified by the empty string and should be
				// translated to zone name for the endpoint entry.
				if r.Name == "" {
					name = zone.Name
				}

				log.Infof(r.Name)

				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(name, string(r.Type), endpoint.TTL(r.TTL), "1.1.1.1"))
			}
		}
	}

	return endpoints, nil
}

func (p *vinyldnsProvider) submitChanges(changes []*vinyldnsChange) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}
	return nil
}

func (p *vinyldnsProvider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*vinyldnsChange, 0)

	return p.submitChanges(combinedChanges)
}
