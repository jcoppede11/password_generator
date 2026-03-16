# password_generator

Librería Go para generar contraseñas aleatorias criptográficamente seguras, con evaluador de fortaleza integrado. También incluye una CLI lista para usar.

## Instalación

```bash
go get github.com/jcoppede11/password_generator
```

## Uso como librería

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

| Campo        | Tipo   | Descripción                              | Default |
|--------------|--------|------------------------------------------|---------|
| `Length`     | `int`  | Longitud de la contraseña (máximo: 128)  | —       |
| `UseUpper`   | `bool` | Incluir letras mayúsculas (A–Z)          | `false` |
| `UseLower`   | `bool` | Incluir letras minúsculas (a–z)          | `false` |
| `UseNumbers` | `bool` | Incluir dígitos (0–9)                    | `false` |
| `UseSymbols` | `bool` | Incluir símbolos (!@#$%^&*…)             | `false` |

### StrengthScore

Evalúa la fortaleza de cualquier contraseña y devuelve uno de estos valores:

| Resultado    | Criterio                         |
|--------------|----------------------------------|
| `Débil`      | score < 3                        |
| `Media`      | score 3–4                        |
| `Fuerte`     | score 5                          |
| `Muy fuerte` | score ≥ 6                        |

El score combina la longitud (0–3 puntos) y la variedad de caracteres (0–4 puntos).

## Uso como CLI

```bash
go run cmd/password-generator/main.go -length=16 -uppercase=true -lowercase=true -numbers=true -symbols=true
```

### Flags

| Flag          | Default | Descripción              |
|---------------|---------|--------------------------|
| `-length`     | `12`    | Longitud de la contraseña |
| `-uppercase`  | `true`  | Incluir mayúsculas        |
| `-lowercase`  | `true`  | Incluir minúsculas        |
| `-numbers`    | `true`  | Incluir números           |
| `-symbols`    | `true`  | Incluir símbolos          |

### Ejemplo de salida

```
Contraseña generada: aK3$gTz9@Lpw1Xm#
Fortaleza:           Muy fuerte
```

## Estructura del proyecto

```
.
├── generator.go          # API pública
├── generator_test.go
├── go.mod
├── internal/
│   └── app/
│       └── app.go        # Orquestación CLI
└── cmd/
    └── password-generator/
        └── main.go       # Punto de entrada
```

## Tests

```bash
go test -v
```

## Requisitos

Go 1.24.2 o superior.

## Licencia

[MIT](LICENSE)
