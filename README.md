# Meteor-Busters 

## Overview

**Meteor-Busters** is a 2D action game built with **Go** and the **Ebiten** library. In this game, players must shoot meteors while navigating through obstacles. The gameplay focuses on dynamic meteor spawning, player-controlled shooting, and a scoring system with increasing difficulty as the game progresses.

**Game Screenshot**

![Screenshot 2025-02-18 123840](https://github.com/user-attachments/assets/8d3cb8be-e731-45b2-a62b-6a05a18189b6)

![Screenshot 2025-02-18 123856](https://github.com/user-attachments/assets/cbee51dc-7c65-4874-af30-db18b9a48e0b)

## Features

- **Player Control:** 
  - The player character can rotate and move in different directions.
  - Bullets are dynamically generated based on player input and fired in the specified direction.
  
- **Random Meteor Generation:** 
  - Meteors spawn randomly, and their frequency increases as the game progresses, adding challenge.
  
- **Scoring System:** 
  - Track and display the player's score, with a high score system.
  
- **Timers and Difficulty Progression:** 
  - Timers control meteor spawn rates and game speed, ensuring a progressively challenging gameplay experience.

## Game Assets

- **Sprites & Fonts:** 
  - Utilizes the `embed` package to load and manage game assets, including player and meteor sprites, laser projectiles, and custom fonts.
  
## Technologies Used

- **Go** for game logic and mechanics
- **Ebiten** library for 2D game rendering and interaction
- **Embed** package for efficient game asset handling

## Controls

- **Arrow keys or WASD** → Move the player
- **Spacebar** → Shoot bullets

## Installation

1. Navigate into the project folder:
   ```bash
   cd meteor-busters
   ```
2. Install Go dependencies :
  ```bash
   go mod tidy
  ```
3. Run the game:
   ```bash
   go run main.go
