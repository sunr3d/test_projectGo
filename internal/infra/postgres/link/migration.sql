CREATE TABLE IF NOT EXISTS links (
    id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    link TEXT NOT NULL, -- Использую TEXT т.к. у него ограничение на 1ГБ (ссылка то может быть огромная)
    fake_link TEXT NOT NULL UNIQUE, -- Здесь по идее можно сделать ограничение через VARCHAR т.к. мы сами будем ссылку сокращать нашим сервисом
    erase_time TIMESTAMP NOT NULL
);