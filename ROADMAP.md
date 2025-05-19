# CraftGame's Roadmap

- **[Pre Alpha](#pre-alpha)**
    - 0.0.0b
    - 0.0.0a

## Pre Alpha

### 0.0.0b

* Added mob: Bob
    * Moves and jumps around randomly
    * 100 of them spawn when the world loads
    * Not saved in `world.dat`

### 0.0.0a

* Added air blocks (not fully working yet)
* Added grass blocks
    * Only show up at y = 42
* Added stone blocks
    * Show up from y = 41 and below
* Added 2 light levels for blocks: bright and dark
    * A block is bright if there’s nothing above it
    * If there’s a block on top, it turns dark
* Added a world
    * Size: 256 x 64 x 256
    * Only the bottom 42 layers (256 x 42 x 256) are filled with blocks
    * Blocks can’t go outside the world
    * Players *can* go outside, but they’ll keep falling down
* Added chunks
    * Each one is 16 x 16 blocks
    * The chunk closest to the player loads first
* Added a player entity
    * No model yet
    * Height is 1.62
    * Spawns at a random spot
* Added controls
    * W to move forward
    * S to move backward
    * A to move left
    * D to move right
    * Space to jump
    * R to respawn randomly
    * Enter to save the world
    * Left-click to break a block
    * Right-click to place a block
* Added a blinking white outline on the block you're looking at
    * Helps you see where you're about to place or break a block
* World auto-saves when you quit the game
