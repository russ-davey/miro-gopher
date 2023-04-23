# miro-gopher

API client for accessing the MIRO API

Currently, only supports the `/boards` endpoint, but more to follow soon.

For now the GET, POST & PUT methods are open to use for any other API calls to MIRO.

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
client.Boards.GetAll(BoardQueryParams{
    TeamID: "gophers",
    Sort: SortAlphabetically,
})
```

### Create

```go
client.Boards.Create(CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: Policy{
        SharingPolicy: SharingPolicy{
            Access:                            AccessPrivate,
            InviteToAccountAndBoardLinkAccess: InviteAccessEditor,
            OrganizationAccess:                AccessEdit,
            TeamAccess:                        AccessEdit,
        },
        PermissionsPolicy: PermissionsPolicy{
            SharingAccess:                 AccessBoardOwnersAndCoOwners,
            CopyAccess:                    CopyAccessTeamEditors,
            CollaborationToolsStartAccess: AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
})
```

### Copy

``` go
client.Boards.Copy(CreateBoard{
    Description: "My Board",
    Name:        "MIRO Gopher",
    Policy: Policy{
        SharingPolicy: SharingPolicy{
            Access:                            AccessPrivate,
            InviteToAccountAndBoardLinkAccess: InviteAccessEditor,
            OrganizationAccess:                AccessEdit,
            TeamAccess:                        AccessEdit,
        },
        PermissionsPolicy: PermissionsPolicy{
            SharingAccess:                 AccessBoardOwnersAndCoOwners,
            CopyAccess:                    CopyAccessTeamEditors,
            CollaborationToolsStartAccess: AccessBoardOwnersAndCoOwners,
        },
    },
    TeamID: "gophers",
},
    "3141592")
```

