package coal

type MinerInfo struct {
	Id 				int
	Cost 			int
	Class           string
	Energy_remained *int
}

func NewMinerInfo(id int, cost int, class string, energy *int) MinerInfo {
	return MinerInfo{
		Id: id,
		Cost: cost,
		Class: class,
		Energy_remained: energy,
	}
}

func (mi *MinerInfo) GetId() int {
	return mi.Id;
}

func (mi *MinerInfo) GetCost() int {
	return mi.Cost;
}

func (mi *MinerInfo) GetClass() string {
	return mi.Class;
}

func (mi *MinerInfo) GetEnergy() *int {
	return mi.Energy_remained;
}