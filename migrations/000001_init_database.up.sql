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
    description TEXT,                  -- Описание
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

CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,               -- Заголовок новости
    content TEXT,                      -- Текст новости
    date TEXT(10) NOT NULL,            -- Дата (формат YYYY-MM-DD)
    image_path TEXT                    -- Путь к картинке
);

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login TEXT UNIQUE,       
    pass TEXT                     
);

CREATE INDEX index_users_login ON users(login);

CREATE TABLE auth(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login TEXT,   
    refresh_token TEXT,
    FOREIGN KEY (login) REFERENCES users(login) ON DELETE CASCADE
);

CREATE INDEX index_auth_login ON auth(login);