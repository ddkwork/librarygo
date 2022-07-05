package notes

import (
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

const (
	countKey       = "notecount"
	noteKey        = "note%d"
	noteDeletedKey = "note%ddeleted"
)

type note struct {
	content binding.String
	deleted binding.Bool
}

func (n *note) title() binding.String {
	return newTitleString(n.content)
}

type Notelist struct {
	all  []*note
	Pref fyne.Preferences
}

func (l *Notelist) add() *note {
	key := fmt.Sprintf(noteKey, len(l.all))
	deleteKey := fmt.Sprintf(noteDeletedKey, len(l.all))
	n := &note{
		binding.BindPreferenceString(key, l.Pref),
		binding.BindPreferenceBool(deleteKey, l.Pref),
	}
	l.all = append([]*note{n}, l.all...)
	l.save()
	return n
}

func (l *Notelist) delete(n *note) {
	n.deleted.Set(true)
}

func (l *Notelist) Load() {
	l.all = nil
	count := l.Pref.Int(countKey)
	if count == 0 {
		return
	}

	for i := count - 1; i >= 0; i-- {
		key := fmt.Sprintf(noteKey, i)
		deleteKey := fmt.Sprintf(noteDeletedKey, i)
		content := binding.BindPreferenceString(key, l.Pref)
		deleted := binding.BindPreferenceBool(deleteKey, l.Pref)
		l.all = append(l.all, &note{content, deleted})
	}
}

func (l *Notelist) notes() []*note {
	var visible []*note
	for _, n := range l.all {
		if del, _ := n.deleted.Get(); del {
			continue
		}
		visible = append(visible, n)
	}
	return visible
}

func (l *Notelist) save() {
	l.Pref.SetInt(countKey, len(l.all))
}

type titleString struct {
	binding.String
}

func (t *titleString) Get() (string, error) {
	content, err := t.String.Get()
	if err != nil {
		return "Error", err
	}

	if content == "" {
		return "Untitled", nil
	}

	return strings.SplitN(content, "\n", 2)[0], nil
}

func (t *titleString) Set(string) error {
	return errors.New("cannot set content from title")
}

func newTitleString(in binding.String) binding.String {
	return &titleString{in}
}
