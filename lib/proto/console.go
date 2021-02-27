// This file is generated by "./lib/proto/generate"

package proto

/*

Console

This domain is deprecated - use Runtime or Log instead.

*/

// ConsoleConsoleMessageSource enum
type ConsoleConsoleMessageSource string

const (
	// ConsoleConsoleMessageSourceXML enum const
	ConsoleConsoleMessageSourceXML ConsoleConsoleMessageSource = "xml"

	// ConsoleConsoleMessageSourceJavascript enum const
	ConsoleConsoleMessageSourceJavascript ConsoleConsoleMessageSource = "javascript"

	// ConsoleConsoleMessageSourceNetwork enum const
	ConsoleConsoleMessageSourceNetwork ConsoleConsoleMessageSource = "network"

	// ConsoleConsoleMessageSourceConsoleAPI enum const
	ConsoleConsoleMessageSourceConsoleAPI ConsoleConsoleMessageSource = "console-api"

	// ConsoleConsoleMessageSourceStorage enum const
	ConsoleConsoleMessageSourceStorage ConsoleConsoleMessageSource = "storage"

	// ConsoleConsoleMessageSourceAppcache enum const
	ConsoleConsoleMessageSourceAppcache ConsoleConsoleMessageSource = "appcache"

	// ConsoleConsoleMessageSourceRendering enum const
	ConsoleConsoleMessageSourceRendering ConsoleConsoleMessageSource = "rendering"

	// ConsoleConsoleMessageSourceSecurity enum const
	ConsoleConsoleMessageSourceSecurity ConsoleConsoleMessageSource = "security"

	// ConsoleConsoleMessageSourceOther enum const
	ConsoleConsoleMessageSourceOther ConsoleConsoleMessageSource = "other"

	// ConsoleConsoleMessageSourceDeprecation enum const
	ConsoleConsoleMessageSourceDeprecation ConsoleConsoleMessageSource = "deprecation"

	// ConsoleConsoleMessageSourceWorker enum const
	ConsoleConsoleMessageSourceWorker ConsoleConsoleMessageSource = "worker"
)

// ConsoleConsoleMessageLevel enum
type ConsoleConsoleMessageLevel string

const (
	// ConsoleConsoleMessageLevelLog enum const
	ConsoleConsoleMessageLevelLog ConsoleConsoleMessageLevel = "log"

	// ConsoleConsoleMessageLevelWarning enum const
	ConsoleConsoleMessageLevelWarning ConsoleConsoleMessageLevel = "warning"

	// ConsoleConsoleMessageLevelError enum const
	ConsoleConsoleMessageLevelError ConsoleConsoleMessageLevel = "error"

	// ConsoleConsoleMessageLevelDebug enum const
	ConsoleConsoleMessageLevelDebug ConsoleConsoleMessageLevel = "debug"

	// ConsoleConsoleMessageLevelInfo enum const
	ConsoleConsoleMessageLevelInfo ConsoleConsoleMessageLevel = "info"
)

// ConsoleConsoleMessage Console message.
type ConsoleConsoleMessage struct {

	// Source Message source.
	Source ConsoleConsoleMessageSource `json:"source"`

	// Level Message severity.
	Level ConsoleConsoleMessageLevel `json:"level"`

	// Text Message text.
	Text string `json:"text"`

	// URL (optional) URL of the message origin.
	URL string `json:"url,omitempty"`

	// Line (optional) Line number in the resource that generated this message (1-based).
	Line int `json:"line,omitempty"`

	// Column (optional) Column number in the resource that generated this message (1-based).
	Column int `json:"column,omitempty"`
}

// ConsoleClearMessages Does nothing.
type ConsoleClearMessages struct {
}

// ProtoReq of the command
func (m ConsoleClearMessages) ProtoReq() string { return "Console.clearMessages" }

// Call of the command, sessionID is optional.
func (m ConsoleClearMessages) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ConsoleDisable Disables console domain, prevents further console messages from being reported to the client.
type ConsoleDisable struct {
}

// ProtoReq of the command
func (m ConsoleDisable) ProtoReq() string { return "Console.disable" }

// Call of the command, sessionID is optional.
func (m ConsoleDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ConsoleEnable Enables console domain, sends the messages collected so far to the client by means of the
// `messageAdded` notification.
type ConsoleEnable struct {
}

// ProtoReq of the command
func (m ConsoleEnable) ProtoReq() string { return "Console.enable" }

// Call of the command, sessionID is optional.
func (m ConsoleEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ConsoleMessageAdded Issued when new console message is added.
type ConsoleMessageAdded struct {

	// Message Console message that has been added.
	Message *ConsoleConsoleMessage `json:"message"`
}

// ProtoEvent interface
func (evt ConsoleMessageAdded) ProtoEvent() string {
	return "Console.messageAdded"
}
