import * as Phaser from 'phaser';
import { Entity, EntityFactory, RadishRedData, RadishBlueData, RadishYellowData } from "../entities";
import World from '../world';

function spawnEnemy(world: World, player: Entity) {
    const MIN_DISTANCE = 300;
    const MAX_DISTANCE = 500;

    const angle = Phaser.Math.FloatBetween(0, 2 * Math.PI);
    const distance = Phaser.Math.FloatBetween(MIN_DISTANCE, MAX_DISTANCE);
    const spawnX = player.position.X + distance * Math.cos(angle);
    const spawnY = player.position.Y + distance * Math.sin(angle);

    const enemyDataTypes = [RadishRedData, RadishBlueData, RadishYellowData];

    const randomEnemyData = Phaser.Math.RND.pick(enemyDataTypes);

    const enemy = EntityFactory(world, randomEnemyData);
    enemy.position.X = spawnX;
    enemy.position.Y = spawnY;
    enemy.dropsLoot = true;
    enemy.sprite.sprite?.setPipeline('Light2D');
    world.enemies.push(enemy);
}

export default function EnemySpawner(world: World, player: Entity, spawnRate: number) {
    let spawnInterval: NodeJS.Timeout;

    const startTick = () => {
        spawnInterval = setInterval(() => {
            if (world.enemies.length >= 300) {
                clearInterval(spawnInterval);
            } else {
                spawnEnemy(world, player);
            }
        }, spawnRate * 1000);

        return spawnInterval;
    }

    setInterval(() => {
        if (world.enemies.length < 1000 && !spawnInterval) {
            startTick();
        }
    }, 5000);

    startTick();
}
