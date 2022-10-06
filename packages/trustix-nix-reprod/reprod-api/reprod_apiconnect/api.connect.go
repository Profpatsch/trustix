// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: reprod-api/api.proto

package reprod_apiconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	reprod_api "github.com/nix-community/trustix/packages/trustix-nix-reprod/reprod-api"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ReproducibilityAPIName is the fully-qualified name of the ReproducibilityAPI service.
	ReproducibilityAPIName = "ReproducibilityAPI"
)

// ReproducibilityAPIClient is a client for the ReproducibilityAPI service.
type ReproducibilityAPIClient interface {
	DerivationReproducibility(context.Context, *connect_go.Request[reprod_api.DerivationReproducibilityRequest]) (*connect_go.Response[reprod_api.DerivationReproducibilityResponse], error)
	AttrReproducibilityTimeSeries(context.Context, *connect_go.Request[reprod_api.AttrReproducibilityTimeSeriesRequest]) (*connect_go.Response[reprod_api.AttrReproducibilityTimeSeriesResponse], error)
	SuggestAttribute(context.Context, *connect_go.Request[reprod_api.SuggestsAttributeRequest]) (*connect_go.Response[reprod_api.SuggestAttributeResponse], error)
	Diff(context.Context, *connect_go.Request[reprod_api.DiffRequest]) (*connect_go.Response[reprod_api.DiffResponse], error)
}

// NewReproducibilityAPIClient constructs a client for the ReproducibilityAPI service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReproducibilityAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReproducibilityAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &reproducibilityAPIClient{
		derivationReproducibility: connect_go.NewClient[reprod_api.DerivationReproducibilityRequest, reprod_api.DerivationReproducibilityResponse](
			httpClient,
			baseURL+"/.ReproducibilityAPI/DerivationReproducibility",
			opts...,
		),
		attrReproducibilityTimeSeries: connect_go.NewClient[reprod_api.AttrReproducibilityTimeSeriesRequest, reprod_api.AttrReproducibilityTimeSeriesResponse](
			httpClient,
			baseURL+"/.ReproducibilityAPI/AttrReproducibilityTimeSeries",
			opts...,
		),
		suggestAttribute: connect_go.NewClient[reprod_api.SuggestsAttributeRequest, reprod_api.SuggestAttributeResponse](
			httpClient,
			baseURL+"/.ReproducibilityAPI/SuggestAttribute",
			opts...,
		),
		diff: connect_go.NewClient[reprod_api.DiffRequest, reprod_api.DiffResponse](
			httpClient,
			baseURL+"/.ReproducibilityAPI/Diff",
			opts...,
		),
	}
}

// reproducibilityAPIClient implements ReproducibilityAPIClient.
type reproducibilityAPIClient struct {
	derivationReproducibility     *connect_go.Client[reprod_api.DerivationReproducibilityRequest, reprod_api.DerivationReproducibilityResponse]
	attrReproducibilityTimeSeries *connect_go.Client[reprod_api.AttrReproducibilityTimeSeriesRequest, reprod_api.AttrReproducibilityTimeSeriesResponse]
	suggestAttribute              *connect_go.Client[reprod_api.SuggestsAttributeRequest, reprod_api.SuggestAttributeResponse]
	diff                          *connect_go.Client[reprod_api.DiffRequest, reprod_api.DiffResponse]
}

// DerivationReproducibility calls ReproducibilityAPI.DerivationReproducibility.
func (c *reproducibilityAPIClient) DerivationReproducibility(ctx context.Context, req *connect_go.Request[reprod_api.DerivationReproducibilityRequest]) (*connect_go.Response[reprod_api.DerivationReproducibilityResponse], error) {
	return c.derivationReproducibility.CallUnary(ctx, req)
}

// AttrReproducibilityTimeSeries calls ReproducibilityAPI.AttrReproducibilityTimeSeries.
func (c *reproducibilityAPIClient) AttrReproducibilityTimeSeries(ctx context.Context, req *connect_go.Request[reprod_api.AttrReproducibilityTimeSeriesRequest]) (*connect_go.Response[reprod_api.AttrReproducibilityTimeSeriesResponse], error) {
	return c.attrReproducibilityTimeSeries.CallUnary(ctx, req)
}

// SuggestAttribute calls ReproducibilityAPI.SuggestAttribute.
func (c *reproducibilityAPIClient) SuggestAttribute(ctx context.Context, req *connect_go.Request[reprod_api.SuggestsAttributeRequest]) (*connect_go.Response[reprod_api.SuggestAttributeResponse], error) {
	return c.suggestAttribute.CallUnary(ctx, req)
}

// Diff calls ReproducibilityAPI.Diff.
func (c *reproducibilityAPIClient) Diff(ctx context.Context, req *connect_go.Request[reprod_api.DiffRequest]) (*connect_go.Response[reprod_api.DiffResponse], error) {
	return c.diff.CallUnary(ctx, req)
}

// ReproducibilityAPIHandler is an implementation of the ReproducibilityAPI service.
type ReproducibilityAPIHandler interface {
	DerivationReproducibility(context.Context, *connect_go.Request[reprod_api.DerivationReproducibilityRequest]) (*connect_go.Response[reprod_api.DerivationReproducibilityResponse], error)
	AttrReproducibilityTimeSeries(context.Context, *connect_go.Request[reprod_api.AttrReproducibilityTimeSeriesRequest]) (*connect_go.Response[reprod_api.AttrReproducibilityTimeSeriesResponse], error)
	SuggestAttribute(context.Context, *connect_go.Request[reprod_api.SuggestsAttributeRequest]) (*connect_go.Response[reprod_api.SuggestAttributeResponse], error)
	Diff(context.Context, *connect_go.Request[reprod_api.DiffRequest]) (*connect_go.Response[reprod_api.DiffResponse], error)
}

// NewReproducibilityAPIHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReproducibilityAPIHandler(svc ReproducibilityAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/.ReproducibilityAPI/DerivationReproducibility", connect_go.NewUnaryHandler(
		"/.ReproducibilityAPI/DerivationReproducibility",
		svc.DerivationReproducibility,
		opts...,
	))
	mux.Handle("/.ReproducibilityAPI/AttrReproducibilityTimeSeries", connect_go.NewUnaryHandler(
		"/.ReproducibilityAPI/AttrReproducibilityTimeSeries",
		svc.AttrReproducibilityTimeSeries,
		opts...,
	))
	mux.Handle("/.ReproducibilityAPI/SuggestAttribute", connect_go.NewUnaryHandler(
		"/.ReproducibilityAPI/SuggestAttribute",
		svc.SuggestAttribute,
		opts...,
	))
	mux.Handle("/.ReproducibilityAPI/Diff", connect_go.NewUnaryHandler(
		"/.ReproducibilityAPI/Diff",
		svc.Diff,
		opts...,
	))
	return "/.ReproducibilityAPI/", mux
}

// UnimplementedReproducibilityAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedReproducibilityAPIHandler struct{}

func (UnimplementedReproducibilityAPIHandler) DerivationReproducibility(context.Context, *connect_go.Request[reprod_api.DerivationReproducibilityRequest]) (*connect_go.Response[reprod_api.DerivationReproducibilityResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ReproducibilityAPI.DerivationReproducibility is not implemented"))
}

func (UnimplementedReproducibilityAPIHandler) AttrReproducibilityTimeSeries(context.Context, *connect_go.Request[reprod_api.AttrReproducibilityTimeSeriesRequest]) (*connect_go.Response[reprod_api.AttrReproducibilityTimeSeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ReproducibilityAPI.AttrReproducibilityTimeSeries is not implemented"))
}

func (UnimplementedReproducibilityAPIHandler) SuggestAttribute(context.Context, *connect_go.Request[reprod_api.SuggestsAttributeRequest]) (*connect_go.Response[reprod_api.SuggestAttributeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ReproducibilityAPI.SuggestAttribute is not implemented"))
}

func (UnimplementedReproducibilityAPIHandler) Diff(context.Context, *connect_go.Request[reprod_api.DiffRequest]) (*connect_go.Response[reprod_api.DiffResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ReproducibilityAPI.Diff is not implemented"))
}
