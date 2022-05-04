package models

import (
	"bytes"
	"encoding/json"
	"gorm.io/gorm"
	"html/template"
	"reflect"
)

type LogAction string

const (
	LogActionCreateRouter LogAction = "LogActionCreateRouter"
	LogActionDeleteRouter LogAction = "LogActionUpdateRouter"
	LogActionUpdateRouter LogAction = "LogActionUpdateRoute"

	LogActionCreateMiddleware LogAction = "LogActionCreateMiddleware"
	LogActionDeleteMiddleware LogAction = "LogActionDeleteMiddleware"
	LogActionUpdateMiddleware LogAction = "LogActionUpdateMiddleware"

	LogActionCommit LogAction = "LogActionCommit"
)

var translationKeys = map[LogAction]string{
	LogActionCreateRouter:     "Created router {{.name}}",
	LogActionDeleteRouter:     "Deleted router {{.name}} (#{{.id}})",
	LogActionUpdateRouter:     "Updated router {{.name}} (#{{.id}})",
	LogActionCreateMiddleware: "Created middleware #{{.id}}",
	LogActionDeleteMiddleware: "Deleted middleware #{{.id}}",
	LogActionUpdateMiddleware: "Updated middleware #{{.id}}",
	LogActionCommit:           "Committed Config",
}

type LogEntry struct {
	gorm.Model

	User     string    `json:"user"`
	Action   LogAction `json:"action"`
	Metadata string    `json:"metadata"`
}

func (logEntry *LogEntry) Translate() string {
	t := template.Must(template.New("").Parse(translationKeys[logEntry.Action]))

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(logEntry.Metadata), &metadata); err != nil {
		return ""
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, metadata); err != nil {
		return ""
	}

	return tpl.String()
}

func (logEntry *LogEntry) Validate() bool {

	if len(logEntry.User) == 0 {
		return false
	}

	keys := reflect.ValueOf(translationKeys).MapKeys()
	foundOne := false
	for _, key := range keys {
		if logEntry.Action == key.Interface().(LogAction) {
			foundOne = true
			break
		}
	}
	if !foundOne {
		return false
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(logEntry.Metadata), &metadata); err != nil {
		return false
	}

	return true
}
