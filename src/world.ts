import { Entity } from "./entities";


export default class World {
    scene: Phaser.Scene;
    wiz: Entity;
    enemies: Array<Entity> = [];
    missiles: Array<Entity> = [];
    loot: Array<Entity> = [];

    constructor(scene: Phaser.Scene) {
        this.scene = scene;
    }
}
