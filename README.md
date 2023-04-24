# miro-gopher

API client for accessing the MIRO API

Currently, only supports the `/boards` endpoint, but more to follow soon.

For now the GET, POST & PUT methods are open to use for any other API calls to MIRO.

Minimum required Go version : `1.18`


[![GoDoc](https://godoc.org/github.com/russ-davey/miro-gopher?status.svg)](http://godoc.org/github.com/russ-davey/miro-gopher)
[![Tests](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml/badge.svg?branch=main)](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml)

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

### Create

```go
client.Boards.Create(CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: Policy{
        SharingPolicy: SharingPolicy{
            Access:                            miro.AccessPrivate,
            InviteToAccountAndBoardLinkAccess: miro.InviteAccessEditor,
            OrganizationAccess:                miro.AccessEdit,
            TeamAccess:                        miro.AccessEdit,
        },
        PermissionsPolicy: PermissionsPolicy{
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
client.Boards.Copy(CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: Policy{
        SharingPolicy: SharingPolicy{
            Access:                            miro.AccessPrivate,
            InviteToAccountAndBoardLinkAccess: miro.InviteAccessEditor,
            OrganizationAccess:                miro.AccessEdit,
            TeamAccess:                        miro.AccessEdit,
        },
        PermissionsPolicy: PermissionsPolicy{
            SharingAccess:                 miro.AccessBoardOwnersAndCoOwners,
            CopyAccess:                    miro.CopyAccessTeamEditors,
            CollaborationToolsStartAccess: miro.AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
},
    "3141592")
```

