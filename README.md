# Character Animation with Ebiten

This project demonstrates how to draw and animate a character using the Ebiten library.

## Project Structure

- `main.go`: The main program file containing initialization and game logic.
- `assets/images/runner.png`: Character animation image asset.

## Running the Project

1. Run the project:
   ```sh
   go run main.go
   ```

## Character Animation Details

- The character image size is 256x96, with each character frame being 32x32.
- The image contains three rows of character frames:
  - The first row has 5 frames.
  - The second row has 8 frames.
  - The third row has 4 frames.
- In this example, we only use the frames from the second row.

## Main Program Logic

- `Update` method: Updates the game state.
- `Draw` method: Draws the character animation.
- `Layout` method: Sets the game window size.

## References

- [Ebiten Official Documentation](https://ebitengine.org/en/examples/animation.html)
