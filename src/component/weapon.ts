
export enum WeaponType {
    BLUDGEONING = 0,
    FIRE,
    FROST,
    SHADOW,
    FORCE,
    ACID,
    PIERCING,
    SLASHING
}

export default class WeaponComponent {
    constructor(public speed: number, public power: number, public weaponType: WeaponType) {}
}
