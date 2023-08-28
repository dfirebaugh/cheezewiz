import * as Phaser from 'phaser';

export default class SpriteComponent {
    sprite: Phaser.GameObjects.Sprite;
    scene: Phaser.Scene;
    currentState: string;

    constructor(scene: Phaser.Scene, label: string) {
        this.scene = scene;
        this.sprite = scene.add.sprite(0, 0, label);
    }

    addAnimation(config: Phaser.Types.Animations.Animation) {
        if (this.scene.anims.exists(this.currentState)) return;

        this.sprite.anims.create(Object.assign({
            frames: this.scene.anims.generateFrameNumbers(config.key || this.currentState, {
                start: 0,
                end: config.frames.length - 1,
            }),
            frameRate: 10,
            repeat: -1,
        },config));
    }
}

