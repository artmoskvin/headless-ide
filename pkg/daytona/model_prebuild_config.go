/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daytona

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PrebuildConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrebuildConfig{}

// PrebuildConfig struct for PrebuildConfig
type PrebuildConfig struct {
	Branch string `json:"branch"`
	CommitInterval int32 `json:"commitInterval"`
	Id string `json:"id"`
	Retention int32 `json:"retention"`
	TriggerFiles []string `json:"triggerFiles"`
}

type _PrebuildConfig PrebuildConfig

// NewPrebuildConfig instantiates a new PrebuildConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrebuildConfig(branch string, commitInterval int32, id string, retention int32, triggerFiles []string) *PrebuildConfig {
	this := PrebuildConfig{}
	this.Branch = branch
	this.CommitInterval = commitInterval
	this.Id = id
	this.Retention = retention
	this.TriggerFiles = triggerFiles
	return &this
}

// NewPrebuildConfigWithDefaults instantiates a new PrebuildConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrebuildConfigWithDefaults() *PrebuildConfig {
	this := PrebuildConfig{}
	return &this
}

// GetBranch returns the Branch field value
func (o *PrebuildConfig) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *PrebuildConfig) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *PrebuildConfig) SetBranch(v string) {
	o.Branch = v
}

// GetCommitInterval returns the CommitInterval field value
func (o *PrebuildConfig) GetCommitInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommitInterval
}

// GetCommitIntervalOk returns a tuple with the CommitInterval field value
// and a boolean to check if the value has been set.
func (o *PrebuildConfig) GetCommitIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitInterval, true
}

// SetCommitInterval sets field value
func (o *PrebuildConfig) SetCommitInterval(v int32) {
	o.CommitInterval = v
}

// GetId returns the Id field value
func (o *PrebuildConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrebuildConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrebuildConfig) SetId(v string) {
	o.Id = v
}

// GetRetention returns the Retention field value
func (o *PrebuildConfig) GetRetention() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *PrebuildConfig) GetRetentionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *PrebuildConfig) SetRetention(v int32) {
	o.Retention = v
}

// GetTriggerFiles returns the TriggerFiles field value
func (o *PrebuildConfig) GetTriggerFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TriggerFiles
}

// GetTriggerFilesOk returns a tuple with the TriggerFiles field value
// and a boolean to check if the value has been set.
func (o *PrebuildConfig) GetTriggerFilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggerFiles, true
}

// SetTriggerFiles sets field value
func (o *PrebuildConfig) SetTriggerFiles(v []string) {
	o.TriggerFiles = v
}

func (o PrebuildConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrebuildConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["branch"] = o.Branch
	toSerialize["commitInterval"] = o.CommitInterval
	toSerialize["id"] = o.Id
	toSerialize["retention"] = o.Retention
	toSerialize["triggerFiles"] = o.TriggerFiles
	return toSerialize, nil
}

func (o *PrebuildConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"branch",
		"commitInterval",
		"id",
		"retention",
		"triggerFiles",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPrebuildConfig := _PrebuildConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrebuildConfig)

	if err != nil {
		return err
	}

	*o = PrebuildConfig(varPrebuildConfig)

	return err
}

type NullablePrebuildConfig struct {
	value *PrebuildConfig
	isSet bool
}

func (v NullablePrebuildConfig) Get() *PrebuildConfig {
	return v.value
}

func (v *NullablePrebuildConfig) Set(val *PrebuildConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePrebuildConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePrebuildConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrebuildConfig(val *PrebuildConfig) *NullablePrebuildConfig {
	return &NullablePrebuildConfig{value: val, isSet: true}
}

func (v NullablePrebuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrebuildConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

