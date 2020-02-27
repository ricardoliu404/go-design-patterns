package prototype

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtCloner() ShirtCloner {
	shirtCache := new(ShirtCache)
	return shirtCache
}
