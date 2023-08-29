import * as Phaser from 'phaser';

export default class SpriteComponent {
    sprite: Phaser.GameObjects.Sprite;
    scene: Phaser.Scene;

    constructor(scene: Phaser.Scene, label: string) {
        this.scene = scene;
        this.sprite = scene.add.sprite(0, 0, label);
    }

    addAnimation(config: Phaser.Types.Animations.Animation) {
        if (!config.key) {
            console.error("Animation config must have a 'key' property.");
            return;
        }

        if (this.scene.anims.exists(config.key)) {
            console.error("Animation already exists.", config.key);
            return;
        }

        const animationConfig = Object.assign({
            frameRate: 10,
            repeat: -1,
        }, config);

        this.sprite?.anims.create(animationConfig);
    }
}

