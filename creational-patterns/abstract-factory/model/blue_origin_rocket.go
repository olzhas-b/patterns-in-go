package model

type BlueOriginRocket struct {
	Rocket
}

func (blueOriR *BlueOriginRocket) NewBlueOriginRocket() *BlueOriginRocket {
	return &BlueOriginRocket{}
}

// here we can overide the function GetName which implemented by Rocket
func (blueOriR *BlueOriginRocket) GetName() string {
	return blueOriR.Name
}
