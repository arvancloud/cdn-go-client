/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the StatusCodeReportCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportCharts{}

// StatusCodeReportCharts struct for StatusCodeReportCharts
type StatusCodeReportCharts struct {
	StatusCode *StatusCodeReportChartsStatusCode `json:"status_code,omitempty"`
}

// NewStatusCodeReportCharts instantiates a new StatusCodeReportCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportCharts() *StatusCodeReportCharts {
	this := StatusCodeReportCharts{}
	return &this
}

// NewStatusCodeReportChartsWithDefaults instantiates a new StatusCodeReportCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportChartsWithDefaults() *StatusCodeReportCharts {
	this := StatusCodeReportCharts{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *StatusCodeReportCharts) GetStatusCode() StatusCodeReportChartsStatusCode {
	if o == nil || IsNil(o.StatusCode) {
		var ret StatusCodeReportChartsStatusCode
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportCharts) GetStatusCodeOk() (*StatusCodeReportChartsStatusCode, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *StatusCodeReportCharts) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given StatusCodeReportChartsStatusCode and assigns it to the StatusCode field.
func (o *StatusCodeReportCharts) SetStatusCode(v StatusCodeReportChartsStatusCode) {
	o.StatusCode = &v
}

func (o StatusCodeReportCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableStatusCodeReportCharts struct {
	value *StatusCodeReportCharts
	isSet bool
}

func (v NullableStatusCodeReportCharts) Get() *StatusCodeReportCharts {
	return v.value
}

func (v *NullableStatusCodeReportCharts) Set(val *StatusCodeReportCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportCharts(val *StatusCodeReportCharts) *NullableStatusCodeReportCharts {
	return &NullableStatusCodeReportCharts{value: val, isSet: true}
}

func (v NullableStatusCodeReportCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

