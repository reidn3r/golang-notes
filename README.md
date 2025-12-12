
# Golang Notes

Este repositório reúne exemplos e anotações dos principais conceitos e recursos da linguagem Go (Golang) que estudei até o momento. O objetivo é servir como referência rápida e prática sobre o aprendizado, útil para revisar tópicos fundamentais da linguagem.

---

## Índice

1. [Hello World](#hello-world)
2. [Funções Anônimas em Go](#funções-anônimas-em-go)
3. [Arrays e Slices](#arrays-e-slices)
   - [Diferença entre Array e Slice](#diferença-entre-array-e-slice)
   - [O que é um Slice por baixo dos panos?](#o-que-é-um-slice-por-baixo-dos-panos)
   - [Consumo de memória de um slice vazio](#consumo-de-memória-de-um-slice-vazio)
   - [Como funciona o slice internamente?](#como-funciona-o-slice-internamente)
4. [Maps](#maps-em-go-detalhes-e-funcionamento)
5. [Slice (Reforço e Prática)](#slice-reforço-e-prática)
6. [Defer e Panic](#defer-e-panic)
7. [Estruturas de Dados](#estruturas-de-dados)
8. [Design Patterns](#design-patterns)

---

## 1. Hello World

Primeiro contato com a linguagem Go, estrutura básica de um programa e execução do clássico "Hello, World!".

---

## 2. Funções Anônimas em Go

Funções anônimas são funções sem nome atribuídas a uma variável ou executadas diretamente. Em Go, funções podem ser atribuídas a variáveis, passadas como argumento e retornadas de outras funções.

**Exemplo:**

```go
f := func(x, y int) int {
    return x + y
}
fmt.Println(f(2, 3)) // Saída: 5
```

Funções anônimas são úteis para lógica rápida, callbacks e closures (funções que "capturam" variáveis do escopo externo). Elas podem ser executadas imediatamente:

```go
result := func(msg string) string {
    return "Olá, " + msg
}("mundo")
fmt.Println(result) // Saída: Olá, mundo
```

---

## 3. Arrays e Slices

- Diferenças entre arrays e slices em Go.
- Como declarar, inicializar e manipular ambos.
- Vantagens dos slices e exemplos de uso prático.

### Diferença entre Array e Slice

**Array:**
- Estrutura de tamanho fixo, definido na declaração (ex: `[5]int`).
- O tamanho faz parte do tipo, ou seja, `[3]int` e `[5]int` são tipos diferentes.
- Armazenado de forma contígua na memória.

**Slice:**
- Estrutura dinâmica, "vista" sobre um array.
- Pode crescer ou diminuir de tamanho.
- Muito mais flexível para uso cotidiano.

#### O que é um Slice por baixo dos panos?
Um slice é uma estrutura composta por três campos:
- **Ponteiro** para o início de um array subjacente
- **Tamanho** (length): quantidade de elementos visíveis
- **Capacidade** (capacity): quantidade máxima de elementos até o final do array subjacente

```go
type SliceHeader struct {
    Data uintptr // ponteiro para o array
    Len  int     // tamanho
    Cap  int     // capacidade
}
```
Esses campos podem ser acessados via o pacote `reflect`.

#### Consumo de memória de um slice vazio
Um slice vazio (`var s []int` ou `s := []int{}`) ocupa apenas o espaço dos três campos do header (ponteiro, tamanho, capacidade). Em arquiteturas 64 bits, isso normalmente equivale a 24 bytes (3 x 8 bytes).

O array subjacente só é alocado quando elementos são inseridos (ex: via `append`).

#### Como funciona o slice internamente?
Quando você faz um slice de um array, o slice aponta para o array original, mas com um intervalo de índices definido. Operações como `append` podem criar um novo array subjacente se a capacidade for excedida. Isso torna slices eficientes para manipulação de subconjuntos de dados sem cópia desnecessária, mas é importante entender que alterações via slice podem afetar o array original (quando dentro da capacidade).

**Resumo:**
- Slices são leves, flexíveis e eficientes para manipulação de coleções.
- Arrays são úteis quando o tamanho é fixo e conhecido em tempo de compilação.

---

## 4. Maps em Go: Detalhes e Funcionamento

O `map` em Go é uma estrutura de dados associativa (dicionário/hash table) que armazena pares chave-valor. É altamente eficiente para buscas, inserções e remoções.

### Como criar um map
```go
mapper := make(map[string]int) // map vazio
mapper["foo"] = 42            // inserção
```
Também é possível inicializar com valores:
```go
mapper := map[string]int{"a": 1, "b": 2}
```

### Como funciona em memória
Um map em Go é implementado como uma hash table. Internamente, ele mantém buckets (baldes) para distribuir as chaves de acordo com o hash. O map é um ponteiro para uma estrutura interna, e seu header ocupa 8 bytes em arquiteturas 64 bits. Um map vazio (`make(map[string]int)`) ocupa apenas o header, sem buckets alocados, até que um elemento seja inserido.

### Consumo de memória de um map vazio
Um map vazio ocupa cerca de 8 bytes (header) em 64 bits. Buckets e arrays internos só são alocados quando o primeiro elemento é inserido.

### Iteração e acesso
**Verificar existência:**
```go
val, exists := mapper["key"]
if exists {
    // chave encontrada, val contém o valor
} else {
    // chave não existe, val é o valor zero do tipo (ex: 0 para int, "" para string)
}
```

**Apenas verificar existência:**
```go
_, exists := mapper["key"]
```

**Iterar sobre todos os elementos:**
```go
for key, value := range mapper {
    fmt.Println(key, value)
}
```
Ou apenas valores:
```go
for _, value := range mapper {
    fmt.Println(value)
}
```

### O que é retornado se a chave não existe?
Ao acessar uma chave inexistente, o map retorna o valor zero do tipo do valor. Por exemplo, para `map[string]int`, retorna 0; para `map[string]string`, retorna "".

### Remover elementos
```go
delete(mapper, "key")
```

### Resumo de propriedades
- Busca, inserção e remoção são operações O(1) na média.
- Não há garantia de ordem na iteração.
- Maps não podem ser comparados (exceto com `nil`).
- Maps são referências: ao passar para funções, não há cópia dos dados.

---


## 5. Slice (Reforço e Prática)

Exploração mais aprofundada dos slices, operações avançadas e exemplos práticos.

### Criação de Slices
Você pode criar slices de diferentes formas:

- A partir de um array:
    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    s := arr[1:4] // [2 3 4]
    ```
- Usando `make` para definir tamanho e capacidade:
    ```go
    s := make([]int, 0, 10) // slice vazio, capacidade 10
    ```
- Literal:
    ```go
    s := []int{1, 2, 3}
    ```

### Gerenciamento em Memória
Um slice é apenas um header (ponteiro, tamanho, capacidade). O array subjacente só é alocado quando necessário (ex: ao usar `append`).

### Consumo de Memória de um Slice Vazio
Um slice vazio (`var s []int` ou `s := []int{}`) ocupa apenas o header (24 bytes em 64 bits). O array subjacente só é alocado quando elementos são inseridos.

### Iteração e Acesso
Você pode iterar sobre slices de várias formas:

- Índice e valor:
    ```go
    for idx, val := range slice {
            fmt.Printf("index: %d, valor: %d\n", idx, val)
    }
    ```
- Apenas índice:
    ```go
    for i := range slice {
            fmt.Println(i)
    }
    ```
- Apenas valor:
    ```go
    for _, v := range slice {
            fmt.Println(v)
    }
    ```

### Acesso Fora do Limite
Tentar acessar um índice fora do limite do slice causa panic em tempo de execução.

### Valores Padrão
Ao criar um slice de inteiros, todos os valores são inicializados com zero.

### Resumo de Propriedades
- Slices são dinâmicos e flexíveis.
- O header do slice é leve, mas o array subjacente pode crescer conforme necessário.
- Slices compartilham o array subjacente, então alterações em um slice podem afetar outros slices derivados do mesmo array.

---

## 6. Defer e Panic

- Uso do `defer` para adiar a execução de funções.
- Tratamento de erros com `panic` e recuperação com `recover`.
- Exemplos práticos de controle de fluxo e limpeza de recursos.

---

## 7. Estruturas de Dados

- Implementação de algumas estruturas clássicas:
  - **List**: listas encadeadas e suas operações básicas.
  - **Set**: conjuntos para garantir unicidade de elementos.
  - **Tree**: árvores e exemplos de navegação.

---

## 8. Design Patterns

- Exemplos de padrões de projeto implementados em Go:
  - **Builder**
  - **Singleton**
  - **Strategy**

---