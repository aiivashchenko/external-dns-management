/*
 * Copyright 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 *
 */

package ovh

import (
	"k8s.io/client-go/util/flowcontrol"

	"github.com/gardener/external-dns-management/pkg/controller/provider/ovh/client"
	"github.com/gardener/external-dns-management/pkg/dns/provider"
	"github.com/gardener/external-dns-management/pkg/dns/provider/raw"
)

type Access interface {
	ListZones(consume func(zone client.DNSZone) (bool, error)) error
	ListRecords(zoneId string, consume func(record client.DNSRecord) (bool, error)) error

	raw.Executor
}

type access struct {
	client      *client.Client
	metrics     provider.Metrics
	rateLimiter flowcontrol.RateLimiter
}

func NewAccess(config *client.Config, metrics provider.Metrics, rateLimiter flowcontrol.RateLimiter) (Access, error) {
	client, err := client.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &access{client: client, metrics: metrics, rateLimiter: rateLimiter}, nil
}

func (a access) ListZones(consume func(zone client.DNSZone) (bool, error)) error {
	panic("implement me")
}

func (a access) ListRecords(zoneId string, consume func(record client.DNSRecord) (bool, error)) error {
	panic("implement me")
}

func (a access) CreateRecord(r raw.Record) error {
	panic("implement me")
}

func (a access) UpdateRecord(r raw.Record) error {
	panic("implement me")
}

func (a access) DeleteRecord(r raw.Record) error {
	panic("implement me")
}

func (a access) NewRecord(fqdn, rtype, value string, zone provider.DNSHostedZone, ttl int64) raw.Record {
	panic("implement me")
}
