package common

var Department = map[int]string{
	1: "Information Technology",
	2: "Business",
	3: "Graphic and Digital Design",
	4: "Marketing",
	5: "Event Management",
}

type department struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetListDepartment() []department {
	return []department{
		{
			1, "Information Technology",
		},
		{
			2, "Business",
		},
		{
			3, "Graphic and Digital Design",
		},
		{
			4, "Marketing",
		},
		{
			5, "Event Management",
		},
	}
}
