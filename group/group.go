package group

type Element interface {
  Equiv(o Element) bool
}

type Member struct {
  Next *Member
  Value Element
}

type Represent struct {
  Next *Represent
  First *Member
}

type Group struct {
  First *Represent
}

func (m *Member) equiv(o *Member) bool {
  return m.Value.Equiv(o.Value)
}

func (m *Member) add(o *Member) {
  if m.Next != nil {
    m.Next.add(o)
  } else {
    m.Next = o
  }
}

func (r *Represent) add(o *Member) {
  if r.First.equiv(o) {
    r.First.add(o)
  } else if (r.Next != nil) {
    r.Next.add(o)
  } else {
    r.Next = &Represent{First: o}
  }
}

func (g *Group) Add(e Element) {
   m := &Member{Value: e}
   if g.First != nil {
     g.First.add(m)
   } else {
     g.First = &Represent{First: m}
   }
}
