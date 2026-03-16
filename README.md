# password_generator

Go library for generating cryptographically secure random passwords, with a built-in strength evaluator. Also includes a ready-to-use CLI.

## Installation

```bash
go get github.com/jcoppede11/password_generator
```

## Usage as a library

```go
import "github.com/jcoppede11/password_generator"

password, err := generator.Generate(generator.Options{
    Length:     16,
    UseUpper:   true,
    UseLower:   true,
    UseNumbers: true,
    UseSymbols: true,
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(password)
fmt.Println(generator.StrengthScore(password))
```

### Options

| Field        | Type   | Description                              | Default |
|--------------|--------|------------------------------------------|---------|
| `Length`     | `int`  | Password length (maximum: 128)           | —       |
| `UseUpper`   | `bool` | Include uppercase letters (A–Z)          | `false` |
| `UseLower`   | `bool` | Include lowercase letters (a–z)          | `false` |
| `UseNumbers` | `bool` | Include digits (0–9)                     | `false` |
| `UseSymbols` | `bool` | Include symbols (!@#$%^&*…)              | `false` |

### StrengthScore

Evaluates the strength of any password and returns one of these values:

| Result        | Criterion                        |
|---------------|----------------------------------|
| `Weak`        | score < 3                        |
| `Medium`      | score 3–4                        |
| `Strong`      | score 5                          |
| `Very strong` | score ≥ 6                        |

The score combines length (0–3 points) and character variety (0–4 points).

## Usage as a CLI

```bash
go run cmd/password-generator/main.go -length=16 -uppercase=true -lowercase=true -numbers=true -symbols=true
```

### Flags

| Flag          | Default | Description               |
|---------------|---------|---------------------------|
| `-length`     | `12`    | Password length            |
| `-uppercase`  | `true`  | Include uppercase letters  |
| `-lowercase`  | `true`  | Include lowercase letters  |
| `-numbers`    | `true`  | Include numbers            |
| `-symbols`    | `true`  | Include symbols            |

### Example output

```
Generated password: aK3$gTz9@Lpw1Xm#
Strength:           Very strong
```

## Project structure

```
.
├── generator.go          # Public API
├── generator_test.go
├── go.mod
├── internal/
│   └── app/
│       └── app.go        # CLI orchestration
└── cmd/
    └── password-generator/
        └── main.go       # Entry point
```

## Tests

```bash
go test -v
```

## Requirements

Go 1.24.2 or higher.

## License

[MIT](LICENSE)
