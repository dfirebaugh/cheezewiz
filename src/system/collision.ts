import { State } from "../component/state";
import { Entity } from "../entities/entity";
import CombatSystem from "./combat";

export default class CollisionSystem {
    static checkCollision(entityA: Entity, entityB: Entity): boolean {
        return (
            entityA.position?.X < entityB.position?.X + entityB.size?.width &&
            entityA.position?.X + entityA.size?.width > entityB.position?.X &&
            entityA.position?.Y < entityB.position?.Y + entityB.size?.height &&
            entityA.position?.Y + entityA.size?.height > entityB.position?.Y
        );
    }

    static update(scene: Phaser.Scene, entity: Entity, collidables: Array<Entity>) {
        for (let collidable of collidables) {
            if (collidable.state.current == State.Dead) continue;

            if (entity !== collidable && CollisionSystem.checkCollision(entity, collidable)) {
                CombatSystem.exchange(scene, entity, collidable)
                CombatSystem.exchange(scene, collidable, entity)
            }
        }
    }
}
