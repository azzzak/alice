package alice

import (
	"encoding/json"
)

const (
	// NENameType именованная сущность с ФИО.
	NENameType = "YANDEX.FIO"

	// NELocationType именованная сущность с топонимом.
	NELocationType = "YANDEX.GEO"

	// NEDataTimeType именованная сущность с датой и временем.
	NEDataTimeType = "YANDEX.DATETIME"

	// NENumberType именованная сущность с числом.
	NENumberType = "YANDEX.NUMBER"
)

//Entity структура прототипа именованной сущности в запросе.
type Entity struct {
	Tokens struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"tokens"`
	Type  string           `json:"type"`
	Value *json.RawMessage `json:"value"`
}

// NEName структура типа NENameType.
type NEName struct {
	Start, End     int
	FirstName      string `json:"first_name,omitempty"`
	PatronymicName string `json:"patronymic_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
}

// NELocation структура типа NELocationType.
type NELocation struct {
	Start, End  int
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	Airport     string `json:"airport,omitempty"`
}

// NEDateTime структура типа NEDataTimeType.
type NEDateTime struct {
	Start, End        int
	Year              int  `json:"year,omitempty"`
	YearIsRelative    bool `json:"year_is_relative,omitempty"`
	Month             int  `json:"month,omitempty"`
	MonthIsRelative   bool `json:"month_is_relative,omitempty"`
	Day               int  `json:"day,omitempty"`
	DayIsRelative     bool `json:"day_is_relative,omitempty"`
	Hour              int  `json:"hour,omitempty"`
	HourIsRelative    bool `json:"hour_is_relative,omitempty"`
	Minute            int  `json:"minute,omitempty"`
	MinuterIsRelative bool `json:"minute_is_relative,omitempty"`
}

// NENumber структура типа NENumberType.
type NENumber float32

func (NEName) netype()     {}
func (NELocation) netype() {}
func (NEDateTime) netype() {}
func (NENumber) netype()   {}

type netype interface {
	netype()
}

// Entities возвращает необработанные именованные сущности из запроса.
func (req *Request) Entities() (Entities, error) {
	var entities Entities
	entities, err := unmarshalEntities(req.Request.NLU.Entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// Names возвращает готовый к использованию массив именованных сущностей с ФИО.
func (e Entities) Names() []NEName {
	var a []NEName
	for _, v := range e[NENameType] {
		d := v.Value.(*NEName)
		d.Start, d.End = v.Start, v.End
		a = append(a, *d)
	}
	return a
}

// Locations возвращает готовый к использованию массив именованных сущностей с топонимами.
func (e Entities) Locations() []NELocation {
	var a []NELocation
	for _, v := range e[NELocationType] {
		d := v.Value.(*NELocation)
		d.Start, d.End = v.Start, v.End
		a = append(a, *d)
	}
	return a
}

// DatesTimes возвращает готовый к использованию массив именованных сущностей с датами и временем.
func (e Entities) DatesTimes() []NEDateTime {
	var a []NEDateTime
	for _, v := range e[NEDataTimeType] {
		d := v.Value.(*NEDateTime)
		d.Start, d.End = v.Start, v.End
		a = append(a, *d)
	}
	return a
}

// NumberWrapper обертка для чисел из именованных сущностей.
type NumberWrapper struct {
	Start, End int
	Value      float32
}

// Numbers возвращает готовый к использованию массив именованных сущностей с числами.
func (e Entities) Numbers() []NumberWrapper {
	var a []NumberWrapper
	for _, v := range e[NENumberType] {
		n := v.Value.(*NENumber)
		d := NumberWrapper{
			Start: v.Start,
			End:   v.End,
			Value: float32(*n),
		}
		a = append(a, d)
	}
	return a
}

type newrapper struct {
	Start, End int
	Value      netype
}

// Entities контейнер для передачи необработанных именованных сущностей.
type Entities map[string][]newrapper

func unmarshalEntities(e []Entity) (Entities, error) {
	m := make(Entities)
	for _, v := range e {
		h := holder(v.Type)
		if err := json.Unmarshal(*v.Value, h); err != nil {
			return nil, err
		}
		wrapper := newrapper{
			Start: v.Tokens.Start,
			End:   v.Tokens.End,
			Value: h,
		}
		m[v.Type] = append(m[v.Type], wrapper)
	}
	return m, nil
}

func holder(t string) netype {
	switch t {
	case NENameType:
		return new(NEName)
	case NELocationType:
		return new(NELocation)
	case NEDataTimeType:
		return new(NEDateTime)
	case NENumberType:
		return new(NENumber)
	}
	return nil
}
