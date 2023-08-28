
class IDManager {
    static count: number = 0;

    static nextID(): number {
        this.count += 1;
        return this.count
    }
}

export default class IDComponent {
    public id: number
    constructor() {
        this.id = IDManager.nextID()
    }
}
