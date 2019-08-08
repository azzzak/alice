package alice_test

import (
	"net/http"

	"github.com/azzzak/alice"
)

func Example() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("привет")
		}
		return resp.Text(req.OriginalUtterance())
	})
}
