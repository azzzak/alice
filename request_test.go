package alice_test

import (
	"encoding/json"
	"testing"

	"github.com/azzzak/alice"
	"github.com/stretchr/testify/assert"
)

func TestRequest_Locale(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "ru-RU",
		}, {
			name:    "",
			request: getReq(1),
			want:    "ru-RU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Locale(); got != tt.want {
				t.Errorf("Request.Locale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Timezone(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "UTC",
		}, {
			name:    "",
			request: getReq(1),
			want:    "UTC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Timezone(); got != tt.want {
				t.Errorf("Request.Timezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_ClientID(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
		}, {
			name:    "",
			request: getReq(1),
			want:    "ru.yandex.searchplugin/7.16 (none none; android 4.4.2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.ClientID(); got != tt.want {
				t.Errorf("Request.ClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_HasScreen(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    true,
		}, {
			name:    "",
			request: getReq(1),
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.HasScreen(); got != tt.want {
				t.Errorf("Request.HasScreen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Command(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "съешь еще этих мягких французских булок",
		}, {
			name:    "",
			request: getReq(1),
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Command(); got != tt.want {
				t.Errorf("Request.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_OriginalUtterance(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "съешь еще этих мягких французских булок",
		}, {
			name:    "",
			request: getReq(1),
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.OriginalUtterance(); got != tt.want {
				t.Errorf("Request.OriginalUtterance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Text(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "съешь еще этих мягких французских булок",
		}, {
			name:    "",
			request: getReq(1),
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Text(); got != tt.want {
				t.Errorf("Request.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Type(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "SimpleUtterance",
		}, {
			name:    "",
			request: getReq(1),
			want:    "ButtonPressed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Type(); got != tt.want {
				t.Errorf("Request.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_DangerousContext(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    false,
		}, {
			name:    "",
			request: getReq(1),
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.DangerousContext(); got != tt.want {
				t.Errorf("Request.DangerousContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Payload(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    nil,
			wantErr: true,
		}, {
			name:    "",
			request: getReq(1),
			want:    map[string]interface{}{"msg": "ok"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			got, err := req.Payload()
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Payload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRequest_PayloadString(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
		wantErr bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "",
			wantErr: true,
		}, {
			name:    "",
			request: getReq(2),
			want:    "msg",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			got, err := req.PayloadString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.PayloadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRequest_Tokens(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    []string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    []string{"съешь", "еще", "этих", "мягких", "французских", "булок"},
		}, {
			name:    "",
			request: getReq(1),
			want:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			got := req.Tokens()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRequest_IsNewSession(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    true,
		}, {
			name:    "",
			request: getReq(1),
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.IsNewSession(); got != tt.want {
				t.Errorf("Request.IsNewSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_MessageID(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    int
	}{
		{
			name:    "",
			request: getReq(0),
			want:    0,
		}, {
			name:    "",
			request: getReq(1),
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.MessageID(); got != tt.want {
				t.Errorf("Request.MessageID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SessionID(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "e19e8eee-ae065e8-36e3f907-567a814b",
		}, {
			name:    "",
			request: getReq(1),
			want:    "eeb9fa7f-940e2502-1fbf9dfb-9448a1a9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.SessionID(); got != tt.want {
				t.Errorf("Request.SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SkillID(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "e03f8d5b-35ef-4d57-9450-b721ca17a6c3",
		}, {
			name:    "",
			request: getReq(1),
			want:    "e03f8d5b-35ef-4d57-9450-b721ca17a6c3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.SkillID(); got != tt.want {
				t.Errorf("Request.SkillID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_UserID(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0",
		}, {
			name:    "",
			request: getReq(1),
			want:    "03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.UserID(); got != tt.want {
				t.Errorf("Request.UserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Ver(t *testing.T) {
	tests := []struct {
		name    string
		request *alice.Request
		want    string
	}{
		{
			name:    "",
			request: getReq(0),
			want:    "1.0",
		}, {
			name:    "",
			request: getReq(1),
			want:    "1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			if got := req.Ver(); got != tt.want {
				t.Errorf("Request.Ver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_StateSession(t *testing.T) {
	tests := map[string]struct {
		request *alice.Request
		want    interface{}
	}{
		"when state is empty 0": {
			request: getReq(0),
			want:    nil,
		},
		"when state is empty 1":{
			request: getReq(1),
			want:    nil,
		},
		"when state is empty 2":{
			request: getReq(2),
			want:    nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req := tt.request
			got := req.StateSession("");
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Request.StateSession() = %v, want %v", got, tt.want)
			}

		})
	}
	t.Run("when state is struct", func(t *testing.T) {
		req := getReq(3)
		want := make(map[string]interface{})
		want["int_value"] = 42

		assert.Equal(t, 42.0, req.StateSession("int_value"))
		assert.Equal(t, "exampleString", req.StateSession("string_value"))
		assert.Equal(t, []interface{}{1.0,2.0,3.0,4.0}, req.StateSession("array_value"))
		assert.Equal(t, map[string]interface{}{"one":"one"}, req.StateSession("struct_value"))
		stateJson, err := req.StateSessionAsJson()
		if assert.NoError(t, err) {
			assert.Equal(t, `{"array_value":[1,2,3,4],"int_value":42,"string_value":"exampleString","struct_value":{"one":"one"}}`, stateJson)
		}
	})
}

func getReq(n int) *alice.Request {
	source := []string{`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"account_linking":{},"payments":{},"screen":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"command":"съешь еще этих мягких французских булок","nlu":{"entities":[],"tokens":["съешь","еще","этих","мягких","французских","булок"]},"original_utterance":"съешь еще этих мягких французских булок","type":"SimpleUtterance"},"session":{"message_id":0,"new":true,"session_id":"e19e8eee-ae065e8-36e3f907-567a814b","skill_id":"e03f8d5b-35ef-4d57-9450-b721ca17a6c3","user_id":"03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0"},"version":"1.0"}`,

		`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"account_linking":{},"payments":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"nlu":{"entities":[],"tokens":[]},"payload":{"msg":"ok"},"type":"ButtonPressed"},"session":{"message_id":1,"new":false,"session_id":"eeb9fa7f-940e2502-1fbf9dfb-9448a1a9","skill_id":"e03f8d5b-35ef-4d57-9450-b721ca17a6c3","user_id":"03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0"},"version":"1.0"}`,

		`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"account_linking":{},"payments":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"nlu":{"entities":[],"tokens":[]},"payload":"msg","type":"ButtonPressed"},"session":{"message_id":1,"new":false,"session_id":"eeb9fa7f-940e2502-1fbf9dfb-9448a1a9","skill_id":"e03f8d5b-35ef-4d57-9450-b721ca17a6c3","user_id":"03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0"},"version":"1.0"}`,

		`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"account_linking":{},"payments":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"nlu":{"entities":[],"tokens":[]},"payload":"msg","type":"ButtonPressed"},"session":{"message_id":1,"new":false,"session_id":"eeb9fa7f-940e2502-1fbf9dfb-9448a1a9","skill_id":"e03f8d5b-35ef-4d57-9450-b721ca17a6c3","user_id":"03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0"},"state":{"session":{"array_value":[1,2,3,4],"int_value":42,"string_value":"exampleString","struct_value":{"one":"one"}}},"version":"1.0"}`,
	}

	var req = new(alice.Request)
	json.Unmarshal([]byte(source[n]), req)
	return req
}
