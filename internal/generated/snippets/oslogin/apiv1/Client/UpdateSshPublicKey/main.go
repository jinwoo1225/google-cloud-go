// Copyright 2021 Google LLC
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

// [START oslogin_v1_generated_OsLoginService_UpdateSshPublicKey_sync]

package main

import (
	"context"

	oslogin "cloud.google.com/go/oslogin/apiv1"
	osloginpb "google.golang.org/genproto/googleapis/cloud/oslogin/v1"
)

func main() {
	// import osloginpb "google.golang.org/genproto/googleapis/cloud/oslogin/v1"

	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osloginpb.UpdateSshPublicKeyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateSshPublicKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END oslogin_v1_generated_OsLoginService_UpdateSshPublicKey_sync]