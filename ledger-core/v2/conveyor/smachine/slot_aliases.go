// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package smachine

/* ------- Slot-dependant aliases and mappings ------------- */

type slotIdKey SlotID

type slotAliasesValue struct {
	//owner *Slot
	keys []interface{}
}

//type slotAliasesMap struct {
//	//owner *Slot
//	keys map[interface{}
//}

// ONLY to be used by a holder of a slot
func (s *Slot) registerBoundAlias(k, v interface{}) bool {
	if k == nil {
		panic("illegal value")
	}

	m := &s.machine.localRegistry
	if _, loaded := m.LoadOrStore(k, v); loaded {
		return false
	}

	var key interface{} = slotIdKey(s.GetSlotID())

	switch isa, ok := m.Load(key); {
	case !ok:
		isa, _ = m.LoadOrStore(key, &slotAliasesValue{ /* owner:s */ })
		fallthrough
	default:
		sa := isa.(*slotAliasesValue)
		sa.keys = append(sa.keys, k)
		s.slotFlags |= slotHasAliases
	}

	if sar := s.machine.config.SlotAliasRegistry; sar != nil {
		if ga, ok := k.(globalAliasKey); ok {
			return sar.PublishAlias(ga.key, v.(SlotAliasValue))
		}
	}

	return true
}

// ONLY to be used by a holder of a slot
func (s *Slot) unregisterBoundAlias(k interface{}) bool {
	if k == nil {
		panic("illegal value")
	}

	switch keyExists, wasUnpublished, _ := s.machine.unpublishUnbound(k); {
	case !keyExists:
		return false
	case wasUnpublished:
		return true
	}
	return s.machine._unregisterSlotBoundAlias(s.GetSlotID(), k)
}

// ONLY to be used by a holder of a slot
func (m *SlotMachine) unregisterBoundAliases(id SlotID) {
	mm := &m.localRegistry // SAFE for concurrent use
	var key interface{} = slotIdKey(id)

	if isa, ok := mm.Load(key); ok {
		sa := isa.(*slotAliasesValue)
		mm.Delete(key)

		sar := m.config.SlotAliasRegistry
		for _, k := range sa.keys {
			mm.Delete(k)

			if sar != nil {
				if ga, ok := k.(globalAliasKey); ok {
					sar.UnpublishAlias(ga.key)
				}
			}
		}
	}
}

// ONLY to be used by a holder of a slot
func (m *SlotMachine) _unregisterSlotBoundAlias(slotID SlotID, k interface{}) bool {
	var key interface{} = slotIdKey(slotID)

	if isa, loaded := m.localRegistry.Load(key); loaded {
		sa := isa.(*slotAliasesValue)

		for i, kk := range sa.keys {
			if k == kk {
				m.localRegistry.Delete(k)
				if sar := m.config.SlotAliasRegistry; sar != nil {
					if ga, ok := k.(globalAliasKey); ok {
						sar.UnpublishAlias(ga.key)
					}
				}

				switch last := len(sa.keys) - 1; {
				case last == 0:
					m.localRegistry.Delete(key)
				case i < last:
					copy(sa.keys[i:], sa.keys[i+1:])
					fallthrough
				default:
					sa.keys = sa.keys[:last]
				}
				return true
			}
		}
	}
	return false
}
