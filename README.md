# miro-gopher

API client for accessing the MIRO API

Supports all non-enterprise plan MIRO API endpoints. _Enterprise plan endpoints will be added at a later date._

For anything not covered there's the `GET`, `POST`, `POST Multipart`, `PUT`, `PATCH` & `DELETE` methods that are open to use for any other API calls to MIRO.

Minimum required Go version : `1.18`


[![GoDoc](https://godoc.org/github.com/russ-davey/miro-gopher?status.svg)](http://godoc.org/github.com/russ-davey/miro-gopher)
[![Tests](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml/badge.svg?branch=main)](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/russ-davey/miro-gopher)](https://goreportcard.com/report/github.com/russ-davey/miro-gopher)

[![Tag](https://img.shields.io/github/v/tag/russ-davey/miro-gopher?style=plastic)](https://github.com/russ-davey/miro-gopher/tags)
[![GoDoc](https://img.shields.io/github/go-mod/go-version/russ-davey/miro-gopher?style=plastic)](https://go.dev/doc/go1.2)
[![License](https://img.shields.io/badge/License-MIT%202.0-blue.svg?style=plastic)](https://opensource.org/licenses/MIT)

![gopher.png](gopher.png)

## Installation

```
go get github.com/russ-davey/miro-gopher/miro
```

## Basic Usage

All interaction starts with a `miro.Client`. Create one with your MIRO token:

```Go
import "github.com/russ-davey/miro-gopher/miro"

client := miro.NewClient(token)
```

---
## Using the Native Functions
If there is part of MIRO API that you would like to access, but it isn't currently supported by this package,
then you can use the `Get`, `Post`, `Put`, `Patch`, & `Delete` functions to natively access them instead:

```go
client := NewClient(os.Getenv("MIRO_TOKEN"))

response := make(map[string]interface{})

err := client.Get("https://api.miro.com/v2/boards/3141592/items/16180339887", &response)
if err != nil {
    fmt.Printf("error: %v", err)
} else {
    jsonData, _ := json.Marshal(response)
    fmt.Printf("MIRO API Response: %s\n", jsonData)
}
```
---
## Using a Customised HTTP client
By default, this package will use a fine-tuned HTTP client, but you may want to use your own.

```go
client := NewClient(os.Getenv("MIRO_TOKEN"))

client.HTTPClient = &http.Client{Timeout: 500 * time.Millisecond}
```

---
## /boards API Methods

### Get

```go
client.Boards.Get("3141592")
```

### GetALL

```go
client.Boards.GetAll()
```

or with query parameters:

```go
client.Boards.GetAll(miro.BoardSearchParams{
    TeamID: "gophers",
    Sort: miro.SortAlphabetically,
})
```

or when there are more than 20 boards, then use the iterator:

```go
iter, err := client.Boards.GetAll(BoardSearchParams{TeamID: "gophers", Limit: "50"})
if err != nil {
    log.Fatalf("error: %v", err)
}

for {
    boards, err := iter.GetNext()
    if err == IteratorDone {
        break
    }

    for _, board := range boards.Data {
        fmt.Println(board.Name)
    }
}
```

### Create

```go
client.Boards.Create(miro.CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: miro.Policy{
        SharingPolicy: miro.SharingPolicy{
            Access:                            miro.AccessPrivate,
            InviteToAccountAndBoardLinkAccess: miro.InviteAccessEditor,
            TeamAccess:                        miro.AccessEdit,
        },
        PermissionsPolicy: miro.PermissionsPolicy{
            SharingAccess:                 miro.AccessBoardOwnersAndCoOwners,
            CopyAccess:                    miro.CopyAccessTeamEditors,
            CollaborationToolsStartAccess: miro.AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
})
```

### Copy

```go
client.Boards.Copy(miro.CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: miro.Policy{
        SharingPolicy: miro.SharingPolicy{
            Access:                            miro.AccessPrivate,
            InviteToAccountAndBoardLinkAccess: miro.InviteAccessEditor,
            TeamAccess:                        miro.AccessEdit,
        },
        PermissionsPolicy: miro.PermissionsPolicy{
            SharingAccess:                 miro.AccessBoardOwnersAndCoOwners,
            CopyAccess:                    miro.CopyAccessTeamEditors,
            CollaborationToolsStartAccess: miro.AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
},
    "3141592")
```

### Update
```go
client.Boards.Update(miro.CreateBoard{
    Description: "My New Board",
    Name:        "New MIRO Gopher",
    Policy: miro.Policy{
        SharingPolicy: miro.SharingPolicy{
            Access:                            miro.AccessPrivate,
            InviteToAccountAndBoardLinkAccess: miro.InviteAccessEditor,
            TeamAccess:                        miro.AccessEdit,
        },
        PermissionsPolicy: miro.PermissionsPolicy{
            SharingAccess:                 miro.AccessBoardOwnersAndCoOwners,
            CopyAccess:                    miro.CopyAccessTeamEditors,
            CollaborationToolsStartAccess: miro.AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
},
"3141592")
```

### Delete

```go
client.Boards.Delete("3141592")
```
---