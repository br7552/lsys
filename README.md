# lsys
CLI and API for drawing fractals with ASCII characters using L-systems

This is based on a description of a program called **lsys** in the book *Computational Beauty of Nature*.

[Demo](https://br7552.github.io/lsys/)

### L-Systems
An L-system consists of an alphabet of symbols, an initial axiom string, and a set of production rules for expanding symbols into strings.

```
Alphabet: F, -, +
Axiom: F++F++F
Rules: F->F-F++F-F
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
\- | Rotate the cursor left
\+ | Rotate the cursor right
\[ | Store the cursor's position and direction on the stack
] | Restore the cursor's position and direction from the stack

```
Axiom: F++F++F
Rules: F->F-F++F-F
Depth (iterations): 3 
                                     /\
                                 ___/ ___/
                                  \      /
                             \    \     \    \
                             /\   /      \   /\
                         ___/ ___/       ___/ ___/
                          \                      /
                          \                     \
                          /                      \
                         /___                 /___
                             \                /
             \               \               \               \
             /\              /                \              /\
         ___/ ___/       ___/                 ___/       ___/ ___/
          \      /        \                      /        \      /
     \    \     \    \    \                     \    \    \     \    \
     /\   /      \   /\   /                      \   /\   /      \   /\
 \__/ ___/       ___/ ___/                       ___/ ___/       ___/ ___/
  \                                                                      /
  \                                                                     \
  /                                                                      \
 /___                                                                 /___
     \                                                                /
     \                                                               \
     /                                                                \
 ___/                                                                 ___/
  \                                                                      /
  \                                                                     \
  /                                                                      \
 /___ /___                                                       /___ /___
     \/   \                                                      /   \/
     \    \                                                     \    \
          /                                                      \
         /___                                                 /___
             \                                                /
             \                                               \
             /                                                \
         ___/                                                 ___/
          \                                                      /
     \    \                                                     \    \
     /\   /                                                      \   /\
 ___/ ___/                                                       ___/ ___/
  \                                                                      /
  \                                                                     \
  /                                                                      \
 /___                                                                 /___
     \                                                                /
     \                                                               \
     /                                                                \
 ___/                                                                 ___/
  \                                                                      /
  \                                                                     \
  /                                                                      \
 /___ /___       /___ /___                       /___ /___       /___ /___
     \/   \      /   \/   \                      /   \/   \      /   \/
     \    \     \    \    \                     \    \    \     \    \
          /      \        /                      \        /      \
         /___ /___       /___                 /___       /___ /___
             \/              \                /              \/
             \               \               \               \
                             /                \
                         ___/                 ___/
                          \                      /
                          \                     \
                          /                      \
                         /___ /___       /___ /___
                             \/   \      /   \/
                             \    \     \    \
                                  /      \
                                 /___ /___
                                     \/
                                     \
```

