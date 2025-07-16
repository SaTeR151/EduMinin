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

type CoursesTitle struct {
	Title string `json:"title"`
}

type News struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Text        string `json:"text"`
	Date        string `json:"date"`
	Image       string `json:"image"`
}

type Reviews struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Text   string `json:"text"`
	Photo  string `json:"photo"`
}

type Events struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Date  string `json:"date"`
	Image string `json:"image"`
}

type UserData struct {
	Login string `json:"username"`
	Pass  string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
}

type LK struct {
	Fio     string   `json:"fio"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Photo   string   `json:"photo"`
	Courses []Course `json:"courses"`
}

type SaveUserCourse struct {
	Id string `json:"id"`
}

type ChangingData struct {
	Changing string `json:"changing"`
}
