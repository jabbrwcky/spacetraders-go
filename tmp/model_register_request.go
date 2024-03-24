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

// checks if the RegisterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterRequest{}

// RegisterRequest struct for RegisterRequest
type RegisterRequest struct {
	Faction FactionSymbol `json:"faction"`
	// Your desired agent symbol. This will be a unique name used to represent your agent, and will be the prefix for your ships.
	Symbol string `json:"symbol"`
	// Your email address. This is used if you reserved your call sign between resets.
	Email *string `json:"email,omitempty"`
}

type _RegisterRequest RegisterRequest

// NewRegisterRequest instantiates a new RegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterRequest(faction FactionSymbol, symbol string) *RegisterRequest {
	this := RegisterRequest{}
	this.Faction = faction
	this.Symbol = symbol
	return &this
}

// NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterRequestWithDefaults() *RegisterRequest {
	this := RegisterRequest{}
	return &this
}

// GetFaction returns the Faction field value
func (o *RegisterRequest) GetFaction() FactionSymbol {
	if o == nil {
		var ret FactionSymbol
		return ret
	}

	return o.Faction
}

// GetFactionOk returns a tuple with the Faction field value
// and a boolean to check if the value has been set.
func (o *RegisterRequest) GetFactionOk() (*FactionSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Faction, true
}

// SetFaction sets field value
func (o *RegisterRequest) SetFaction(v FactionSymbol) {
	o.Faction = v
}

// GetSymbol returns the Symbol field value
func (o *RegisterRequest) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *RegisterRequest) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *RegisterRequest) SetSymbol(v string) {
	o.Symbol = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RegisterRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RegisterRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RegisterRequest) SetEmail(v string) {
	o.Email = &v
}

func (o RegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["faction"] = o.Faction
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

func (o *RegisterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"faction",
		"symbol",
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

	varRegisterRequest := _RegisterRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterRequest)

	if err != nil {
		return err
	}

	*o = RegisterRequest(varRegisterRequest)

	return err
}

type NullableRegisterRequest struct {
	value *RegisterRequest
	isSet bool
}

func (v NullableRegisterRequest) Get() *RegisterRequest {
	return v.value
}

func (v *NullableRegisterRequest) Set(val *RegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterRequest(val *RegisterRequest) *NullableRegisterRequest {
	return &NullableRegisterRequest{value: val, isSet: true}
}

func (v NullableRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

