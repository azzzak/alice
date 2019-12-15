package alice

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/azzzak/alice/effects"
)

// Response структура исходящего сообщения.
type Response struct {
	Response *struct {
		Text       string   `json:"text"`
		TTS        string   `json:"tts,omitempty"`
		Card       *Card    `json:"card,omitempty"`
		Buttons    []Button `json:"buttons,omitempty"`
		EndSession bool     `json:"end_session"`
	} `json:"response,omitempty"`

	StartAccountLinking *struct{} `json:"start_account_linking,omitempty"`

	Session struct {
		MessageID int    `json:"message_id"`
		SessionID string `json:"session_id"`
		UserID    string `json:"user_id"`
	} `json:"session"`

	Version string `json:"version"`
}

func (resp *Response) clean() *Response {
	resp.Response = &struct {
		Text       string   `json:"text"`
		TTS        string   `json:"tts,omitempty"`
		Card       *Card    `json:"card,omitempty"`
		Buttons    []Button `json:"buttons,omitempty"`
		EndSession bool     `json:"end_session"`
	}{}
	resp.StartAccountLinking = nil
	return resp
}

func (resp *Response) prepareResponse(req *Request) *Response {
	resp.Session.MessageID = req.Session.MessageID
	resp.Session.SessionID = req.Session.SessionID
	resp.Session.UserID = req.Session.UserID
	resp.Version = "1.0"
	return resp
}

// StartAuthorization начать создание связки аккаунтов и показать пользователю карточку авторизации.
func (resp *Response) StartAuthorization() *Response {
	// resp.Response = Response{}
	resp.StartAccountLinking = &struct{}{}
	resp.Response = nil
	return resp
}

// Text добавляет строку к текстовому ответу. Если передано несколько строк, они будут разделены пробелом.
func (resp *Response) Text(s ...string) *Response {
	resp.Response.Text += strings.Join(s, " ")
	return resp
}

// RandomText добавляет к текстовому ответу случайную строку из числа предложенных.
func (resp *Response) RandomText(s ...string) *Response {
	ix := rand.Intn(len(s))
	resp.Response.Text += s[ix]
	return resp
}

// Space добавляет пробел к текстовому ответу.
func (resp *Response) Space() *Response {
	resp.Response.Text += " "
	return resp
}

// S синоним метода Space().
func (resp *Response) S() *Response {
	return resp.Space()
}

// ResetText обнуляет текстовый ответ.
func (resp *Response) ResetText() *Response {
	resp.Response.Text = ""
	return resp
}

// TTS добавляет строку к TTS. Если передано несколько строк, они будут разделены пробелом.
func (resp *Response) TTS(tts ...string) *Response {
	resp.Response.TTS += strings.Join(tts, " ")
	return resp
}

// TextWithTTS добавляет строки к текстовому ответу и к TTS.
func (resp *Response) TextWithTTS(s, tts string) *Response {
	return resp.Text(s).TTS(tts)
}

// Pause добавляет паузу к TTS (знак "-").
func (resp *Response) Pause(n int) *Response {
	if n < 1 {
		return resp
	}

	p := strings.Repeat("- ", n)
	resp.Response.TTS += " " + p
	return resp
}

// Effect накладывает звуковой эффект на TTS.
func (resp *Response) Effect(effect string) *Response {
	e := fmt.Sprintf("<speaker effect=\"%s\">", effect)
	resp.Response.TTS += e
	return resp
}

// NoEffect отключает наложенный на TTS эффект.
func (resp *Response) NoEffect() *Response {
	return resp.Effect(effects.NoEffect)
}

// Sound добавляет к TTS звук из библиотеки Алисы.
func (resp *Response) Sound(sound string) *Response {
	sound = strings.TrimSuffix(sound, ".opus")
	s := fmt.Sprintf("<speaker sound=\"%s.opus\">", sound)
	resp.Response.TTS += s
	return resp
}

// CustomSound добавляет к TTS собственный звук.
func (resp *Response) CustomSound(skill, sound string) *Response {
	sound = strings.TrimSuffix(sound, ".opus")
	s := fmt.Sprintf("<speaker audio='dialogs-upload/%s/%s.opus'>", skill, sound)
	resp.Response.TTS += s
	return resp
}

// ResetTTS обнуляет TTS.
func (resp *Response) ResetTTS() *Response {
	resp.Response.TTS = ""
	return resp
}

// Button структура кнопки.
type Button struct {
	Title   string      `json:"title"`
	URL     string      `json:"url,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Hide    bool        `json:"hide,omitempty"`
}

// NewButton создает кнопку. Payload может быть проигнорирован. Если задано больше одного payload используется только первый.
func NewButton(title, url string, hide bool, payload ...interface{}) Button {
	var p interface{}
	if len(payload) > 0 {
		p = payload[0]
	}

	return Button{
		Title:   title,
		URL:     url,
		Hide:    hide,
		Payload: p,
	}
}

// Button создает кнопку и добавляет ее в ответ. Payload может быть проигнорирован. Если задано больше одного payload используется только первый.
func (resp *Response) Button(title, url string, hide bool, payload ...interface{}) *Response {
	var p interface{}
	if len(payload) > 0 {
		p = payload[0]
	}

	resp.Buttons(NewButton(title, url, hide, p))
	return resp
}

// Buttons добавляет одну или несколько кнопок в ответ.
func (resp *Response) Buttons(buttons ...Button) *Response {
	resp.Response.Buttons = append(resp.Response.Buttons, buttons...)
	return resp
}

// EndSession флаг о закрытии сессии.
func (resp *Response) EndSession() *Response {
	resp.Response.EndSession = true
	return resp
}

// Phrase структура фразы.
type Phrase struct {
	Text string
	TTS  string
}

// NewPhrase создает фразу с текстом и TTS.
func NewPhrase(text, tts string) Phrase {
	return Phrase{
		Text: text,
		TTS:  tts,
	}
}

// Phrase добавляет к тексту и TTS ответа данные фразы.
func (resp *Response) Phrase(p ...Phrase) *Response {
	ix := rand.Intn(len(p))
	resp.Response.Text += p[ix].Text
	resp.Response.TTS += p[ix].TTS
	return resp
}

// RandomPhrase добавляет к тексту и TTS ответа данные случайной фразы из числа предложенных.
func (resp *Response) RandomPhrase(p ...Phrase) *Response {
	ix := rand.Intn(len(p))
	resp.Response.Text += p[ix].Text
	resp.Response.TTS += p[ix].TTS
	return resp
}
