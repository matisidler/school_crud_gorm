package handlers

type StudyPlan struct {
	Driver     string `json:"driver"`
	FirstYear  string `json:"fy"`
	SecondYear string `json:"sy"`
	ThirdYear  string `json:"ty"`
}

type Subject struct {
	Driver string `json:"driver"`
	Name   string `json:"name"`
}

type Career struct {
	Driver string `json:"driver"`
	Name   string `json:"name"`
	Spid   int    `json:"spid"`
}

type Teacher struct {
	Driver    string  `json:"driver"`
	Name      string  `json:"name"`
	Salary    float32 `json:"salary"`
	SubjectId int     `json:"subid"`
	CareerId  int     `json:"carid"`
	Mail      string  `json:"mail"`
	Phone     string  `json:"phone"`
}

type Student struct {
	Driver   string `json:"driver"`
	Name     string `json:"name"`
	CareerId int    `json:"carid"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
}
