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

// checks if the ShipRefine201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRefine201ResponseData{}

// ShipRefine201ResponseData struct for ShipRefine201ResponseData
type ShipRefine201ResponseData struct {
	Cargo ShipCargo `json:"cargo"`
	Cooldown Cooldown `json:"cooldown"`
	// Goods that were produced by this refining process.
	Produced []ShipRefine201ResponseDataProducedInner `json:"produced"`
	// Goods that were consumed during this refining process.
	Consumed []ShipRefine201ResponseDataProducedInner `json:"consumed"`
}

type _ShipRefine201ResponseData ShipRefine201ResponseData

// NewShipRefine201ResponseData instantiates a new ShipRefine201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRefine201ResponseData(cargo ShipCargo, cooldown Cooldown, produced []ShipRefine201ResponseDataProducedInner, consumed []ShipRefine201ResponseDataProducedInner) *ShipRefine201ResponseData {
	this := ShipRefine201ResponseData{}
	this.Cargo = cargo
	this.Cooldown = cooldown
	this.Produced = produced
	this.Consumed = consumed
	return &this
}

// NewShipRefine201ResponseDataWithDefaults instantiates a new ShipRefine201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRefine201ResponseDataWithDefaults() *ShipRefine201ResponseData {
	this := ShipRefine201ResponseData{}
	return &this
}

// GetCargo returns the Cargo field value
func (o *ShipRefine201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *ShipRefine201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetCooldown returns the Cooldown field value
func (o *ShipRefine201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *ShipRefine201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetProduced returns the Produced field value
func (o *ShipRefine201ResponseData) GetProduced() []ShipRefine201ResponseDataProducedInner {
	if o == nil {
		var ret []ShipRefine201ResponseDataProducedInner
		return ret
	}

	return o.Produced
}

// GetProducedOk returns a tuple with the Produced field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseData) GetProducedOk() ([]ShipRefine201ResponseDataProducedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Produced, true
}

// SetProduced sets field value
func (o *ShipRefine201ResponseData) SetProduced(v []ShipRefine201ResponseDataProducedInner) {
	o.Produced = v
}

// GetConsumed returns the Consumed field value
func (o *ShipRefine201ResponseData) GetConsumed() []ShipRefine201ResponseDataProducedInner {
	if o == nil {
		var ret []ShipRefine201ResponseDataProducedInner
		return ret
	}

	return o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value
// and a boolean to check if the value has been set.
func (o *ShipRefine201ResponseData) GetConsumedOk() ([]ShipRefine201ResponseDataProducedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Consumed, true
}

// SetConsumed sets field value
func (o *ShipRefine201ResponseData) SetConsumed(v []ShipRefine201ResponseDataProducedInner) {
	o.Consumed = v
}

func (o ShipRefine201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRefine201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cargo"] = o.Cargo
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["produced"] = o.Produced
	toSerialize["consumed"] = o.Consumed
	return toSerialize, nil
}

func (o *ShipRefine201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cargo",
		"cooldown",
		"produced",
		"consumed",
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

	varShipRefine201ResponseData := _ShipRefine201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipRefine201ResponseData)

	if err != nil {
		return err
	}

	*o = ShipRefine201ResponseData(varShipRefine201ResponseData)

	return err
}

type NullableShipRefine201ResponseData struct {
	value *ShipRefine201ResponseData
	isSet bool
}

func (v NullableShipRefine201ResponseData) Get() *ShipRefine201ResponseData {
	return v.value
}

func (v *NullableShipRefine201ResponseData) Set(val *ShipRefine201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRefine201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRefine201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRefine201ResponseData(val *ShipRefine201ResponseData) *NullableShipRefine201ResponseData {
	return &NullableShipRefine201ResponseData{value: val, isSet: true}
}

func (v NullableShipRefine201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRefine201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


