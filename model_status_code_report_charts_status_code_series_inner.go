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

// checks if the StatusCodeReportChartsStatusCodeSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportChartsStatusCodeSeriesInner{}

// StatusCodeReportChartsStatusCodeSeriesInner struct for StatusCodeReportChartsStatusCodeSeriesInner
type StatusCodeReportChartsStatusCodeSeriesInner struct {
	Name *string `json:"name,omitempty"`
	Data []float64 `json:"data,omitempty"`
}

// NewStatusCodeReportChartsStatusCodeSeriesInner instantiates a new StatusCodeReportChartsStatusCodeSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportChartsStatusCodeSeriesInner() *StatusCodeReportChartsStatusCodeSeriesInner {
	this := StatusCodeReportChartsStatusCodeSeriesInner{}
	return &this
}

// NewStatusCodeReportChartsStatusCodeSeriesInnerWithDefaults instantiates a new StatusCodeReportChartsStatusCodeSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportChartsStatusCodeSeriesInnerWithDefaults() *StatusCodeReportChartsStatusCodeSeriesInner {
	this := StatusCodeReportChartsStatusCodeSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) GetData() []float64 {
	if o == nil || IsNil(o.Data) {
		var ret []float64
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) GetDataOk() ([]float64, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []float64 and assigns it to the Data field.
func (o *StatusCodeReportChartsStatusCodeSeriesInner) SetData(v []float64) {
	o.Data = v
}

func (o StatusCodeReportChartsStatusCodeSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportChartsStatusCodeSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStatusCodeReportChartsStatusCodeSeriesInner struct {
	value *StatusCodeReportChartsStatusCodeSeriesInner
	isSet bool
}

func (v NullableStatusCodeReportChartsStatusCodeSeriesInner) Get() *StatusCodeReportChartsStatusCodeSeriesInner {
	return v.value
}

func (v *NullableStatusCodeReportChartsStatusCodeSeriesInner) Set(val *StatusCodeReportChartsStatusCodeSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportChartsStatusCodeSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportChartsStatusCodeSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportChartsStatusCodeSeriesInner(val *StatusCodeReportChartsStatusCodeSeriesInner) *NullableStatusCodeReportChartsStatusCodeSeriesInner {
	return &NullableStatusCodeReportChartsStatusCodeSeriesInner{value: val, isSet: true}
}

func (v NullableStatusCodeReportChartsStatusCodeSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportChartsStatusCodeSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

