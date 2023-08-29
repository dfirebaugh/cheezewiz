import { State } from "../component/state";
import { WeaponType } from "../component/weapon";
import { Entity, EntityFactory, NachoMissileData } from "../entities";

let current = 0

function spawnMissile(scene: Phaser.Scene & { missiles: Array<Entity> }, caster: Entity, label: string, speed: number, power: number, weaponType: WeaponType) {
    if (!speed) return;

    const startTick = () => {
        const interval = setInterval(() => {
            current += speed
            const missile = EntityFactory(scene, NachoMissileData)
            missile.position.X = caster.position.X
            missile.position.Y = caster.position.Y

            missile.state.setState(State.Walking)

            scene.missiles.push(missile)
        }, speed * 100);

        return interval;
    }

    startTick();
}

export default function MissileSpawner(scene: Phaser.Scene & { missiles: Array<Entity> }, caster: Entity, label: string, speed: number, power: number, weaponType: WeaponType) {
    spawnMissile(scene, caster, label, speed, power, weaponType)
}
