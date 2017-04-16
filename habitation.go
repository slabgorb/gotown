package townomatic

type Habitation struct {
	Residents []*Being
}

type Area struct {
	Habitation
	Name string
}

type Dwelling struct {
	Habitation
	Location *Area
}

func (h *Habitation) Resident(b *Being) (int, bool) {
	found := false
	index := 0
	for i, r := range h.Residents {
		if r == b {
			found = true
			index = i
			break
		}
	}
	return index, found
}

func (h *Habitation) Add(b *Being) (ok bool) {
	_, found := h.Resident(b)
	if found {
		return false
	}
	h.Residents = append(h.Residents, b)
	return true
}

func (h *Habitation) Remove(b *Being) (ok bool) {
	index, found := h.Resident(b)
	if !found {
		return false
	}
	h.Residents = append(h.Residents[:index], h.Residents[index+1:]...)
	return true
}

func (h *Habitation) Population() int {
	return len(h.Residents)
}
