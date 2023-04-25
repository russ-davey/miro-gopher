# miro-gopher

API client for accessing the MIRO API

Currently, only supports the `/boards` endpoint, but more to follow soon.

For now the `GET`, `POST`, `PUT`, `PATCH` & `DELETE` methods are open to use for any other API calls to MIRO.

Minimum required Go version : `1.18`


[![GoDoc](https://godoc.org/github.com/russ-davey/miro-gopher?status.svg)](http://godoc.org/github.com/russ-davey/miro-gopher)
[![Tests](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml/badge.svg?branch=main)](https://github.com/russ-davey/miro-gopher/actions/workflows/miro-gopher.yml)

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