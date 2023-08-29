
export enum State {
    Dead = 0,
    Idle,
    Hurt,
    Attacking,
    Walking,
}

export const StateMapping = {
    "dead": 0,
    "idle": 1,
    "hurt": 2,
    "attacking": 3,
    "walk": 4,
}

export default class StateComponent {
    constructor(public states: Array<State>, public current: State) { }

    toString(): string {
        switch (this.current) {
            case State.Dead:
                return "dead"
            case State.Idle:
                return "idle"
            case State.Hurt:
                return "hurt"
            case State.Walking:
                return "walk"
            case State.Attacking:
                return "attack"
            default:
                return "idle"
        }
    }

    setState(state: State) {
        this.current = state;
    }
}
