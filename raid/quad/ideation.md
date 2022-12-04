#   MIND
##  INSTRUCTIONS
Write a function QuadA that prints a valid rectangle with a given width of `x` and height of `y`.
The function must draw the rectangles as in the examples.
If `x` and `y` are positive numbers, the program should print the rectangles as seen in the examples, otherwise, the function should print nothing.
Make sure you submit all the necessary files to run the program.

##  NOTES
+   We must print 5 kind of rectangles
+   Those kind of rectangles is differences by the character used for drawing them
    +   quadA draw a rectangle with
        +   `o` on the corners
        +   `-` on the row edges
        +   `|` on the column edges
    +   quadB draw a rectangle with
        +   `\` on the top right corner and bottom left corner
        +   `/` on the top left corner and bottom right corner
        +   `*` on the edge
    +   quadC draw a rectangle with
        +   `A` on the top corners
        +   `C` on the bottom corners
        +   `B` on the edges
    +   quadD draw a rectangle with
        +   `A` on the left corners
        +   `C` on the right corners
        +   `B` on the edges
    +   quadD draw a rectangle with
        +   `A` on the top left corner and bottom right corner
        +   `C` on the top right corner and bottom left corner
        +   `B` on the edges

##  IMPLEMENTATION
We need :
+   A basic function for drawing rectangle
+   A function for drawing line
    +   Take care of the corner
    +   Take care of the edge
    +   Take care of the void space
+   Reuse the function with the different quad symbol
    +   variable for each symbol
        +   top left corner symbol
        +   bottom left corner symbol
        +   top right corner symbol
        +   top right corner symbol
        +   bottom right corner symbol
        +   row edge symbol
        +   column edge symbol

##  PSEUDO CODE

```
cornerTopRight
cornerTopLeft
cornerBottomRight
cornerBottomLeft
edgeLine
edgeColumn

afficher rectangle
    pour i = 0 -> max_ligne
        pour j = 0 -> max_colonne
            si i = 0
                afficher premier ligne
            si i > 0 et i < max_ligne
                affiche les ligne avec espace vide
            si i = max_ligne
                affiche la dernière ligne
        end_pour
    end_pour

afficher une ligne
    afficher startCharacter
    pour i = 1 -> max_colonne - 1
        afficher middleCharacter
    end_pour
    afficher endCharacter
```

##  OPEN SOURCE
Baccalauréat - 2018
Lycée Thuriaf BANTSANTSA
Gabon