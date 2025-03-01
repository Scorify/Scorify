// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/audit"
	"github.com/scorify/scorify/pkg/ent/status"
)

type AuditLogInput struct {
	FromTime *time.Time     `json:"from_time,omitempty"`
	ToTime   *time.Time     `json:"to_time,omitempty"`
	Actions  []audit.Action `json:"actions,omitempty"`
	Message  *string        `json:"message,omitempty"`
	IP       *string        `json:"ip,omitempty"`
	Users    []uuid.UUID    `json:"users,omitempty"`
}

type AuditLogQueryInput struct {
	FromTime *time.Time     `json:"from_time,omitempty"`
	ToTime   *time.Time     `json:"to_time,omitempty"`
	Actions  []audit.Action `json:"actions,omitempty"`
	Message  *string        `json:"message,omitempty"`
	IP       *string        `json:"ip,omitempty"`
	Users    []uuid.UUID    `json:"users,omitempty"`
	Limit    *int           `json:"limit,omitempty"`
	Offset   *int           `json:"offset,omitempty"`
}

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	URL  string    `json:"url"`
}

type InjectSubmissionByUser struct {
	User        *ent.User               `json:"user"`
	Submissions []*ent.InjectSubmission `json:"submissions"`
}

type KothCheckScore struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	User       *ent.User `json:"user,omitempty"`
	Host       *string   `json:"host,omitempty"`
	Error      *string   `json:"error,omitempty"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

type KothScoreboard struct {
	Round  *ent.Round        `json:"round"`
	Checks []*KothCheckScore `json:"checks"`
	Scores []*Score          `json:"scores"`
}

type LoginOutput struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Expires  int    `json:"expires"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Secure   bool   `json:"secure"`
	HTTPOnly bool   `json:"httpOnly"`
}

type MinionStatusSummary struct {
	Total   int `json:"total"`
	Up      int `json:"up"`
	Down    int `json:"down"`
	Unknown int `json:"unknown"`
}

type Notification struct {
	Message string           `json:"message"`
	Type    NotificationType `json:"type"`
}

type RubricFieldInput struct {
	Name  string  `json:"name"`
	Score int     `json:"score"`
	Notes *string `json:"notes,omitempty"`
}

type RubricInput struct {
	Fields []*RubricFieldInput `json:"fields"`
	Notes  *string             `json:"notes,omitempty"`
}

type RubricTemplateFieldInput struct {
	Name     string `json:"name"`
	MaxScore int    `json:"max_score"`
}

type RubricTemplateInput struct {
	Fields   []*RubricTemplateFieldInput `json:"fields"`
	MaxScore int                         `json:"max_score"`
}

type SchemaField struct {
	Name    string          `json:"name"`
	Type    SchemaFieldType `json:"type"`
	Default *string         `json:"default,omitempty"`
	Enum    []string        `json:"enum,omitempty"`
}

type Score struct {
	User  *ent.User `json:"user"`
	Score int       `json:"score"`
}

type Scoreboard struct {
	Teams    []*ent.User     `json:"teams"`
	Checks   []*ent.Check    `json:"checks"`
	Round    *ent.Round      `json:"round"`
	Statuses [][]*ent.Status `json:"statuses"`
	Scores   []*Score        `json:"scores"`
}

type Source struct {
	Name   string         `json:"name"`
	Schema []*SchemaField `json:"schema"`
}

type StatusesQueryInput struct {
	FromTime  *time.Time      `json:"from_time,omitempty"`
	ToTime    *time.Time      `json:"to_time,omitempty"`
	FromRound *int            `json:"from_round,omitempty"`
	ToRound   *int            `json:"to_round,omitempty"`
	Minions   []uuid.UUID     `json:"minions,omitempty"`
	Checks    []uuid.UUID     `json:"checks,omitempty"`
	Users     []uuid.UUID     `json:"users,omitempty"`
	Statuses  []status.Status `json:"statuses,omitempty"`
	Limit     *int            `json:"limit,omitempty"`
	Offset    *int            `json:"offset,omitempty"`
}

type Subscription struct {
}

type EngineState string

const (
	EngineStatePaused   EngineState = "paused"
	EngineStateWaiting  EngineState = "waiting"
	EngineStateRunning  EngineState = "running"
	EngineStateStopping EngineState = "stopping"
)

var AllEngineState = []EngineState{
	EngineStatePaused,
	EngineStateWaiting,
	EngineStateRunning,
	EngineStateStopping,
}

func (e EngineState) IsValid() bool {
	switch e {
	case EngineStatePaused, EngineStateWaiting, EngineStateRunning, EngineStateStopping:
		return true
	}
	return false
}

func (e EngineState) String() string {
	return string(e)
}

func (e *EngineState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EngineState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EngineState", str)
	}
	return nil
}

func (e EngineState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationType string

const (
	NotificationTypeDefault NotificationType = "default"
	NotificationTypeError   NotificationType = "error"
	NotificationTypeInfo    NotificationType = "info"
	NotificationTypeSuccess NotificationType = "success"
	NotificationTypeWarning NotificationType = "warning"
)

var AllNotificationType = []NotificationType{
	NotificationTypeDefault,
	NotificationTypeError,
	NotificationTypeInfo,
	NotificationTypeSuccess,
	NotificationTypeWarning,
}

func (e NotificationType) IsValid() bool {
	switch e {
	case NotificationTypeDefault, NotificationTypeError, NotificationTypeInfo, NotificationTypeSuccess, NotificationTypeWarning:
		return true
	}
	return false
}

func (e NotificationType) String() string {
	return string(e)
}

func (e *NotificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationType", str)
	}
	return nil
}

func (e NotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SchemaFieldType string

const (
	SchemaFieldTypeString SchemaFieldType = "string"
	SchemaFieldTypeInt    SchemaFieldType = "int"
	SchemaFieldTypeBool   SchemaFieldType = "bool"
)

var AllSchemaFieldType = []SchemaFieldType{
	SchemaFieldTypeString,
	SchemaFieldTypeInt,
	SchemaFieldTypeBool,
}

func (e SchemaFieldType) IsValid() bool {
	switch e {
	case SchemaFieldTypeString, SchemaFieldTypeInt, SchemaFieldTypeBool:
		return true
	}
	return false
}

func (e SchemaFieldType) String() string {
	return string(e)
}

func (e *SchemaFieldType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SchemaFieldType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SchemaFieldType", str)
	}
	return nil
}

func (e SchemaFieldType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
