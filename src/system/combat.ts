import { State } from "../component/state";
import { Entity } from "../entities";
import { displayDamage, displayHealthGain } from "./combatText";

function mitigateDamage(entity: Entity, damage: number): number {
    if (!entity.defense) {
        return damage;
    }

    return damage - (damage * entity.defense.mitigation)
}

const baseMeleeAttack = 10;

function takeDamage(scene: Phaser.Scene, entity: Entity) {
    if (entity.isDestroyed) return;

    const attackPower = entity.weapon?.power || baseMeleeAttack;
    const damage = mitigateDamage(entity, attackPower);
    entity.health.current -= damage;

    entity.state.setState(State.Hurt);

    const currentTime = scene.time.now;

    // Check if the entity is still within its invulnerability duration
    if (currentTime - entity.health.lastHitTime < entity.health.invulnerabilityDuration) {
        entity.health.invulnerable = true;
    } else {
        entity.health.invulnerable = false;
        entity.health.lastHitTime = currentTime;
    }

    displayDamage(scene, entity, damage);

    if (entity.health?.current <= 0) {
        entity.health.current = 0;
        entity.state.setState(State.Dead);
        if (!!entity.destroyable) {
            entity.destroy();
        }
    }
}

export function HealthRegen(scene: Phaser.Scene, entity: Entity) {
    if (!entity.health.regenRate) return;

    const startTick = () => {
        const interval = setInterval(() => {
            entity.health.current += entity.health.regenRate;

            if (entity.health.current > entity.health.max) entity.health.current = entity.health.max;
            displayHealthGain(scene, entity, entity.health.regenRate);
        }, 2000); // 2000 milliseconds = 2 seconds

        return interval;
    }

    startTick();
}

export default function CombatSystem(scene: Phaser.Scene, attacker: Entity, defender: Entity) {
    const currentTime = scene.time.now;

    // Check if invulnerability duration has passed for the defender
    if (defender.health.invulnerable && currentTime - defender.health.lastHitTime > defender.health.invulnerabilityDuration) {
        defender.health.invulnerable = false;
    }


    // If the defender is invulnerable, exit early without applying damage
    if (!defender.health.invulnerable) {
        takeDamage(scene, defender);
    }
}
