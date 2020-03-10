# active-campaign-sdk-go #
[![Build Status](https://travis-ci.com/benkrig/active-campaign-sdk-go.svg?token=zD75aqrV8gE1Q1ghw6yU&branch=master)](https://travis-ci.com/benkrig/active-campaign-sdk-go)

**active-campaign-sdk-go** provides access to the [Active Campaign API V3](https://developers.activecampaign.com/reference) for Go. Currently, it's heavily under development.
## Usage ##

```go
package main

import "github.com/benkrig/active-campaign-sdk-go" 
```
Construct a new client, then use the services available within the client to access the Active Campaign API.

```go
package main

import (
    ac "github.com/benkrig/active-campaign-sdk-go"
    "os"
) 

func main() {
    baseURL := os.Getenv("YOUR_BASE_URL_KEY")
    token := os.Getenv("YOUR_TOKEN_KEY")

    a, err := ac.NewClient(&ac.ClientOpts{
            BaseUrl: baseURL, 
            Token: token,
        },
    )
    if err != nil {
        panic(err)
    }

    c := ac.CreateContactRequest{
        Email: "test@email.com",
        FirstName: "testf",
        LastName: "testl",
        Phone: "1234567890",
    }

    contact, _, err := a.Contacts.Create(&c)
    if err != nil {
        panic(err)
    }
}
```

## Code structure

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github) and [andygrunwald/go-jira](https://github.com/andygrunwald/go-jira).

Everything is based around the Client. The Client contains various services for resources found in the Active Campaign API, like Contacts, or Automations. Each service implements actions for its respective resource(s).

## Contribution ##

PR's are always welcome! The SDK is still being heavily developed and is missing many entities.

It doesn't matter if you are not able to write code.
Creating issues or holding talks and helping other people use the SDK is contribution as well!
A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).