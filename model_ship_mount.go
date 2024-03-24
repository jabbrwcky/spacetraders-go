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

// checks if the ShipMount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipMount{}

// ShipMount A mount is installed on the exterier of a ship.
type ShipMount struct {
	// Symbo of this mount.
	Symbol string `json:"symbol"`
	// Name of this mount.
	Name string `json:"name"`
	// Description of this mount.
	Description *string `json:"description,omitempty"`
	// Mounts that have this value, such as mining lasers, denote how powerful this mount's capabilities are.
	Strength *int32 `json:"strength,omitempty"`
	// Mounts that have this value denote what goods can be produced from using the mount.
	Deposits []string `json:"deposits,omitempty"`
	Requirements ShipRequirements `json:"requirements"`
}

type _ShipMount ShipMount

// NewShipMount instantiates a new ShipMount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipMount(symbol string, name string, requirements ShipRequirements) *ShipMount {
	this := ShipMount{}
	this.Symbol = symbol
	this.Name = name
	this.Requirements = requirements
	return &this
}

// NewShipMountWithDefaults instantiates a new ShipMount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipMountWithDefaults() *ShipMount {
	this := ShipMount{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipMount) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipMount) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipMount) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ShipMount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipMount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipMount) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ShipMount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ShipMount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ShipMount) SetDescription(v string) {
	o.Description = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *ShipMount) GetStrength() int32 {
	if o == nil || IsNil(o.Strength) {
		var ret int32
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMount) GetStrengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *ShipMount) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given int32 and assigns it to the Strength field.
func (o *ShipMount) SetStrength(v int32) {
	o.Strength = &v
}

// GetDeposits returns the Deposits field value if set, zero value otherwise.
func (o *ShipMount) GetDeposits() []string {
	if o == nil || IsNil(o.Deposits) {
		var ret []string
		return ret
	}
	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMount) GetDepositsOk() ([]string, bool) {
	if o == nil || IsNil(o.Deposits) {
		return nil, false
	}
	return o.Deposits, true
}

// HasDeposits returns a boolean if a field has been set.
func (o *ShipMount) HasDeposits() bool {
	if o != nil && !IsNil(o.Deposits) {
		return true
	}

	return false
}

// SetDeposits gets a reference to the given []string and assigns it to the Deposits field.
func (o *ShipMount) SetDeposits(v []string) {
	o.Deposits = v
}

// GetRequirements returns the Requirements field value
func (o *ShipMount) GetRequirements() ShipRequirements {
	if o == nil {
		var ret ShipRequirements
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ShipMount) GetRequirementsOk() (*ShipRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ShipMount) SetRequirements(v ShipRequirements) {
	o.Requirements = v
}

func (o ShipMount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipMount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !IsNil(o.Deposits) {
		toSerialize["deposits"] = o.Deposits
	}
	toSerialize["requirements"] = o.Requirements
	return toSerialize, nil
}

func (o *ShipMount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"requirements",
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

	varShipMount := _ShipMount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipMount)

	if err != nil {
		return err
	}

	*o = ShipMount(varShipMount)

	return err
}

type NullableShipMount struct {
	value *ShipMount
	isSet bool
}

func (v NullableShipMount) Get() *ShipMount {
	return v.value
}

func (v *NullableShipMount) Set(val *ShipMount) {
	v.value = val
	v.isSet = true
}

func (v NullableShipMount) IsSet() bool {
	return v.isSet
}

func (v *NullableShipMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipMount(val *ShipMount) *NullableShipMount {
	return &NullableShipMount{value: val, isSet: true}
}

func (v NullableShipMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


