import * as Phaser from 'phaser';
import { Entity } from "../entities/entity";
import { State } from '../component/state';
import { XPComponent } from '../component';
import World from '../world';

export default function InputSystem(world: World, entity: Entity) {
    const keys = {
        up: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W),
        left: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A),
        down: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S),
        right: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D),
        space: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.SPACE),
        shift: world.scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.SHIFT)
    };

    if (entity.health?.current <= 0) {
        entity.state.setState(State.Dead);
        return
    }

    const speed = 1
    if (keys.space.isDown) {
        if (!entity.xp) {
            entity.xp = new XPComponent(world, 1)
        }
        entity.xp.xp++;
    }

    if (keys.left.isDown) {
        entity.position.X -= speed
        entity.state.setState(State.Walking);


        if (entity.sprite?.sprite?.flipX) {
            entity.sprite?.sprite?.setFlipX(false)
        }
    }
    if (keys.right.isDown) {
        entity.position.X += speed
        entity.state.setState(State.Walking);
        if (!entity.sprite?.sprite?.flipX) {
            entity.sprite?.sprite?.setFlipX(true)
        }
    }
    if (keys.up.isDown) {
        entity.state.setState(State.Walking);
        entity.position.Y -= speed
    }
    if (keys.down.isDown) {
        entity.state.setState(State.Walking);
        entity.position.Y += speed
    }

    if (!keys.up.isDown && !keys.down.isDown && !keys.left.isDown && !keys.right.isDown && entity.state.current !== State.Dead) {
        entity.state.setState(State.Idle);
    }
}
