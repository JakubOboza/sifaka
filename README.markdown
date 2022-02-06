# Sifaka

Sifaka is a tool to monitor your x509 certificates or simply websites certificates expirey date.

If your organisation handles certs in non automatics renwal way, needs calendars and from time to time someone simply forgets to request new cert it can be a pain. With sifaka you just point it to your https web page or public cert in a file and it will notify you 29 days before it expires.

Sifaka interface is mostly CLI but it also runs a server than cant notify you via Slack (version v0.0.1).

It will notify You 29 days before certificate will expire. (It will repeat notifications every 4 hours until you fix it ;) )

#### The Name

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Propithecus_coquereli_02.jpg/1024px-Propithecus_coquereli_02.jpg" alt="sifaka picture from wikipedia" style="height:240px;"/>

*A sifaka is a lemur of the genus Propithecus from the family Indriidae within the order Primates. The name of their family is an onomatopoeia of their characteristic "shi-fak" alarm call.*

Source: Wikipedia: [https://en.wikipedia.org/wiki/Sifaka](https://en.wikipedia.org/wiki/Sifaka)

# Usage

**!important** If you plan use any command other than `check` before sure `data` directory exists and run at least once `sifaka server` it will create database file and run all migrations.

Command you can use

* `sifaka check` - standalone command to check cert expiration date from --url= url or --file= file
* `sifaka add` - checks cert and adds it to sifaka database for tracking
* `sifaka list` - lists all certs in database with expirations in CSV format
* `sifaka remove` - removes by --id= (from list) cert from sifaka database
* `sifaka server` - runs the app server, periodic checks and notifications and hosts on --port= simple website that lists the certificates

#### Example

To add cert to sifaka database:

via website url
```
./bin/sifaka add -u https://google.com
```
or via file
```
./bin/sifaka add -f sample.cer
```

To run server on a selected port eg. 7788

```
./bin/sifaka server -p 7788
```

to list all certs in sifaka db

```
./bin/sifaka list
```

to remove cert from db via id

```
./bin/sifaka remove --id=69420
```


# Install

By default sifaka stores its data in `data/sifaka.db` where the app exists.
Be sure that the directory `data` exists and is owned in a way sifaka can make files in it. That is the only requirement for the app to run.

#### from binary

Via releases. Go to [https://github.com/JakubOboza/sifaka/releases](https://github.com/JakubOboza/sifaka/releases)

1. Download the binary
2. Untar it eg `tar -zxvf sifaka.tar.gz`
3. run it `./sifaka server`
4. visit [http://localhost:6123](http://localhost:6123)

#### from source

You will need go at least 1.16, sqlite3 and gcc in your system to build sifaka.

1. clone repo `git clone https://github.com/JakubOboza/sifaka`
2. cd sifaka
3. make build
4. run it `./bin/sifaka server`
4. visit [http://localhost:6123](http://localhost:6123)
