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
	"fmt"
)

// ActivityLevel The activity level of a trade good. If the good is an import, this represents how strong consumption is. If the good is an export, this represents how strong the production is for the good. When activity is strong, consumption or production is near maximum capacity. When activity is weak, consumption or production is near minimum capacity.
type ActivityLevel string

// List of ActivityLevel
const (
	ACTIVITYLEVEL_WEAK ActivityLevel = "WEAK"
	ACTIVITYLEVEL_GROWING ActivityLevel = "GROWING"
	ACTIVITYLEVEL_STRONG ActivityLevel = "STRONG"
	ACTIVITYLEVEL_RESTRICTED ActivityLevel = "RESTRICTED"
)

// All allowed values of ActivityLevel enum
var AllowedActivityLevelEnumValues = []ActivityLevel{
	"WEAK",
	"GROWING",
	"STRONG",
	"RESTRICTED",
}

func (v *ActivityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActivityLevel(value)
	for _, existing := range AllowedActivityLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActivityLevel", value)
}

// NewActivityLevelFromValue returns a pointer to a valid ActivityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivityLevelFromValue(v string) (*ActivityLevel, error) {
	ev := ActivityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActivityLevel: valid values are %v", v, AllowedActivityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivityLevel) IsValid() bool {
	for _, existing := range AllowedActivityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActivityLevel value
func (v ActivityLevel) Ptr() *ActivityLevel {
	return &v
}

type NullableActivityLevel struct {
	value *ActivityLevel
	isSet bool
}

func (v NullableActivityLevel) Get() *ActivityLevel {
	return v.value
}

func (v *NullableActivityLevel) Set(val *ActivityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLevel(val *ActivityLevel) *NullableActivityLevel {
	return &NullableActivityLevel{value: val, isSet: true}
}

func (v NullableActivityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

