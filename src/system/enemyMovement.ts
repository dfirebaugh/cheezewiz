import { State } from "../component/state";
import { Entity } from "../entities/entity";
import { MoveAway, MoveTo, MoveToward } from "./movement";

const MAP_WIDTH = 800;
const MAP_HEIGHT = 600;
const MAP_LEFT_BOUND = 0;
const MAP_RIGHT_BOUND = MAP_WIDTH;
const MAP_TOP_BOUND = 0;
const MAP_BOTTOM_BOUND = MAP_HEIGHT;
const DISTANCE_THRESHOLD = 500;

function isOutsideMap(entity: Entity): boolean {
    return (
        entity.position.X < MAP_LEFT_BOUND ||
        entity.position.X > MAP_RIGHT_BOUND ||
        entity.position.Y < MAP_TOP_BOUND ||
        entity.position.Y > MAP_BOTTOM_BOUND
    );
}

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

function moveToMap(entity: Entity) {
    const center = {
        X: MAP_WIDTH / 2,
        Y: MAP_HEIGHT / 2
    };
    const dirXToCenter = center.X - entity.position.X;
    const dirYToCenter = center.Y - entity.position.Y;
    const distanceToCenter = Math.sqrt(dirXToCenter * dirXToCenter + dirYToCenter * dirYToCenter);
    const unitX = dirXToCenter / distanceToCenter;
    const unitY = dirYToCenter / distanceToCenter;

    MoveTo(entity, unitX, unitY);
}

function defaultMovement(entity: Entity, player: Entity) {
    const dirXToPlayer = player.position.X - entity.position.X;
    const dirYToPlayer = player.position.Y - entity.position.Y;
    const distanceToPlayer = Math.sqrt(dirXToPlayer * dirXToPlayer + dirYToPlayer * dirYToPlayer);

    if (distanceToPlayer > DISTANCE_THRESHOLD) {
        const unitX = dirXToPlayer / distanceToPlayer;
        const unitY = dirYToPlayer / distanceToPlayer;

        MoveToward(entity, player, entity.speed.value)
        return;
    }

    if (isOutsideMap(entity)) {
        moveToMap(entity)
        return
    }
    const center = {
        X: 800 / 2,
        Y: 600 / 2
    }
    const dirXToCenter = center.X - entity.position.X;
    const dirYToCenter = center.Y - entity.position.Y;
    const distanceToCenter = Math.sqrt(dirXToCenter * dirXToCenter + dirYToCenter * dirYToCenter);

    if (distanceToCenter > 400) {
        MoveToward(entity, player, entity.speed.value)
        return;
    }

    if (distanceToCenter < 100) {
        MoveAway(entity, player, entity.speed.value)
        return;
    }

    if (distanceToCenter >= 100 && distanceToCenter <= 400) {
        entity.velocity.VX = 0;
        entity.velocity.VY = 0;
        return;
    }
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
