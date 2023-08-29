import { Entity } from "../entities/entity";
import { HealthBar } from '../component';

function renderHealthBar(entity: Entity) {
    if (!entity.health) return;
    if (!entity.health?.current && !entity.health?.max) return;
    if (!entity.health.graphics) return;
    if (!entity.health || !entity.position) return;

    if (!entity.healthBar) {
        entity.healthBar = new HealthBar()
    }

    entity.healthBar.draw(entity)
}

export default function RenderSystem(entity: Entity) {
    renderHealthBar(entity);
    entity.sprite?.sprite?.setPosition(entity.position.X, entity.position.Y);

    console.log(entity.state.toString())
    if (entity.sprite.sprite.anims.currentAnim?.key !== entity.state?.toString()) {
        entity.sprite.sprite.play(entity.state?.toString());
    }
}
