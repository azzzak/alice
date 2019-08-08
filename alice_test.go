package alice_test

import (
	"context"
	"testing"

	"github.com/azzzak/alice"
	"github.com/stretchr/testify/assert"
)

func TestAutoPong(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args bool
		want alice.Options
	}{
		{
			name: "",
			args: true,
			want: alice.Options{
				AutoPong: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := alice.Options{}
			opt := alice.AutoPong(tt.args)
			opt(&conf)
			assert.Equal(t, tt.want, conf)
		})
	}
}

func TestTimeout(t *testing.T) {
	tests := []struct {
		name string
		args int
		want alice.Options
	}{
		{
			name: "",
			args: 2500,
			want: alice.Options{
				Timeout: 2500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := alice.Options{}
			opt := alice.Timeout(tt.args)
			opt(&conf)
			assert.Equal(t, tt.want, conf)
		})
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want alice.Options
	}{
		{
			name: "",
			args: true,
			want: alice.Options{
				Debug: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := alice.Options{}
			opt := alice.Debug(tt.args)
			opt(&conf)
			assert.Equal(t, tt.want, conf)
		})
	}
}

func TestKit_Init(t *testing.T) {
	type fields struct {
		Req  *alice.Request
		Resp *alice.Response
		Ctx  context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   *alice.Request
		want1  *alice.Response
	}{
		{
			name: "",
			fields: fields{
				Req:  getReq(0),
				Resp: getResp(0),
			},
			want:  getReq(0),
			want1: getResp(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := alice.Kit{
				Req:  tt.fields.Req,
				Resp: tt.fields.Resp,
			}
			got, got1 := k.Init()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestPlural(t *testing.T) {
	type args struct {
		n        int
		singular string
		plural1  string
		plural2  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				n:        0,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылок",
		}, {
			name: "",
			args: args{
				n:        1,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылка",
		}, {
			name: "",
			args: args{
				n:        2,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылки",
		}, {
			name: "",
			args: args{
				n:        5,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылок",
		}, {
			name: "",
			args: args{
				n:        15,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылок",
		}, {
			name: "",
			args: args{
				n:        21,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылка",
		}, {
			name: "",
			args: args{
				n:        105,
				singular: "бутылка",
				plural1:  "бутылки",
				plural2:  "бутылок",
			},
			want: "бутылок",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alice.Plural(tt.args.n, tt.args.singular, tt.args.plural1, tt.args.plural2); got != tt.want {
				t.Errorf("PluralForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
