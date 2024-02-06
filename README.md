# tictactoe

This is a tic-tac-toe engine that can be played via the command line. It uses the minimax algorithm and always searches to the max depth before making its move. So, the engine will play perfectly but be inneficient.

# why

Some questions you might have:
- [why make this?](#why-make-this)
- [why leave it so inneficient](#why-leave-it-so-inneficient)
- [why choose Go](#why-choose-go)

## why make this
I created this as a stepping stone to creating a chess engine. The general idea is very close to a chess engine:
1. Represent the game state
2. Have a method to apply a move on the game state
3. Generate available moves from a game state
4. Evaluate a move list for the "best" move

## why leave it so inneficient
I do not want to improve this engine for two reasons:
1. I hope that its current form is simple enough for others to understand and possibly encourage some of them to make something.
2. I would like to create a chess engine and would prefer to dedicate my time improving upon it, rather than this.

## why choose Go
I enjoy using the language; its simple and easy to understand. I like that it is opinionated and often provides only "one way" to do something.

Further, it offers a great developer experience. Everything is cross platform for free. It provides:
- build system
- package manager
- dependency manager
- unit test framework
all to the developer. And all of those things are standardized by the language so there's one way to do them. Once you become familiar with managing these systems for a project, everything is transferrable to any other Go project.

Lastly, Go provides all of the above while still being relatively fast and compliling to native executables. I think its a fantastic experience when you don't absolutely need to squeeze all the performance out of a system. This was difficult for me to give in to because I am traditionally a C/C++ programmer who always thought that the performance shouldn't take a back seat (I'm looking at you, Java from the 00's). But now I've come to realize that performance doesn't matter if your choice of tools prevents you from building the thing in the first place.