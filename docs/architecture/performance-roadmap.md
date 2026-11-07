# Roadmap de Otimizações de Desempenho

Este documento registra as decisões de arquitetura e o planejamento de otimizações de performance planejadas para o motor de animação do MotionGo.

---

## 1. Contexto e Oportunidades Identificadas

Atualmente:
1. **Recálculo Total da Árvore de Nós**: A cada frame, `BaseNode.Update()` recalcula as matrizes locais e globais de **todos** os nós da cena, mesmo que a maioria esteja estática.
2. **Varredura Linear de Tracks**: `Engine.Tick()` percorre todas as tracks de animação registradas no slice `Tracks []func(frame int)`, inclusive aquelas que já terminaram ou ainda não começaram.
3. **Reflexão de Tipos por Amostra**: A função genérica `Linear[T]()` executa um `type switch` a cada interpolação para descobrir o tipo do dado animado.

---

## 2. Estratégias de Otimização

### A. Invalidação Seletiva por Flags de Modificação (*Dirty Flags*)
Em vez de recalcular todas as matrizes a cada frame:
- Adicionar referências de pai e flags `isLocalDirty` e `isBranchDirty` em `BaseNode`.
- Criar método `MarkDirty()` que sinaliza os nós ancestrais até encontrar um ramo já marcado.
- No método `Update(parentGlobal)`:
  - Pular subárvores inteiras cujo ramo esteja limpo e cujo pai não tenha mudado de posição.
  - Recalcular a matriz local somente se a transformação do próprio nó mudou.
  - Recalcular a matriz global apenas se a local mudou ou se o nó pai foi atualizado.

### B. Agendamento de Animações Ativas
Substituir a lista opaca de funções por uma estrutura com metadados de ciclo de vida (`StartFrame`, `EndFrame`):
- **Tracks Futuras**: Fila ordenada pelo `StartFrame`.
- **Tracks Ativas**: Pool de tracks sendo executadas no frame atual.
- **Tracks Concluídas**: Removidas do processamento após executarem o frame final (`EndFrame`).

Com isso, o custo por frame no `Tick()` dependerá apenas da quantidade de animações em execução naquele instante, não do total acumulado ao longo de um vídeo/timeline.

### C. Especialização de Interpoladores e Redução de Alocações
- Resolver o `type switch` de tipos genéricos na etapa de configuração do `Builder[T]` (uma única vez), armazenando diretamente a função de cálculo otimizada para `float32`, `Vector3`, `Quaternion` ou `Color`.
- Reutilizar slices e buffers para evitar alocações de heap dentro do render loop.

---

## 3. Fases de Implementação

| Fase | Foco | Ganho Esperado |
| :--- | :--- | :--- |
| **Fase 1** | *Dirty Flags* no Grafo de Cena | Redução drástica do uso de CPU em cenas com muitos objetos estáticos ou semi-estáticos. |
| **Fase 2** | Agendador de Tracks Ativas | Custo constante de CPU independente da duração total da cena e do número de tracks passadas. |
| **Fase 3** | Otimização de Interpoladores & Benchmarks | Eliminação de alocações e overhead de reflexão/interfaces dentro do loop de desenho. |
