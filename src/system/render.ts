import { Entity } from "../entities/entity";
import { HealthBar } from '../component';

function renderHealthBar(entity: Entity) {
    if (!entity.health) return;
    if (!entity.health?.current && !entity.health?.max) return;
    if (!entity.health.graphics) return;
    if (!entity.health || !entity.position) return;
    if (entity.health.disableHealthBar) return;

    if (!entity.healthBar) {
        entity.healthBar = new HealthBar()
    }

    entity.healthBar.draw(entity)
}

function renderXPBar(entity: Entity) {
    if (!entity.xp || !entity.xp.bar) return;

    entity.xp.bar.draw(entity)
}

export default function RenderSystem(entity: Entity) {
    renderHealthBar(entity);
    renderXPBar(entity);

    entity.sprite?.sprite?.setPosition(entity.position.X, entity.position.Y);

    if (entity.rotation !== undefined && entity.sprite?.sprite) {
        entity.sprite.sprite.rotation = entity.rotation + Math.PI;
    }

    if (entity.isDestroyed) {
        entity.sprite?.sprite?.destroy()
        return
    }

    if (entity.sprite.sprite?.anims?.currentAnim?.key !== entity.state?.toString()) {
        entity.sprite?.sprite?.play(entity.state?.toString());
    }
}
