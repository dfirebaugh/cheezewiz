import * as Phaser from 'phaser';
import { Entity } from "../entities/entity";
import { HealthBar } from '../component';

export default class RenderSystem {
    static renderHealthBar(entity: Entity) {
        if (!entity.health) return;
        if (!entity.health?.current && !entity.health?.max) return;
        if (!entity.health.graphics) return;
        if (!entity.health || !entity.position) return;

        if (!entity.healthBar) {
            entity.healthBar = new HealthBar()
        }

        entity.healthBar.draw(entity)
    }

    static update(entity: Entity) {
        RenderSystem.renderHealthBar(entity);
        entity.sprite?.sprite?.setPosition(entity.position.X, entity.position.Y);
        if (!entity.state?.toString()) return;
        if (!entity.sprite?.sprite.anims.exists(entity.state?.toString())) return;


        entity.sprite?.sprite?.play(entity.state?.toString());
    }
}
