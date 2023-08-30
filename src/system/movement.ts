import { Entity } from "../entities/entity";

export default function MovementSystem(entity: Entity) {
    if (!entity.position || !entity.velocity) {
        return
    }

    entity.position.X += entity.velocity.VX;
    entity.position.Y += entity.velocity.VY;
}

export function MoveAway(source: Entity, target: Entity, speed: number) {
    const distanceX = source.position.X - target.position.X;
    const distanceY = source.position.Y - target.position.Y;

    const newX = source.position.X + distanceX * speed;
    const newY = source.position.Y + distanceY * speed;

    source.position.X = newX;
    source.position.Y = newY;
}

export function MoveToward(source: Entity, target: Entity, speed: number) {
    const distanceX = target.position.X - source.position.X;
    const distanceY = target.position.Y - source.position.Y;

    const newX = source.position.X + distanceX * speed;
    const newY = source.position.Y + distanceY * speed;

    source.position.X = newX;
    source.position.Y = newY;
}
