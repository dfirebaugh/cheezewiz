import * as Phaser from 'phaser';
import { Entity, EntityFactory, RadishRedData } from "../entities";

function spawnEnemy(scene: Phaser.Scene & { enemies: Array<Entity> }, player: Entity) {
    const MIN_DISTANCE = 300;
    const MAX_DISTANCE = 500;

    const angle = Phaser.Math.FloatBetween(0, 2 * Math.PI);

    const distance = Phaser.Math.FloatBetween(MIN_DISTANCE, MAX_DISTANCE);

    const spawnX = player.position.X + distance * Math.cos(angle);
    const spawnY = player.position.Y + distance * Math.sin(angle);

    const enemy = EntityFactory(scene, RadishRedData);
    enemy.position.X = spawnX;
    enemy.position.Y = spawnY;
    enemy.sprite.sprite?.setPipeline('Light2D');
    scene.enemies.push(enemy);
}

export default function EnemySpawner(scene: Phaser.Scene & { enemies: Array<Entity> }, player: Entity, spawnRate: number) {
    let spawnInterval: NodeJS.Timeout;

    const startTick = () => {
        spawnInterval = setInterval(() => {
            if (scene.enemies.length >= 1000) {
                clearInterval(spawnInterval);
            } else {
                spawnEnemy(scene, player);
            }
        }, spawnRate * 1000);

        return spawnInterval;
    }

    setInterval(() => {
        if (scene.enemies.length < 1000 && !spawnInterval) {
            startTick();
        }
    }, 5000);

    startTick();
}
