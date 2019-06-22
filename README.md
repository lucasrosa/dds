# Requirements
Remember to set export GO111MODULE=on; 

# Building and deploying

#### How to build
```make build```
#### How to deploy
```make deploy``` or ```sls deploy```
#### Deploying a single function
```sls deploy -f clientget```

# Business Logic
## Entities
- domain

## Public
- Domain
    - GetDeceptiveScore
    https://godoc.org/golang.org/x/text/search
    https://github.com/dzstudio/similar-text
    https://github.com/agext/levenshtein
    

### Logic

#### Cleanup
- Remove tld
- Convert to lowercase
- Convert from punycode

#### Detection
- Consider natural language matching (bahía == bahia)
- Contains numbers
- Contains company/brand name
- Contains suspiscious word (christmas, world cup)
- Is similar to existing phishing
- Contains prefix/suffix (login, signin, verify...)

