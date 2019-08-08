package alice_test

import (
	"encoding/json"
	"testing"

	"github.com/azzzak/alice"
	"github.com/stretchr/testify/assert"
)

func TestNewImageButton(t *testing.T) {
	test0 := alice.ImageButton{
		Title:   "кнопка",
		URL:     "https://google.com",
		Payload: nil,
	}

	test1 := alice.ImageButton{
		Title:   "кнопка2",
		URL:     "https://google.com",
		Payload: `{"msg":"ok"}`,
	}

	type args struct {
		title   string
		url     string
		payload []interface{}
	}
	tests := []struct {
		name string
		args args
		want alice.ImageButton
	}{
		{
			name: "",
			args: args{
				title:   "кнопка",
				url:     "https://google.com",
				payload: nil,
			},
			want: test0,
		}, {
			name: "",
			args: args{
				title:   "кнопка2",
				url:     "https://google.com",
				payload: []interface{}{`{"msg":"ok"}`},
			},
			want: test1,
		}, {
			name: "",
			args: args{
				title:   "кнопка2",
				url:     "https://google.com",
				payload: []interface{}{`{"msg":"ok"}`, `{"msg":"nope"}`},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, alice.NewImageButton(tt.args.title, tt.args.url, tt.args.payload...))
		})
	}
}

func TestResponse_BigImage(t *testing.T) {
	source_0 := `{"response":{"text":"","card":{"image_id":"123","title":"название","description":"описание","type":"BigImage"},"end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`
	var test0 = new(alice.Response)
	json.Unmarshal([]byte(source_0), test0)

	source_1 := `{"response":{"text":"","card":{"image_id":"456","title":"название2","description":"описание2","button":{"url":"https://google.com"},"type":"BigImage"},"end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`
	var test1 = new(alice.Response)
	json.Unmarshal([]byte(source_1), test1)

	type args struct {
		id     string
		title  string
		desc   string
		button []alice.ImageButton
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
				id:     "123",
				title:  "название",
				desc:   "описание",
				button: nil,
			},
			want: test0,
		}, {
			name:     "",
			response: getResp(0),
			args: args{
				id:    "456",
				title: "название2",
				desc:  "описание2",
				button: []alice.ImageButton{
					{
						Title: "",
						URL:   "https://google.com",
					},
				},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			got := resp.BigImage(tt.args.id, tt.args.title, tt.args.desc, tt.args.button...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGallery_Add(t *testing.T) {
	test0 := &alice.List{
		Images: []alice.Image{
			{
				ImageID:     "123",
				Title:       "название",
				Description: "описание",
				Button:      nil,
			},
		},
	}

	test1 := &alice.List{
		Images: []alice.Image{
			{
				ImageID:     "456",
				Title:       "название2",
				Description: "описание2",
				Button: &alice.ImageButton{
					Title:   "кнопка",
					URL:     "https://google.com",
					Payload: "data",
				},
			},
		},
	}

	type args struct {
		id     string
		title  string
		desc   string
		button []alice.ImageButton
	}
	tests := []struct {
		name string
		args args
		want *alice.List
	}{
		{
			name: "",
			args: args{
				id:     "123",
				title:  "название",
				desc:   "описание",
				button: nil,
			},
			want: test0,
		}, {
			name: "",
			args: args{
				id:    "456",
				title: "название2",
				desc:  "описание2",
				button: []alice.ImageButton{
					{
						Title:   "кнопка",
						URL:     "https://google.com",
						Payload: "data",
					},
				},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &alice.List{}
			got := g.Add(tt.args.id, tt.args.title, tt.args.desc, tt.args.button...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGallery_AddImages(t *testing.T) {
	test0 := &alice.List{
		Images: []alice.Image{
			{
				ImageID:     "123",
				Title:       "название",
				Description: "описание",
				Button:      nil,
			}, {
				ImageID:     "456",
				Title:       "название2",
				Description: "описание2",
				Button: &alice.ImageButton{
					Title:   "кнопка",
					URL:     "https://google.com",
					Payload: "data",
				},
			},
		},
	}

	type args struct {
		images []alice.Image
	}
	tests := []struct {
		name string
		args args
		want *alice.List
	}{
		{
			name: "",
			args: args{
				images: []alice.Image{{
					ImageID:     "123",
					Title:       "название",
					Description: "описание",
					Button:      nil,
				}, {
					ImageID:     "456",
					Title:       "название2",
					Description: "описание2",
					Button: &alice.ImageButton{
						Title:   "кнопка",
						URL:     "https://google.com",
						Payload: "data",
					},
				}},
			},
			want: test0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &alice.List{}
			got := g.AddImages(tt.args.images...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_List(t *testing.T) {
	source := `{"response":{"text":"","card":{"type":"ItemsList","header":{"text":"заголовок"},"items":[{"image_id":"123","title":"название","description":"описание"},{"image_id":"456","title":"название2","description":"описание2","button":{"title":"кнопка","url":"https://google.com","payload":"data"}}],"footer":{"text":"подвал","button":{"title":"кнопка2","url":"https://ya.ru","payload":"msg"}}},"end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`
	var test0 = new(alice.Response)
	json.Unmarshal([]byte(source), test0)

	source1 := `{"response":{"text":"","card":{"type":"ItemsList","header":{"text":"заголовок"},"items":[{"image_id":"123"},{"image_id":"123"},{"image_id":"123"},{"image_id":"123"},{"image_id":"123"}],"footer":{"text":"подвал"}},"end_session":true},"session":{"message_id":0,"session_id":"","user_id":""},"version":"1.0"}`
	var test1 = new(alice.Response)
	json.Unmarshal([]byte(source1), test1)

	type args struct {
		header string
		footer string
		g      alice.List
		button []alice.ImageButton
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
				header: "заголовок",
				footer: "подвал",
				g: alice.List{
					Images: []alice.Image{{
						ImageID:     "123",
						Title:       "название",
						Description: "описание",
						Button:      nil,
					}, {
						ImageID:     "456",
						Title:       "название2",
						Description: "описание2",
						Button: &alice.ImageButton{
							Title:   "кнопка",
							URL:     "https://google.com",
							Payload: "data",
						},
					}},
				},
				button: []alice.ImageButton{
					{
						Title:   "кнопка2",
						URL:     "https://ya.ru",
						Payload: "msg",
					},
				},
			},
			want: test0,
		}, {
			name:     "",
			response: getResp(0),
			args: args{
				header: "заголовок",
				footer: "подвал",
				g: alice.List{
					Images: []alice.Image{{
						ImageID:     "123",
						Title:       "",
						Description: "",
					}, {
						ImageID:     "123",
						Title:       "",
						Description: "",
					}, {
						ImageID:     "123",
						Title:       "",
						Description: "",
					}, {
						ImageID:     "123",
						Title:       "",
						Description: "",
					}, {
						ImageID:     "123",
						Title:       "",
						Description: "",
					}, {
						ImageID:     "123",
						Title:       "",
						Description: "",
					},
					},
				},
			},
			want: test1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.response
			got := resp.List(tt.args.header, tt.args.footer, tt.args.g, tt.args.button...)
			assert.Equal(t, tt.want, got)
		})
	}
}
