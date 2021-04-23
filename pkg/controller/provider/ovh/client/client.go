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

package client

import (
	api "github.com/ovh/go-ovh/ovh"
	"github.com/pkg/errors"
	"k8s.io/klog"
)

const (
	COVHEUEndpoint        = "ovh-eu"
	COVHCAEndpoint        = "ovh-ca"
	COVHUSEndpoint        = "ovh-us"
	CKimsufiEUEndpoint    = "kimsufi-eu"
	CKimsufiCAEndpoint    = "kimsufi-ca"
	CSoYouStartEUEndpoint = "soyoustart-eu"
	CSoYouStartCAEndpoint = "soyoustart-ca"

	COVHCZSubsidiary = "CZ"
	COVHDESubsidiary = "DE"
	COVHESSubsidiary = "ES"
	COVHEUSubsidiary = "EU"
	COVHFISubsidiary = "FI"
	COVHFRSubsidiary = "FR"
	COVHGBSubsidiary = "GB"
	COVHIESubsidiary = "IE"
	COVHITSubsidiary = "IT"
	COVHLTSubsidiary = "LT"
	COVHMASubsidiary = "MA"
	COVHNLSubsidiary = "NL"
	COVHPLSubsidiary = "PL"
	COVHPTSubsidiary = "PT"
	COVHSNSubsidiary = "SN"
	COVHTNSubsidiary = "TN"
)

var CRequiredPermissions = map[string][]string{
	"/domain":   api.ReadWrite,
	"/domain/*": api.ReadWrite,
}

type Config struct {
	Endpoint          string
	ApplicationKey    string
	ApplicationSecret string
	ConsumerKey       string
}

type Client struct {
	ovh *api.Client
}

func NewClient(config *Config) (*Client, error) {
	endpointClient, err := api.NewClient(config.Endpoint, config.ApplicationKey, config.ApplicationSecret, config.ConsumerKey)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get ovh")
	}

	if config.ConsumerKey != "" {
		return &Client{
			ovh: endpointClient,
		}, nil
	}

	ckRequest := endpointClient.NewCkRequest()
	for k, v := range CRequiredPermissions {
		ckRequest.AddRules(v, k)
	}

	response, err := ckRequest.Do()
	if err != nil {
		return nil, errors.Wrap(err, "unable to perform consumer key request")
	}
	klog.Infof("to validate consumer key: %s visit URL: %s", response.ConsumerKey, response.ValidationURL)

	return &Client{
		ovh: endpointClient,
	}, nil
}
