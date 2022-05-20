package component

type Level struct {
	EntityMap map[int]interface{}
	Entities  []int
}

func (l *Level) Add(entity interface{}) {
	uuid := 0
	l.Entities = append(l.Entities, uuid)
	l.EntityMap[uuid] = entity
}
