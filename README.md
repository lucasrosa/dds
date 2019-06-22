**Feature**: Get the domain score  
    As a domain analyst  
    I want to know the deceptiveness score of a domain  
    To decide whether or not it should be evaluated by a real person

**Rules**

- There is a list of keywords
- The input is a domain name
- The score is 100 if the domain matches a keyword exactly and the keyword is a brand name
- The score is 90 if the domain matches a keyword exactly and the keyword is a word from the dictionary
- The score is 80 if the domain matches a keyword by changing/adding/deleting some letters and the keyword is a brand name
- The score is 70 if the domain matches a keyword by changing/adding/deleting some letters and the keyword is a word from the dictionary
- The score is 50 if the domain contains numbers
- The score is 0 if the domain is not in any of the categories listed above

**Scenario**: testing a deceptive domain  
    **Given** the domain "atendimentobb.com"  
    **When** I run the script on it  
    **Then** it should result in a score of 100 

**Scenario**: testing a domain that contains a number  
    **Given** the domain "meufutebol2019.com"  
    **When** I run the script on it  
    **Then** it should result in a score of 50  

## Requirements
Remember to set ```export GO111MODULE=on;```

## Building and deploying

#### How to build
```make build```
#### How to deploy
```make deploy``` or ```sls deploy```
#### Deploying a single function
```sls deploy -f clientget```

## Business Logic
### Entities
- domain

### Public
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

