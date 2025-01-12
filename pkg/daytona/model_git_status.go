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

// checks if the GitStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitStatus{}

// GitStatus struct for GitStatus
type GitStatus struct {
	Ahead *int32 `json:"ahead,omitempty"`
	Behind *int32 `json:"behind,omitempty"`
	BranchPublished *bool `json:"branchPublished,omitempty"`
	CurrentBranch string `json:"currentBranch"`
	FileStatus []FileStatus `json:"fileStatus"`
}

type _GitStatus GitStatus

// NewGitStatus instantiates a new GitStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitStatus(currentBranch string, fileStatus []FileStatus) *GitStatus {
	this := GitStatus{}
	this.CurrentBranch = currentBranch
	this.FileStatus = fileStatus
	return &this
}

// NewGitStatusWithDefaults instantiates a new GitStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitStatusWithDefaults() *GitStatus {
	this := GitStatus{}
	return &this
}

// GetAhead returns the Ahead field value if set, zero value otherwise.
func (o *GitStatus) GetAhead() int32 {
	if o == nil || IsNil(o.Ahead) {
		var ret int32
		return ret
	}
	return *o.Ahead
}

// GetAheadOk returns a tuple with the Ahead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitStatus) GetAheadOk() (*int32, bool) {
	if o == nil || IsNil(o.Ahead) {
		return nil, false
	}
	return o.Ahead, true
}

// HasAhead returns a boolean if a field has been set.
func (o *GitStatus) HasAhead() bool {
	if o != nil && !IsNil(o.Ahead) {
		return true
	}

	return false
}

// SetAhead gets a reference to the given int32 and assigns it to the Ahead field.
func (o *GitStatus) SetAhead(v int32) {
	o.Ahead = &v
}

// GetBehind returns the Behind field value if set, zero value otherwise.
func (o *GitStatus) GetBehind() int32 {
	if o == nil || IsNil(o.Behind) {
		var ret int32
		return ret
	}
	return *o.Behind
}

// GetBehindOk returns a tuple with the Behind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitStatus) GetBehindOk() (*int32, bool) {
	if o == nil || IsNil(o.Behind) {
		return nil, false
	}
	return o.Behind, true
}

// HasBehind returns a boolean if a field has been set.
func (o *GitStatus) HasBehind() bool {
	if o != nil && !IsNil(o.Behind) {
		return true
	}

	return false
}

// SetBehind gets a reference to the given int32 and assigns it to the Behind field.
func (o *GitStatus) SetBehind(v int32) {
	o.Behind = &v
}

// GetBranchPublished returns the BranchPublished field value if set, zero value otherwise.
func (o *GitStatus) GetBranchPublished() bool {
	if o == nil || IsNil(o.BranchPublished) {
		var ret bool
		return ret
	}
	return *o.BranchPublished
}

// GetBranchPublishedOk returns a tuple with the BranchPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitStatus) GetBranchPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.BranchPublished) {
		return nil, false
	}
	return o.BranchPublished, true
}

// HasBranchPublished returns a boolean if a field has been set.
func (o *GitStatus) HasBranchPublished() bool {
	if o != nil && !IsNil(o.BranchPublished) {
		return true
	}

	return false
}

// SetBranchPublished gets a reference to the given bool and assigns it to the BranchPublished field.
func (o *GitStatus) SetBranchPublished(v bool) {
	o.BranchPublished = &v
}

// GetCurrentBranch returns the CurrentBranch field value
func (o *GitStatus) GetCurrentBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentBranch
}

// GetCurrentBranchOk returns a tuple with the CurrentBranch field value
// and a boolean to check if the value has been set.
func (o *GitStatus) GetCurrentBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBranch, true
}

// SetCurrentBranch sets field value
func (o *GitStatus) SetCurrentBranch(v string) {
	o.CurrentBranch = v
}

// GetFileStatus returns the FileStatus field value
func (o *GitStatus) GetFileStatus() []FileStatus {
	if o == nil {
		var ret []FileStatus
		return ret
	}

	return o.FileStatus
}

// GetFileStatusOk returns a tuple with the FileStatus field value
// and a boolean to check if the value has been set.
func (o *GitStatus) GetFileStatusOk() ([]FileStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileStatus, true
}

// SetFileStatus sets field value
func (o *GitStatus) SetFileStatus(v []FileStatus) {
	o.FileStatus = v
}

func (o GitStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ahead) {
		toSerialize["ahead"] = o.Ahead
	}
	if !IsNil(o.Behind) {
		toSerialize["behind"] = o.Behind
	}
	if !IsNil(o.BranchPublished) {
		toSerialize["branchPublished"] = o.BranchPublished
	}
	toSerialize["currentBranch"] = o.CurrentBranch
	toSerialize["fileStatus"] = o.FileStatus
	return toSerialize, nil
}

func (o *GitStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentBranch",
		"fileStatus",
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

	varGitStatus := _GitStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitStatus)

	if err != nil {
		return err
	}

	*o = GitStatus(varGitStatus)

	return err
}

type NullableGitStatus struct {
	value *GitStatus
	isSet bool
}

func (v NullableGitStatus) Get() *GitStatus {
	return v.value
}

func (v *NullableGitStatus) Set(val *GitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitStatus(val *GitStatus) *NullableGitStatus {
	return &NullableGitStatus{value: val, isSet: true}
}

func (v NullableGitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


