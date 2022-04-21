# housePriceCrowler

Crowler for otodom.pl website.
Currently set to Poznan city.

Collects, stores to Firebase, and informs telegram bot on newly created available houses for rent.

Config:

- create your telegram bot using BotFather
- set variables in teleBot.go file

  var botToken = "###"
  
  var chatIds = []string{"###", "##"}
  
  
- upload Firestore json access file to gcp folder.
