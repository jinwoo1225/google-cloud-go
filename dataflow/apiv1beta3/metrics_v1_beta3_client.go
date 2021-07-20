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

package dataflow

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newMetricsV1Beta3ClientHook clientHook

// MetricsV1Beta3CallOptions contains the retry settings for each method of MetricsV1Beta3Client.
type MetricsV1Beta3CallOptions struct {
	GetJobMetrics            []gax.CallOption
	GetJobExecutionDetails   []gax.CallOption
	GetStageExecutionDetails []gax.CallOption
}

func defaultMetricsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetricsV1Beta3CallOptions() *MetricsV1Beta3CallOptions {
	return &MetricsV1Beta3CallOptions{
		GetJobMetrics:            []gax.CallOption{},
		GetJobExecutionDetails:   []gax.CallOption{},
		GetStageExecutionDetails: []gax.CallOption{},
	}
}

// internalMetricsV1Beta3Client is an interface that defines the methods availaible from Dataflow API.
type internalMetricsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetJobMetrics(context.Context, *dataflowpb.GetJobMetricsRequest, ...gax.CallOption) (*dataflowpb.JobMetrics, error)
	GetJobExecutionDetails(context.Context, *dataflowpb.GetJobExecutionDetailsRequest, ...gax.CallOption) *StageSummaryIterator
	GetStageExecutionDetails(context.Context, *dataflowpb.GetStageExecutionDetailsRequest, ...gax.CallOption) *WorkerDetailsIterator
}

// MetricsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Dataflow Metrics API lets you monitor the progress of Dataflow
// jobs.
type MetricsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalMetricsV1Beta3Client

	// The call options for this service.
	CallOptions *MetricsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *MetricsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetJobMetrics request the job status.
//
// To request the status of a job, we recommend using
// projects.locations.jobs.getMetrics with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.getMetrics is not recommended, as you can only request the
// status of jobs that are running in us-central1.
func (c *MetricsV1Beta3Client) GetJobMetrics(ctx context.Context, req *dataflowpb.GetJobMetricsRequest, opts ...gax.CallOption) (*dataflowpb.JobMetrics, error) {
	return c.internalClient.GetJobMetrics(ctx, req, opts...)
}

// GetJobExecutionDetails request detailed information about the execution status of the job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *MetricsV1Beta3Client) GetJobExecutionDetails(ctx context.Context, req *dataflowpb.GetJobExecutionDetailsRequest, opts ...gax.CallOption) *StageSummaryIterator {
	return c.internalClient.GetJobExecutionDetails(ctx, req, opts...)
}

// GetStageExecutionDetails request detailed information about the execution status of a stage of the
// job.
//
// EXPERIMENTAL.  This API is subject to change or removal without notice.
func (c *MetricsV1Beta3Client) GetStageExecutionDetails(ctx context.Context, req *dataflowpb.GetStageExecutionDetailsRequest, opts ...gax.CallOption) *WorkerDetailsIterator {
	return c.internalClient.GetStageExecutionDetails(ctx, req, opts...)
}

// metricsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MetricsV1Beta3Client
	CallOptions **MetricsV1Beta3CallOptions

	// The gRPC API client.
	metricsV1Beta3Client dataflowpb.MetricsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMetricsV1Beta3Client creates a new metrics v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Dataflow Metrics API lets you monitor the progress of Dataflow
// jobs.
func NewMetricsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*MetricsV1Beta3Client, error) {
	clientOpts := defaultMetricsV1Beta3GRPCClientOptions()
	if newMetricsV1Beta3ClientHook != nil {
		hookOpts, err := newMetricsV1Beta3ClientHook(ctx, clientHookParams{})
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
	client := MetricsV1Beta3Client{CallOptions: defaultMetricsV1Beta3CallOptions()}

	c := &metricsV1Beta3GRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		metricsV1Beta3Client: dataflowpb.NewMetricsV1Beta3Client(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *metricsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *metricsV1Beta3GRPCClient) GetJobMetrics(ctx context.Context, req *dataflowpb.GetJobMetricsRequest, opts ...gax.CallOption) (*dataflowpb.JobMetrics, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetJobMetrics[0:len((*c.CallOptions).GetJobMetrics):len((*c.CallOptions).GetJobMetrics)], opts...)
	var resp *dataflowpb.JobMetrics
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsV1Beta3Client.GetJobMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsV1Beta3GRPCClient) GetJobExecutionDetails(ctx context.Context, req *dataflowpb.GetJobExecutionDetailsRequest, opts ...gax.CallOption) *StageSummaryIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetJobExecutionDetails[0:len((*c.CallOptions).GetJobExecutionDetails):len((*c.CallOptions).GetJobExecutionDetails)], opts...)
	it := &StageSummaryIterator{}
	req = proto.Clone(req).(*dataflowpb.GetJobExecutionDetailsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.StageSummary, string, error) {
		var resp *dataflowpb.JobExecutionDetails
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.metricsV1Beta3Client.GetJobExecutionDetails(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetStages(), resp.GetNextPageToken(), nil
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

func (c *metricsV1Beta3GRPCClient) GetStageExecutionDetails(ctx context.Context, req *dataflowpb.GetStageExecutionDetailsRequest, opts ...gax.CallOption) *WorkerDetailsIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetStageExecutionDetails[0:len((*c.CallOptions).GetStageExecutionDetails):len((*c.CallOptions).GetStageExecutionDetails)], opts...)
	it := &WorkerDetailsIterator{}
	req = proto.Clone(req).(*dataflowpb.GetStageExecutionDetailsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.WorkerDetails, string, error) {
		var resp *dataflowpb.StageExecutionDetails
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.metricsV1Beta3Client.GetStageExecutionDetails(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetWorkers(), resp.GetNextPageToken(), nil
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

// StageSummaryIterator manages a stream of *dataflowpb.StageSummary.
type StageSummaryIterator struct {
	items    []*dataflowpb.StageSummary
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.StageSummary, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StageSummaryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StageSummaryIterator) Next() (*dataflowpb.StageSummary, error) {
	var item *dataflowpb.StageSummary
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StageSummaryIterator) bufLen() int {
	return len(it.items)
}

func (it *StageSummaryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkerDetailsIterator manages a stream of *dataflowpb.WorkerDetails.
type WorkerDetailsIterator struct {
	items    []*dataflowpb.WorkerDetails
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.WorkerDetails, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkerDetailsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkerDetailsIterator) Next() (*dataflowpb.WorkerDetails, error) {
	var item *dataflowpb.WorkerDetails
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkerDetailsIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkerDetailsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}