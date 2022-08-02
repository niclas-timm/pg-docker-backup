# Postgres Docker backup

A simple tool for creating backups from Postgres Docker containers and upload them to AWS S3. Written in Go.

# Features

- Create gzipped backup files from a Docker Postgres container.
- Upload the backup file to S3.
- Configuration via `yml` files.
- Only keep a predefined number of backups in the cloud. Delete all excess backups
- Get notified via email and/or Slack when something goes wrong.
- Import the latest backup from the cloud into your database with a single command.

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

Add the settings of your choice in `config.yml`.

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

Now we are ready to go!

# Configuration

If you followed the quick start guide and executed the `setup.sh` script, you should have two yml files in your project:

- config.yml
- config.default.yml

If you don't have `config.yml` just duplicate `config.defautl.yml` and name it `config.yml`.

`config.default.yml` contains some default configurations that you don't need to worry about. But please also don't delete it ;).

You can use `config.yml` to override the default config. For example, you can define the maximum number of backups will be stored in S3 and what the subdirectory these backups will be stored in (e.g., `backups/`). You can also define if you want to be notified about errors via slack and/or email (see below for details).

# Notifications

If you want to be notified about any errors via email or Slack, you must do 2 things:

1. Go to `config.yml` and set `notificatios.slack.enabled` and/or `notifications.email.enabled` to true
2. Populate the corresponing environment variables to your `.env` file.

Now you will be notified via email and/or Slack about errors during the backup process. That way you can rest assured that the backups were successful if you receive no notifications and can act right away if the backups fail.

# Backing up a single table

Sometimes you might not want to create a backup of the whole database, but only a single table. You can easily accomplish this by adding the `--table` flag to your import command. For example

```
db-backup backup --container="<POSTGRES_CONTAINER_NAME>"  --username="<POSTGRES_USERNAME>" --database="<DATABASE_NAME>" --table="public.users"
```

will create a backup of the `users` table only.

**Attention:** Notice the "public" keyword before the acutal name of the table. This is required.

# Import a database

Creating backups in useless if you can' import these backups. But don't worry, we've got you covered. All you have to do is execute the following command:

```
db-backup import --container="<POSTGRES_CONTAINER_NAME>"  --username="<POSTGRES_USERNAME>" --database="<DATABASE_NAME>"
```

This will by default donwload and import the latest dump into your database. If you want to import another database (not the latest one), you can do the following

```
db-backup import --container="<POSTGRES_CONTAINER_NAME>"  --username="<POSTGRES_USERNAME>" --database="<DATABASE_NAME>" --filename="<NAME_OF_THE_BACKUP_FILE>"
```

Notice the `--filename` flag in the end. It must match the name of a backup file in your S§ without the subdirectory name that you configured in `config.yml`. So for example if the backup lives under `backups/my_backup.sql.gz`, only do `--filename="my_backup.sql.gz"`.

# Under the hood

This little backup tool tries to use as few dependencies as possible. The major dependencies are:

- [cobra](https://github.com/spf13/cobra) for creating the command line arguments.
- [AWS SDK V2](https://github.com/aws/aws-sdk-go-v2) for interacting with AWS S3.
- [Go Dotenv](github.com/joho/godotenv) for parsing the `.env` file.
- [Go yaml V3](https://github.com/go-yaml/yaml) for parsing `.yml` files.

The backups are created by executing `docker exec` shell commands via Go. You can see this under `db/dump.go`, for example.
