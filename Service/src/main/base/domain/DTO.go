package domain

type ServiceList struct {
	Name      string
	GruopList []*Gruop
}

type Gruop struct {
	Name        string
	ServiceList []*Service
}
