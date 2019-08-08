# Руководство

Разберем простейший пример:

```Go
package main

import "github.com/azzzak/alice"

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

Функция `ListenForWebhook` регистрирует обработчик, который принимает входящие пакеты от Алисы. Функции передается путь и настройки.

В простейшем случае путь будет `/`, но выбор осмысленного имени позволит запускать на одном домене несколько навыков.

```Go
alice.ListenForWebhook("/hook")
```

Настройки задаются с помощью соответствующих функций. Например, чтобы изменить время ожидания ответа (по умолчанию это 3000 мс) есть функция `alice.Timeout`.

```Go
alice.ListenForWebhook("/hook", alice.Timeout(2500))
```

Следующая строка запускает горутину в которой работает сервер.

```Go
go http.ListenAndServe(":3000", nil)
```

Вместо http-сервера можно запустить https-сервер.

```Go
go http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
```

Однако разумно «спрятать» приложение за веб-сервером (вроде nginx), оставив последнему работу с сертификатами.

Функция `ListenForWebhook` возвращает канал, который получает входящие пакеты. За манипуляции с ними отвечает метод `Loop`, принимающий функцию-обработчик.

```Go
func(k alice.Kit) *alice.Response{...}
```

Разработчик наполняет тело функции кодом, который будет выполняться для каждой новой порции данных. Результат работы функции станет ответом Алисе.

Метод `Init()` получает входящий пакет и заготовку исходящего из данных запроса.

```Go
req, resp := k.Init()
```

С помощью различных методов из структуры запроса можно получить текст реплики пользователя, узнать даннные, извлеченные из запроса и пр. В примере используется два таких метода: один помогает выяснить было ли сообщение первым в сессии, другой — получает реплику пользователя.

```Go
req.IsNewSession()
...
req.OriginalUtterance()
```

Структура ответа уже готова к использованию. Для конструирования ответа доступны методы, которые можно объединять в цепочки.

Полный список методов для работы с запросом и ответом есть в [документации](https://godoc.org/github.com/azzzak/alice).

После того, как ответ подготовлен, его нужно отправить Алисе. В примере мы видим два варианта. Если это первое сообщение в сессии пользователь получит приветствие.

```Go
return resp.Text("привет")
```

Во всех остальных случаях это будет повторение его собственной реплики.

```Go
return resp.Text(req.OriginalUtterance())
```