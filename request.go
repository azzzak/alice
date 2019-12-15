package alice

import (
	"errors"
	"strings"
)

const (
	// SimpleUtterance простая реплика.
	SimpleUtterance = "SimpleUtterance"

	// ButtonPressed нажатие кнопки.
	ButtonPressed = "ButtonPressed"
)

// Request структура входящего сообщения.
type Request struct {
	Meta struct {
		Locale     string `json:"locale"`
		Timezone   string `json:"timezone"`
		ClientID   string `json:"client_id"`
		Interfaces struct {
			AccountLinking *struct{} `json:"account_linking"`
			Screen         *struct{} `json:"screen"`
		} `json:"interfaces"`
	} `json:"meta"`

	LinkingComplete *struct{} `json:"account_linking_complete_event,omitenpty"`

	Request struct {
		Command           string `json:"command"`
		OriginalUtterance string `json:"original_utterance"`
		Type              string `json:"type"`
		Markup            struct {
			DangerousContext *bool `json:"dangerous_context,omitempty"`
		} `json:"markup,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
		NLU     struct {
			Tokens   []string `json:"tokens"`
			Entities []Entity `json:"entities,omitempty"`
		} `json:"nlu"`
	} `json:"request"`

	Session struct {
		New       bool   `json:"new"`
		MessageID int    `json:"message_id"`
		SessionID string `json:"session_id"`
		SkillID   string `json:"skill_id"`
		UserID    string `json:"user_id"`
	} `json:"session"`

	Version string `json:"version"`
	Bearer  string
}

func (req *Request) clean() *Request {
	req.Meta.Interfaces = struct {
		AccountLinking *struct{} `json:"account_linking"`
		Screen         *struct{} `json:"screen"`
	}{
		nil,
		nil,
	}
	req.LinkingComplete = nil
	req.Request.Command = ""
	req.Request.OriginalUtterance = ""
	req.Request.Payload = nil
	req.Request.Markup = struct {
		DangerousContext *bool `json:"dangerous_context,omitempty"`
	}{
		nil,
	}
	req.Request.NLU = struct {
		Tokens   []string `json:"tokens"`
		Entities []Entity `json:"entities,omitempty"`
	}{
		[]string{},
		[]Entity{},
	}
	req.Bearer = ""
	return req
}

// Locale язык в формате POSIX.
func (req *Request) Locale() string {
	return req.Meta.Locale
}

// Timezone название часового пояса.
func (req *Request) Timezone() string {
	return req.Meta.Timezone
}

// ClientID идентификатор клиентского устройства или приложения. Не рекомендуется использовать.
func (req *Request) ClientID() string {
	return req.Meta.ClientID
}

// CanAccountLinking поддерживает ли устройство пользователя создание связки аккаунтов.
func (req *Request) CanAccountLinking() bool {
	return req.Meta.Interfaces.AccountLinking != nil
}

// HasScreen имеет ли устройство пользователя экран.
func (req *Request) HasScreen() bool {
	return req.Meta.Interfaces.Screen != nil
}

// IsLinkingComplete связка аккаунтов создана.
func (req *Request) IsLinkingComplete() bool {
	return req.LinkingComplete != nil
}

// Command реплика пользователя, преобразованная Алисой. В частности, текст очищается от знаков препинания, а числительные преобразуются в числа.
func (req *Request) Command() string {
	return req.Request.Command
}

// OriginalUtterance реплика пользователя без изменений.
func (req *Request) OriginalUtterance() string {
	return req.Request.OriginalUtterance
}

// Text синоним метода OriginalUtterance().
func (req *Request) Text() string {
	return req.OriginalUtterance()
}

// Type тип запроса (реплика или нажатие кнопки).
func (req *Request) Type() string {
	return req.Request.Type
}

// DangerousContext флаг опасной реплики.
func (req *Request) DangerousContext() bool {
	if req.Request.Markup.DangerousContext != nil {
		return *req.Request.Markup.DangerousContext
	}
	return false
}

// Payload возвращает map[string]interface{} с данными, которые были переданы в Payload кнопки. Подходит для Payload, оформленного в виде json-объекта. Если Payload простая строка следует использовать метод PayloadString(). Если в запросе нет Payload возвращается ошибка.
func (req *Request) Payload() (map[string]interface{}, error) {
	if req.Request.Payload != nil {
		return req.Request.Payload.(map[string]interface{}), nil
	}
	return nil, errors.New("Payload is nil")
}

// PayloadString возвращает строку, которая была передана в Payload кнопки. Для сложных объектов следует использовать метод Payload(). Если в запросе нет Payload возвращается ошибка.
func (req *Request) PayloadString() (string, error) {
	if req.Request.Payload != nil {
		return req.Request.Payload.(string), nil
	}
	return "", errors.New("Payload is nil")
}

// Tokens массив слов из реплики.
func (req *Request) Tokens() []string {
	return req.Request.NLU.Tokens
}

// IsNewSession отправлена реплика в рамках нового разговора или уже начатого.
func (req *Request) IsNewSession() bool {
	return req.Session.New
}

// MessageID счетчик сообщений в рамках сессии.
func (req *Request) MessageID() int {
	return req.Session.MessageID
}

// SessionID идентификатор сессии.
func (req *Request) SessionID() string {
	return req.Session.SessionID
}

// SkillID идентификатор навыка.
func (req *Request) SkillID() string {
	return req.Session.SkillID
}

// UserID идентификатор пользователя.
func (req *Request) UserID() string {
	return req.Session.UserID
}

// Ver версия протокола.
func (req *Request) Ver() string {
	return req.Version
}

// AuthToken токен, полученный при связке аккаунтов.
func (req *Request) AuthToken() string {
	if req.Bearer != "" {
		return strings.TrimPrefix(req.Bearer, "Bearer ")
	}
	return ""
}
