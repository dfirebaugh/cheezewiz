import * as Phaser from 'phaser';
import { Entity } from '../entities';

export class HealthBar {
    barWidth = 32;
    barHeight = 3;
    barOffsetY = 18;

    calculateHealthPercentage(entity: Entity): number {
        return entity.health.current / entity.health.max;
    }

    drawBackground(entity: Entity) {
        entity.health.graphics.fillStyle(0xfff, 0.8);
        entity.health.graphics.fillRect(entity.position.X - this.barWidth / 2, entity.position.Y + this.barOffsetY, this.barWidth, this.barHeight);
    }

    draw(entity: Entity) {
        this.drawBackground(entity)
        entity.health.graphics.clear()
        entity.health.graphics.fillStyle(0xff0000);
        entity.health.graphics.fillRect(entity.position.X - this.barWidth / 2, entity.position.Y + this.barOffsetY, this.barWidth * this.calculateHealthPercentage(entity), this.barHeight);
    }
}

export default class HealthComponent {
    graphics?: Phaser.GameObjects.Graphics;
    scene?: Phaser.Scene;
    entity?: Entity;

    invulnerabilityDuration = 500; // half a sec
    invulnerable: boolean = false;
    lastHitTime: number = 0;

    constructor(scene: Phaser.Scene, entity: Entity, public max: number, public current: number, public regenRate: number) {
        this.scene = scene;
        this.entity = entity;
        this.graphics = new Phaser.GameObjects.Graphics(scene);
        scene.add.existing(this.graphics);
        this.regen();
    }

    regen() {
        if (!this.regenRate) return;

        const startTick = () => {
            const interval = setInterval(() => {
                this.current += this.regenRate

                if (this.current > this.max) this.current = this.max;

                this.displayHealthGain(this.scene, this.entity, this.regenRate)
            }, 2000); // 2000 milliseconds = 2 seconds

            return interval;
        }

        startTick();
    }

    displayDamage(scene: Phaser.Scene, entity: Entity, damage: number) {
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

    displayHealthGain(scene: Phaser.Scene, entity: Entity, gain: number) {
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
}
