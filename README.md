# Kobutor

**Kobutor** is a mailing service for sending emails to anyone from anyone! So far, it is using the
[Sendgrid](https://sendgrid.com/) email delivery service api. I plan to integrate other services in the future as well.
Send an email with a simple ```POST``` request or even from your terminal!

## Running the app

1. [Download the app](https://github.com/Anondo/kobutor-bin.git)
1. ```./kobutor server``` to start the server or,
1. ```./kobutor send -f '{"name":"Anondo", "email":"aanondos@gmail.com"}' -t '{"name":"Ahmad Ananda","email":"ananda.anabil@pathao.com"}' -s "This is a test message" -b "<h4>Testing the kobutor service</h4>" --type=html ```

**Important Note**: Before you can run the app, you need to generate a key for using the api of the sendgrid and add it to the key field under ```sendgrid``` in the ```kobutor_config.yml``` file. Also this app uses ```basic_auth``` as authorization so you need to add the credentials manually in the config file under the ```auth``` field. A default ```username=uname password=secret``` key-value pair is already provided. Follow that pattern. 

## Payload
```javascript
{
"from":{
  "name": "Ahmad Ananda",
  "email": "ananda.anabil@pathao.com"
},
"to":{
  "name": "Anondo",
  "email": "aanondos@gmail.com"
},
"subject": "Testing Kobutor",
"body": "This is a test email",
"type": "text/plain",
"attachments": [
  {
    "type": "text/text",
    "filename": "hello.txt",
    "disposition": "attachment",
    "content": "Hello World"
  }
]

}

```

## Commands

### Server
This command starts an http server on the default port ```8080```. You can pass the flag ```-p``` or ```--port``` to
change the port to run the server on.

### Send
This command enables the user to send an email without any server or third-party tools rather, this lets the user send email from the terminal.
There are some must provide flags though.
1. ```-f or --from```: A json containing two fields ```name``` & ```email``` about the sender
1. ```-t or --to```: A json containing two fields ```name``` & ```email``` about the receiver
1. ```-s or --subject``` : The subject of the email which is a string
1. ```-b or --body``` : The main body of the email which is also a string
1. ```-p or --type``` : The type of the body being sent. This is also string

No ```attachment``` features available for the ```send``` command so far. For the sake of simplicity.
