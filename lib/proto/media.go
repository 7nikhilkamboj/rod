// This file is generated by "./lib/proto/generate"

package proto

/*

Media

This domain allows detailed inspection of media elements

*/

// MediaPlayerID Players will get an ID that is unique within the agent context.
type MediaPlayerID string

// MediaTimestamp ...
type MediaTimestamp float64

// MediaPlayerMessageLevel enum
type MediaPlayerMessageLevel string

const (
	// MediaPlayerMessageLevelError enum const
	MediaPlayerMessageLevelError MediaPlayerMessageLevel = "error"

	// MediaPlayerMessageLevelWarning enum const
	MediaPlayerMessageLevelWarning MediaPlayerMessageLevel = "warning"

	// MediaPlayerMessageLevelInfo enum const
	MediaPlayerMessageLevelInfo MediaPlayerMessageLevel = "info"

	// MediaPlayerMessageLevelDebug enum const
	MediaPlayerMessageLevelDebug MediaPlayerMessageLevel = "debug"
)

// MediaPlayerMessage Have one type per entry in MediaLogRecord::Type
// Corresponds to kMessage
type MediaPlayerMessage struct {

	// Level Keep in sync with MediaLogMessageLevel
	// We are currently keeping the message level 'error' separate from the
	// PlayerError type because right now they represent different things,
	// this one being a DVLOG(ERROR) style log message that gets printed
	// based on what log level is selected in the UI, and the other is a
	// representation of a media::PipelineStatus object. Soon however we're
	// going to be moving away from using PipelineStatus for errors and
	// introducing a new error type which should hopefully let us integrate
	// the error log level into the PlayerError type.
	Level MediaPlayerMessageLevel `json:"level"`

	// Message ...
	Message string `json:"message"`
}

// MediaPlayerProperty Corresponds to kMediaPropertyChange
type MediaPlayerProperty struct {

	// Name ...
	Name string `json:"name"`

	// Value ...
	Value string `json:"value"`
}

// MediaPlayerEvent Corresponds to kMediaEventTriggered
type MediaPlayerEvent struct {

	// Timestamp ...
	Timestamp MediaTimestamp `json:"timestamp"`

	// Value ...
	Value string `json:"value"`
}

// MediaPlayerErrorType enum
type MediaPlayerErrorType string

const (
	// MediaPlayerErrorTypePipelineError enum const
	MediaPlayerErrorTypePipelineError MediaPlayerErrorType = "pipeline_error"

	// MediaPlayerErrorTypeMediaError enum const
	MediaPlayerErrorTypeMediaError MediaPlayerErrorType = "media_error"
)

// MediaPlayerError Corresponds to kMediaError
type MediaPlayerError struct {

	// Type ...
	Type MediaPlayerErrorType `json:"type"`

	// ErrorCode When this switches to using media::Status instead of PipelineStatus
	// we can remove "errorCode" and replace it with the fields from
	// a Status instance. This also seems like a duplicate of the error
	// level enum - there is a todo bug to have that level removed and
	// use this instead. (crbug.com/1068454)
	ErrorCode string `json:"errorCode"`
}

// MediaEnable Enables the Media domain
type MediaEnable struct {
}

// ProtoReq of the command
func (m MediaEnable) ProtoReq() string { return "Media.enable" }

// Call of the command, sessionID is optional.
func (m MediaEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MediaDisable Disables the Media domain.
type MediaDisable struct {
}

// ProtoReq of the command
func (m MediaDisable) ProtoReq() string { return "Media.disable" }

// Call of the command, sessionID is optional.
func (m MediaDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// MediaPlayerPropertiesChanged This can be called multiple times, and can be used to set / override /
// remove player properties. A null propValue indicates removal.
type MediaPlayerPropertiesChanged struct {

	// PlayerID ...
	PlayerID MediaPlayerID `json:"playerId"`

	// Properties ...
	Properties []*MediaPlayerProperty `json:"properties"`
}

// ProtoEvent interface
func (evt MediaPlayerPropertiesChanged) ProtoEvent() string {
	return "Media.playerPropertiesChanged"
}

// MediaPlayerEventsAdded Send events as a list, allowing them to be batched on the browser for less
// congestion. If batched, events must ALWAYS be in chronological order.
type MediaPlayerEventsAdded struct {

	// PlayerID ...
	PlayerID MediaPlayerID `json:"playerId"`

	// Events ...
	Events []*MediaPlayerEvent `json:"events"`
}

// ProtoEvent interface
func (evt MediaPlayerEventsAdded) ProtoEvent() string {
	return "Media.playerEventsAdded"
}

// MediaPlayerMessagesLogged Send a list of any messages that need to be delivered.
type MediaPlayerMessagesLogged struct {

	// PlayerID ...
	PlayerID MediaPlayerID `json:"playerId"`

	// Messages ...
	Messages []*MediaPlayerMessage `json:"messages"`
}

// ProtoEvent interface
func (evt MediaPlayerMessagesLogged) ProtoEvent() string {
	return "Media.playerMessagesLogged"
}

// MediaPlayerErrorsRaised Send a list of any errors that need to be delivered.
type MediaPlayerErrorsRaised struct {

	// PlayerID ...
	PlayerID MediaPlayerID `json:"playerId"`

	// Errors ...
	Errors []*MediaPlayerError `json:"errors"`
}

// ProtoEvent interface
func (evt MediaPlayerErrorsRaised) ProtoEvent() string {
	return "Media.playerErrorsRaised"
}

// MediaPlayersCreated Called whenever a player is created, or when a new agent joins and receives
// a list of active players. If an agent is restored, it will receive the full
// list of player ids and all events again.
type MediaPlayersCreated struct {

	// Players ...
	Players []MediaPlayerID `json:"players"`
}

// ProtoEvent interface
func (evt MediaPlayersCreated) ProtoEvent() string {
	return "Media.playersCreated"
}
