// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDowntimeMatchResponse Response for retrieving all downtime matches for a monitor.
type MonitorDowntimeMatchResponse struct {
	// An array of downtime matches.
	Data []MonitorDowntimeMatchResponseData `json:"data,omitempty"`
	// Pagination metadata returned by the API.
	Meta *DowntimeMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorDowntimeMatchResponse instantiates a new MonitorDowntimeMatchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDowntimeMatchResponse() *MonitorDowntimeMatchResponse {
	this := MonitorDowntimeMatchResponse{}
	return &this
}

// NewMonitorDowntimeMatchResponseWithDefaults instantiates a new MonitorDowntimeMatchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDowntimeMatchResponseWithDefaults() *MonitorDowntimeMatchResponse {
	this := MonitorDowntimeMatchResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponse) GetData() []MonitorDowntimeMatchResponseData {
	if o == nil || o.Data == nil {
		var ret []MonitorDowntimeMatchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponse) GetDataOk() (*[]MonitorDowntimeMatchResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []MonitorDowntimeMatchResponseData and assigns it to the Data field.
func (o *MonitorDowntimeMatchResponse) SetData(v []MonitorDowntimeMatchResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponse) GetMeta() DowntimeMeta {
	if o == nil || o.Meta == nil {
		var ret DowntimeMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponse) GetMetaOk() (*DowntimeMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DowntimeMeta and assigns it to the Meta field.
func (o *MonitorDowntimeMatchResponse) SetMeta(v DowntimeMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorDowntimeMatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDowntimeMatchResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []MonitorDowntimeMatchResponseData `json:"data,omitempty"`
		Meta *DowntimeMeta                      `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
