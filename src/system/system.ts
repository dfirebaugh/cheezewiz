import { Entity } from "../entities/entity";

export abstract class System {
    abstract update(entity: Entity): void;
}
