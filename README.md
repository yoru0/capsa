# Capsa

A terminal-based implementation of classic **Capsa** card game.

---

## Game Overview
**Capsa** is a multiplayer card game where players take turns playing valid card combinations to beat the previous combo. The player who runs out of cards first wins, while the last player to play all their cards is considered the loser.

---

## Rules

### Suits (Ascending Order)
- Diamonds ♦
- Clubs ♣
- Hearts ♥
- Spades ♠

### Ranks (Ascending Order)
- 3 < 4 < 5 < 6 < 7 < 8 < 9 < 10 < J < Q < K < A < 2

### Opening Rule
- At the start of the game, all players must reveal and discard any cards with the rank **3**.
- The **first player** to play is the one who has **3♠ (Three of Spades)**.

### Turn Flow
1. Players take turns in a clockwise order.
2. On your turn, you may either:
   - Play a valid combo that **beats** the last combo, or
   - **Skip** your turn.
3. A round ends when **3 players skip consecutively**, resetting the combo.
4. The player who plays after a round reset can play **any valid combo**.

### Winning Rule
- When a player runs out of cards (wins), they are removed from the game.
- The next player in order may play **any valid card or combo** — they are **not required** to follow the winning player's last combo.
- The game continues until only one player remains.

---

## Combos

- **Single**: One card
- **Pair**: Two cards of the same rank
- **Triple**: Three cards of the same rank
- **Straight**: Five cards in sequence (mixed suits)
- **Flush**: Five cards of the same suit
- **Full House**: Three of a kind + a pair
- **Four of a Kind**: Four cards of the same rank + any single card
- **Straight Flush**: Straight of same suit

---

## Technical

- Language: **Go (Golang)**
- Terminal-based interface
- Local multiplayer for up to 4 players
- In-memory game state

---

## Getting Started

```bash
git clone https://github.com/yoru0/capsa.git
cd capsa
go run main.go