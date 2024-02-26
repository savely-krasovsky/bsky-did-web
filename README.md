# didweb

> ⚠️WIP

`didweb` is a simple utility that helps to migrate you DID to [Bluesky PDS](https://github.com/bluesky-social/pds).

## Compiling
```
go build -o didweb main.go
```

## Example
```bash
PRIVKEY=$(didweb genkey)
PUBKEY=$(echo -n $PRIVKEY | didweb pubkey)
didweb gendid --handle alice.domain.tld --pubkey $PUBKEY
# upload this did to your .well-known directory
# now you can try to sign up
didweb sign --privkey $PRIVKEY --iss did:web:alice.domain.tld --aud did:web:pds.domain.tld --exp 180 | didweb createAccount --pds https://pds.domain.tld --handle alice.domain.tld --invite pds-domain-tld-invite-code --email alice@domain.tld --password password123
# now you will get new JWT token to complete registration
```
