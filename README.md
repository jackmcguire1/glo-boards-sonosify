# glo-boards-sonosify
![GIT KRAKEN](https://cdn.worldvectorlogo.com/logos/gitkraken.svg)
![SONOS](https://cdn.worldvectorlogo.com/logos/sonos.svg)


**glo-boards-sonosify** is built on top of the amazing NodeJS component [Node-Sonos-HTTP-API][sonos-api].

[git]:      https://git-scm.com/
[golang]:   https://golang.org/
[modules]:  https://github.com/golang/go/wiki/Modules
[sonos]:    https://github.com/jishi/node-sonos-http-api
[npm]:      https://www.npmjs.com/get-npm
[aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html
[aws-sam-cli]: https://github.com/awslabs/aws-sam-cli
[node]:  http://nodejs.org
[sonos-api]: https://github.com/jishi/node-sonos-http-api
[git-kraken-pat]: https://app.gitkraken.com/pats
[time-zone]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

## About
This project has been developed for the purposes of the https://www.gitkraken.com/glo-api-contest

**glo-boards-sonosify** is a application designed to notify a user at 8:45 UTC in the morning
regarding how many Glo Board Cards they have due that day and any modifications to their cards
that occured since the last working day.

This application will be especially useful to users
who work remotely and/or colloborate with users whom are based in different time zones.

## Installation
This section describes how you can build and deploy
the glo-boards-sonosify application.

### Requirements
- a AWS account
- Create AWS S3 bucket 'glo-boards-sonosify'
- [Node][node]
- [Npm][npm]
- [Git][git]
- [Go 1.11][golang]+
- [AWS CLI][aws-cli]
- [AWS SAM CLI][aws-sam-cli]
- [Node-Sonos-HTTP-API][sonos-api]

### Sonos HTTP API

Clone and setup [node-sonos-http-api][sonos-api] on your machine or a device of your choice
- `git clone https://github.com/jishi/node-sonos-http-api.git `
- `npm install --production`
- `npm start`

ensure you have the port :5050 exposed publicly within your network and that the sonos API 
   is accessible via the internet by typing http://{{publicIP}}:5005/zones.
<br>
<br>
Any FAQs or troubleshooting please consult [node-sonos-http-api's][sonos-api] README file.

### Configuration
please configure the following settings in env.json

- your [PAT][git-kraken-pat] Glo Boards API TOKEN with (card and boards read permissions)
- your [time zone][time-zone] (in tz database format)
- your working hours e.g. 9am-5pm

### build

You will need to activate [Modules][modules] for your version of Go, generally
by invoking `go` with the support `GO111MODULE=on` environment variable set.

#### deploy
1. `aws cloudformation package --template-file template.yml --output-template-file packaged.yml --s3-bucket glo-boards-sonosify`

2. `aws cloudformation deploy --template-file ./packaged.yml --stack-name glo-boards-sonosify --capabilities CAPABILITY_IAM`





