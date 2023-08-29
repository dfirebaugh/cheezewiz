import * as Phaser from 'phaser';

import {
    Entity,
    EntityFactory,
    WizData,
    RadishRedData,
} from './entities/'

import {
    CollisionSystem,
    EnemyMovementSystem,
    InputSystem,
    LightSystem,
    RenderSystem,
} from './system/';
import { LoadAssets } from './assetLoader';

const ScreenWidth = 640;
const ScreenHeight = 512;

export default class Demo extends Phaser.Scene {
    constructor() {
        super('GameScene');
    }

    wiz: Entity;
    enemies: Array<Entity>;

    preload() {
        LoadAssets(this)
    }

    create() {
        const map = this.make.tilemap({ key: 'kitchen' })
        const tileset = map.addTilesetImage('kitchen floor', 'kitchen_tiles')
        map.createLayer('Tile Layer 1', tileset, 0, 0).setPipeline('Light2D');

        this.wiz = EntityFactory(this, WizData)

        this.cameras.main.startFollow(this.wiz.sprite.sprite);
        this.cameras.main.setLerp(0.1, 0.1);
        this.cameras.main.setBounds(0, 0, 5088, ScreenHeight);

        this.enemies = [];

        for (let i = 0; i < 10; i++) {
            this.enemies.push(EntityFactory(this, RadishRedData))
        }

        function getRandomInt(min, max) {
            min = Math.ceil(min);
            max = Math.floor(max);
            return Math.floor(Math.random() * (max - min + 1)) + min;
        }

        this.enemies.forEach(e => {
            e.position.X = getRandomInt(0, 800)
            e.position.Y = getRandomInt(0, 600)
        })
    }

    update() {
        this.enemies.forEach(e => {
            EnemyMovementSystem.update(e, this.wiz)
            RenderSystem.update(e)
        })

        CollisionSystem.update(this, this.wiz, this.enemies)
        InputSystem.update(this, this.wiz)
        RenderSystem.update(this.wiz)
        LightSystem.update(this.wiz)
    }
}

const config = {
    type: Phaser.AUTO,
    backgroundColor: '#000',
    width: ScreenWidth,
    height: ScreenHeight,
    scene: Demo
};

const game = new Phaser.Game(config);
