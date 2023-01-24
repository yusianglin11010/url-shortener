# url shortener

An url shortener server realized by golang and redis. Followed the guide from [Let's build a URL shortener in Go - Part I](https://www.eddywm.com/lets-build-a-url-shortener-in-go/). Besides, I added extra feature for deployment, a docker file and docker-compose for quick start in any environment.

## Built with
[![Docker][docker-shield]][docker-url]
[![Golang][golang-shield]][golang-url]
[![Redis][redis-shield]][redis-url]

## Quick Start

1. build go server's image
```shell
docker build -t go-url-server .
```
2. run service on
```shell
docker-compose up
```
3. shorten your url
```shell
curl -X POST {your-host}/url -d '{"long_url": "https:google.com"}'
# expected output
# {your-host}/Lhr4BWAi
```
4. check if server redirect {your-host}/Lhr4BWAi to https://google.com

## How It Works

### Convert original URL to another shorter identical string 
- Hash Original URL With Sha256 
- Encode the Hashed Result With Base58
   - This server encoded URL with base58 for improving user experience. It avoided tje usage of the following confusing characters: "O", "0", "I", "l", "+" and "/".

### Redirect the generated shorter url to the original one
- When user hit the host IP with shortened url, server would lookup the original url in redis and redirect the user to original url. 

## Roadmap
- [x] Add go server Dockerfile
- [x] Add docker-compose
- [ ] Check if generated duplicated URL

[docker-shield]: https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white
[docker-url]: https://www.docker.com/
[golang-url]:  https://go.dev/
[golang-shield]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[redis-url]: https://redis.io/
[redis-shield]: https://img.shields.io/badge/redis-DC382D?style=for-the-badge&logo=redis&logoColor=white






