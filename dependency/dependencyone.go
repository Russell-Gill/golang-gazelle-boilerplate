package done

type BasicStruct struct {
	Field      string
	OtherField string
}

func (b *BasicStruct) BasicStructOperation() string {
	return b.Field
}

type Dependency interface {
	BasicStructOperation() string
}
