import * as Phaser from "phaser";
import { Entity } from "../entities";
import World from "../world";

export function displayDamage(world: World, entity: Entity, damage: number) {
    if (entity.health.disableHealthBar) return;

    const damageText = world.scene.add.text(entity.position.X, entity.position.Y, `-${damage}`, {
        fontSize: '12px',
        color: '#ff0000',
        stroke: '#000',
        strokeThickness: 2
    });

    const randomAngle = Phaser.Math.Between(0, 360) * (Math.PI / 180);

    const offsetX = 30 * Math.cos(randomAngle);
    const offsetY = 30 * Math.sin(randomAngle);


    world.scene.tweens.add({
        targets: damageText,
        x: damageText.x + offsetX,
        y: damageText.y + offsetY,
        alpha: 0,
        duration: 1000,
        onComplete: () => {
            damageText.destroy();
        }
    });
}

export function displayHealthGain(world: World, entity: Entity, gain: number) {
    const healthGainText = world.scene.add.text(entity.position.X, entity.position.Y, `+${gain}`, {
        fontSize: '12px',
        color: '#00ff00',
        stroke: '#000',
        strokeThickness: 2
    });

    const randomAngle = Phaser.Math.Between(0, 360) * (Math.PI / 180);

    const offsetX = 30 * Math.cos(randomAngle);
    const offsetY = 30 * Math.sin(randomAngle);

    world.scene.tweens.add({
        targets: healthGainText,
        x: healthGainText.x + offsetX,
        y: healthGainText.y + offsetY,
        alpha: 0,
        duration: 1000,
        onComplete: () => {
            healthGainText.destroy();
        }
    });
}
