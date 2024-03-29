import {
    PositionComponent,
    SpriteComponent,
    VelocityComponent,
    RigidBodyComponent,
    HealthComponent,
    SizeComponent,
    StateComponent,
    SpeedComponent,
    IDComponent,
    HealthBar,
    DefenseComponent,
    Condition,
    WeaponComponent,
    XPComponent,
} from "../component";
import LightComponent from "../component/light";
import { StateMapping } from "../component/state";
import World from "../world";
// import WeaponComponent from "../component/weapon";

export class Entity {
    id: IDComponent;
    tag: string;
    position?: PositionComponent;
    velocity?: VelocityComponent;
    rigidBody?: RigidBodyComponent;
    health?: HealthComponent;
    size?: SizeComponent;
    state?: StateComponent;
    sprite?: SpriteComponent;
    speed?: SpeedComponent;
    healthBar?: HealthBar;
    light?: LightComponent;
    defense?: DefenseComponent;
    conditions?: Array<Condition>;
    weapon?: WeaponComponent;
    // weapons?: Array<WeaponComponent>;
    destroyable?: boolean;
    isDestroyed: boolean = false;
    rotation?: number;
    xp?: XPComponent;
    dropsLoot?: boolean;

    constructor(tag: string) {
        this.tag = tag;
        this.id = new IDComponent();
    }

    destroy() {
        this.isDestroyed = true
    }
}

export function EntityFactory(world: World, fileData: any): Entity {
    const entity = new Entity(fileData.tag)

    if (fileData.position?.X && fileData.position?.Y)
        entity.position = new PositionComponent(fileData.position?.X, fileData.position?.Y);

    if (fileData.velocity?.VX, fileData.velocity?.VY)
        entity.velocity = new VelocityComponent(fileData.velocity?.VX, fileData.velocity?.VY);

    if (fileData.rigidBody?.X, fileData.rigidBody?.Y, fileData.rigidBody?.W, fileData.rigidBody?.H)
        entity.rigidBody = new RigidBodyComponent(fileData.rigidBody?.X, fileData.rigidBody?.Y, fileData.rigidBody?.W, fileData.rigidBody?.H);

    if (fileData.health?.max, fileData.health?.hp) {
        entity.health = new HealthComponent(world, entity, fileData.health?.max, fileData.health?.hp, fileData.health?.regenRate);
        entity.health.invulnerabilityDuration = fileData.health?.invulnerabilityDuration
        entity.health.disableHealthBar = fileData.health?.disableHealthBar
    }

    if (fileData.size?.height, fileData.size?.width)
        entity.size = new SizeComponent(fileData.size?.height, fileData.size?.width);

    if (fileData.state?.states, fileData.state?.current)
        entity.state = new StateComponent(fileData.state?.states, fileData.state?.current)

    if (fileData.speed)
        entity.speed = new SpeedComponent(fileData.speed)

    if (fileData.light)
        entity.light = new LightComponent(world, fileData.light?.intensity, fileData.light?.color, fileData.light?.radius)

    if (fileData.animations) {
        const states = fileData.animations?.map(e => StateMapping[e.name])
        if (states.length > 0)
            entity.state = new StateComponent(states, StateMapping[states[0]])

        entity.sprite = new SpriteComponent(world, fileData.tag)

        fileData.animations?.forEach(e => {
            if (entity.sprite.sprite.anims.exists(e.name)) return;

            entity.sprite.addAnimation({
                key: e.name,
                frames: entity.sprite.sprite.anims.generateFrameNumbers(e.textureName, {
                    start: 0,
                    end: e.frameCount - 1,
                }),
                frameRate: 10,
                repeat: -1,
            })
        });
    }

    if (fileData.weapon) {
        entity.weapon = new WeaponComponent(fileData.weapon?.label, fileData.weapon?.speed, fileData.weapon?.power, fileData.weapon?.type);
    }

    if (fileData.defense) {
        entity.defense = new DefenseComponent(fileData.defense);
    }

    if (fileData.destroyable) {
        entity.destroyable = fileData.destroyable
    }

    if (fileData.xp) {
        entity.xp = new XPComponent(world, fileData.xp)
    }

    return entity
}
