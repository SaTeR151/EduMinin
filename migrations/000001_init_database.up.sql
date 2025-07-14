CREATE TABLE courses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,               -- Название курса
    description TEXT,                  -- Описание
    academic_hours INTEGER,            -- Кол-во академ часов
    price INTEGER,                        -- Стоимость
    image TEXT                   -- Путь к картинке
);

-- Таблица новостей
CREATE TABLE news (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,               -- Заголовок новости
    content TEXT,                      -- Текст новости
    date TEXT(10) NOT NULL,            -- Дата (формат YYYY-MM-DD)
    image_path TEXT                    -- Путь к картинке
);

-- Таблица отзывов
CREATE TABLE reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_name TEXT NOT NULL,         -- Имя того кто оставил
    date TEXT(10) NOT NULL,            -- Дата (формат YYYY-MM-DD)
    review_text TEXT NOT NULL,         -- Текст отзыва
    photo_path TEXT                    -- Путь к фото
);