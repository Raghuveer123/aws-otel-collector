// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansCompute A compute rule to compute metrics or timeseries.
type SpansCompute struct {
	// An aggregation function.
	Aggregation SpansAggregationFunction `json:"aggregation"`
	// The time buckets' size (only used for type=timeseries)
	// Defaults to a resolution of 150 points.
	Interval *string `json:"interval,omitempty"`
	// The metric to use.
	Metric *string `json:"metric,omitempty"`
	// The type of compute.
	Type *SpansComputeType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansCompute instantiates a new SpansCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansCompute(aggregation SpansAggregationFunction) *SpansCompute {
	this := SpansCompute{}
	this.Aggregation = aggregation
	var typeVar SpansComputeType = SPANSCOMPUTETYPE_TOTAL
	this.Type = &typeVar
	return &this
}

// NewSpansComputeWithDefaults instantiates a new SpansCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansComputeWithDefaults() *SpansCompute {
	this := SpansCompute{}
	var typeVar SpansComputeType = SPANSCOMPUTETYPE_TOTAL
	this.Type = &typeVar
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *SpansCompute) GetAggregation() SpansAggregationFunction {
	if o == nil {
		var ret SpansAggregationFunction
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *SpansCompute) GetAggregationOk() (*SpansAggregationFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *SpansCompute) SetAggregation(v SpansAggregationFunction) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SpansCompute) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansCompute) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SpansCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *SpansCompute) SetInterval(v string) {
	o.Interval = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SpansCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SpansCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SpansCompute) SetMetric(v string) {
	o.Metric = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpansCompute) GetType() SpansComputeType {
	if o == nil || o.Type == nil {
		var ret SpansComputeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansCompute) GetTypeOk() (*SpansComputeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpansCompute) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SpansComputeType and assigns it to the Type field.
func (o *SpansCompute) SetType(v SpansComputeType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *SpansAggregationFunction `json:"aggregation"`
		Interval    *string                   `json:"interval,omitempty"`
		Metric      *string                   `json:"metric,omitempty"`
		Type        *SpansComputeType         `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "interval", "metric", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
	o.Interval = all.Interval
	o.Metric = all.Metric
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
