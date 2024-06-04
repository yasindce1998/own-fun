# Headline  
DIY English Cricket Word Game (CLI)

## Introduction

I remember playing a fun cricket-inspired word game in school. It involved guessing words based on the first letter, similar to the classic cricket sport we all know and love. 

Now that I'm learning Go, I'm excited to bring this game to life as a text-based (CLI) application. 

## Gameplay

Just like cricket with overs, this game has "overs" where you throw letters instead of balls. For example, a 2-over game means guessing words that start with 12 different letters.

Throwing an "a" means you need to guess a word starting with "a," like "aeroplane" (score: 9 characters). You keep guessing words until the "over" (set of letters) is complete.

This is a great way to test your vocabulary and reflexes!

## Implementation
### Single Team or Two Team
    - If Two Player is selected, Your opponent team will be human.
    - If Single Team is selected, Your opponent team will be system.
### For Single Team
    - Let's see the implementation for Single Team (You vs. System)
    - Before tossing the coin you have to choose heads or tails
    - If the tossed coin has heads, they can choose bat or bowl
    - For this let's say single team(You) chose to bowl
This is How the Game play begins.
Now Let's see how the game looks like in CLI
```
$ ec --start
Welcome to English cricket
Choose your team:
1. Single Team (You vs. System)
2. Two Team (You vs. Human)
Input: 1
You selected Single Team
Choose heads or tails for coin toss
1. Heads
2. Tails
Input: 1
You selected Heads
Tossing...
Tossed coin is Heads
Choose bat or bowl
1. Bat
2. Bowl
Input: 1
You chose to Bat
Select How many overs do you want
1. One
2. Two
3. Three
Input: 1
You selected One over so 6 balls so 6 letters
Now system will start the bowling
System: D
You: Determination
Score: 13
System: E
You: Elephant
Score: 8
This game will go on until the six ball finish.
You chose single team so your oppoent will be system.




**#englishcricket #golang #programming #fun**

