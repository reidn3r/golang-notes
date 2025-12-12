# Golang Notes

Este repositório reúne exemplos e anotações dos principais conceitos e recursos da linguagem Go (Golang) que estudei até o momento. O objetivo é servir como referência rápida e prática sobre o aprendendizado, útil para revisar tópicos fundamentais da linguagem.

## Conteúdo Abordado

1. **Hello World**
   - Primeiro contato com a linguagem Go, estrutura básica de um programa e execução do clássico "Hello, World!".

### Funções Anônimas em Go

**Funções anônimas:** Funções sem nome atribuídas a uma variável ou executadas diretamente. Em Go, funções podem ser atribuídas a variáveis, passadas como argumento e retornadas de outras funções.

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

2. **Arrays e Slices**
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

3. **Maps**
   - Estrutura de dados do tipo dicionário (map).
   - Criação, inserção, remoção e iteração de elementos.
   - Casos de uso comuns.

4. **Slice (Reforço e Prática)**
   - Exploração mais aprofundada dos slices.
   - Operações avançadas, como append, copy e slicing.

5. **Defer e Panic**
   - Uso do `defer` para adiar a execução de funções.
   - Tratamento de erros com `panic` e recuperação com `recover`.
   - Exemplos práticos de controle de fluxo e limpeza de recursos.

6. **Estruturas de Dados**
   - Implementação de algumas estruturas clássicas:
     - **List**: listas encadeadas e suas operações básicas.
     - **Set**: conjuntos para garantir unicidade de elementos.
     - **Tree**: árvores e exemplos de navegação.

7. **Design Patterns**
   - Exemplos de padrões de projeto implementados em Go:
     - **Builder**
     - **Singleton**
     - **Strategy**

---