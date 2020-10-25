type ParkingSystem struct {
    Count [4]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    init:=[4]int{0,big,medium,small}
    return ParkingSystem{
        Count:init,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if this.Count[carType]<=0{
        return false
    }else{
        this.Count[carType]--
    }
    return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */