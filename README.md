# Blackjack-go
El código implementa un juego de blackjack en Go en la consola de comandos. En el juego, el jugador juega contra el dealer. El objetivo del juego es tener una mano con un valor más cercano a 21 sin pasarse de 21.

El juego comienza con la baraja de cartas mezclada y se reparten dos cartas a cada uno, una a cada jugador. El jugador puede elegir seguir pidiendo cartas adicionales o quedarse con las cartas que tiene en la mano actualmente.

Una vez que el jugador decide quedarse, le toca al dealer jugar su mano. El dealer debe seguir ciertas reglas al jugar, tomando cartas hasta que su mano tenga un valor de al menos 17. Luego, se compara el valor de la mano del jugador con la del dealer para determinar el ganador. Si el jugador tiene una mano con un valor mayor a la del dealer, entonces gana.

El código implementa las reglas del juego, incluyendo el seguimiento del valor de la mano, el reparto de cartas, la decisión de seguir pidiendo cartas adicionales o quedarse, y la determinación del ganador.

## Commands
### test
    go run main.go
### build
    go build
### run (after build)
    ./blackjack
