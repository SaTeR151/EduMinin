## Созданные ендпониты:
1. /main - основная страница (index.html)
2. /api/news [GET] - получение списка новостей
{
  "data":
  {
    Id    string `json:"id"`
  	Title string `json:"title"`
  	Text  string `json:"text"`
  	Date  string `json:"date"`
  	Image string `json:"image"`
   },
   {
    Id    string `json:"id"`
  	Title string `json:"title"`
  	Text  string `json:"text"`
  	Date  string `json:"date"`
  	Image string `json:"image"`
   },
}
3. /api/news/ [POST] - сохранение новости (защищен)
{
    Id    string `json:"id"`
  	Title string `json:"title"`
  	Text  string `json:"text"`
  	Date  string `json:"date"`
  	Image string `json:"image"`
}
4. /api/news/?id=... [DELETE] - удаление новости по id (защищен)
5. /api/courses [GET] - получение списка курсов
{
  "data":
  {
  	Id            string `json:"id"`
  	Title         string `json:"title"`
  	Description   string `json:"description"`
  	AcademicHours string `json:"academic_hour"`
  	Price         string `json:"price"`
  	Image         string `json:"image"`
   },
   {
  	Id            string `json:"id"`
  	Title         string `json:"title"`
  	Description   string `json:"description"`
  	AcademicHours string `json:"academic_hour"`
  	Price         string `json:"price"`
  	Image         string `json:"image"`
   },
}
6. /api/courses/title [GET] - получение заголовков курса (для записи)
{
  "data":
  {
    Title string `json:"title"`  
  },
  {
    Title string `json:"title"`
  },
}
8. /api/courses/ [POST] - сохранение курса (защищен)
{
	Id            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	AcademicHours string `json:"academic_hour"`
	Price         string `json:"price"`
	Image         string `json:"image"`
}
9. /api/courses/?id=... [DELETE] - удаление курс по id (защищен)

10. /api/reviews [GET] - все отзывы
Тело для отзывов:
{
  Id     string `json:"id"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Text   string `json:"text"`
	Photo  string `json:"photo"`
}
11. /api/reviews/ [POST] (защищен)
12. /api/reviews/ [DELETE] (защищен)

13. /api/events [GET] - всё расписание

14.  /api/reviews/ [POST] (защищен)
15.  /api/reviews/ [DELETE] (защищен)

## Про авторизацию:
При попытке получить доступ к защищенной странице проверяется авторизация. В случае успеха - страница загружается. 
Основные коды ошибок:
- 404 в случае отсутствия тела запроса при регистрации/авторизации и в случае отсутствия куки при получении защищенной страницы
- 401 - если неправильные куки при получении защищенной страницы; если неправильные данные при авторизации
- 500 - если я рукожоп и где-то накосячил

Эндпоинты: 
16. /api/auth/register
17. /api/auth/signup
18. /api/auth/logout (щашишенный)



## Будущая структура каталогов:
- cmd - для запуска приложения
- internal - файлы приложения
- migration - для sql-файлов БД
- web - для статики (http, css, js) 

## URLs для отпраки запросов:
```
http://localhost:8080/
```

## Ссылка пример с AJAX. Тебе нужен только первый участок кода:
```
https://qna.habr.com/q/441823
```
