import * as Phaser from 'phaser';
import { Entity } from "../entities";
import World from '../world';

export class XPBar {
    barWidth: number;
    barHeight = 10;
    screen: { width: number, height: number };
    gradientTextureKey: string = 'gradientTexture';
    xpBarSprite: Phaser.GameObjects.Sprite;
    xpBarShadowSprite: Phaser.GameObjects.Sprite;
    xpBarBackground: Phaser.GameObjects.Rectangle;
    levelText: Phaser.GameObjects.Text;

    constructor(world: World, screenWidth: number, screenHeight: number) {
        this.screen = { width: screenWidth, height: screenHeight };
        this.barWidth = this.screen.width * 0.8;

        const canvas = document.createElement('canvas');
        canvas.width = this.barWidth;
        canvas.height = this.barHeight;
        const ctx = canvas.getContext('2d');
        const gradient = ctx.createLinearGradient(0, 0, this.barWidth, 0);
        gradient.addColorStop(0, '#FF0000');
        gradient.addColorStop(0.15, '#FF7F00');
        gradient.addColorStop(0.3, '#FFFF00');
        gradient.addColorStop(0.5, '#7FFF00');
        gradient.addColorStop(0.65, '#00FF00');
        gradient.addColorStop(0.8, '#00FF7F');
        gradient.addColorStop(1, '#00FFFF');
        ctx.fillStyle = gradient;
        ctx.fillRect(0, 0, this.barWidth, this.barHeight);
        world.scene.textures.addCanvas(this.gradientTextureKey, canvas);

        this.xpBarBackground = world.scene.add.rectangle((this.screen.width - this.barWidth) / 2, 10, this.barWidth, this.barHeight, 0x000000);
        this.xpBarBackground.setOrigin(0, 0);
        this.xpBarBackground.setScrollFactor(0);
        this.xpBarBackground.setStrokeStyle(2, 0xFFFFFF);

        this.xpBarSprite = world.scene.add.sprite((this.screen.width - this.barWidth) / 2, 10, this.gradientTextureKey);
        this.xpBarSprite.setOrigin(0, 0);
        this.xpBarSprite.setScrollFactor(0);

        this.xpBarShadowSprite = world.scene.add.sprite((this.screen.width - this.barWidth) / 2 + 2, 12, this.gradientTextureKey);
        this.xpBarShadowSprite.setOrigin(0, 0);
        this.xpBarShadowSprite.setScrollFactor(0);
        this.xpBarShadowSprite.setTint(0xD3D3D3);
        this.xpBarShadowSprite.setAlpha(0.5);
        this.xpBarShadowSprite.setDepth(999);

        this.levelText = world.scene.add.text((this.screen.width - this.barWidth) / 2 - 3, 10, `Level: 0`, {
            fontSize: '10px',
            color: '#fff',
            shadow: {
                offsetX: 2,
                offsetY: 2,
                color: '#D3D3D3',
                blur: 2,
                fill: true
            }
        });
        this.levelText.setOrigin(1, 0);
        this.levelText.setScrollFactor(0);

        this.xpBarShadowSprite.setDepth(1000);
        this.xpBarBackground.setDepth(1001);
        this.xpBarSprite.setDepth(1002);
        this.levelText.setDepth(1003);
    }

    draw(entity: Entity) {
        const cropWidth = this.barWidth * entity.xp.calculateXPPercentage();
        this.xpBarSprite.setCrop(0, 0, cropWidth, this.barHeight);
        this.xpBarShadowSprite.setCrop(0, 0, cropWidth, this.barHeight);
        this.levelText.setText(`Level: ${entity.xp.calculateLevel()}`);
    }
}

export default class XPComponent {
    graphics?: Phaser.GameObjects.Graphics;
    world?: World;
    bar: XPBar;
    public level: number;

    constructor(world: World, public xp: number) {
        this.bar = new XPBar(world, 640, 512)
        this.world = world;
        this.graphics = new Phaser.GameObjects.Graphics(world.scene);
    }

    calculateXPPercentage(): number {
        return (this.xp % 100 / 100);
    }

    calculateLevel(): number {
        return Math.floor(this.xp / 100);
    }
}
