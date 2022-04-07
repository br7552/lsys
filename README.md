# lsys
CLI and API for drawing fractals with ASCII characters using L-systems

This is based on a description of a program called **lsys** in the book *Computational Beauty of Nature*.

### L-Systems
An L-system consists of an alphabet of symbols, an initial axiom string, and a set of production rules for expanding symbols into strings.

```
Alphabet: F, -, +
Axiom: F++F++F
Rules: F=F-F++F-F
Iterations:
0: F++F++F
1: F-F++F-F++F-F++F-F++F-F++F-F
2: F++F++F-F++F++F++F++F++F-F++F++F++F++F++F-F++F++F++F++F++F-F++F++F++F++F++F-F++F++F++F++F++F-F++F++F
```

The resulting strings have a self-similar structure and can be used to draw fractal forms.

### Drawing
When included in an input L-system, the following symbols have a special meaning for drawing:
Symbol | Meaning
--- | ---
F | Draw a line in the direction the cursor is facing
G | Advance the cursor foward without drawing a line
\- | Rotate the cursor left
\+ | Rotate the cursor right
\[ | Store the cursor's position and direction on the stack
] | Restore the cursor's position and direction from the stack

```
Axiom: F++F++G
Rules: F=F-F++F-F+[G] G=[G]G
Depth (iterations): 6
                       /|\
                  ____/ | ____|
                  |     |     |
                  |     \     |
                  |    / \    |
              \   |   \   /   \   \
             / \ /   / \ / \   \ / \
        ____/   /___/___|_______/   ____|
        |\     /        |        \     /|
        | ____/         |         ____/ |
        | |   |         |         |   | |
        | |   |         |         |   | \
       /  |   |         |         |   |  \
      /   |____         |         |____   /
       \ /     \        |        /     \ /
        /       \       \       /       |
       /|        \     / \     /        |\
  ____/ |         \___|   /___/         | ____|
  |     |         |   |\ /|   |         |     |
  |     |         |   | | |   |         |     |
  |    /|         |   | | |   |         |\    |
  |   / |         |____ | |___\         | \   \
 /   / \|        / \   \|/   / \        |/ \   \
/____   \________   /___|___/   ________/   ____/
 \   \ /|        \ /   /|\   \ /        |\ /   /
  \   \ |         ____| | ____\         | /   |
  |    \|         |   | | |   |         |/    |
  |     |         |   | | |   |         |     |
  |     |         |   |/ \|   |         |     |
  |____ |         /____   /___\         | /____
       \|        /     \ /     \        |/
        |       /       \       \       /
       / \     /        |        \     / \
      /   ____|         |         ____|   /
       \  |   |         |         |   |  /
        \ |   |         |         |   | |
        | |   |         |         |   | |
        | /____         |         /____ |
        |/     \        |        /     \|
        |____   /___________/____   /____
             \ / \   \ / \ /   / \ /
              \   \   /   \   |   \
                  |    \ /    |
                  |     \     |
                  |     |     |
                  |____ | /____
                       \|/
                        |
```



### API Endpoint
Make a POST request to /v1/fractals with a JSON object containing the following fields:
Name | JSON Type | Description | Required
--- | --- | --- | ---
axiom | string | The L-system's axiom | *
rules | object | The L-system's rules | *
depth | number | Number of iterations |
angle | number | Number of degrees to rotate the cursor |
startAngle | number | The starting angle for the cursor |
Step | number | The length of drawn lines |
Width | number | Canvas width |
Height | number | Canvas height |
