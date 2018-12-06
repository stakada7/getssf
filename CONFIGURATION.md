# 1. prepare config files.
- conf/account.yaml
- conf/keys.yaml

## conf/account.yaml
you should create keys:
- track:
- follow:

track is twitter account name.
e.g.) stakada7

follow is twitter account id.
e.g.) 1234567890

## conf/keys.yaml
you should create keys:
- TWITTER_CONSUMER_KEY:
- TWITTER_CONSUMER_SECRET:
- TWITTER_OAUTH_TOKEN:
- TWITTER_OAUTH_TOKEN_SECRET:

TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET, TWITTER_OAUTH_TOKEN, TWITTER_OAUTH_TOKEN_SECRET is twitter app's configuration.

- TWITTER_OAUTH_TOKEN2:
- TWITTER_OAUTH_TOKEN_SECRET2:

If you want to use multiple connections between twitter api, you should create key. (TWITTER_OAUTH_TOKEN{X}, TWITTER_OAUTH_TOKEN_SECRET{X})
