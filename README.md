# Problema do Jantar dos Filósofos

O Problema do Jantar dos Filósofos é um clássico exemplo de um problema de concorrência em sistemas de computação. Ele foi formulado por Edsger Dijkstra em 1965 e ilustra a questão de sincronização entre múltiplos processos que precisam compartilhar recursos limitados.

## Cenário

O cenário envolve:

- **Cinco filósofos** que se sentam em volta de uma mesa circular.
- **Cinco garfos**, um entre cada par de filósofos.
- Cada filósofo alterna entre pensar e comer. Para comer, um filósofo precisa de **dois garfos**, um de cada lado.

## Regras

- **Comer**: Um filósofo pode comer apenas se tiver ambos os garfos à sua disposição (um à sua esquerda e um à sua direita).
- **Pensar**: Um filósofo pode pensar quando não está comendo.
- **Garfos**: Os garfos não têm identidade; um garfo é indistinguível do outro. Um filósofo deve pegar dois garfos ao mesmo tempo para comer.
- **Garçom**: Em algumas variações do problema, pode-se introduzir um garçom que controla o acesso aos garfos, permitindo que um filósofo pegue dois garfos de cada vez.

## Problemas a serem abordados

- **Deadlock**: É possível que todos os filósofos peguem um garfo e, em seguida, fiquem esperando indefinidamente pelo segundo garfo, resultando em um estado de deadlock.
- **Starvation**: Mesmo que não haja deadlock, um filósofo pode ficar esperando por um longo tempo, enquanto outros filósofos comem, resultando em fome.
- **Concorrência**: O problema demonstra a necessidade de uma coordenação cuidadosa entre processos concorrentes para evitar conflitos e garantir que todos os filósofos possam comer.

## Objetivos

O objetivo do problema é encontrar uma solução que permita que todos os filósofos comam sem entrar em deadlock ou starvation, gerenciando o acesso aos garfos de maneira eficiente. Isso pode ser alcançado através de várias abordagens, incluindo:

- Uso de semáforos.
- Algoritmos de mutex.
- Estruturas de dados que controlam o acesso de forma a evitar conflitos.

## Importância

O problema é amplamente utilizado para ensinar conceitos de programação concorrente e sistemas operacionais, pois oferece uma forma de explorar as dificuldades de gerenciar recursos compartilhados e a necessidade de sincronização em sistemas multi-threaded.

