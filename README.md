## Линтер логов для Go

Этот репозиторий содержит набор правил для проверки лог‑сообщений на Go:

- **формат** сообщения (первая буква в нижнем регистре);
- **язык** сообщения (только английский);
- отсутствие **спецсимволов и эмодзи**;
- отсутствие **чувствительных данных** (ключевые слова, e‑mail, номера карт, IP, значения переменных `token/password/...`).

Правила можно использовать:

- как **отдельный анализатор** (`go run ./cmd ./...`);
- как **плагин для `golangci-lint`** (тип `goplugin`).

---

## Требования

- Go **1.22+**;
- `golangci-lint` **IDEs / CI** (для использования как плагин);
- для сборки плагина (`.so`) нужен **Linux / WSL / macOS**:
  - нужен установленный C‑компилятор (`gcc` / `clang`).

---

## Структура проекта

- `cmd/main.go` — CLI‑обёртка над анализатором (`singlechecker`);
- `rules/` — правила линтера:
  - `rules.go` — основной анализатор логов (`LogAnalyzer`);
  - `nonEnglish.go` — проверка, что в сообщении нет нелатинских букв;
  - `specialSymbols.go` — проверка на спецсимволы и эмодзи;
  - `sensitiveData.go` — проверка на наличие чувствительных данных;
  - `sensitiveIdentifier.go` — поиск чувствительных переменных (`token`, `password` и т.п.);
- `plugin/example.go` — entrypoint плагина для `golangci-lint`;
- `testdata/` — пример кода для проверки.

---

## Запуск как отдельного анализатора

Из корня репозитория:

```bash
go run ./cmd ./...
```

или с указанием конкретного пакета:

```bash
go run ./cmd ./testdata
```

Анализатор пройдётся по указанным пакетам и выведет все найденные нарушения.

---

## Сборка и использование как плагина `golangci-lint`

### 1. Сборка `.so` плагина

На Linux / WSL / macOS:

```bash
cd /path/to/golangci-lint-plugin

# при необходимости
sudo apt update && sudo apt install -y build-essential

go build -buildmode=plugin -o ./plugin/logrules.so ./plugin
```

В результате будет создан файл `plugin/logrules.so`.

### 2. Настройка `.golangci.yaml`

В проекте, где вы хотите использовать линтер, добавьте `.golangci.yaml`, например:

```yaml
version: "2"

linters:
  enable:
    - logrules

  settings:
    custom:
      logrules:
        type: goplugin
        path: ./plugin/logrules.so
        description: "Checks log format, language and sensitive data"
        original-url: github.com/snickers78/golangci-lint-plugin
```

- `logrules` — имя кастомного линтера в `golangci-lint`;
- `path` — путь до собранного `.so`.

### 3. Запуск `golangci-lint`

Из корня проекта с `.golangci.yaml`:

```bash
golangci-lint run ./...
```

Если всё настроено корректно, вы увидите репорты с идентификатором вашего кастомного линтера (например, `logrules`) с сообщениями о нарушениях.

---

## Как работают правила

- **Регистр первой буквы**  
  Сообщение считается некорректным, если первая буква в сообщении — заглавная. Правильный формат:

  ```go
  log.Println("starting server on port 8080")
  ```

- **Язык (только английский)**  
  Если в строке встречаются буквенные символы не из латинского алфавита (например, кириллица), будет репорт:

  > log message must contain only english letters

- **Спецсимволы и эмодзи**  
  Разрешены буквы, цифры, пробелы. Эмодзи и прочие символы считаются нарушением.

- **Чувствительные данные**  
  линтер ищет:
  - ключевые слова (`password`, `secret`, `apikey`, `sessionid`, и др.);
  - e‑mail адреса;
  - «похожие на» номера банковских карт;
  - IP‑адреса;
  - идентификаторы переменных с подстроками `token`, `password`, `secret`, `apikey` и т.п. в аргументах логов.

  Примеры:

  ```go
  log.Println("token validated")          // допустимо
  log.Println("token: " + token)         // будет подсвечено как чувствительная информация
  slog.Info("contact me at a@b.com")     // чувствительные данные (email)
  ```

---

## Разработка и тестирование

Запуск юнит‑тестов для правил:

```bash
go test ./rules/...
```

Рекомендуется запускать:

- `go test ./...` перед коммитом.

