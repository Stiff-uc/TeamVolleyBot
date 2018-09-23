# PoGoPollrBot
This is a modified version of telegram bot pollrBot that is used for Pokemon Go
raid polls. Poll will automatically create raid times after user gives the raid
starting time. Raid starting time is added as is and following raid times are
rounded to the next 0/5 minutes.

Original pollrBot:
https://github.com/jheuel/pollrBot/

The bot uses inline queries and feedback to inline queries, which have to be
enabled with the telegram [@BotFather](https://telegram.me/BotFather).

## Usage

Add your api key to env.list and build with docker.

```
docker build --tag pollrbot .
docker run -p 8443:8443 --env-file env.list \
  -v $(pwd)/db:/db -u "$(id -u):$(id -g)" pollrbot
```
