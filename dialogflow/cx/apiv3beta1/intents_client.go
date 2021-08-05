// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newIntentsClientHook clientHook

// IntentsCallOptions contains the retry settings for each method of IntentsClient.
type IntentsCallOptions struct {
	ListIntents  []gax.CallOption
	GetIntent    []gax.CallOption
	CreateIntent []gax.CallOption
	UpdateIntent []gax.CallOption
	DeleteIntent []gax.CallOption
}

func defaultIntentsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIntentsCallOptions() *IntentsCallOptions {
	return &IntentsCallOptions{
		ListIntents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalIntentsClient is an interface that defines the methods availaible from Dialogflow API.
type internalIntentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListIntents(context.Context, *cxpb.ListIntentsRequest, ...gax.CallOption) *IntentIterator
	GetIntent(context.Context, *cxpb.GetIntentRequest, ...gax.CallOption) (*cxpb.Intent, error)
	CreateIntent(context.Context, *cxpb.CreateIntentRequest, ...gax.CallOption) (*cxpb.Intent, error)
	UpdateIntent(context.Context, *cxpb.UpdateIntentRequest, ...gax.CallOption) (*cxpb.Intent, error)
	DeleteIntent(context.Context, *cxpb.DeleteIntentRequest, ...gax.CallOption) error
}

// IntentsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Intents.
type IntentsClient struct {
	// The internal transport-dependent client.
	internalClient internalIntentsClient

	// The call options for this service.
	CallOptions *IntentsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IntentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IntentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IntentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListIntents returns the list of all intents in the specified agent.
func (c *IntentsClient) ListIntents(ctx context.Context, req *cxpb.ListIntentsRequest, opts ...gax.CallOption) *IntentIterator {
	return c.internalClient.ListIntents(ctx, req, opts...)
}

// GetIntent retrieves the specified intent.
func (c *IntentsClient) GetIntent(ctx context.Context, req *cxpb.GetIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	return c.internalClient.GetIntent(ctx, req, opts...)
}

// CreateIntent creates an intent in the specified agent.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *IntentsClient) CreateIntent(ctx context.Context, req *cxpb.CreateIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	return c.internalClient.CreateIntent(ctx, req, opts...)
}

// UpdateIntent updates the specified intent.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *IntentsClient) UpdateIntent(ctx context.Context, req *cxpb.UpdateIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	return c.internalClient.UpdateIntent(ctx, req, opts...)
}

// DeleteIntent deletes the specified intent.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *IntentsClient) DeleteIntent(ctx context.Context, req *cxpb.DeleteIntentRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteIntent(ctx, req, opts...)
}

// intentsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type intentsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IntentsClient
	CallOptions **IntentsCallOptions

	// The gRPC API client.
	intentsClient cxpb.IntentsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIntentsClient creates a new intents client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Intents.
func NewIntentsClient(ctx context.Context, opts ...option.ClientOption) (*IntentsClient, error) {
	clientOpts := defaultIntentsGRPCClientOptions()
	if newIntentsClientHook != nil {
		hookOpts, err := newIntentsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IntentsClient{CallOptions: defaultIntentsCallOptions()}

	c := &intentsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		intentsClient:    cxpb.NewIntentsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *intentsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *intentsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *intentsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *intentsGRPCClient) ListIntents(ctx context.Context, req *cxpb.ListIntentsRequest, opts ...gax.CallOption) *IntentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIntents[0:len((*c.CallOptions).ListIntents):len((*c.CallOptions).ListIntents)], opts...)
	it := &IntentIterator{}
	req = proto.Clone(req).(*cxpb.ListIntentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Intent, string, error) {
		resp := &cxpb.ListIntentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.intentsClient.ListIntents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIntents(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *intentsGRPCClient) GetIntent(ctx context.Context, req *cxpb.GetIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIntent[0:len((*c.CallOptions).GetIntent):len((*c.CallOptions).GetIntent)], opts...)
	var resp *cxpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.GetIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) CreateIntent(ctx context.Context, req *cxpb.CreateIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIntent[0:len((*c.CallOptions).CreateIntent):len((*c.CallOptions).CreateIntent)], opts...)
	var resp *cxpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.CreateIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) UpdateIntent(ctx context.Context, req *cxpb.UpdateIntentRequest, opts ...gax.CallOption) (*cxpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "intent.name", url.QueryEscape(req.GetIntent().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateIntent[0:len((*c.CallOptions).UpdateIntent):len((*c.CallOptions).UpdateIntent)], opts...)
	var resp *cxpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.UpdateIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) DeleteIntent(ctx context.Context, req *cxpb.DeleteIntentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIntent[0:len((*c.CallOptions).DeleteIntent):len((*c.CallOptions).DeleteIntent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.intentsClient.DeleteIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// IntentIterator manages a stream of *cxpb.Intent.
type IntentIterator struct {
	items    []*cxpb.Intent
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Intent, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IntentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IntentIterator) Next() (*cxpb.Intent, error) {
	var item *cxpb.Intent
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IntentIterator) bufLen() int {
	return len(it.items)
}

func (it *IntentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
