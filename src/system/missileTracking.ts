import { Entity } from "../entities/entity";

export default function MissileTrackingSystem(entity: Entity, enemies: Array<Entity>) {
    if (!entity.position || !entity.velocity) {
        return;
    }

    const MAX_TRAVEL_DISTANCE = 800;
    const MAX_TRACKING_DISTANCE = 500;

    if (!entity.weapon.distanceTraveled) {
        entity.weapon.distanceTraveled = 0;
    }

    let closestEnemy: Entity | null = null;
    let closestDistance: number = Infinity;

    for (let enemy of enemies) {
        if (enemy.isDestroyed) continue;

        const dx = enemy.position.X - entity.position.X;
        const dy = enemy.position.Y - entity.position.Y;
        const distance = Math.sqrt(dx * dx + dy * dy);

        if (distance < closestDistance) {
            closestDistance = distance;
            closestEnemy = enemy;
        }
    }

    if (closestEnemy && closestDistance <= MAX_TRACKING_DISTANCE) {
        const dirX = closestEnemy.position.X - entity.position.X;
        const dirY = closestEnemy.position.Y - entity.position.Y;

        const distance = Math.sqrt(dirX * dirX + dirY * dirY);
        const unitX = dirX / distance;
        const unitY = dirY / distance;

        const speed = entity.speed?.value || 1;
        entity.velocity.VX = unitX * speed;
        entity.velocity.VY = unitY * speed;

        entity.rotation = Math.atan2(unitY, unitX);
    } else {
        entity.rotation = Math.atan2(entity.velocity.VY, entity.velocity.VX);
    }

    const prevX = entity.position.X;
    const prevY = entity.position.Y;
    entity.position.X += entity.velocity.VX;
    entity.position.Y += entity.velocity.VY;

    const dx = entity.position.X - prevX;
    const dy = entity.position.Y - prevY;
    const frameDistance = Math.sqrt(dx * dx + dy * dy);

    entity.weapon.distanceTraveled += frameDistance;

    if (entity.weapon.distanceTraveled > MAX_TRAVEL_DISTANCE) {
        entity.isDestroyed = true;
    }
}
