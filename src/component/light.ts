import World from "../world";


export default class LightComponent {
    world: World;
    light: Phaser.GameObjects.Light;

    constructor(world: World, public intensity: number, public color: number, public radius: number) {
        this.world = world;

        world.scene.lights.enable();
        world.scene.lights.setAmbientColor(color);
        this.light = world.scene.lights.addLight(0, 0, radius).setIntensity(intensity);
    }
}
