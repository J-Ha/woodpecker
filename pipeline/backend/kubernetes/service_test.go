// Copyright 2022 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	expected := `
	{
	  "metadata": {
	    "name": "bar",
	    "namespace": "foo",
	    "creationTimestamp": null
	  },
	  "spec": {
	    "ports": [
	      {
	        "port": 1234,
			"protocol": "UDP",
	        "targetPort": 1234
	      },
	      {
	        "port": 2345,
			"protocol": "UDP",
	        "targetPort": 2345
	      },
	      {
	        "port": 3456,
			"protocol": "TCP",
	        "targetPort": 3456
		},
		{
		  "port": 4567,
		  "protocol": "TCP",
		  "targetPort": 4567
		},
		{
		  "port": 5678,
		  "protocol": "SCTP",
		  "targetPort": 5678
		},
		{
		  "port": 6789,
		  "protocol": "SCTP",
		  "targetPort": 6789
		},
		{
		  "port": 7891,
		  "protocol": "TCP",
		  "targetPort": 7891
		}
	    ],
	    "selector": {
	      "step": "bar"
	    },
	    "type": "ClusterIP"
	  },
	  "status": {
	    "loadBalancer": {}
	  }
	}`

	s, _ := Service("foo", "bar", []string{"1234/udp", "2345/UDP", "3456/tcp", "4567/TCP", "5678/sctp", "6789/SCTP", "7891"})
	j, err := json.Marshal(s)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(j))
}
