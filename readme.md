cors для post запросов не делал.
Оболочка shell в контейнера упорно отказывается запускать исполняемый файл, 
чтоб проект работал нужно выполнить следующие команды по порядку.

1. docker compose up —build
2. docker rm $(docker ps -aq) 
3. docker run -p 8080:8080 -ti autstore bash 
    -> ./autstore # запустить исполняемый в контейнере уже с помощью оболочки bash 


Посмотреть базу данных: 
    docker exec -i autstore-db psql -U user -d score