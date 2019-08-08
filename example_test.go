package alice_test

import (
	"fmt"
	"net/http"

	"github.com/azzzak/alice"
)

func ExampleListenForWebhook() {
	updates := alice.ListenForWebhook("/hook", alice.Debug(true))
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		_, resp := k.Init()
		return resp.Text("ok")
	})
}

func ExampleRequest_Entities() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()

		entities, _ := req.Entities()
		for _, v := range entities.Names() {
			fmt.Printf("%+v\n", v)
		}
		for _, v := range entities.Locations() {
			fmt.Printf("%+v\n", v)
		}
		for _, v := range entities.DatesTimes() {
			fmt.Printf("%+v\n", v)
		}
		for _, v := range entities.Numbers() {
			fmt.Printf("%+v\n", v)
		}

		return resp.Text("ok")
	})
}

func ExampleRequest_Payload() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()

		// payload вида {"msg":"ok","n":42}
		payload, err := req.Payload()
		if err == nil {
			fmt.Println(payload["msg"].(string))
			fmt.Println(payload["n"].(int))
		}

		return resp.Text("ok")
	})
}

func ExampleRequest_PayloadString() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()

		payload, err := req.PayloadString()
		if err == nil {
			fmt.Println(payload)
		}

		return resp.Text("ok")
	})
}

func ExampleResponse_BigImage() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		_, resp := k.Init()

		return resp.BigImage("111111/00000000000000000000", "image", "desc")
	})
}

func ExampleResponse_BigImage_button() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		_, resp := k.Init()

		b := alice.NewImageButton("кнопка", "https://yandex.ru")
		return resp.BigImage("111111/00000000000000000000", "image", "desc", b)
	})
}

func ExampleResponse_List() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	updates.Loop(func(k alice.Kit) *alice.Response {
		_, resp := k.Init()

		var list alice.List
		list.Add("111111/00000000000000000000", "image1", "desc1")
		list.Add("222222/00000000000000000000", "image2", "desc2")
		list.Add("333333/00000000000000000000", "image3", "desc3")

		return resp.Text("список").List("список", "подвал", list)
	})
}

func ExamplePlural() {
	num := 5
	fmt.Printf("%d %s пива %s на столе", num,
		alice.Plural(num, "бутылка", "бутылки", "бутылок"),
		alice.Plural(num, "стояла", "стояли", "стояло"))
	// Output: 5 бутылок пива стояло на столе
}
