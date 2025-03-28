package structx

type Pair[K, V any] struct {
	Key   K
	Value V
}

type Triple[X, Y, Z any] struct {
	First  X
	Second Y
	Third  Z
}

type NoCopy struct{}

func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
