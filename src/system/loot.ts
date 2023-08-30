import { Entity, EntityFactory, JellyBeanRainbowData } from "../entities";
import World from "../world";
import { CheckCollision } from "./collision";
import { MoveAway, MoveToward } from "./movement";

export function DropJellyBean(world: World, dropper: Entity) {
    const jb = EntityFactory(world, JellyBeanRainbowData)
    jb.destroyable = true;
    jb.position = dropper.position
    jb.sprite.sprite?.setPipeline('Light2D');

    MoveAway(jb, world.wiz, 0.4)

    world.loot.push(jb)
}

export function LootMovement(world: World, loot: Entity) {
    const playerPosition = world.wiz.position;
    const lootPosition = loot.position;

    const direction = {
        x: playerPosition.X - lootPosition.X,
        y: playerPosition.Y - lootPosition.Y
    };

    const distance = Math.sqrt(direction.x * direction.x + direction.y * direction.y);

    if (distance > 100) return;
    const movementSpeed = 0.04;

    MoveToward(loot, world.wiz, movementSpeed)
}

export function LootSystem(world: World, loot: Entity) {
    LootMovement(world, loot);
    LootCollision(world, loot);
}

export function LootCollision(world: World, loot: Entity) {
    const player = world.wiz;

    if (CheckCollision(player, loot)) {
        player.xp.xp++;
        loot.destroy();
    }
}
