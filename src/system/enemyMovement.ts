import { State } from "../component/state";
import { Entity } from "../entities/entity";

export default class EnemyMovementSystem {
    static trackPlayer(entity: Entity, player: Entity) {
        if (player.state.current == State.Dead) {
            EnemyMovementSystem.defaultMovement(entity);
            return;
        }
        const dirX = player.position.X - entity.position.X;
        const dirY = player.position.Y - entity.position.Y;

        const distance = Math.sqrt(dirX * dirX + dirY * dirY);

        const unitX = dirX / distance;
        const unitY = dirY / distance;

        entity.velocity.VX = unitX * entity.speed.value;
        entity.velocity.VY = unitY * entity.speed.value;
    }

    static defaultMovement(entity: Entity) {
        const center = {
            X: 800 / 2,
            Y: 600 / 2
        }
        const dirXToCenter = center.X - entity.position.X;
        const dirYToCenter = center.Y - entity.position.Y;
        const distanceToCenter = Math.sqrt(dirXToCenter * dirXToCenter + dirYToCenter * dirYToCenter);
        const unitX = dirXToCenter / distanceToCenter;
        const unitY = dirYToCenter / distanceToCenter;
    
        if (distanceToCenter > 400) {
            EnemyMovementSystem.moveTo(entity, unitX, unitY);
            return;
        }
    
        if (distanceToCenter < 100) {
            EnemyMovementSystem.moveAwayFrom(entity, center.X, center.Y); // Corrected to move away from the center
            return;
        }
    
        if (distanceToCenter >= 100 && distanceToCenter <= 400) {
            entity.state.current = State.Idle;
            entity.velocity.VX = 0;
            entity.velocity.VY = 0;
            return;
        }
    }
    

    static moveTo(entity: Entity, unitX: number, unitY: number) {
        entity.velocity.VX = unitX * entity.speed.value;
        entity.velocity.VY = unitY * entity.speed.value;
    }

    static moveAwayFrom(entity: Entity, unitX: number, unitY: number) {
        entity.velocity.VX = -entity.velocity.VX;
        entity.velocity.VY = -entity.velocity.VX;
    }

    static update(entity: Entity, player: Entity) {
        if (!entity.position || !entity.velocity || !player.position) {
            return;
        }

        if (entity.health?.current <= 0) {
            entity.state.current = State.Dead
            return
        }

        const dirXToPlayer = player.position.X - entity.position.X;
        const dirYToPlayer = player.position.Y - entity.position.Y;
        const distanceToPlayer = Math.sqrt(dirXToPlayer * dirXToPlayer + dirYToPlayer * dirYToPlayer);
        const trackingThreshold = 300;

        if (distanceToPlayer <= trackingThreshold) {
            EnemyMovementSystem.trackPlayer(entity, player);
        } else {
            EnemyMovementSystem.defaultMovement(entity)
        }

        entity.state.current = State.Walking;
        entity.position.X += entity.velocity.VX;
        entity.position.Y += entity.velocity.VY;
    }
}
