# relaxdays-hackathon-cc-vol1-7-backend-storage

This project was created in the Relaxdays Code Challenge Vol. 1. See https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite for more information. My participant ID in the challenge was: CC-VOL1-7

## How to run this project

You can get a running version of this code by using:

```bash
git clone https://github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage.git
cd relaxdays-hackathon-cc-vol1-7-backend-storage
docker build -t relaxdays-hackathon-cc-vol1-7-backend-storage .
docker run -p 8080:8080 -it relaxdays-hackathon-cc-vol1-7-backend-storage
```

Afterwards you can access http://127.0.0.1:8080/ which will redirect you to the swagger UI.

By default, data is stored in a SQLite database in memory. To persist data you can provide a DSN:
```bash
docker run -v "$(pwd)/output:/output" -p 8080:8080 -it relaxdays-hackathon-cc-vol1-7-backend-storage -dsn "file:/output/output.sqlite?cache=shared"
```

## OpenAPI definitions

OpenAPI definitions are available as [YAML](https://github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/blob/master/docs/swagger.yaml) and [JSON](https://github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/blob/master/docs/swagger.json).
