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

package endpoints

import (
	"fmt"
	"net/http"

	"gitlab.com/ais8/cortex/pkg/operator/resources/job/batchapi"
	"gitlab.com/ais8/cortex/pkg/operator/schema"
	"gitlab.com/ais8/cortex/pkg/types/spec"
	"gitlab.com/ais8/cortex/pkg/types/userconfig"
	"github.com/gorilla/mux"
)

func StopBatchJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiName := vars["apiName"]
	jobID, err := getRequiredQueryParam("jobID", r)
	if err != nil {
		respondError(w, r, err)
		return
	}

	err = batchapi.StopJob(spec.JobKey{
		APIName: apiName,
		ID:      jobID,
		Kind:    userconfig.BatchAPIKind,
	})
	if err != nil {
		respondError(w, r, err)
		return
	}

	respond(w, schema.DeleteResponse{
		Message: fmt.Sprintf("stopped job %s", jobID),
	})
}
