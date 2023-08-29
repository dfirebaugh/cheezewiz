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
} from "../component";
import LightComponent from "../component/light";
import { StateMapping } from "../component/state";

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

    constructor(tag: string) {
        this.tag = tag;
        this.id = new IDComponent();
    }
}

export function EntityFactory(scene: Phaser.Scene, fileData: any): Entity {
    const entity = new Entity(fileData.tag)

    if (fileData.position?.X && fileData.position?.Y)
        entity.position = new PositionComponent(fileData.position?.X, fileData.position?.Y);

    if (fileData.velocity?.VX, fileData.velocity?.VY)
        entity.velocity = new VelocityComponent(fileData.velocity?.VX, fileData.velocity?.VY);

    if (fileData.rigidBody?.X, fileData.rigidBody?.Y, fileData.rigidBody?.W, fileData.rigidBody?.H)
        entity.rigidBody = new RigidBodyComponent(fileData.rigidBody?.X, fileData.rigidBody?.Y, fileData.rigidBody?.W, fileData.rigidBody?.H);

    if (fileData.health?.max, fileData.health?.hp) {
        entity.health = new HealthComponent(scene, entity, fileData.health?.max, fileData.health?.hp, fileData.health?.regenRate);
        entity.health.invulnerabilityDuration = fileData.health?.invulnerabilityDuration
    }

    if (fileData.size?.height, fileData.size?.width)
        entity.size = new SizeComponent(fileData.size?.height, fileData.size?.width);

    if (fileData.state?.states, fileData.state?.current)
        entity.state = new StateComponent(fileData.state?.states, fileData.state?.current)

    if (fileData.speed)
        entity.speed = new SpeedComponent(fileData.speed)

    if (fileData.light)
        entity.light = new LightComponent(scene, fileData.light?.intensity, fileData.light?.color, fileData.light?.radius)

    if (fileData.animations) {
        const states = fileData.animations?.map(e => StateMapping[e.name])
        if (states.length > 0)
            entity.state = new StateComponent(states, StateMapping[states[0]])

        entity.sprite = new SpriteComponent(scene, fileData.tag)

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

    return entity
}
