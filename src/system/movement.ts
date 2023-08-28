import { Entity } from "../entities/entity";

export default class MovementSystem {
    static update(entity: Entity) {
        if (!entity.position || !entity.velocity) {
            return
        }

        entity.position.X += entity.velocity.VX;
        entity.position.Y += entity.velocity.VY;
    }
}
