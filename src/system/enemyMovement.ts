import { State } from "../component/state";
import { Entity } from "../entities/entity";

function trackPlayer(entity: Entity, player: Entity) {
    if (player.state.current == State.Dead) {
        defaultMovement(entity, player);
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

function defaultMovement(entity: Entity, player: Entity) {
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
        moveTo(entity, unitX, unitY);
        return;
    }

    if (distanceToCenter < 100) {
        moveAwayFrom(entity, player.position.X, player.position.Y); // Corrected to move away from the center
        return;
    }

    if (distanceToCenter >= 100 && distanceToCenter <= 400) {
        entity.velocity.VX = 0;
        entity.velocity.VY = 0;
        return;
    }
}


function moveTo(entity: Entity, unitX: number, unitY: number) {
    entity.velocity.VX = unitX * entity.speed.value;
    entity.velocity.VY = unitY * entity.speed.value;
}

function moveAwayFrom(entity: Entity, unitX: number, unitY: number) {
    entity.velocity.VX = -unitX * entity.velocity.VX;
    entity.velocity.VY = -unitY * entity.velocity.VX;
}

export default function EnemyMovementSystem(entity: Entity, player: Entity) {
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
        trackPlayer(entity, player);
    } else {
        defaultMovement(entity, player)
    }

    entity.state.current = State.Walking;
    entity.position.X += entity.velocity.VX;
    entity.position.Y += entity.velocity.VY;
}
