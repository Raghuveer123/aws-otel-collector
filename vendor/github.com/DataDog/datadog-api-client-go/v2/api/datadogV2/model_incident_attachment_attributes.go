// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// IncidentAttachmentAttributes - The attributes object for an attachment.
type IncidentAttachmentAttributes struct {
	IncidentAttachmentPostmortemAttributes *IncidentAttachmentPostmortemAttributes
	IncidentAttachmentLinkAttributes       *IncidentAttachmentLinkAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentAttachmentPostmortemAttributesAsIncidentAttachmentAttributes is a convenience function that returns IncidentAttachmentPostmortemAttributes wrapped in IncidentAttachmentAttributes.
func IncidentAttachmentPostmortemAttributesAsIncidentAttachmentAttributes(v *IncidentAttachmentPostmortemAttributes) IncidentAttachmentAttributes {
	return IncidentAttachmentAttributes{IncidentAttachmentPostmortemAttributes: v}
}

// IncidentAttachmentLinkAttributesAsIncidentAttachmentAttributes is a convenience function that returns IncidentAttachmentLinkAttributes wrapped in IncidentAttachmentAttributes.
func IncidentAttachmentLinkAttributesAsIncidentAttachmentAttributes(v *IncidentAttachmentLinkAttributes) IncidentAttachmentAttributes {
	return IncidentAttachmentAttributes{IncidentAttachmentLinkAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentAttachmentAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentAttachmentPostmortemAttributes
	err = json.Unmarshal(data, &obj.IncidentAttachmentPostmortemAttributes)
	if err == nil {
		if obj.IncidentAttachmentPostmortemAttributes != nil && obj.IncidentAttachmentPostmortemAttributes.UnparsedObject == nil {
			jsonIncidentAttachmentPostmortemAttributes, _ := json.Marshal(obj.IncidentAttachmentPostmortemAttributes)
			if string(jsonIncidentAttachmentPostmortemAttributes) == "{}" { // empty struct
				obj.IncidentAttachmentPostmortemAttributes = nil
			} else {
				match++
			}
		} else {
			obj.IncidentAttachmentPostmortemAttributes = nil
		}
	} else {
		obj.IncidentAttachmentPostmortemAttributes = nil
	}

	// try to unmarshal data into IncidentAttachmentLinkAttributes
	err = json.Unmarshal(data, &obj.IncidentAttachmentLinkAttributes)
	if err == nil {
		if obj.IncidentAttachmentLinkAttributes != nil && obj.IncidentAttachmentLinkAttributes.UnparsedObject == nil {
			jsonIncidentAttachmentLinkAttributes, _ := json.Marshal(obj.IncidentAttachmentLinkAttributes)
			if string(jsonIncidentAttachmentLinkAttributes) == "{}" { // empty struct
				obj.IncidentAttachmentLinkAttributes = nil
			} else {
				match++
			}
		} else {
			obj.IncidentAttachmentLinkAttributes = nil
		}
	} else {
		obj.IncidentAttachmentLinkAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentAttachmentPostmortemAttributes = nil
		obj.IncidentAttachmentLinkAttributes = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentAttachmentAttributes) MarshalJSON() ([]byte, error) {
	if obj.IncidentAttachmentPostmortemAttributes != nil {
		return json.Marshal(&obj.IncidentAttachmentPostmortemAttributes)
	}

	if obj.IncidentAttachmentLinkAttributes != nil {
		return json.Marshal(&obj.IncidentAttachmentLinkAttributes)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentAttachmentAttributes) GetActualInstance() interface{} {
	if obj.IncidentAttachmentPostmortemAttributes != nil {
		return obj.IncidentAttachmentPostmortemAttributes
	}

	if obj.IncidentAttachmentLinkAttributes != nil {
		return obj.IncidentAttachmentLinkAttributes
	}

	// all schemas are nil
	return nil
}
