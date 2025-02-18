Meteor-Busters 
Overview
Meteor-Busters is a 2D action game built with Go and the Ebiten library. In this game, players must shoot meteors while navigating through obstacles. The gameplay focuses on dynamic meteor spawning, player-controlled shooting, and a scoring system with increasing difficulty as the game progresses.


Features
Player Control:
The player character can rotate and move in different directions.
Bullets are dynamically generated based on player input and fired in the specified direction.
Random Meteor Generation:
Meteors spawn randomly, and their frequency increases as the game progresses, adding challenge.
Scoring System:
Track and display the player's score, with a high score system.
Timers and Difficulty Progression:
Timers control meteor spawn rates and game speed, ensuring a progressively challenging gameplay experience.
Game Assets
Sprites & Fonts:
Utilizes the embed package to load and manage game assets, including player and meteor sprites, laser projectiles, and custom fonts.
Technologies Used
Go for game logic and mechanics
Ebiten library for 2D game rendering and interaction
Embed package for efficient game asset handling
Installation
Clone the repository:

bash
Copy
Edit
git clone https://github.com/your-username/meteor-busters.git
Navigate into the project folder:

bash
Copy
Edit
cd meteor-busters
Install Go dependencies (if any):

bash
Copy
Edit
go mod tidy
Run the game:

bash
Copy
Edit
go run main.go
Controls
Arrow keys or WASD to move the player.
Spacebar to shoot bullets.
Screenshots

Contributing
If you have any suggestions or improvements, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.
