import { Entity } from "../entities/entity";

export default class BoundarySystem {
    static isOverRightBounds(entity: Entity, boundary: Entity) {
        return entity.position?.X + entity.rigidBody?.X < boundary.position?.X + boundary.rigidBody?.X
    }
    static isOverLeftBounds(entity: Entity, boundary: Entity) {
        return entity.position?.X + entity.rigidBody?.X > boundary.position?.X + boundary.rigidBody?.W
    }
    static isOverLowerBounds(entity: Entity, boundary: Entity) {
        return entity.position?.Y  + entity.rigidBody?.Y > boundary.position?.X + boundary.rigidBody?.H
    }
    static isOverUpperBounds(entity: Entity, boundary: Entity) {
        return entity.position?.Y + entity.rigidBody?.Y < boundary.position?.Y + boundary.rigidBody?.Y
    }

    static update(entity: Entity, boundary: Entity) {
        if (!entity.position || !entity.velocity || !entity.rigidBody) {
            return
        }

        if (BoundarySystem.isOverRightBounds(entity, boundary)) {
            entity.velocity.VX = -entity.velocity.VX
        }
        if (BoundarySystem.isOverLeftBounds(entity, boundary)) {
            entity.velocity.VX = -entity.velocity.VX
        }
        if (BoundarySystem.isOverLowerBounds(entity, boundary)) {
            entity.velocity.VY = -entity.velocity.VY
        }
        if (BoundarySystem.isOverUpperBounds(entity, boundary)) {
            entity.velocity.VY = -entity.velocity.VY
        }
    }
}
