import * as Phaser from 'phaser';

import {
    Entity,
    EntityFactory,
    WizData,
    RadishRedData,
    NachoMissileData,
} from './entities/'

import {
    CollisionSystem,
    EnemyMovementSystem,
    InputSystem,
    LightSystem,
    MissileTrackingSystem,
    RenderSystem,
} from './system/';
import { LoadAssets } from './assetLoader';

import MissileSpawner from './system/missileSpawner';
import { WeaponType } from './component/weapon';
import { HealthRegen } from './system/combat';
import EnemySpawner from './system/enemySpawner';

const ScreenWidth = 640;
const ScreenHeight = 512;
const SceneWidth = 5088
const SceneHeight = 512

export default class Demo extends Phaser.Scene {
    constructor() {
        super('GameScene');
    }

    wiz: Entity;
    enemies: Array<Entity>;
    missiles: Array<Entity>;

    preload() {
        LoadAssets(this)
    }

    create() {
        const map = this.make.tilemap({ key: 'kitchen' })
        const tileset = map.addTilesetImage('kitchen floor', 'kitchen_tiles')
        map.createLayer('Tile Layer 1', tileset, 0, 0).setPipeline('Light2D');

        this.wiz = EntityFactory(this, WizData)
        HealthRegen(this, this.wiz)

        this.missiles = []
        MissileSpawner(this, this.wiz, 'nachomissile', 20, 20, WeaponType.EXPLOSIVE)

        this.cameras.main.startFollow(this.wiz.sprite.sprite);
        this.cameras.main.setLerp(0.1, 0.1);
        this.cameras.main.setBounds(0, 0, SceneWidth, ScreenHeight);

        this.enemies = [];
        EnemySpawner(this, this.wiz, .2)
    }

    update() {
        LightSystem(this.wiz)
        this.enemies.forEach(e => {
            EnemyMovementSystem(e, this.wiz)
            RenderSystem(e)
        })

        CollisionSystem(this, this.wiz, this.enemies)
        InputSystem(this, this.wiz)
        RenderSystem(this.wiz)

        this.missiles.forEach(missile => {
            if (missile?.isDestroyed) return;

            CollisionSystem(this, missile, this.enemies)
            MissileTrackingSystem(missile, this.enemies)
            LightSystem(missile)
            RenderSystem(missile)
        })
    }
}

const config = {
    type: Phaser.AUTO,
    backgroundColor: '#000',
    width: ScreenWidth,
    height: ScreenHeight,
    scene: Demo,
};

const game = new Phaser.Game(config);
