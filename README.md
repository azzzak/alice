# alice [![GoDoc](https://godoc.org/github.com/azzzak/alice?status.svg)](https://godoc.org/github.com/azzzak/alice) [![Go Report Card](https://goreportcard.com/badge/github.com/azzzak/alice)](https://goreportcard.com/report/github.com/azzzak/alice)

Библиотека для создания навыков, расширяющих функциональность голосового помощника [Алиса](https://alice.yandex.ru). Упрощает разработку навыков, оставляя возможность тонкой настройки. Преимущества библиотеки:

- поддержка [связки аккаунтов](https://yandex.ru/dev/dialogs/alice/doc/auth/about-account-linking-docpage/)
- объединение методов в цепочки при конструировании ответа
- вспомогательные методы для оживления диалога
- автоответ на служебные ping-пакеты

## Установка или обновление

`go get -u github.com/azzzak/alice`

## Пример

Простейший навык — говорит "привет", после чего повторяет каждую реплику пользователя.

```Go
package main

import (
  "net/http"

  "github.com/azzzak/alice"
)

func main() {
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
```

Детальный разбор примера смотрите в [руководстве](manual/README.md).

## Цепочки методов

Позволяют конструировать ответ со всеми возможностями, которые предлагает Алиса: кнопки, картинки и звуки.

**Пример.** Ответ с текстом и TTS с паузой и звуком из библиотеки Алисы:

```Go
resp.Text("творог").
  TTS("твор+ог").
  Pause(3).
  Sound(sounds.Harp1)
```

**Пример.** Ответ со случайно выбранной строкой и двумя кнопками:

```Go
resp.RandomText("привет", "алоха").
  Button("хай", "", false).
  Button("отстань", "", false)
```

**Пример.** При любом _num_ ответ остается согласованным:

```Go
resp.Text(fmt.Sprintf("%d %s пива %s на столе", num,
  alice.Plural(num, "бутылка", "бутылки", "бутылок"),
  alice.Plural(num, "стояла", "стояли", "стояло")).
  Sound(sounds.ThingsGlass1)
```

## Навыки на базе библиотеки

[![Дневник здоровья](images/health.png)](https://dialogs.yandex.ru/store/skills/dd5bb5ec-dnevnik-zdorov-ya) \
[**Дневник здоровья**](https://dialogs.yandex.ru/store/skills/dd5bb5ec-dnevnik-zdorov-ya)

[![Полезный Ждун](images/zhdun.jpeg)](https://dialogs.yandex.ru/store/skills/16ff4b52-poleznyj-zhdu) \
[**Полезный Ждун**](https://dialogs.yandex.ru/store/skills/16ff4b52-poleznyj-zhdu)

Присылайте свои навыки на azzzak@yandex.ru
