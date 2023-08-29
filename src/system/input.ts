import * as Phaser from 'phaser';
import { Entity } from "../entities/entity";
import { State } from '../component/state';

export default function InputSystem(scene: Phaser.Scene, entity: Entity) {
    const keys = {
        up: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W),
        left: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A),
        down: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S),
        right: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D),
        space: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.SPACE),
        shift: scene.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.SHIFT)
    };

    if (entity.health?.current <= 0) {
        entity.state.setState(State.Dead);
        return
    }

    const speed = 1
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
}
