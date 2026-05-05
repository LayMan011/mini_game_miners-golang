package equipment

var (
	classPickaxes = "pickaxes"
	classVentilation = "ventilation"
	classTrolleys    = "trolleys"
)

var equipments = NewEquipments(); 

var pickaxes = NewEquipment(3000);
var ventilation = NewEquipment(15000);
var trolleys = NewEquipment(50000);