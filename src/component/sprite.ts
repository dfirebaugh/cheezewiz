import * as Phaser from 'phaser';
import World from '../world';

export default class SpriteComponent {
    sprite: Phaser.GameObjects.Sprite;
    world: World;

    constructor(world: World, label: string) {
        this.world = world;
        this.sprite = world.scene.add.sprite(0, 0, label);
    }

    addAnimation(config: Phaser.Types.Animations.Animation) {
        if (!config.key) {
            console.error("Animation config must have a 'key' property.");
            return;
        }

        if (this.world.scene.anims.exists(config.key)) {
            console.error("Animation already exists.", config.key);
            return;
        }

        const animationConfig = Object.assign({
            frameRate: 10,
            repeat: -1,
        }, config);

        this.sprite?.anims.create(animationConfig);
    }

    flash(duration: number = 250, repeat: number = 4) {
        if (this.sprite) {
            this.world.scene.tweens.add({
                targets: this.sprite,
                alpha: { from: 1, to: 0 },
                duration: duration / (2 * repeat),
                yoyo: true,
                repeat: repeat - 1
            });
        }
    }
}

