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
    disableHealthBar?: boolean;

    invulnerabilityDuration = 500; // half a sec
    invulnerable: boolean = false;
    lastHitTime: number = 0;

    constructor(scene: Phaser.Scene, entity: Entity, public max: number, public current: number, public regenRate: number) {
        this.scene = scene;
        this.entity = entity;
        this.graphics = new Phaser.GameObjects.Graphics(scene);
        scene.add.existing(this.graphics);
    }
}
