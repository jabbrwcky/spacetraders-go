/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InstallMount201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallMount201Response{}

// InstallMount201Response struct for InstallMount201Response
type InstallMount201Response struct {
	Data InstallMount201ResponseData `json:"data"`
}

type _InstallMount201Response InstallMount201Response

// NewInstallMount201Response instantiates a new InstallMount201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallMount201Response(data InstallMount201ResponseData) *InstallMount201Response {
	this := InstallMount201Response{}
	this.Data = data
	return &this
}

// NewInstallMount201ResponseWithDefaults instantiates a new InstallMount201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallMount201ResponseWithDefaults() *InstallMount201Response {
	this := InstallMount201Response{}
	return &this
}

// GetData returns the Data field value
func (o *InstallMount201Response) GetData() InstallMount201ResponseData {
	if o == nil {
		var ret InstallMount201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InstallMount201Response) GetDataOk() (*InstallMount201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InstallMount201Response) SetData(v InstallMount201ResponseData) {
	o.Data = v
}

func (o InstallMount201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallMount201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *InstallMount201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varInstallMount201Response := _InstallMount201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstallMount201Response)

	if err != nil {
		return err
	}

	*o = InstallMount201Response(varInstallMount201Response)

	return err
}

type NullableInstallMount201Response struct {
	value *InstallMount201Response
	isSet bool
}

func (v NullableInstallMount201Response) Get() *InstallMount201Response {
	return v.value
}

func (v *NullableInstallMount201Response) Set(val *InstallMount201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallMount201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallMount201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallMount201Response(val *InstallMount201Response) *NullableInstallMount201Response {
	return &NullableInstallMount201Response{value: val, isSet: true}
}

func (v NullableInstallMount201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallMount201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


