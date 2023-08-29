import * as Phaser from "phaser";

import { State } from "../component/state";
import { Entity } from "../entities/entity";
import CombatSystem from "./combat";

function checkCollision(entityA: Entity, entityB: Entity): boolean {
    return (
        entityA.position?.X < entityB.position?.X + entityB.size?.width &&
        entityA.position?.X + entityA.size?.width > entityB.position?.X &&
        entityA.position?.Y < entityB.position?.Y + entityB.size?.height &&
        entityA.position?.Y + entityA.size?.height > entityB.position?.Y
    );
}

export default function CollisionSystem(scene: Phaser.Scene, entity: Entity, collidables: Array<Entity>) {
    if (entity?.isDestroyed) return;

    for (let collidable of collidables) {
        if (collidable.isDestroyed) continue;
        if (collidable.state.current == State.Dead) continue;


        if (entity !== collidable && checkCollision(entity, collidable)) {
            CombatSystem(scene, entity, collidable)
            CombatSystem(scene, collidable, entity)
        }
    }
}
