package SquadCache

type ByteView struct {
	b []byte
}

func (b ByteView) Len() int {
	return len(b.b)
}

func (b ByteView) ByteSlice() []byte {
	c := make([]byte, b.Len())
	copy(c, b.b)
	return c
}