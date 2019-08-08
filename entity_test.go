package alice

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Entities(t *testing.T) {
	test0 := map[string][]newrapper{
		"YANDEX.DATETIME": {
			{
				Start: 4,
				End:   5,
				Value: &NEDateTime{
					Start:             0,
					End:               0,
					Year:              0,
					YearIsRelative:    false,
					MonthIsRelative:   false,
					Month:             0,
					Day:               -1,
					DayIsRelative:     true,
					Hour:              0,
					HourIsRelative:    false,
					Minute:            0,
					MinuterIsRelative: false,
				},
			},
		},
		"YANDEX.FIO": []newrapper{
			{
				Start: 2,
				End:   5,
				Value: &NEName{
					Start:          0,
					End:            0,
					FirstName:      "валентин",
					PatronymicName: "петрович",
					LastName:       "вчера",
				},
			}, {
				Start: 5,
				End:   7,
				Value: &NEName{
					Start:          0,
					End:            0,
					FirstName:      "сергей",
					PatronymicName: "",
					LastName:       "иванов",
				},
			},
		},
		"YANDEX.GEO": []newrapper{
			{
				Start: 8,
				End:   10,
				Value: &NELocation{
					Start:       0,
					End:         0,
					Country:     "",
					City:        "саранск",
					Street:      "",
					HouseNumber: "",
					Airport:     "",
				},
			},
		},
		"YANDEX.NUMBER": []newrapper{
			{
				Start: 10,
				End:   11,
				Value: func(i NENumber) *NENumber { return &i }(NENumber(5)),
			},
		},
	}

	tests := []struct {
		name    string
		request *Request
		want    Entities
		wantErr bool
	}{
		{
			name:    "",
			request: getReq(0),
			want:    Entities(test0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.request
			got, err := req.Entities()
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Entities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEntities_Names(t *testing.T) {
	req := getReq(0)
	entities, _ := req.Entities()

	names := []NEName{
		{
			Start:          2,
			End:            5,
			FirstName:      "валентин",
			PatronymicName: "петрович",
			LastName:       "вчера",
		}, {
			Start:     5,
			End:       7,
			FirstName: "сергей",
			LastName:  "иванов",
		},
	}

	tests := []struct {
		name string
		e    Entities
		want []NEName
	}{
		{
			name: "",
			e:    entities,
			want: names,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Names()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEntities_Locations(t *testing.T) {
	req := getReq(0)
	entities, _ := req.Entities()

	locations := []NELocation{
		{
			Start:   8,
			End:     10,
			Country: "",
			City:    "саранск",
		},
	}

	tests := []struct {
		name string
		e    Entities
		want []NELocation
	}{
		{
			name: "",
			e:    entities,
			want: locations,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Locations()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEntities_DatesTimes(t *testing.T) {
	req := getReq(0)
	entities, _ := req.Entities()

	dt := []NEDateTime{
		{
			Start:         4,
			End:           5,
			Day:           -1,
			DayIsRelative: true,
		},
	}

	tests := []struct {
		name string
		e    Entities
		want []NEDateTime
	}{
		{
			name: "",
			e:    entities,
			want: dt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.DatesTimes()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEntities_Numbers(t *testing.T) {
	req := getReq(0)
	entities, _ := req.Entities()

	num := []NumberWrapper{
		{
			Start: 10,
			End:   11,
			Value: 5,
		},
	}

	tests := []struct {
		name string
		e    Entities
		want []NumberWrapper
	}{
		{
			name: "",
			e:    entities,
			want: num,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Numbers()
			assert.Equal(t, tt.want, got)
		})
	}
}

func getReq(n int) *Request {
	source := []string{`{"meta":{"client_id":"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)","interfaces":{"account_linking":{},"payments":{},"screen":{}},"locale":"ru-RU","timezone":"UTC"},"request":{"command":"по словам валентина петровича, вчера сергей иванов отвез в саранск пять мешков дерьма","nlu":{"entities":[{"tokens":{"end":5,"start":2},"type":"YANDEX.FIO","value":{"first_name":"валентин","last_name":"вчера","patronymic_name":"петрович"}},{"tokens":{"end":5,"start":4},"type":"YANDEX.DATETIME","value":{"day":-1,"day_is_relative":true}},{"tokens":{"end":7,"start":5},"type":"YANDEX.FIO","value":{"first_name":"сергей","last_name":"иванов"}},{"tokens":{"end":10,"start":8},"type":"YANDEX.GEO","value":{"city":"саранск"}},{"tokens":{"end":11,"start":10},"type":"YANDEX.NUMBER","value":5}],"tokens":["по","словам","валентина","петровича","вчера","сергей","иванов","отвез","в","саранск","5","мешков","д"]},"original_utterance":"по словам валентина петровича, вчера сергей иванов отвез в саранск пять мешков дерьма","type":"SimpleUtterance"},"session":{"message_id":7,"new":false,"session_id":"ed0cba5b-e68516f8-b73d36cc-ce696176","skill_id":"e03f8d5b-35ef-4d57-9450-b721ca17a6c3","user_id":"03B1D487CAA1C7EBF80A195491B78ACA0AC9934CDFB12A29D063A8329BC42BF0"},"version":"1.0"}`}

	var req = new(Request)
	json.Unmarshal([]byte(source[n]), req)
	return req
}
