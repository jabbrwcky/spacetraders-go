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

// checks if the CreateShipSystemScan201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipSystemScan201ResponseData{}

// CreateShipSystemScan201ResponseData struct for CreateShipSystemScan201ResponseData
type CreateShipSystemScan201ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	// List of scanned systems.
	Systems []ScannedSystem `json:"systems"`
}

type _CreateShipSystemScan201ResponseData CreateShipSystemScan201ResponseData

// NewCreateShipSystemScan201ResponseData instantiates a new CreateShipSystemScan201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipSystemScan201ResponseData(cooldown Cooldown, systems []ScannedSystem) *CreateShipSystemScan201ResponseData {
	this := CreateShipSystemScan201ResponseData{}
	this.Cooldown = cooldown
	this.Systems = systems
	return &this
}

// NewCreateShipSystemScan201ResponseDataWithDefaults instantiates a new CreateShipSystemScan201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipSystemScan201ResponseDataWithDefaults() *CreateShipSystemScan201ResponseData {
	this := CreateShipSystemScan201ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *CreateShipSystemScan201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *CreateShipSystemScan201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *CreateShipSystemScan201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetSystems returns the Systems field value
func (o *CreateShipSystemScan201ResponseData) GetSystems() []ScannedSystem {
	if o == nil {
		var ret []ScannedSystem
		return ret
	}

	return o.Systems
}

// GetSystemsOk returns a tuple with the Systems field value
// and a boolean to check if the value has been set.
func (o *CreateShipSystemScan201ResponseData) GetSystemsOk() ([]ScannedSystem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Systems, true
}

// SetSystems sets field value
func (o *CreateShipSystemScan201ResponseData) SetSystems(v []ScannedSystem) {
	o.Systems = v
}

func (o CreateShipSystemScan201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShipSystemScan201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["systems"] = o.Systems
	return toSerialize, nil
}

func (o *CreateShipSystemScan201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"systems",
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

	varCreateShipSystemScan201ResponseData := _CreateShipSystemScan201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateShipSystemScan201ResponseData)

	if err != nil {
		return err
	}

	*o = CreateShipSystemScan201ResponseData(varCreateShipSystemScan201ResponseData)

	return err
}

type NullableCreateShipSystemScan201ResponseData struct {
	value *CreateShipSystemScan201ResponseData
	isSet bool
}

func (v NullableCreateShipSystemScan201ResponseData) Get() *CreateShipSystemScan201ResponseData {
	return v.value
}

func (v *NullableCreateShipSystemScan201ResponseData) Set(val *CreateShipSystemScan201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipSystemScan201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipSystemScan201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipSystemScan201ResponseData(val *CreateShipSystemScan201ResponseData) *NullableCreateShipSystemScan201ResponseData {
	return &NullableCreateShipSystemScan201ResponseData{value: val, isSet: true}
}

func (v NullableCreateShipSystemScan201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipSystemScan201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


