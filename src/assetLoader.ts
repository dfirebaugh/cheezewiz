
export function LoadAssets(scene: Phaser.Scene) {
    scene.load.spritesheet('cheezewiz', 'assets/cheezewiz.png', { frameHeight: 32, frameWidth: 32 })
    scene.load.spritesheet('nachomissile', 'assets/nachomissile.png', { frameHeight: 32, frameWidth: 32 })
    scene.load.spritesheet('cheezewiz-damaged', 'assets/cheezewiz-damaged.png', { frameHeight: 32, frameWidth: 32 })
    scene.load.spritesheet('cheezewiz-death', 'assets/cheezewiz-death.png', { frameHeight: 32, frameWidth: 32 })
    scene.load.spritesheet('radishred', 'assets/radishred.png', { frameHeight: 32, frameWidth: 32 })

    scene.load.image('kitchen_tiles', 'assets/kitchen_floor.png')
    scene.load.tilemapTiledJSON('kitchen', 'assets/kitchen1.json')
}
