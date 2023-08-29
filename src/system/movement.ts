import { Entity } from "../entities/entity";

export default function MovementSystem(entity: Entity) {
    if (!entity.position || !entity.velocity) {
        return
    }

    entity.position.X += entity.velocity.VX;
    entity.position.Y += entity.velocity.VY;
}
