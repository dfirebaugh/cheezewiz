import { Entity } from "../entities/entity";

export default class LightSystem {
    static update(entity: Entity) {
        if (!entity.light) return;

        entity.light.light.x = entity.position.X
        entity.light.light.y = entity.position.Y
    }
}
