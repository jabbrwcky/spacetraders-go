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

// checks if the GetStatus200ResponseAnnouncementsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseAnnouncementsInner{}

// GetStatus200ResponseAnnouncementsInner struct for GetStatus200ResponseAnnouncementsInner
type GetStatus200ResponseAnnouncementsInner struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

type _GetStatus200ResponseAnnouncementsInner GetStatus200ResponseAnnouncementsInner

// NewGetStatus200ResponseAnnouncementsInner instantiates a new GetStatus200ResponseAnnouncementsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseAnnouncementsInner(title string, body string) *GetStatus200ResponseAnnouncementsInner {
	this := GetStatus200ResponseAnnouncementsInner{}
	this.Title = title
	this.Body = body
	return &this
}

// NewGetStatus200ResponseAnnouncementsInnerWithDefaults instantiates a new GetStatus200ResponseAnnouncementsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseAnnouncementsInnerWithDefaults() *GetStatus200ResponseAnnouncementsInner {
	this := GetStatus200ResponseAnnouncementsInner{}
	return &this
}

// GetTitle returns the Title field value
func (o *GetStatus200ResponseAnnouncementsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseAnnouncementsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetStatus200ResponseAnnouncementsInner) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *GetStatus200ResponseAnnouncementsInner) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseAnnouncementsInner) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *GetStatus200ResponseAnnouncementsInner) SetBody(v string) {
	o.Body = v
}

func (o GetStatus200ResponseAnnouncementsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseAnnouncementsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *GetStatus200ResponseAnnouncementsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"body",
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

	varGetStatus200ResponseAnnouncementsInner := _GetStatus200ResponseAnnouncementsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatus200ResponseAnnouncementsInner)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseAnnouncementsInner(varGetStatus200ResponseAnnouncementsInner)

	return err
}

type NullableGetStatus200ResponseAnnouncementsInner struct {
	value *GetStatus200ResponseAnnouncementsInner
	isSet bool
}

func (v NullableGetStatus200ResponseAnnouncementsInner) Get() *GetStatus200ResponseAnnouncementsInner {
	return v.value
}

func (v *NullableGetStatus200ResponseAnnouncementsInner) Set(val *GetStatus200ResponseAnnouncementsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseAnnouncementsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseAnnouncementsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseAnnouncementsInner(val *GetStatus200ResponseAnnouncementsInner) *NullableGetStatus200ResponseAnnouncementsInner {
	return &NullableGetStatus200ResponseAnnouncementsInner{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseAnnouncementsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseAnnouncementsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

