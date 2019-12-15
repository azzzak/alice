package alice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Kit структура для передачи данных в главный цикл.
type Kit struct {
	Req  *Request
	Resp *Response
	// Ctx позволяет получить сигнал об истечении периода ожидания ответа.
	Ctx context.Context

	c chan<- *Response
}

// Init получает входящий пакет и заготовку исходящего из данных запроса.
func (k Kit) Init() (*Request, *Response) {
	return k.Req, k.Resp
}

// Options структура с настройками.
type Options struct {
	AutoPong bool
	Timeout  time.Duration
	Debug    bool
}

// AutoPong автоматический ответ на технические сообщения ping, проверяющие работоспособность навыка. По умолчанию включено.
func AutoPong(b bool) func(*Options) {
	return func(opts *Options) {
		opts.AutoPong = b
	}
}

// Timeout таймаут обработки запроса в миллисекундах. По истечении запрос перестает обрабатываться и навык отправляет ошибку. Значение по умолчанию 3000 мс — официальное время ожидания ответа навыка.
func Timeout(t int) func(*Options) {
	return func(opts *Options) {
		if t < 1 {
			log.Fatalln("Timeout must be positive integer")
		}
		opts.Timeout = time.Duration(t)
	}
}

// Debug показывает в консоле содержимое входящих и исходящих пакетов. По умолчанию отключено.
func Debug(b bool) func(*Options) {
	return func(opts *Options) {
		opts.Debug = b
	}
}

// Stream канал, передающий данные в основной цикл.
type Stream <-chan Kit

// Handler сигнатура функции, передаваемой методу Loop().
type Handler func(k Kit) *Response

// Loop отвечает за работу главного цикла.
func (updates Stream) Loop(f Handler) {
	for kit := range updates {
		go func(k Kit) {
			k.c <- f(k)
			close(k.c)
		}(kit)
	}
}

// ListenForWebhook регистрирует обработчик входящих пакетов.
func ListenForWebhook(hook string, opts ...func(*Options)) Stream {
	conf := Options{
		AutoPong: true,
		Timeout:  3000,
		Debug:    false,
	}
	for _, opt := range opts {
		opt(&conf)
	}

	stream := make(chan Kit, 1)
	http.HandleFunc(hook, webhook(conf, stream))

	return stream
}

func webhook(conf Options, stream chan<- Kit) http.HandlerFunc {
	reqPool := sync.Pool{
		New: func() interface{} {
			return new(Request)
		},
	}

	respPool := sync.Pool{
		New: func() interface{} {
			return new(Response)
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		ctx, cancel := context.WithTimeout(r.Context(), conf.Timeout*time.Millisecond)
		defer cancel()

		if conf.Debug {
			requestDump, err := httputil.DumpRequest(r, true)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(requestDump))
		}

		req := reqPool.Get().(*Request)
		defer reqPool.Put(req)

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(req.clean()); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := respPool.Get().(*Response)
		resp.clean().prepareResponse(req)
		defer respPool.Put(resp)

		if conf.AutoPong {
			if req.Type() == SimpleUtterance && req.Text() == "ping" {
				if md, err := json.Marshal(resp.Text("pong")); err == nil {
					w.Header().Set("Content-Type", "application/json")
					w.Write(md)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		req.Bearer = r.Header.Get("Authorization")

		back := make(chan *Response)
		stream <- Kit{
			Req:  req,
			Resp: resp,
			Ctx:  ctx,

			c: back,
		}

		var response *Response
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			w.WriteHeader(http.StatusInternalServerError)
			return
		case response = <-back:
		}

		writer := io.Writer(w)

		if conf.Debug {
			var buf bytes.Buffer
			writer = io.MultiWriter(w, &buf)
			defer func() {
				fmt.Printf("\n%s\n\n", buf.String())
			}()
		}

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		if err := encoder.Encode(&response); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// Plural помогает согласовать слово с числительным.
func Plural(n int, singular, plural1, plural2 string) string {
	switch n % 100 {
	case 11, 12, 13, 14:
		return plural2
	}
	switch n % 10 {
	case 0, 5, 6, 7, 8, 9:
		return plural2
	case 1:
		return singular
	default:
		return plural1
	}
}
