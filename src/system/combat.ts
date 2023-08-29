import { State } from "../component/state";
import { Entity } from "../entities";

function takeDamage(scene: Phaser.Scene, entity: Entity) {
    const currentTime = scene.time.now;
    const damage = 15; // This can be adjusted based on attacker's power and defender's defense
    entity.health.current -= damage;
    entity.state.setState(State.Hurt);

    // Set the defender to be invulnerable and record the hit time
    entity.health.invulnerable = true;
    entity.health.lastHitTime = currentTime;

    entity.health.displayDamage(scene, entity, damage);

    if (entity.health?.current <= 0) {
        entity.state.setState(State.Dead);
    }
}

export default function CombatSystem(scene: Phaser.Scene, attacker: Entity, defender: Entity) {
    const currentTime = scene.time.now;

    // Check if invulnerability duration has passed for the defender
    if (defender.health.invulnerable && currentTime - defender.health.lastHitTime > defender.health.invulnerabilityDuration) {
        defender.health.invulnerable = false;
    }

    if (attacker.state.current == State.Dead || defender.state.current == State.Dead) return;

    // If the defender is invulnerable, exit early without applying damage
    if (defender.health.invulnerable) return;

    // take into account attacker's ability power or something
    // consider the defenders defense score?
    takeDamage(scene, defender)
    takeDamage(scene, attacker)
}
