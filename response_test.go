package alice_test

import (
	"encoding/json"
	"testing"

	"github.com/azzzak/alice"
	"github.com/stretchr/testify/assert"
)

func TestResponse_Text(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"

	test1 := getResp(0)
	test1.Response.Text = "раз два три"

	test2 := getResp(0)
	test2.Response.Text = ""

	tests := []struct {
		name     string
		response *alice.Response
		args     []string
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     []string{"тест"},
			want:     test0,
		},
		{
			name:     "",
			response: getResp(0),
			args:     []string{"раз", "два", "три"},
			want:     test1,
		}, {
			name:     "",
			response: getResp(0),
			args:     []string{""},
			want:     test2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Text(tt.args...))
		})
	}
}

func TestResponse_RandomText(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"

	test1 := getResp(0)
	test1.Response.Text = ""

	tests := []struct {
		name     string
		response *alice.Response
		args     []string
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     []string{"тест"},
			want:     test0,
		}, {
			name:     "",
			response: getResp(0),
			args:     []string{""},
			want:     test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.RandomText(tt.args...))
		})
	}
}

func TestResponse_Space(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = " "

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Space())
		})
	}
}

func TestResponse_S(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = " "

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.S())
		})
	}
}

func TestResponse_ResetText(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: test0,
			want:     getResp(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.ResetText())
		})
	}
}

func TestResponse_TTS(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = "тест"

	test1 := getResp(0)
	test1.Response.TTS = "раз два три"

	test2 := getResp(0)
	test2.Response.TTS = ""

	tests := []struct {
		name     string
		response *alice.Response
		args     []string
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     []string{"тест"},
			want:     test0,
		},
		{
			name:     "",
			response: getResp(0),
			args:     []string{"раз", "два", "три"},
			want:     test1,
		}, {
			name:     "",
			response: getResp(0),
			args:     []string{""},
			want:     test2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.TTS(tt.args...))
		})
	}
}

func TestResponse_TextWithTTS(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"
	test0.Response.TTS = "ттс"

	type args struct {
		s   string
		tts string
	}
	tests := []struct {
		name     string
		response *alice.Response
		args     args
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     args{"тест", "ттс"},
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.TextWithTTS(tt.args.s, tt.args.tts))
		})
	}
}

func TestResponse_Pause(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = " - - - - - "

	tests := []struct {
		name     string
		response *alice.Response
		args     int
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     5,
			want:     test0,
		},
		{
			name:     "",
			response: getResp(0),
			args:     0,
			want:     getResp(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Pause(tt.args))
		})
	}
}

func TestResponse_Effect(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = `<speaker effect="pulse">`

	tests := []struct {
		name     string
		response *alice.Response
		args     string
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     "pulse",
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Effect(tt.args))
		})
	}
}

func TestResponse_NoEffect(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = `<speaker effect="-">`

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.NoEffect())
		})
	}
}

func TestResponse_Sound(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = `<speaker sound="alice-sounds-things-chainsaw-1.opus">`

	tests := []struct {
		name     string
		response *alice.Response
		args     string
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     "alice-sounds-things-chainsaw-1.opus",
			want:     test0,
		}, {
			name:     "",
			response: getResp(0),
			args:     "alice-sounds-things-chainsaw-1",
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Sound(tt.args))
		})
	}
}

func TestResponse_CustomSound(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = `<speaker audio='dialogs-upload/e03f8d5b-35ef-4d57-9450-b721ca17a6c3/sound.opus'>`

	type args struct {
		skill string
		sound string
	}
	tests := []struct {
		name     string
		response *alice.Response
		args     args
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args:     args{"e03f8d5b-35ef-4d57-9450-b721ca17a6c3", "sound.opus"},
			want:     test0,
		}, {
			name:     "",
			response: getResp(0),
			args:     args{"e03f8d5b-35ef-4d57-9450-b721ca17a6c3", "sound"},
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.CustomSound(tt.args.skill, tt.args.sound))
		})
	}
}

func TestResponse_ResetTTS(t *testing.T) {
	test0 := getResp(0)
	test0.Response.TTS = "ттс"

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: test0,
			want:     getResp(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.ResetTTS())
		})
	}
}

func TestNewButton(t *testing.T) {
	test0 := alice.Button{
		Title:   "кнопка",
		URL:     "https://google.com",
		Hide:    false,
		Payload: nil,
	}

	test1 := alice.Button{
		Title:   "кнопка2",
		URL:     "https://google.com",
		Hide:    true,
		Payload: `{"msg":"ok"}`,
	}

	type args struct {
		title   string
		url     string
		hide    bool
		payload []interface{}
	}
	tests := []struct {
		name string
		args args
		want alice.Button
	}{
		{
			name: "",
			args: args{
				title:   "кнопка",
				url:     "https://google.com",
				hide:    false,
				payload: nil,
			},
			want: test0,
		}, {
			name: "",
			args: args{
				title:   "кнопка2",
				url:     "https://google.com",
				hide:    true,
				payload: []interface{}{`{"msg":"ok"}`},
			},
			want: test1,
		}, {
			name: "",
			args: args{
				title:   "кнопка2",
				url:     "https://google.com",
				hide:    true,
				payload: []interface{}{`{"msg":"ok"}`, `{"msg":"nope"}`},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, alice.NewButton(tt.args.title, tt.args.url, tt.args.hide, tt.args.payload...))
		})
	}
}

func TestResponse_Button(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "кнопка"
	test0.Response.EndSession = false

	test1 := getResp(0)
	test1.Response.Text = "кнопка2"

	type args struct {
		title   string
		url     string
		hide    bool
		payload []interface{}
	}
	tests := []struct {
		name     string
		response *alice.Response
		args     args
		want     *alice.Response
	}{
		{
			name:     "",
			response: test0,
			args: args{
				title:   "кнопка",
				url:     "https://google.com",
				hide:    false,
				payload: nil,
			},
			want: getResp(1),
		},
		{
			name:     "",
			response: test1,
			args: args{
				title:   "кнопка2",
				url:     "",
				hide:    true,
				payload: []interface{}{"msg", "nope"},
			},
			want: getResp(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Button(tt.args.title, tt.args.url, tt.args.hide, tt.args.payload...))
		})
	}
}

func TestResponse_Buttons(t *testing.T) {
	b0 := alice.Button{
		Title:   "кнопка",
		URL:     "https://google.com",
		Hide:    false,
		Payload: nil,
	}
	b1 := alice.Button{
		Title:   "кнопка2",
		Hide:    true,
		Payload: nil,
	}
	b2 := alice.Button{
		Title:   "кнопка2",
		Hide:    true,
		Payload: nil,
	}
	test0 := getResp(0)
	test0.Response.Buttons = []alice.Button{b0, b1, b2}

	b3, b4 := b0, b1
	b3.Payload = "msg"
	b4.Payload = "ok"
	test1 := getResp(0)
	test1.Response.Buttons = []alice.Button{b3, b4}

	type args struct {
		buttons []alice.Button
	}
	tests := []struct {
		name     string
		response *alice.Response
		args     args
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args: args{
				buttons: []alice.Button{b0, b1, b2},
			},
			want: test0,
		}, {
			name:     "",
			response: getResp(0),
			args: args{
				buttons: []alice.Button{b3, b4},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Buttons(tt.args.buttons...))
		})
	}
}

func TestResponse_EndSession(t *testing.T) {
	test0 := getResp(0)
	test0.Response.EndSession = true

	tests := []struct {
		name     string
		response *alice.Response
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			want:     test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.EndSession())
		})
	}
}

func TestNewPhrase(t *testing.T) {
	type args struct {
		text string
		tts  string
	}
	tests := []struct {
		name string
		args args
		want alice.Phrase
	}{
		{
			name: "",
			args: args{
				text: "текст",
				tts:  "ттс",
			},
			want: alice.Phrase{
				Text: "текст",
				TTS:  "ттс",
			},
		},
		{
			name: "",
			args: args{
				text: "",
				tts:  "",
			},
			want: alice.Phrase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alice.NewPhrase(tt.args.text, tt.args.tts)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_Phrase(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"
	test0.Response.TTS = "ттс"

	test1 := getResp(0)
	test1.Response.Text = ""
	test1.Response.TTS = ""

	tests := []struct {
		name     string
		response *alice.Response
		args     alice.Phrase
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args: alice.Phrase{
				Text: "тест",
				TTS:  "ттс",
			},
			want: test0,
		}, {
			name:     "",
			response: getResp(0),
			args:     alice.Phrase{},
			want:     test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.Phrase(tt.args))
		})
	}
}

func TestResponse_RandomPhrase(t *testing.T) {
	test0 := getResp(0)
	test0.Response.Text = "тест"
	test0.Response.TTS = "ттс"

	test1 := getResp(0)
	test1.Response.Text = ""
	test1.Response.TTS = ""

	tests := []struct {
		name     string
		response *alice.Response
		args     []alice.Phrase
		want     *alice.Response
	}{
		{
			name:     "",
			response: getResp(0),
			args: []alice.Phrase{
				{
					Text: "тест",
					TTS:  "ттс",
				},
			},
			want: test0,
		}, {
			name:     "",
			response: getResp(0),
			args: []alice.Phrase{
				{},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			assert.Equal(t, tt.want, resp.RandomPhrase(tt.args...))
		})
	}
}

func getResp(n int) *alice.Response {
	source := []string{`{"response":{"text":"","end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`,

		`{"response":{"text":"кнопка","buttons":[{"title":"кнопка","url":"https://google.com"}],"end_session":false},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`,

		`{"response":{"text":"кнопка2","buttons":[{"title":"кнопка2","hide":true,"payload":"msg"}],"end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`,

		`{"response":{"text":"тест","buttons":[{"title":"кнопка","url":"https://google.com"},{"title":"кнопка2","hide":true},{"title":"кнопка3","hide":true}],"end_session":false},"session":{"message_id":1,"session_id":"","user_id":""},"version":"1.0"}`,

		`{"response":{"text":"тест","buttons":[{"title":"кнопка","url":"https://google.com","payload":"msg"},{"title":"кнопка2","payload":"ok","hide":true}],"end_session":false},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`}

	var resp = new(alice.Response)
	json.Unmarshal([]byte(source[n]), resp)
	return resp
}
