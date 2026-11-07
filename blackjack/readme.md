# Goal

Write a programme to test Go knowledge of structs, types, methods, input handling through a blackjack game.

# Pin down the rules and the scope

- the goal is to have a higher hand than the dealer, without going bust
- hand > 21, automatically bust
- hand must be >= 10
- K, Q, J all values 10
- A can value 1 or 11, depending on the context, as long as it favoures the players

# Write a fake console session before any code

- Begin new game? [y/n]
- Hand: [K ❤️, Q ♦️] (value: 20)
- Dealer's hand: [A ♣️, ??]
- Dealer's hand: [A ♣️, 4♠️]
- Draw? [y/n]
- Blackjack!
- Bust!
- Dealer wins! 

# Find the “things” and the “actions”

The nouns in your rules become types and the verbs become functions or methods

## Card

## Deck

- shuffle()

## Player

## Dealer

## Game

# Describe the flow as steps

1. User is asked to start game
2. The used gets 2 cards
3. The dealer gets 2 cards, but only one is shown to the user
4. The user is asked if stay or draw one more card
5. Step 4 repeats until user decides to stay, hits Blackjack or gets busted
6. If the user stays, the dealer will reveal it's hidden card and pull until win, bust or at least draws
7. Regardless of the outcome, the player is asked to play again

# Separate the game logic from the console

Good developers keep game rules (pure logic: hand values, who wins) separate from input/output (reading keys, printing). Then the rules are easy to test without typing anything. In Go, that could be a blackjack package for the logic and main.go for the console.

# List the edge cases and the riskiest part

1. A, A = 11
2. A, A, A = 12
3. A, A, A, A = 13
4. A, A, A, A, K = 14

# Cut the work into small milestones, each one runnable
1. CMD interaction
2. Generate Deck
3. Generate Hand
4. Generate Player
5. Generate Dealer
6. Hand-Player
7. Hand-Dealer
8. Game

# Notes

- insert a 500ms delay when the dealers acts to simulate thinking

# Decide what “done” means
First simple version.

# Future

- Include balance logic.
- Include double logic.