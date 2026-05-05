package coal

import "errors"

func NewLittleMiner() *Miner {
    return NewMiner("little", 5, 30, 1, 3, 0)
}

func NewNormalMiner() *Miner {
    return NewMiner("normal", 50, 45, 3, 2, 0)
}

func NewBigMiner() *Miner {
    return NewMiner("big", 450, 60, 10, 1, 3)
}

func NewMinersType(class string) (*Miner, error) {
	switch class {
	case classLittleMiner:
		return NewLittleMiner(), nil
	case classNormalMiner:
		return NewNormalMiner(), nil
	case classBigMiner:
		return NewBigMiner(), nil
	default:
		return &Miner{}, errors.New("there is no such miner");
	}
}