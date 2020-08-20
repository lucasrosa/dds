# Deceptive Domain Score
![coverage](https://img.shields.io/badge/coverage-94%25-brightgreen.svg)
![build](https://img.shields.io/badge/build-passing-brightgreen.svg)

## Features
**Feature**: Get the domain score  
    As a domain analyst  
    I want to know the deceptiveness score of a domain  
    To decide whether or not it should be evaluated by a real person

**Rules**

- The domain will be compared to a list of keywords
- The input is a domain name
- Natural language must be matched (vísa == visa)
- The score is 100 if part of the domain that matches a keyword exactly and the keyword is a brand name
- The score is 90 if part of the domain matches a keyword exactly and the keyword is a common word
- The score is 80 if part of the domain is similar and the keyword is a brand name
- The score is 70 if part of the domain is similar and the keyword is a common word
- The score is 50 if the domain contains numbers
- The score is 0 if the domain is not in any of the categories listed above

**Scenario**: testing a deceptive domain  
    **Given** the domain "**visa-card-help.com**"  
    **When** I run the script on it  
    **Then** it should result in a score of **100** 

**Scenario**: testing a domain that contains a number  
    **Given** the domain "**nfl-mega-league2019.com**"  
    **When** I run the script on it  
    **Then** it should result in a score of **50**  

# Code
## Business Logic
### Entities
- domain
- keyword

### Public
- Domain
    - GetDeceptiveScore

## Requirements
### For the core business logic
- Go 1.12.x
- Remember to set ```export GO111MODULE=on;```

### For the AWS Serverless adapter
- Serverless Framework >=1.28.0
- An AWS account

## Running tests
```
// In the root folder of the repository run
go test ./...

// To run tests showing the test coverage
go test ./... -cover
```

## Building and deploying the AWS Serverless adapter

#### How to build
```make build```
#### How to deploy
```make deploy``` or ```sls deploy```
#### Deploying a single function
```sls deploy -f clientget```      

### TODO

- [ ] Create an architecture diagram
- [ ] Create a flow diagram for how the score is defined
- [ ] Remove tld (cleanup)
- [ ] Convert to lowercase (cleanup)
- [ ] Convert from punycode (cleanup)
- [ ] Different check for prefix/suffix (login, signin, verify...)
- [ ] Artificial intelligence (Machine learning) detection
- [ ] Create Go Benchmarks :revolving_hearts:
- [ ] Create Go Examples
- [X] Run each algorithm in a separate Go Routine :heart_eyes: (based on Go Benchmarks, this is and slower option)

