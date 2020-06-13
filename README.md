# game_of_life
A console version of Conway's Game of Life in honour of John Horton Conway.
I've created the logic for the rules below:
For each 'step':
Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with two or three live neighbours lives on to the next generation.
Any live cell with more than three live neighbours dies, as if by overpopulation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

WHAT NEXT?
Next I will use a gui library to make the frontend a bit prettier.
