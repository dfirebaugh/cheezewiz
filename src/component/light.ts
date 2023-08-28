

export default class LightComponent {
    scene: Phaser.Scene;
    light: Phaser.GameObjects.Light;

    constructor(scene: Phaser.Scene, public intensity: number, public color: number, public radius: number){
        this.scene = scene;

        scene.lights.enable();
        scene.lights.setAmbientColor(color);
        this.light = scene.lights.addLight(0, 0, radius).setIntensity(intensity);
    }
}
