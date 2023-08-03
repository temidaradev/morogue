package game

import (
	"encoding/json"
	"errors"

	"github.com/kettek/morogue/id"
)

// ObjectType is a key that represents an object's type. This is used
// for safely marshalling/unmarshalling the Object interface.
type ObjectType string

// Object is an interface intended for location objects. This includes
// players, items, and enemies.
type Object interface {
	Type() ObjectType
	SetWID(id.WID) // SetWID should only be called by the world.
	GetWID() id.WID
	SetPosition(Position)
	GetPosition() Position
}

func CreateObjectFromArchetype(a Archetype) Object {
	switch a.(type) {
	case CharacterArchetype:
		return &Character{
			Archetype: a.GetID(),
		}
	case WeaponArchetype:
		return &Weapon{
			Archetype: a.GetID(),
		}
	case ArmorArchetype:
		return &Armor{
			Archetype: a.GetID(),
		}
	}
	return nil
}

// ObjectWrapper wraps an Object interface for json marshal and unmarshal.
type ObjectWrapper struct {
	Type ObjectType      `json:"t"`
	Data json.RawMessage `json:"d"`
}

// Object returns the wrapped Object. To add additional types, they must be
// handled here.
func (ow ObjectWrapper) Object() (Object, error) {
	switch ow.Type {
	case (Item{}).Type():
		var o *Item
		if err := json.Unmarshal(ow.Data, &o); err != nil {
			return nil, err
		}
		return o, nil
	case (Character{}).Type():
		var c *Character
		if err := json.Unmarshal(ow.Data, &c); err != nil {
			return nil, err
		}
		return c, nil
	case (Weapon{}).Type():
		var w *Weapon
		if err := json.Unmarshal(ow.Data, &w); err != nil {
			return nil, err
		}
		return w, nil
	case (Armor{}).Type():
		var a *Armor
		if err := json.Unmarshal(ow.Data, &a); err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, errors.New("unknown object type: " + string(ow.Type))
}
