# game_of_life
A console version of Conway's Game of Life in honour of John Horton Conway.
I've created the logic for the rules below:

For each 'step':
- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

Main menu:
- press n for new generated world
- press q to quit
- press g for a lower res frontend GUI (rather than the standard console output)
- press enter for one cycle, hold enter down to run dozens a second.

When you are in the lower res GUI bit:
- press n to refresh the grid 
- press q to go back to the Main menu.

What i plan to do next:
Add a higher res GUI - i've managed to bungle in the canvas from 'termui' but i'm sure there's probably a better way.
