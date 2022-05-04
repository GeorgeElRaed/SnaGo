package entitymanager

import (
	"github.com/faiface/pixel/pixelgl"
)

type Entity interface {
	Init()
	Update()
	Render(win *pixelgl.Window)
}

type EntityManager struct {
	Entities []Entity
}

func Init() *EntityManager {
	return &EntityManager{Entities: make([]Entity, 0)}
}

func (em *EntityManager) forEach(action func(Entity)) {
	for _, e := range em.Entities {
		action(e)
	}
}

func (em *EntityManager) Add(entity Entity) {
	em.Entities = append(em.Entities, entity)
}

func (em *EntityManager) Init() {
	em.forEach(Entity.Init)
}

func (em *EntityManager) Update() {
	em.forEach(Entity.Update)
}

func (em *EntityManager) Render(win *pixelgl.Window) {
	em.forEach(func(e Entity) { e.Render(win) })
}
