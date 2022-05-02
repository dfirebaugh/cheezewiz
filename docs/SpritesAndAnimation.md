# Animation
We are using [ganim8](https://github.com/yohamta/ganim8) to help easily make animations from spritesheets.

To make an animation, place a png of the sprite sheet in the `assets` directory.

The png should be in a grid layout (like the below image).
Animations are accessed by this function: `grid.GetFrames("1-5", 5)`
The first argument is the range in the column.  The second argument is the row.
So, it may make sense to have an action per row (e.g. the sequence for a jump might exist on one row and the sequence for an attack could exist on another)
> note that we should be able to trigger multiple rows in a single action

<p align="center">
  <img src="https://github.com/yohamta/ganim8/blob/master/examples/gif/example.gif?raw=true" />
</p>

<img src="https://github.com/yohamta/ganim8/blob/master/examples/assets/images/Character_Monster_Slime_Blue.png?raw=true" />

> images are from [ganim8](https://github.com/yohamta/ganim8) example
