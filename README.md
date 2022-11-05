# go.discord-publisher
<a href="https://www.repostatus.org/#wip"><img src="https://www.repostatus.org/badges/latest/wip.svg" alt="Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public." /></a><br>

This service is a part of the repo: [microservices-training-ground](https://github.com/JacekKorta/microservices-training-ground)<br>
This is simply a discord message publisher. In the first step it reads a message from the queue. The service reads the tags added by the previous service. As the final step messages are published on the discord channel with various topics depending on the receiving tags.   

### How to run?

You should run this service via docker compose in main repo [microservices-training-ground](https://github.com/JacekKorta/microservices-training-ground)

Create env file:
```bash
cp .env.example .env
```

Additional info about env variables:

DISCORD_WEBHOOK - a webhook for your channel where you will receive messages from the service. More info here: [discord api docs](https://discord.com/developers/docs/intro).
