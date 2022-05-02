# Scripting
We have an embedded javascript engine for scripting.

We can define javascript functions that are available from the global scope in the script that will call go functions.

To do this, define a function in the `cheezewiz/vm` package.

Within javascript you should be able to do basic control flow that will be triggered on an action.

## Building Levels

We can also configure levels with these scripts.
You will need to add the `Systems` that are in effect to the scene (aka level).

A basic example of a scene with simple player movement enabled.
```javascript

/**
 * load optional systems/entities into the scene
 */
function init() {
    addPlayerControlSystem();
    addMovementSystem();
    addRenderSystem();
    addAnimationSystem();
    addPlayerEntity();
}

init();

```

## Building actions
tbd...
