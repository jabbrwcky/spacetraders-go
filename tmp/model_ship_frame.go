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

// checks if the ShipFrame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipFrame{}

// ShipFrame The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.
type ShipFrame struct {
	// Symbol of the frame.
	Symbol string `json:"symbol"`
	// Name of the frame.
	Name string `json:"name"`
	// Description of the frame.
	Description string `json:"description"`
	// Condition is a range of 0 to 100 where 0 is completely worn out and 100 is brand new.
	Condition *int32 `json:"condition,omitempty"`
	// The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed.
	ModuleSlots int32 `json:"moduleSlots"`
	// The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed.
	MountingPoints int32 `json:"mountingPoints"`
	// The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount.
	FuelCapacity int32 `json:"fuelCapacity"`
	Requirements ShipRequirements `json:"requirements"`
}

type _ShipFrame ShipFrame

// NewShipFrame instantiates a new ShipFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipFrame(symbol string, name string, description string, moduleSlots int32, mountingPoints int32, fuelCapacity int32, requirements ShipRequirements) *ShipFrame {
	this := ShipFrame{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	this.ModuleSlots = moduleSlots
	this.MountingPoints = mountingPoints
	this.FuelCapacity = fuelCapacity
	this.Requirements = requirements
	return &this
}

// NewShipFrameWithDefaults instantiates a new ShipFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipFrameWithDefaults() *ShipFrame {
	this := ShipFrame{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipFrame) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipFrame) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ShipFrame) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipFrame) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipFrame) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipFrame) SetDescription(v string) {
	o.Description = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ShipFrame) GetCondition() int32 {
	if o == nil || IsNil(o.Condition) {
		var ret int32
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetConditionOk() (*int32, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ShipFrame) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given int32 and assigns it to the Condition field.
func (o *ShipFrame) SetCondition(v int32) {
	o.Condition = &v
}

// GetModuleSlots returns the ModuleSlots field value
func (o *ShipFrame) GetModuleSlots() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModuleSlots
}

// GetModuleSlotsOk returns a tuple with the ModuleSlots field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetModuleSlotsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleSlots, true
}

// SetModuleSlots sets field value
func (o *ShipFrame) SetModuleSlots(v int32) {
	o.ModuleSlots = v
}

// GetMountingPoints returns the MountingPoints field value
func (o *ShipFrame) GetMountingPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MountingPoints
}

// GetMountingPointsOk returns a tuple with the MountingPoints field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetMountingPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountingPoints, true
}

// SetMountingPoints sets field value
func (o *ShipFrame) SetMountingPoints(v int32) {
	o.MountingPoints = v
}

// GetFuelCapacity returns the FuelCapacity field value
func (o *ShipFrame) GetFuelCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FuelCapacity
}

// GetFuelCapacityOk returns a tuple with the FuelCapacity field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetFuelCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FuelCapacity, true
}

// SetFuelCapacity sets field value
func (o *ShipFrame) SetFuelCapacity(v int32) {
	o.FuelCapacity = v
}

// GetRequirements returns the Requirements field value
func (o *ShipFrame) GetRequirements() ShipRequirements {
	if o == nil {
		var ret ShipRequirements
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ShipFrame) GetRequirementsOk() (*ShipRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ShipFrame) SetRequirements(v ShipRequirements) {
	o.Requirements = v
}

func (o ShipFrame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipFrame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["moduleSlots"] = o.ModuleSlots
	toSerialize["mountingPoints"] = o.MountingPoints
	toSerialize["fuelCapacity"] = o.FuelCapacity
	toSerialize["requirements"] = o.Requirements
	return toSerialize, nil
}

func (o *ShipFrame) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"description",
		"moduleSlots",
		"mountingPoints",
		"fuelCapacity",
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

	varShipFrame := _ShipFrame{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipFrame)

	if err != nil {
		return err
	}

	*o = ShipFrame(varShipFrame)

	return err
}

type NullableShipFrame struct {
	value *ShipFrame
	isSet bool
}

func (v NullableShipFrame) Get() *ShipFrame {
	return v.value
}

func (v *NullableShipFrame) Set(val *ShipFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableShipFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableShipFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipFrame(val *ShipFrame) *NullableShipFrame {
	return &NullableShipFrame{value: val, isSet: true}
}

func (v NullableShipFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


