import * as Phaser from "phaser";
import { Entity } from "../entities";

export function displayDamage(scene: Phaser.Scene, entity: Entity, damage: number) {
    if (entity.health.disableHealthBar) return;

    // Create a text object to display the damage
    const damageText = scene.add.text(entity.position.X, entity.position.Y, `-${damage}`, {
        fontSize: '12px',
        color: '#ff0000',
        stroke: '#000',
        strokeThickness: 2
    });

    const randomAngle = Phaser.Math.Between(0, 360) * (Math.PI / 180);

    const offsetX = 30 * Math.cos(randomAngle);
    const offsetY = 30 * Math.sin(randomAngle);


    scene.tweens.add({
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

export function displayHealthGain(scene: Phaser.Scene, entity: Entity, gain: number) {
    // Create a text object to display the health gain
    const healthGainText = scene.add.text(entity.position.X, entity.position.Y, `+${gain}`, {
        fontSize: '12px',
        color: '#00ff00', // Green color
        stroke: '#000',
        strokeThickness: 2
    });

    const randomAngle = Phaser.Math.Between(0, 360) * (Math.PI / 180);

    const offsetX = 30 * Math.cos(randomAngle);
    const offsetY = 30 * Math.sin(randomAngle);

    scene.tweens.add({
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
