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

// checks if the ShipNavRouteWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipNavRouteWaypoint{}

// ShipNavRouteWaypoint The destination or departure of a ships nav route.
type ShipNavRouteWaypoint struct {
	// The symbol of the waypoint.
	Symbol string `json:"symbol"`
	Type WaypointType `json:"type"`
	// The symbol of the system.
	SystemSymbol string `json:"systemSymbol"`
	// Position in the universe in the x axis.
	X int32 `json:"x"`
	// Position in the universe in the y axis.
	Y int32 `json:"y"`
}

type _ShipNavRouteWaypoint ShipNavRouteWaypoint

// NewShipNavRouteWaypoint instantiates a new ShipNavRouteWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipNavRouteWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32) *ShipNavRouteWaypoint {
	this := ShipNavRouteWaypoint{}
	this.Symbol = symbol
	this.Type = type_
	this.SystemSymbol = systemSymbol
	this.X = x
	this.Y = y
	return &this
}

// NewShipNavRouteWaypointWithDefaults instantiates a new ShipNavRouteWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipNavRouteWaypointWithDefaults() *ShipNavRouteWaypoint {
	this := ShipNavRouteWaypoint{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipNavRouteWaypoint) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipNavRouteWaypoint) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipNavRouteWaypoint) SetSymbol(v string) {
	o.Symbol = v
}

// GetType returns the Type field value
func (o *ShipNavRouteWaypoint) GetType() WaypointType {
	if o == nil {
		var ret WaypointType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipNavRouteWaypoint) GetTypeOk() (*WaypointType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipNavRouteWaypoint) SetType(v WaypointType) {
	o.Type = v
}

// GetSystemSymbol returns the SystemSymbol field value
func (o *ShipNavRouteWaypoint) GetSystemSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemSymbol
}

// GetSystemSymbolOk returns a tuple with the SystemSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipNavRouteWaypoint) GetSystemSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemSymbol, true
}

// SetSystemSymbol sets field value
func (o *ShipNavRouteWaypoint) SetSystemSymbol(v string) {
	o.SystemSymbol = v
}

// GetX returns the X field value
func (o *ShipNavRouteWaypoint) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ShipNavRouteWaypoint) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ShipNavRouteWaypoint) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *ShipNavRouteWaypoint) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ShipNavRouteWaypoint) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ShipNavRouteWaypoint) SetY(v int32) {
	o.Y = v
}

func (o ShipNavRouteWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipNavRouteWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["type"] = o.Type
	toSerialize["systemSymbol"] = o.SystemSymbol
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *ShipNavRouteWaypoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"type",
		"systemSymbol",
		"x",
		"y",
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

	varShipNavRouteWaypoint := _ShipNavRouteWaypoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipNavRouteWaypoint)

	if err != nil {
		return err
	}

	*o = ShipNavRouteWaypoint(varShipNavRouteWaypoint)

	return err
}

type NullableShipNavRouteWaypoint struct {
	value *ShipNavRouteWaypoint
	isSet bool
}

func (v NullableShipNavRouteWaypoint) Get() *ShipNavRouteWaypoint {
	return v.value
}

func (v *NullableShipNavRouteWaypoint) Set(val *ShipNavRouteWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableShipNavRouteWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableShipNavRouteWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipNavRouteWaypoint(val *ShipNavRouteWaypoint) *NullableShipNavRouteWaypoint {
	return &NullableShipNavRouteWaypoint{value: val, isSet: true}
}

func (v NullableShipNavRouteWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipNavRouteWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


