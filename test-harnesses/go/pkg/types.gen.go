// Package pkg provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package pkg

import (
	"encoding/json"
	"fmt"
)

const (
	BotTokenAuthScopes = "BotTokenAuth.Scopes"
	V1Oauth2Scopes     = "V1Oauth2.Scopes"
)

// Defines values for ConversationsListChannelIsChannel.
const (
	ConversationsListChannelIsChannelTrue ConversationsListChannelIsChannel = true
)

// Defines values for ConversationsListConversationType.
const (
	ConversationsListConversationTypeIm ConversationsListConversationType = "im"

	ConversationsListConversationTypeNpim ConversationsListConversationType = "npim"

	ConversationsListConversationTypePrivateChannel ConversationsListConversationType = "private_channel"

	ConversationsListConversationTypePublicChannel ConversationsListConversationType = "public_channel"
)

// Defines values for ConversationsListGroupIsGroup.
const (
	ConversationsListGroupIsGroupTrue ConversationsListGroupIsGroup = true
)

// Defines values for ConversationsListImIsIm.
const (
	ConversationsListImIsImTrue ConversationsListImIsIm = true
)

// Defines values for ConversationsListMpimIsMpim.
const (
	ConversationsListMpimIsMpimTrue ConversationsListMpimIsMpim = true
)

// Public slack channel
type ConversationsListChannel struct {
	Created            float32                           `json:"created"`
	Creator            string                            `json:"creator"`
	Id                 string                            `json:"id"`
	IsArchived         bool                              `json:"is_archived"`
	IsChannel          ConversationsListChannelIsChannel `json:"is_channel"`
	IsExtShared        bool                              `json:"is_ext_shared"`
	IsGeneral          bool                              `json:"is_general"`
	IsGroup            *bool                             `json:"is_group,omitempty"`
	IsIm               *bool                             `json:"is_im,omitempty"`
	IsMember           bool                              `json:"is_member"`
	IsMpim             *bool                             `json:"is_mpim,omitempty"`
	IsOrgShared        bool                              `json:"is_org_shared"`
	IsPendingExtShared bool                              `json:"is_pending_ext_shared"`
	IsPrivate          bool                              `json:"is_private"`
	IsShared           bool                              `json:"is_shared"`
	Name               string                            `json:"name"`
	NameNormalized     string                            `json:"name_normalized"`
	NumMembers         *float32                          `json:"num_members,omitempty"`
	PendingShared      []map[string]interface{}          `json:"pending_shared"`
	PreviousNames      *[]string                         `json:"previous_names,omitempty"`
	Purpose            struct {
		Creator string  `json:"creator"`
		LastSet float32 `json:"last_set"`
		Value   string  `json:"value"`
	} `json:"purpose"`
	Topic struct {
		Creator string  `json:"creator"`
		LastSet float32 `json:"last_set"`
		Value   string  `json:"value"`
	} `json:"topic"`
	Unlinked float32 `json:"unlinked"`
	Updated  float32 `json:"updated"`
}

// ConversationsListChannelIsChannel defines model for ConversationsListChannel.IsChannel.
type ConversationsListChannelIsChannel bool

// ConversationsListConversationType defines model for conversations.list_ConversationType.
type ConversationsListConversationType string

// Schema for error response from conversations.list method
type ConversationsListErrorResponseBody struct {
	Error                string                 `json:"error"`
	Ok                   bool                   `json:"ok"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Private slack channel
type ConversationsListGroup struct {
	Created            float32                       `json:"created"`
	Creator            string                        `json:"creator"`
	Id                 string                        `json:"id"`
	IsArchived         bool                          `json:"is_archived"`
	IsChannel          *bool                         `json:"is_channel,omitempty"`
	IsExtShared        bool                          `json:"is_ext_shared"`
	IsGeneral          bool                          `json:"is_general"`
	IsGroup            ConversationsListGroupIsGroup `json:"is_group"`
	IsIm               *bool                         `json:"is_im,omitempty"`
	IsMember           bool                          `json:"is_member"`
	IsMpim             *bool                         `json:"is_mpim,omitempty"`
	IsOpen             *bool                         `json:"is_open,omitempty"`
	IsOrgShared        bool                          `json:"is_org_shared"`
	IsPendingExtShared bool                          `json:"is_pending_ext_shared"`
	IsPrivate          bool                          `json:"is_private"`
	IsShared           bool                          `json:"is_shared"`
	Name               string                        `json:"name"`
	NameNormalized     string                        `json:"name_normalized"`
	PendingShared      []map[string]interface{}      `json:"pending_shared"`
	Priority           *float32                      `json:"priority,omitempty"`
	Purpose            struct {
		Creator string  `json:"creator"`
		LastSet float32 `json:"last_set"`
		Value   string  `json:"value"`
	} `json:"purpose"`
	Topic struct {
		Creator string  `json:"creator"`
		LastSet float32 `json:"last_set"`
		Value   string  `json:"value"`
	} `json:"topic"`
	Unlinked float32 `json:"unlinked"`
}

// ConversationsListGroupIsGroup defines model for ConversationsListGroup.IsGroup.
type ConversationsListGroupIsGroup bool

// Single party instant message
type ConversationsListIm struct {
	Created       float32                 `json:"created"`
	Id            string                  `json:"id"`
	IsIm          ConversationsListImIsIm `json:"is_im"`
	IsOrgShared   bool                    `json:"is_org_shared"`
	IsUserDeleted bool                    `json:"is_user_deleted"`
	Priority      float32                 `json:"priority"`
	User          string                  `json:"user"`
}

// ConversationsListImIsIm defines model for ConversationsListIm.IsIm.
type ConversationsListImIsIm bool

// Multi party instant message
type ConversationsListMpim struct {
	Created   float32                     `json:"created"`
	Creator   string                      `json:"creator"`
	Id        string                      `json:"id"`
	IsChannel *bool                       `json:"is_channel,omitempty"`
	IsGroup   *bool                       `json:"is_group,omitempty"`
	IsIm      *bool                       `json:"is_im,omitempty"`
	IsMpim    ConversationsListMpimIsMpim `json:"is_mpim"`
	Name      string                      `json:"name"`
}

// ConversationsListMpimIsMpim defines model for ConversationsListMpim.IsMpim.
type ConversationsListMpimIsMpim bool

// Schema for successful response from conversations.list method
type ConversationsListResponseBody struct {
	Channels         []interface{} `json:"channels"`
	Ok               bool          `json:"ok"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Schema for error response from conversations.members method
type ConversationsMembersErrorResponseBody struct {
	Error                string                 `json:"error"`
	Ok                   bool                   `json:"ok"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Schema for successful response from conversations.members method
type ConversationsMembersResponseBody struct {
	Members          []string `json:"members"`
	Ok               bool     `json:"ok"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Schema for error response from users.list method
type UsersListErrorResponseBody struct {
	Error                string                 `json:"error"`
	Ok                   bool                   `json:"ok"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// UsersListMember defines model for users.list_Member.
type UsersListMember struct {
	Color             *string                `json:"color,omitempty"`
	Deleted           bool                   `json:"deleted"`
	Has2fa            *bool                  `json:"has_2fa,omitempty"`
	Id                string                 `json:"id"`
	IsAdmin           *bool                  `json:"is_admin,omitempty"`
	IsAppUser         *bool                  `json:"is_app_user,omitempty"`
	IsBot             bool                   `json:"is_bot"`
	IsOwner           *bool                  `json:"is_owner,omitempty"`
	IsPrimaryOwner    *bool                  `json:"is_primary_owner,omitempty"`
	IsRestricted      *bool                  `json:"is_restricted,omitempty"`
	IsUltraRestricted *bool                  `json:"is_ultra_restricted,omitempty"`
	Name              string                 `json:"name"`
	Profile           UsersListMemberProfile `json:"profile"`
	RealName          *string                `json:"real_name,omitempty"`
	TeamId            string                 `json:"team_id"`
	Tz                *string                `json:"tz,omitempty"`
	TzLabel           *string                `json:"tz_label,omitempty"`
	TzOffset          *float32               `json:"tz_offset,omitempty"`
	Updated           float32                `json:"updated"`
}

// UsersListMemberProfile defines model for users.list_MemberProfile.
type UsersListMemberProfile struct {
	AvatarHash            string  `json:"avatar_hash"`
	DisplayName           string  `json:"display_name"`
	DisplayNameNormalized string  `json:"display_name_normalized"`
	Email                 *string `json:"email,omitempty"`
	FirstName             *string `json:"first_name,omitempty"`
	Image1024             *string `json:"image_1024,omitempty"`
	Image192              *string `json:"image_192,omitempty"`
	Image24               *string `json:"image_24,omitempty"`
	Image32               *string `json:"image_32,omitempty"`
	Image48               *string `json:"image_48,omitempty"`
	Image512              *string `json:"image_512,omitempty"`
	Image72               *string `json:"image_72,omitempty"`
	ImageOriginal         *string `json:"image_original,omitempty"`
	LastName              *string `json:"last_name,omitempty"`
	Phone                 *string `json:"phone,omitempty"`
	RealName              *string `json:"real_name,omitempty"`
	RealNameNormalized    string  `json:"real_name_normalized"`
	Skype                 *string `json:"skype,omitempty"`
	StatusEmoji           string  `json:"status_emoji"`
	StatusText            string  `json:"status_text"`
	Team                  string  `json:"team"`
	Title                 *string `json:"title,omitempty"`
}

// Schema for successful response from users.list method
type UsersListResponseBody struct {
	CacheTs          float32           `json:"cache_ts"`
	Members          []UsersListMember `json:"members"`
	Ok               bool              `json:"ok"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ConversationsListParams defines parameters for ConversationsList.
type ConversationsListParams struct {
	// Set to `true` to exclude archived channels from the list
	ExcludeArchived *bool `json:"exclude_archived,omitempty"`

	// Mix and match channel types by providing a comma-separated list of any combination of `public_channel`, `private_channel`, `mpim`, `im`
	Types *[]ConversationsListConversationType `json:"types,omitempty"`

	// The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn't been reached. Must be an integer no larger than 1000.
	Limit *int `json:"limit,omitempty"`

	// Paginate through collections of data by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request's `response_metadata`. Default value fetches the first "page" of the collection. See [pagination](/docs/pagination) for more detail.
	Cursor *string `json:"cursor,omitempty"`
}

// ConversationsMembersParams defines parameters for ConversationsMembers.
type ConversationsMembersParams struct {
	// ID of the conversation to retrieve members for
	Channel *string `json:"channel,omitempty"`

	// The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn't been reached.
	Limit *int `json:"limit,omitempty"`

	// Paginate through collections of data by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request's `response_metadata`. Default value fetches the first "page" of the collection. See [pagination](/docs/pagination) for more detail.
	Cursor *string `json:"cursor,omitempty"`
}

// UsersListParams defines parameters for UsersList.
type UsersListParams struct {
	// Authentication token. Requires scope: `users:read`
	Token *string `json:"token,omitempty"`

	// The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn't been reached. Providing no `limit` value will result in Slack attempting to deliver you the entire result set. If the collection is too large you may experience `limit_required` or HTTP 500 errors.
	Limit *int `json:"limit,omitempty"`

	// Paginate through collections of data by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request's `response_metadata`. Default value fetches the first "page" of the collection. See [pagination](/docs/pagination) for more detail.
	Cursor *string `json:"cursor,omitempty"`

	// Set this to `true` to receive the locale for users. Defaults to `false`
	IncludeLocale *bool `json:"include_locale,omitempty"`
}

// Getter for additional properties for ConversationsListErrorResponseBody. Returns the specified
// element and whether it was found
func (a ConversationsListErrorResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ConversationsListErrorResponseBody
func (a *ConversationsListErrorResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ConversationsListErrorResponseBody to handle AdditionalProperties
func (a *ConversationsListErrorResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["error"]; found {
		err = json.Unmarshal(raw, &a.Error)
		if err != nil {
			return fmt.Errorf("error reading 'error': %w", err)
		}
		delete(object, "error")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ConversationsListErrorResponseBody to handle AdditionalProperties
func (a ConversationsListErrorResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["error"], err = json.Marshal(a.Error)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'error': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ConversationsListResponseBody. Returns the specified
// element and whether it was found
func (a ConversationsListResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ConversationsListResponseBody
func (a *ConversationsListResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ConversationsListResponseBody to handle AdditionalProperties
func (a *ConversationsListResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["channels"]; found {
		err = json.Unmarshal(raw, &a.Channels)
		if err != nil {
			return fmt.Errorf("error reading 'channels': %w", err)
		}
		delete(object, "channels")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if raw, found := object["response_metadata"]; found {
		err = json.Unmarshal(raw, &a.ResponseMetadata)
		if err != nil {
			return fmt.Errorf("error reading 'response_metadata': %w", err)
		}
		delete(object, "response_metadata")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ConversationsListResponseBody to handle AdditionalProperties
func (a ConversationsListResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["channels"], err = json.Marshal(a.Channels)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'channels': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	object["response_metadata"], err = json.Marshal(a.ResponseMetadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_metadata': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ConversationsMembersErrorResponseBody. Returns the specified
// element and whether it was found
func (a ConversationsMembersErrorResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ConversationsMembersErrorResponseBody
func (a *ConversationsMembersErrorResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ConversationsMembersErrorResponseBody to handle AdditionalProperties
func (a *ConversationsMembersErrorResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["error"]; found {
		err = json.Unmarshal(raw, &a.Error)
		if err != nil {
			return fmt.Errorf("error reading 'error': %w", err)
		}
		delete(object, "error")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ConversationsMembersErrorResponseBody to handle AdditionalProperties
func (a ConversationsMembersErrorResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["error"], err = json.Marshal(a.Error)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'error': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ConversationsMembersResponseBody. Returns the specified
// element and whether it was found
func (a ConversationsMembersResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ConversationsMembersResponseBody
func (a *ConversationsMembersResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ConversationsMembersResponseBody to handle AdditionalProperties
func (a *ConversationsMembersResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["members"]; found {
		err = json.Unmarshal(raw, &a.Members)
		if err != nil {
			return fmt.Errorf("error reading 'members': %w", err)
		}
		delete(object, "members")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if raw, found := object["response_metadata"]; found {
		err = json.Unmarshal(raw, &a.ResponseMetadata)
		if err != nil {
			return fmt.Errorf("error reading 'response_metadata': %w", err)
		}
		delete(object, "response_metadata")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ConversationsMembersResponseBody to handle AdditionalProperties
func (a ConversationsMembersResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["members"], err = json.Marshal(a.Members)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'members': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	object["response_metadata"], err = json.Marshal(a.ResponseMetadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_metadata': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for UsersListErrorResponseBody. Returns the specified
// element and whether it was found
func (a UsersListErrorResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UsersListErrorResponseBody
func (a *UsersListErrorResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UsersListErrorResponseBody to handle AdditionalProperties
func (a *UsersListErrorResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["error"]; found {
		err = json.Unmarshal(raw, &a.Error)
		if err != nil {
			return fmt.Errorf("error reading 'error': %w", err)
		}
		delete(object, "error")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UsersListErrorResponseBody to handle AdditionalProperties
func (a UsersListErrorResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["error"], err = json.Marshal(a.Error)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'error': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for UsersListResponseBody. Returns the specified
// element and whether it was found
func (a UsersListResponseBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UsersListResponseBody
func (a *UsersListResponseBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UsersListResponseBody to handle AdditionalProperties
func (a *UsersListResponseBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["cache_ts"]; found {
		err = json.Unmarshal(raw, &a.CacheTs)
		if err != nil {
			return fmt.Errorf("error reading 'cache_ts': %w", err)
		}
		delete(object, "cache_ts")
	}

	if raw, found := object["members"]; found {
		err = json.Unmarshal(raw, &a.Members)
		if err != nil {
			return fmt.Errorf("error reading 'members': %w", err)
		}
		delete(object, "members")
	}

	if raw, found := object["ok"]; found {
		err = json.Unmarshal(raw, &a.Ok)
		if err != nil {
			return fmt.Errorf("error reading 'ok': %w", err)
		}
		delete(object, "ok")
	}

	if raw, found := object["response_metadata"]; found {
		err = json.Unmarshal(raw, &a.ResponseMetadata)
		if err != nil {
			return fmt.Errorf("error reading 'response_metadata': %w", err)
		}
		delete(object, "response_metadata")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UsersListResponseBody to handle AdditionalProperties
func (a UsersListResponseBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["cache_ts"], err = json.Marshal(a.CacheTs)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'cache_ts': %w", err)
	}

	object["members"], err = json.Marshal(a.Members)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'members': %w", err)
	}

	object["ok"], err = json.Marshal(a.Ok)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ok': %w", err)
	}

	object["response_metadata"], err = json.Marshal(a.ResponseMetadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'response_metadata': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
