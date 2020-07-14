# twitter-quarantineTime

### it changes your twitter name to display how many days you've been in quarantine
basically it adds `stayed at home for %d days` to your name

# Required flags/vars

| Flag name         | Env name                    | Description                                               |
|-------------------|-----------------------------|-----------------------------------------------------------|
| --access-token    | TWITTER_ACCESS_TOKEN        | twitter app access token                                  |
| --access-secret   | TWITTER_ACCESS_TOKEN_SECRET | twitter app access token secret                           |
| --consumer-key    | TWITTER_CONSUMER_KEY        | twitter app consumer key                                  |
| --consumer-secret | TWITTER_CONSUMER_SECRET     | twitter app consumer secret                               |
| --date            | Q_DATE                      | date from where you started the quarantine, ex 03/13/2020 |
| username-prefix   | TWITTER_USER_PREFIX         | the prefix to the username number of days in quarantine   |

---
pd: I had to use a [fork](https://github.com/acidobinario/go-twitter/) from [go-twitter](https://github.com/dghubble/go-twitter) to add the [UpdateUser](https://github.com/acidobinario/go-twitter/commit/915351721ff0ac85000a29ef06215165d62d1424) endpoint 