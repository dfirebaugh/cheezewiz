
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

