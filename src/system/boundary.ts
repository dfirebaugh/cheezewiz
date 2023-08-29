import { Entity } from "../entities/entity";

function isOverRightBounds(entity: Entity, boundary: Entity) {
    return entity.position?.X + entity.rigidBody?.X < boundary.position?.X + boundary.rigidBody?.X
}
function isOverLeftBounds(entity: Entity, boundary: Entity) {
    return entity.position?.X + entity.rigidBody?.X > boundary.position?.X + boundary.rigidBody?.W
}
function isOverLowerBounds(entity: Entity, boundary: Entity) {
    return entity.position?.Y + entity.rigidBody?.Y > boundary.position?.X + boundary.rigidBody?.H
}
function isOverUpperBounds(entity: Entity, boundary: Entity) {
    return entity.position?.Y + entity.rigidBody?.Y < boundary.position?.Y + boundary.rigidBody?.Y
}

export default function BoundarySystem(entity: Entity, boundary: Entity) {
    if (!entity.position || !entity.velocity || !entity.rigidBody) {
        return
    }

    if (isOverRightBounds(entity, boundary)) {
        entity.velocity.VX = -entity.velocity.VX
    }
    if (isOverLeftBounds(entity, boundary)) {
        entity.velocity.VX = -entity.velocity.VX
    }
    if (isOverLowerBounds(entity, boundary)) {
        entity.velocity.VY = -entity.velocity.VY
    }
    if (isOverUpperBounds(entity, boundary)) {
        entity.velocity.VY = -entity.velocity.VY
    }
}
