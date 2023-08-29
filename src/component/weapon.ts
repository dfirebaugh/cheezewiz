
export enum WeaponType {
    BLUDGEONING = 0,
    FIRE,
    FROST,
    SHADOW,
    FORCE,
    ACID,
    PIERCING,
    SLASHING,
    EXPLOSIVE,
}

export default class WeaponComponent {
    public distanceTraveled: number = 0;
    constructor(public label: string, public speed: number, public power: number, public weaponType: WeaponType) { }
}
