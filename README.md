# Go Error Handling

## 🧠 O que são errors?

De forma geral, **erros representam situações inesperadas ou indesejadas** que ocorrem durante a execução de um programa. Eles indicam que algo deu errado e que o fluxo normal precisa ser interrompido, tratado ou redirecionado.

### Em outras linguagens:
- Erros são comumente tratados usando estruturas como `try`, `catch`, `except`, `throw` e afins.
- Essas linguagens utilizam um **mecanismo de exceções** (exceptions) que permite interromper a execução de forma automática quando algo dá errado.
- Por exemplo, em Python:

```python
try:
    resultado = 10 / 0
except ZeroDivisionError as e:
    print("Erro:", e)
```

---

### Em Go:
Em Go, erros **não são exceções**. Eles são **valores retornados** pelas funções:

```go
res, err := dividir(10, 0)
if err != nil {
    fmt.Println("Erro:", err)
    return
}
```

Erros em Go são representados pela interface embutida:

```go
type error interface {
    Error() string
}
```

Qualquer tipo que implemente esse método pode ser retornado como erro. Isso faz com que o tratamento de erros em Go seja **explícito**, **simples** e **controlado pelo desenvolvedor**, sem o uso de exceções.

---

## 🔍 Como Go lida com erros em relação a outras linguagens?

Diferente de linguagens como:

- **Java** (com `try-catch`) ou
- **Python** (com `try-except`),

Go utiliza uma abordagem **explícita e simples**:

```go
res, err := algumaFuncao()
if err != nil {
    // tratar erro
}
```

Essa simplicidade torna o código mais previsível e fácil de seguir. O tratamento de erros em Go evita "ocultar" falhas e incentiva a tratá-las imediatamente.

---

## 🧩 O que é `errors.Is` e `errors.As`

### ✅ `errors.Is`
Verifica se um erro é ou **envolve** um erro específico:

```go
if errors.Is(err, ErrSemPermissao) {
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

---

## 🧬 Pilha de errors (Wrapping)

Com o `fmt.Errorf` e `%w`, é possível **encadear erros** (wrap):

```go
errOrig := errors.New("erro ao abrir arquivo")
err := fmt.Errorf("erro de leitura: %w", errOrig)
```

A pilha de erros permite rastrear a origem de falhas sem perder o contexto intermediário. Ferramentas como `errors.Is` e `errors.As` ajudam a "navegar" por essa cadeia de erros.

---

## ✅ Boas práticas de tratamento de erros em Go

1. **Sempre verifique erros:**
   ```go
   if err != nil {
       return err
   }
   ```

2. **Adicione contexto com `fmt.Errorf` e `%w`**
   ```go
   return fmt.Errorf("falha no repositório: %w", err)
   ```

3. **Crie erros customizados quando necessário**
   ```go
   type ErroNegocio struct {
       Codigo int
       Mensagem string
   }
   ```

4. **Evite panics**: use `panic` apenas em casos irrecuperáveis, como falha em inicialização.
