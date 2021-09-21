FROM postgres:latest

ENV POSTGRES_PASSWORD=Temp1234!1

COPY create_db.sql /docker-entrypoint-initdb.d/create_db.sql