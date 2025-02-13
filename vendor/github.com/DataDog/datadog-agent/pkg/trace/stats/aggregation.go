// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package stats

import (
	"strconv"
	"strings"

	pb "github.com/DataDog/datadog-agent/pkg/proto/pbgo/trace"
	"github.com/DataDog/datadog-agent/pkg/trace/log"
	"github.com/DataDog/datadog-agent/pkg/trace/traceutil"
)

const (
	tagStatusCode  = "http.status_code"
	tagSynthetics  = "synthetics"
	tagPeerService = "peer.service"
	tagSpanKind    = "span.kind"
)

// Aggregation contains all the dimension on which we aggregate statistics.
type Aggregation struct {
	BucketsAggregationKey
	PayloadAggregationKey
}

// BucketsAggregationKey specifies the key by which a bucket is aggregated.
type BucketsAggregationKey struct {
	Service     string
	Name        string
	PeerService string
	Resource    string
	Type        string
	SpanKind    string
	StatusCode  uint32
	Synthetics  bool
}

// PayloadAggregationKey specifies the key by which a payload is aggregated.
type PayloadAggregationKey struct {
	Env         string
	Hostname    string
	Version     string
	ContainerID string
}

func getStatusCode(s *pb.Span) uint32 {
	code, ok := traceutil.GetMetric(s, tagStatusCode)
	if ok {
		// only 7.39.0+, for lesser versions, always use Meta
		return uint32(code)
	}
	strC := traceutil.GetMetaDefault(s, tagStatusCode, "")
	if strC == "" {
		return 0
	}
	c, err := strconv.ParseUint(strC, 10, 32)
	if err != nil {
		log.Debugf("Invalid status code %s. Using 0.", strC)
		return 0
	}
	return uint32(c)
}

// NewAggregationFromSpan creates a new aggregation from the provided span and env
func NewAggregationFromSpan(s *pb.Span, origin string, aggKey PayloadAggregationKey, enablePeerSvcAgg bool) Aggregation {
	synthetics := strings.HasPrefix(origin, tagSynthetics)
	agg := Aggregation{
		PayloadAggregationKey: aggKey,
		BucketsAggregationKey: BucketsAggregationKey{
			Resource:   s.Resource,
			Service:    s.Service,
			Name:       s.Name,
			SpanKind:   s.Meta[tagSpanKind],
			Type:       s.Type,
			StatusCode: getStatusCode(s),
			Synthetics: synthetics,
		},
	}
	if enablePeerSvcAgg {
		agg.PeerService = s.Meta[tagPeerService]
	}
	return agg
}

// NewAggregationFromGroup gets the Aggregation key of grouped stats.
func NewAggregationFromGroup(g *pb.ClientGroupedStats) Aggregation {
	return Aggregation{
		BucketsAggregationKey: BucketsAggregationKey{
			Resource:    g.Resource,
			Service:     g.Service,
			PeerService: g.PeerService,
			Name:        g.Name,
			SpanKind:    g.SpanKind,
			StatusCode:  g.HTTPStatusCode,
			Synthetics:  g.Synthetics,
		},
	}
}
