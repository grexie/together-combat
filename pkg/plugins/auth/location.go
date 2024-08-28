package auth

type Country interface {
	ISO() *string
	Name() *string
}

type CountrySubdivision interface {
	ISO() *string
	Name() *string
}

type Location interface {
	Country() Country
	Subdivisions() *[]CountrySubdivision
	City() *string
	Latitude() *float64
	Longitude() *float64
}