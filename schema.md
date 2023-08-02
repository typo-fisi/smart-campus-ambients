# Schema

So the idea of this is to get the information of the classes through endpoints in this api, however, we must start with the notion of what an ambient is, since it is the base of the endpoints structure:

## `type ambient`

```go
package types

type Ambient struct{
    Category string
    Id string;
    Location struct{
        piso int;
        geometría string;
        pavilion string;
    };
    Description string;
    Fotos []string;
    Tags []string;
}
```

An ambient is basically a room, it may be of type `administrativo`, `salones`, `sshh` o `laboratorios`, its `.Id` field is just `"Category.number"` where the `number` is some arbitrary number given by the campus.

The first API to use should be then `GET /ambients` which shall list all existent ambients with its respective `ID`'s.

## `type AssignmentGroup`

```go
type AssignmentGroup struct {
    AmbientID string;
    Group string;
    Professor Professor;
    Enrolled int;

    Schedules []struct{
        Day string;
        ClassType string;
        From int;
        To int;
    };
}
```

So each ambient (which serves as either a "salon" or "laboratorio") will have a current assignment in it, based on this we can:

1. Get all of some ambient's assignments:
- `GET /ambients/{ID}`
2. Get the data of each one of those assignments:
- `GET /assignmentgroup/{ID}`
3. Get the data of the Professor which is currently at the ambient by the assignment it is currently hosting:
- `GET /professors/{Professor.code}`











