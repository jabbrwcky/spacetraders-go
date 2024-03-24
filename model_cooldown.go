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
	"time"
	"bytes"
	"fmt"
)

// checks if the Cooldown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cooldown{}

// Cooldown A cooldown is a period of time in which a ship cannot perform certain actions.
type Cooldown struct {
	// The symbol of the ship that is on cooldown
	ShipSymbol string `json:"shipSymbol"`
	// The total duration of the cooldown in seconds
	TotalSeconds int32 `json:"totalSeconds"`
	// The remaining duration of the cooldown in seconds
	RemainingSeconds int32 `json:"remainingSeconds"`
	// The date and time when the cooldown expires in ISO 8601 format
	Expiration *time.Time `json:"expiration,omitempty"`
}

type _Cooldown Cooldown

// NewCooldown instantiates a new Cooldown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCooldown(shipSymbol string, totalSeconds int32, remainingSeconds int32) *Cooldown {
	this := Cooldown{}
	this.ShipSymbol = shipSymbol
	this.TotalSeconds = totalSeconds
	this.RemainingSeconds = remainingSeconds
	return &this
}

// NewCooldownWithDefaults instantiates a new Cooldown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCooldownWithDefaults() *Cooldown {
	this := Cooldown{}
	return &this
}

// GetShipSymbol returns the ShipSymbol field value
func (o *Cooldown) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *Cooldown) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *Cooldown) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetTotalSeconds returns the TotalSeconds field value
func (o *Cooldown) GetTotalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSeconds
}

// GetTotalSecondsOk returns a tuple with the TotalSeconds field value
// and a boolean to check if the value has been set.
func (o *Cooldown) GetTotalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSeconds, true
}

// SetTotalSeconds sets field value
func (o *Cooldown) SetTotalSeconds(v int32) {
	o.TotalSeconds = v
}

// GetRemainingSeconds returns the RemainingSeconds field value
func (o *Cooldown) GetRemainingSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingSeconds
}

// GetRemainingSecondsOk returns a tuple with the RemainingSeconds field value
// and a boolean to check if the value has been set.
func (o *Cooldown) GetRemainingSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingSeconds, true
}

// SetRemainingSeconds sets field value
func (o *Cooldown) SetRemainingSeconds(v int32) {
	o.RemainingSeconds = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Cooldown) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cooldown) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Cooldown) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *Cooldown) SetExpiration(v time.Time) {
	o.Expiration = &v
}

func (o Cooldown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cooldown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["totalSeconds"] = o.TotalSeconds
	toSerialize["remainingSeconds"] = o.RemainingSeconds
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	return toSerialize, nil
}

func (o *Cooldown) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipSymbol",
		"totalSeconds",
		"remainingSeconds",
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

	varCooldown := _Cooldown{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCooldown)

	if err != nil {
		return err
	}

	*o = Cooldown(varCooldown)

	return err
}

type NullableCooldown struct {
	value *Cooldown
	isSet bool
}

func (v NullableCooldown) Get() *Cooldown {
	return v.value
}

func (v *NullableCooldown) Set(val *Cooldown) {
	v.value = val
	v.isSet = true
}

func (v NullableCooldown) IsSet() bool {
	return v.isSet
}

func (v *NullableCooldown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCooldown(val *Cooldown) *NullableCooldown {
	return &NullableCooldown{value: val, isSet: true}
}

func (v NullableCooldown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCooldown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


