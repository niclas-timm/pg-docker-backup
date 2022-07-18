# Postgres Docker backup

A simple tool for creating backups from Postgres Docker containers and upload them to AWS S3. Written in Go.

# Features

- Create gzipped backup files from a Docker Postgres container.
- Upload the backup file to S3.
- Delete all but the last 7 backup files in S3 to save storage space.
- Get notified via email and/or Slack when something goes wrong.

# Prerequisits

- Docker installed
- Go installed

# Quick start

Clone the repo

```
https://github.com/NiclasTimmeDev/pg-docker-backup.git;

```

Run setup script

```
chmod 755 ./sripts/setup.sh;
./scripts/setup.sh

```

Populate variables in `.env` (will be created by `./scripts/setup.sh`). Especially the `AWS` variables are important.

Compile the code

```
./scripts/compile.sh

```

Create a backup & store it in S3

```
db-backup backup --container="<POSTGRES_CONTAINER_NAME>"  --username="<POSTGRES_USERNAME>" --database="<DATABASE_NAME>"
```

Import the latest dump from S3

```
db-backup import --container="<POSTGRES_CONTAINER_NAME>"  --username="<POSTGRES_USERNAME>" --database="<DATABASE_NAME>"
```

# Environment variables

In the quickstart guide, we addressed that you have to populate the environment variables before compiling and executing the program.

The most important part is that you populate the AWS values. They will be required to access and read/write to AWS S3.

If you want to be notified about any errors via email or Slack, you must also configure the environment variables for these. Also, you must set `EMAILS_ENABLED` and/or `SLACK_ENABLED` to `true` in order to notifications being sent.

Now we are ready to go!

# Notifications

if you populated the variables in the `.env` file (see section "Environment variables") you will be notified via email and/or Slack about errors during the backup process. That way you can rest assured that the backups were successful if you receive no notifications and can act right away if the backups fail.

# Under the hood

This little backup tool tries to use as few dependencies as possible. The major dependencies are:

- [cobra](https://github.com/spf13/cobra) for creating the command line arguments.
- [AWS SDK V2](https://github.com/aws/aws-sdk-go-v2) for interacting with AWS S3.
- [Go Dotenv](github.com/joho/godotenv) for parsing the `.env` file

The backups are created by executing `docker exec` shell commands via Go. You can see this under `db/dump.go`, for example.
