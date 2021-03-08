/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package reverseproxycache

import (
	"testing"

	"github.com/tricksterproxy/trickster/pkg/backends"
	bo "github.com/tricksterproxy/trickster/pkg/backends/options"
)

func TestReverseProxyCacheClientInterfacing(t *testing.T) {

	// this test ensures the client will properly conform to the
	// Client interface

	c, err := NewClient("test", nil, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var o backends.Backend = c

	if o.Name() != "test" {
		t.Errorf("expected %s got %s", "test", o.Name())
	}

}

func TestNewNewClient(t *testing.T) {
	c, err := NewClient("test", bo.New(), nil, nil)
	if err != nil {
		t.Error(err)
	}
	if c == nil {
		t.Errorf("expected client named %s", "test")
	}
}