
export enum State {
    Dead = 0,
    Idle,
    Walking,
    Attacking
}

export const StateMapping = {
    "dead": 0,
    "idle": 1,
    "walking": 2,
    "attacking": 3,
}

export default class StateComponent {
    constructor(public states: Array<State>, public current: State) { }

    toString(): string {
        switch (this.current) {
            case State.Dead:
                return "dead"
            case State.Idle:
                return "idle"
            case State.Walking:
                return "walking"
            case State.Attacking:
                return "attacking"
            default:
                return "idle"
        }
    }
}
