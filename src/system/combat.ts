import { State } from "../component/state";
import { Entity } from "../entities";
import World from "../world";
import { displayDamage, displayHealthGain } from "./combatText";
import { DropJellyBean } from "./loot";
import { MoveAway } from "./movement";

function mitigateDamage(entity: Entity, damage: number): number {
    if (!entity.defense) {
        return damage;
    }

    return damage - (damage * entity.defense.mitigation)
}

const baseMeleeAttack = 10;

function takeDamage(world: World, entity: Entity) {
    if (entity.isDestroyed) return;

    const attackPower = entity.weapon?.power || baseMeleeAttack;
    const damage = mitigateDamage(entity, attackPower);
    entity.health.current -= damage;

    entity.state.setState(State.Hurt);
    entity.sprite.flash()

    const currentTime = world.scene.time.now;

    if (currentTime - entity.health.lastHitTime < entity.health.invulnerabilityDuration) {
        entity.health.invulnerable = true;
    } else {
        entity.health.invulnerable = false;
        entity.health.lastHitTime = currentTime;
    }

    displayDamage(world, entity, damage);

    if (entity.health?.current <= 0) {
        entity.health.current = 0;
        entity.state.setState(State.Dead);
        if (!!entity.destroyable) {
            entity.destroy();
        }
    }
}

export function HealthRegen(world: World, entity: Entity) {
    if (!entity.health.regenRate) return;

    const startTick = () => {
        const interval = setInterval(() => {
            entity.health.current += entity.health.regenRate;

            if (entity.health.current > entity.health.max) entity.health.current = entity.health.max;
            displayHealthGain(world, entity, entity.health.regenRate);
        }, 2000); // 2000 milliseconds = 2 seconds

        return interval;
    }

    startTick();
}

export default function CombatSystem(world: World, attacker: Entity, defender: Entity) {
    const currentTime = world.scene.time.now;

    if (defender.health.invulnerable && currentTime - defender.health.lastHitTime > defender.health.invulnerabilityDuration) {
        defender.health.invulnerable = false;
    }


    if (!defender.health.invulnerable) {
        takeDamage(world, defender);
        MoveAway(defender, attacker, defender.speed?.value)
    }
}
