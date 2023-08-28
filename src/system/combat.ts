import { State } from "../component/state";
import { Entity } from "../entities";

export default class CombatSystem {
    static exchange(scene: Phaser.Scene, attacker: Entity, defender: Entity) {
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

        const damage = 15; // This can be adjusted based on attacker's power and defender's defense
        defender.health.current -= damage;

        // Set the defender to be invulnerable and record the hit time
        defender.health.invulnerable = true;
        defender.health.lastHitTime = currentTime;

        defender.health.displayDamage(scene, defender, damage);

        if (defender.health?.current <= 0) {
            defender.state.current = State.Dead;
        }
    }
}
