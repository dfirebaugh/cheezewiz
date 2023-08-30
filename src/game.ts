import * as Phaser from 'phaser';

import {
    Entity,
    EntityFactory,
    WizData,
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
import World from './world';
import { DropJellyBean, LootMovement, LootSystem } from './system/loot';

const ScreenWidth = 640;
const ScreenHeight = 512;
const SceneWidth = 5088
const SceneHeight = 512

export default class Demo extends Phaser.Scene {
    world: World;

    constructor() {
        super('GameScene');
        this.world = new World(this)
    }

    preload() {
        LoadAssets(this)
    }

    setupCamera() {
        this.cameras.main.startFollow(this.world.wiz.sprite.sprite);
        this.cameras.main.setLerp(0.1, 0.1);
        this.cameras.main.setBounds(0, 0, SceneWidth, ScreenHeight);
    }

    setupMap() {
        const map = this.make.tilemap({ key: 'kitchen' })
        const tileset = map.addTilesetImage('kitchen floor', 'kitchen_tiles')
        map.createLayer('Tile Layer 1', tileset, 0, 0).setPipeline('Light2D');
    }

    create() {
        this.setupMap();
        this.world.wiz = EntityFactory(this.world, WizData)
        this.setupCamera();
        HealthRegen(this.world, this.world.wiz)
        MissileSpawner(this.world, this.world.wiz, 'nachomissile', 20, 20, WeaponType.EXPLOSIVE)
        EnemySpawner(this.world, this.world.wiz, .2)
    }

    update() {
        LightSystem(this.world.wiz)

        CollisionSystem(this.world, this.world.wiz, this.world.enemies)
        InputSystem(this.world, this.world.wiz)
        RenderSystem(this.world.wiz)

        this.world.enemies.forEach(e => {
            EnemyMovementSystem(e, this.world.wiz)
            RenderSystem(e)

            if (e.isDestroyed) {
                DropJellyBean(this.world, e)
            }
        })
        this.world.enemies = this.world.enemies.filter(e => !e.isDestroyed)

        this.world.missiles.forEach(missile => {
            if (missile?.isDestroyed) return;

            CollisionSystem(this.world, missile, this.world.enemies)
            MissileTrackingSystem(missile, this.world.enemies)
            LightSystem(missile)
            RenderSystem(missile)
        })

        this.world.loot.forEach(loot => {
            LootSystem(this.world, loot)
            LightSystem(loot)
            RenderSystem(loot)
        })
        this.world.loot = this.world.loot.filter(e => !e.isDestroyed)
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
