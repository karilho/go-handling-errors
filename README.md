# Go Error Handling – Guia Explicativo

## 🧠 O que são errors?

De forma geral, **erros representam situações inesperadas ou indesejadas** que ocorrem durante a execução de um programa. Eles indicam que algo deu errado e que o fluxo normal precisa ser interrompido, tratado ou redirecionado.

### Em outras linguagens:
- Erros são comumente tratados usando estruturas como `try`, `catch`, `except`, `throw` e afins.
- Essas linguagens utilizam um **mecanismo de exceções** (exceptions) que permite interromper a execução de forma automática quando algo dá errado.

```python
try:
    resultado = 10 / 0
except ZeroDivisionError as e:
    print("Erro:", e)
```

```java
try {
    int resultado = 10 / 0;
} catch (ArithmeticException e) {
    System.out.println("Erro: " + e.getMessage());
}
```

---

### Em Go:
Em Go, o tratamento de erros é feito de forma diferente. Não existem exceções como em outras linguagens. Em vez disso, Go utiliza uma abordagem baseada em **valores retornados** para indicar erros.
A ideia é que as funções retornem um valor de erro junto com o resultado esperado. Se ocorrer um erro, esse valor de erro será diferente de `nil`, indicando que algo deu errado.

```go
type error interface {
    Error() string
}
```
Qualquer tipo que implemente esse método pode ser retornado como erro. Isso faz com que o tratamento de erros em Go seja **explícito**, **simples** e **controlado pelo desenvolvedor**, sem o uso de exceções.

```go
res, err := dividir(10, 0)
if err != nil {
    fmt.Println("Erro:", err)
    return
}
```

## 🔍 Como criar e tratar erros em Go

## Implementando a interface `error`

Para criar um erro personalizado, basta implementar a interface `error`:

```go
type ErroNegocio struct {
    Codigo int
    Mensagem string
}

func (e ErroNegocio) Error() string {
    return fmt.Sprintf("Código: %d, Mensagem: %s", e.Codigo, e.Mensagem)
}
```

Qualquer tipo que implemente o método `Error() string` automaticamente satisfaz essa interface. Isso significa que você pode criar tipos de erro personalizados com informações extras (como códigos, mensagens, contexto, etc).

### Exemplo de erro customizado

Ao invés de usar `errors.New`, você pode retornar seu próprio erro customizado para fornecer mais informações:

```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, ErroNegocio{
            Codigo:   400,
            Mensagem: "Divisão por zero não é permitida",
        }
    }
    return a / b, nil
}

res, err := dividir(10, 0)
if err != nil {
    fmt.Println("Erro:", err)
}
```

## 🔗 Panic? O que é?

Em Go, `panic` é uma forma de interromper a execução normal do programa quando ocorre um erro irrecuperável. É como um "pânico" que faz o programa parar imediatamente, geralmente usado em situações críticas onde não há como continuar, mas isso não funciona como tratamento de erros.

```go
func dividir(a, b int) int {
    if b == 0 {
        panic("Divisão por zero")
    }
    return a / b
}
```

O  panic é algo análogo a uma exceção em outras linguagens, mas em Go é desencorajado para erros comuns. O uso de `panic` deve ser reservado para situações onde o programa não pode continuar, como falhas de inicialização ou condições irrecuperáveis, como por exemplo falha ao abrir um arquivo essencial ou conexão com banco de dados.


## 🧬 Pilha de errors (Wrapping)

Com o `fmt.Errorf` e `%w`, é possível **encadear erros** (wrap):

```go
errOrig := errors.New("erro ao abrir arquivo")
err := fmt.Errorf("erro de leitura: %w", errOrig)
```

A pilha de erros permite rastrear a origem de falhas sem perder o contexto intermediário. Ferramentas como `errors.Is` e `errors.As` ajudam a "navegar" por essa cadeia de erros.


## 🧩 O que é `errors.Is` e `errors.As`

### ✅ `errors.Is`
Verifica se um erro é ou **envolve** um erro específico:

```go
if errors.Is(err, ErroNegocio) {
    fmt.Println("Erro de permissão")
}
```

Útil quando usamos `fmt.Errorf("...: %w", err)` para adicionar contexto e queremos verificar se o erro original ainda está presente.

### 🧱 `errors.As`
Permite extrair um **tipo específico** de erro:

```go
var e ErroNegocio
if errors.As(err, &e) {
    fmt.Println(e.Codigo, e.Mensagem)
}
```

Isso é poderoso quando usamos structs personalizadas como erros.
