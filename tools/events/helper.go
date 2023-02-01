/*
Copyright 2021 The Kubernetes Authors.

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

package events

import (
	"github.com/yubo/golib/api"
	"github.com/yubo/golib/fields"
)

// GetFieldSelector returns the appropriate field selector based on the API version being used to communicate with the server.
// The returned field selector can be used with List and Watch to filter desired events.
func GetFieldSelector(regardingName string, regardingUID api.UID) (fields.Selector, error) {
	field := fields.Set{}

	if len(regardingName) > 0 {
		field["name"] = regardingName
	}

	if len(regardingUID) > 0 {
		field["uid"] = string(regardingUID)
	}

	return field.AsSelector(), nil
}
