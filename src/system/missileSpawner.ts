import { State } from "../component/state";
import { WeaponType } from "../component/weapon";
import { Entity, EntityFactory, NachoMissileData } from "../entities";
import World from "../world";

let current = 0

function spawnMissile(world: World, caster: Entity, label: string, speed: number, power: number, weaponType: WeaponType) {
    if (!speed) return;

    const startTick = () => {
        const interval = setInterval(() => {
            current += speed
            const missile = EntityFactory(world, NachoMissileData)
            missile.position.X = caster.position.X
            missile.position.Y = caster.position.Y

            missile.state.setState(State.Walking)

            world.missiles.push(missile)
        }, speed * 100);

        return interval;
    }

    startTick();
}

export default function MissileSpawner(world: World, caster: Entity, label: string, speed: number, power: number, weaponType: WeaponType) {
    spawnMissile(world, caster, label, speed, power, weaponType)
}
