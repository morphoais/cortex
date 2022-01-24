/*
Copyright 2021 Cortex Labs, Inc.

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

package asyncapi

import (
	"fmt"

	"gitlab.com/g-ogawa/cortex/pkg/lib/errors"
)

const (
	ErrAPIUpdating = "ayncapi.api_updating"
)

func ErrorAPIUpdating(apiName string) error {
	return errors.WithStack(&errors.Error{
		Kind:    ErrAPIUpdating,
		Message: fmt.Sprintf("%s is updating (override with --force)", apiName),
	})
}
