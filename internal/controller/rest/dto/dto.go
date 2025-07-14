package dto

type Error struct {
	Error string `json:"error"`
}

type Data struct {
	Data interface{} `json:"data"`
}

type Course struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	AcademicHours string `json:"academic_hour"`
	Price         string `json:"price"`
	Image         string `json:"image"`
}

type News struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Date  string `json:"date"`
	Image string `json:"image"`
}

type Reviews struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Text   string `json:"text"`
	Photo  string `json:"photo"`
}
