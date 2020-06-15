# game_of_life
A cellular automaton game engine that can handle 'Conway's Game of Life' in honour of John Horton Conway.
Can also do other different life rules via user input.
Makes use of the awesome 'termui' (https://github.com/gizak/termui) library for the graphics.

- 'Game of Life'
\
![Game of Life](https://raw.githubusercontent.com/tbal999/game_of_life/master/gui.gif)
- 'Day & Night'
\
![Day&Night](https://raw.githubusercontent.com/tbal999/game_of_life/master/blob.gif)

- (sourced from https://en.wikipedia.org/wiki/Life-like_cellular_automaton)
There are 218 = 262,144 possible Life-like rules, only a small fraction of which have been studied in any detail. In the descriptions below, all rules are specified in Golly/RLE format.

Notable Life-like rules:

- B1357/S1357	Replicator	Edward Fredkin's replicating automaton: every pattern is eventually replaced by multiple copies of itself.[2][3][4]
- B2/S	Seeds	All patterns are phoenixes, meaning that every live cell immediately dies, and many patterns lead to explosive chaotic growth. However, some engineered patterns with complex behavior are known.[2][5][6]
- B25/S4		This rule supports a small self-replicating pattern which, when combined with a small glider pattern, causes the glider to bounce back and forth in a pseudorandom walk.[4][7]
- B3/S012345678	Life without Death	Also known as Inkspot or Flakes. Cells that become alive never die. It combines chaotic growth with more structured ladder-like patterns that can be used to simulate arbitrary Boolean circuits.[2][4][8][9]
- B3/S23	Life	Highly complex behavior.[10][11] (Game of Life rules)
- B34/S34	34 Life	Was initially thought to be a stable alternative to Life, until computer simulation found that larger patterns tend to explode. Has many small oscillators and spaceships.[2][12][13]
- B35678/S5678	Diamoeba	Forms large diamonds with chaotically fluctuating boundaries. First studied by Dean Hickerson, who in 1993 offered a $50 prize to find a pattern that fills space with live cells; the prize was won in 1999 by David Bell.[2][4][14]
- B36/S125	2x2	If a pattern is composed of 2x2 blocks, it will continue to evolve in the same form; grouping these blocks into larger powers of two leads to the same behavior, but slower. Has complex oscillators of high periods as well as a small glider.[2][15]
- B36/S23	HighLife	Similar to Life but with a small self-replicating pattern.[2][4][16]
- B3678/S34678	Day & Night	Symmetric under on-off reversal. Has engineered patterns with highly complex behavior.[2][4][17]
- B368/S245	Morley	Named after Stephen Morley; also called Move. Supports very high-period and slow spaceships.[2][4][18]
- B4678/S35678	Anneal	Also called the twisted majority rule. Symmetric under on-off reversal. Approximates the curve-shortening flow on the boundaries between live and dead cells.[19][20][21]

-more info here http://www.mirekw.com/ca/rullex_life.html

What i plan to do next:
- Add a higher res GUI! I've managed to bungle in the canvas from termui but i'm sure there's probably a better way as termui is meant for command line GUI's.
- Refactor the code so it makes more sense
